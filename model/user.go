package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Users is simply a collection of user structs.
type Users []*User

// User represents a user model definition.
type User struct {
	ID         int         `json:"id" gorm:"primary_key"`
	Permission *Permission `json:"permission"`
	Slug       string      `json:"slug" sql:"unique_index"`
	Username   string      `json:"username" sql:"unique_index"`
	Email      string      `json:"email" sql:"unique_index"`
	Password   string      `json:"password,omitempty" sql:"-"`
	Hashword   string      `json:"-"`
	Active     bool        `json:"active" sql:"default:false"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Mods       Mods        `json:"mods,omitempty" gorm:"many2many:user_mods;"`
	Packs      Packs       `json:"packs,omitempty" gorm:"many2many:user_packs;"`
}

// BeforeSave invokes required actions before persisting.
func (u *User) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = slugify.Slugify(u.Username)
			} else {
				u.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", u.Username, i),
				)
			}

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&User{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	if u.Email != "" {
		email, err := govalidator.NormalizeEmail(
			u.Email,
		)

		if err != nil {
			return fmt.Errorf("Failed to normalize email")
		}

		u.Email = email
	}

	if u.Password != "" {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(u.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return fmt.Errorf("Failed to encrypt password")
		}

		u.Hashword = string(encrypt)
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *User) AfterDelete(tx *gorm.DB) error {
	if err := tx.Delete(u.Permission).Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Mods").Clear().Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Packs").Clear().Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *User) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Username, "2", "255") {
		db.AddError(fmt.Errorf("Username should be longer than 2 and shorter than 255"))
	}

	if u.Username != "" {
		notFound := db.Where(
			"username = ?",
			u.Username,
		).Not(
			"id",
			u.ID,
		).First(
			&User{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Username is already present"))
		}
	}

	if !govalidator.IsEmail(u.Email) {
		db.AddError(fmt.Errorf(
			"Email must be a valid email address",
		))
	}

	if u.Email != "" {
		normalized, _ := govalidator.NormalizeEmail(
			u.Email,
		)

		notFound := db.Where(
			"email = ?",
			normalized,
		).Not(
			"id",
			u.ID,
		).First(
			&User{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Email is already present"))
		}
	}

	if db.NewRecord(u) {
		if !govalidator.StringLength(u.Password, "5", "255") {
			db.AddError(fmt.Errorf("Password should be longer than 5 and shorter than 255"))
		}
	}
}
