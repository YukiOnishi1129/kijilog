package view

import (
	"github.com/YukiOnishi1129/kijilog/backend/database/entity"
	"github.com/YukiOnishi1129/kijilog/backend/graph/model"
	"github.com/YukiOnishi1129/kijilog/backend/util/timeutil"
	"strconv"
)

func NewUserFromModel(entity *entity.User) *model.User {
	resUser := model.User{
		ID:        strconv.FormatInt(entity.ID, 10),
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: timeutil.TimeFormat(entity.CreatedAt),
		UpdatedAt: timeutil.TimeFormat(entity.UpdatedAt),
	}

	if entity.ImageURL.Valid {
		imageURL := entity.ImageURL.String
		resUser.ImageURL = &imageURL
	}

	if entity.DeletedAt.Valid {
		deletedAt := timeutil.TimeFormat(entity.DeletedAt.Time)
		resUser.DeletedAt = &deletedAt
	}

	return &resUser
}
