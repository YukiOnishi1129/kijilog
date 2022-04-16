package view

import (
	"github.com/YukiOnishi1129/kijilog/backend/database/entity"
	"github.com/YukiOnishi1129/kijilog/backend/graph/model"
	"github.com/YukiOnishi1129/kijilog/backend/util/timeutil"
	"strconv"
)

func NewTodoFromModel(entity *entity.Todo) *model.Todo {
	resUser := model.User{
		ID:        strconv.FormatInt(entity.R.User.ID, 10),
		Name:      entity.R.User.Name,
		Email:     entity.R.User.Email,
		CreatedAt: timeutil.TimeFormat(entity.R.User.CreatedAt),
		UpdatedAt: timeutil.TimeFormat(entity.R.User.UpdatedAt),
	}
	resTodo := model.Todo{
		ID:        strconv.FormatInt(entity.ID, 10),
		Title:     entity.Title,
		Comment:   entity.Comment,
		CreatedAt: timeutil.TimeFormat(entity.CreatedAt),
		UpdatedAt: timeutil.TimeFormat(entity.UpdatedAt),
	}

	if entity.R.User.ImageURL.Valid {
		imageURL := entity.R.User.ImageURL.String
		resUser.ImageURL = &imageURL
	}

	if entity.R.User.DeletedAt.Valid {
		userDeletedAt := timeutil.TimeFormat(entity.R.User.DeletedAt.Time)
		resUser.DeletedAt = &userDeletedAt
	}

	resTodo.User = &resUser

	if entity.DeletedAt.Valid {
		deletedAt := timeutil.TimeFormat(entity.DeletedAt.Time)
		resTodo.DeletedAt = &deletedAt
	}

	return &resTodo
}
