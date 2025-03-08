package v1

import (
	"encoding/json"
	"net/http"

	"github.com/drexedam/gravatar"
)

func gravatarFor(email string) string {
	return gravatar.New(email).
		Size(64).
		Default(gravatar.Identicon).
		Rating(gravatar.G).
		AvatarURL()
}

func fromSearch(val *string) string {
	if val == nil {
		return ""
	}

	return FromPtr(val)
}

func toPageParams(sort *SortColumnParam, limit *PagingLimitParam, offset *PagingOffsetParam, search *SearchQueryParam) (string, int64, int64, string) {
	sortResult := ""

	if sort != nil {
		sortResult = string(FromPtr(sort))
	}

	limitResult := int64(100)

	if limit != nil {
		limitResult = int64(FromPtr(limit))
	}

	offsetResult := int64(0)

	if offset != nil {
		offsetResult = int64(FromPtr(offset))
	}

	searchResult := ""

	if search != nil {
		searchResult = string(FromPtr(search))
	}

	return sortResult, limitResult, offsetResult, searchResult
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
