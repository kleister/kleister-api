package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

const (
	// UserUsernameMinLength is the minimum length of the username.
	UserUsernameMinLength = "3"

	// UserUsernameMaxLength is the maximum length of the username.
	UserUsernameMaxLength = "255"

	// UserPasswordMinLength is the minimum length of the password.
	UserPasswordMinLength = "3"

	// UserPasswordMaxLength is the maximum length of the password.
	UserPasswordMaxLength = "255"
)

// UserDefaultOrder is the default ordering for user listings.
func UserDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Order(
		"users.username ASC",
	)
}

// Users is simply a collection of user structs.
type Users []*User

// User represents a user model definition.
type User struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	Permission *Permission `json:"permission"`
	Slug       string      `json:"slug" sql:"unique_index"`
	Username   string      `json:"username" sql:"unique_index"`
	Password   string      `json:"password"`
	Email      string      `json:"email"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// BeforeSave invokes required actions before persisting.
func (u *User) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {

		u.Slug = slugify.Slugify(u.Username)
		// Fill the slug with slugified username

	}

	u.Email, _ = govalidator.NormalizeEmail(u.Email)

	return nil
}

// Defaults prefills the struct with some default values.
func (u *User) Defaults() {
	u.Permission = &Permission{}
	u.Permission.Defaults()
}

// Validate does some validation to be able to store the record.
func (u *User) Validate(db *gorm.DB) {
	if u.Username == "" {
		db.AddError(fmt.Errorf("Username is a required attribute"))
	}

	if !govalidator.StringLength(u.Username, UserUsernameMinLength, UserUsernameMaxLength) {
		db.AddError(fmt.Errorf("Username should be longer than 3 characters"))
	}

	if u.Email == "" {
		db.AddError(fmt.Errorf("Email is a required attribute"))
	}

	if !govalidator.IsEmail(u.Email) {
		db.AddError(fmt.Errorf("Email must be a valid email address"))
	}

	if db.NewRecord(u) {
		if !govalidator.StringLength(u.Password, UserPasswordMinLength, UserPasswordMaxLength) {
			db.AddError(fmt.Errorf("Password should be longer than 3 characters"))
		}
	}

	if u.Username != "" {
		count := 1

		db.Where("username = ?", u.Username).Not("id", u.ID).Find(
			&User{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("Username is already present"))
		}
	}

	if u.Email != "" {
		count := 1

		db.Where("email = ?", u.Email).Not("id", u.ID).Find(
			&User{},
		).Count(
			&count,
		)

		if count > 0 {
			db.AddError(fmt.Errorf("Email is already present"))
		}
	}
}
