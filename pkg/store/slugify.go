package store

import (
	"fmt"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

// Slugify generates a slug.
func Slugify(tx *gorm.DB, value, id string) string {
	var slug string

	for i := 0; true; i++ {
		if i == 0 {
			slug = slugify.Slugify(value)
		} else {
			slug = slugify.Slugify(
				fmt.Sprintf("%s-%s", value, uniuri.NewLen(6)),
			)
		}

		var count int64

		query := tx.Where(
			"slug = ?",
			slug,
		)

		if id != "" {
			query = query.Not(
				"id",
				id,
			)
		}

		if err := query.Count(
			&count,
		).Error; err == nil || count == 0 {
			break
		}
	}

	return slug
}
