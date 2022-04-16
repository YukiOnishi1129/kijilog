package initializer

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/YukiOnishi1129/kijilog/backend/database"
	"github.com/YukiOnishi1129/kijilog/backend/graph"
	"github.com/YukiOnishi1129/kijilog/backend/graph/generated"
	"github.com/YukiOnishi1129/kijilog/backend/service"
	"github.com/YukiOnishi1129/kijilog/backend/util/auth"
	awsutil "github.com/YukiOnishi1129/kijilog/backend/util/aws"
	"github.com/YukiOnishi1129/kijilog/backend/util/view"
	"github.com/go-chi/chi"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Init(router *chi.Mux) (*handler.Server, error) {
	db, dbErr := database.Init()
	if dbErr != nil {
		return nil, dbErr
	}

	router.Use(auth.MiddleWare(db))

	awsSession, awsErr := awsutil.Init()
	if awsErr != nil {
		return nil, awsErr
	}

	awsUtil := awsutil.LazyInitTodoService(awsSession)

	userService := service.LazyInitUserService(db, awsUtil)
	todoService := service.LazyInitTodoService(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(userService, todoService)}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var appErr view.AppError
		if errors.As(err, &appErr) {
			return &gqlerror.Error{
				Message: appErr.Msg,
				Extensions: map[string]interface{}{
					"code": appErr.Code,
				},
			}
		}
		return err
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return &gqlerror.Error{
			Extensions: map[string]interface{}{
				"code":  view.ErrorStatusInternalServerError,
				"cause": err,
			},
		}
	})

	return srv, nil
}
