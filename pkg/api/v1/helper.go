package v1

import (
	"encoding/json"
	"net/http"

	"github.com/drexedam/gravatar"
	"github.com/gopad/gopad-api/pkg/model"
)

func gravatarFor(email string) string {
	return gravatar.New(email).
		Size(64).
		Default(gravatar.Identicon).
		Rating(gravatar.G).
		AvatarURL()
}

func toListParams(search, sort, order *string, limit, offset *int) model.ListParams {
	result := model.ListParams{}

	if search != nil {
		result.Search = FromPtr(search)
	}

	if sort != nil {
		result.Sort = FromPtr(sort)
	}

	if order != nil {
		result.Order = FromPtr(order)
	}

	if limit != nil {
		result.Limit = FromPtr(limit)
	} else {
		result.Limit = 50
	}

	if offset != nil {
		result.Offset = FromPtr(offset)
	}

	return result
}

// ToPtr transform input to a pointer.
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr transform input from pointer.
func FromPtr[T any](v *T) T {
	return *v
}

// Notify simply processes a notification.
func Notify(w http.ResponseWriter, notification Notification) {
	w.Header().Set("Content-Type", "application/json")

	if notification.Status != nil {
		w.WriteHeader(FromPtr(notification.Status))
	}

	_ = json.NewEncoder(w).Encode(notification)
}
