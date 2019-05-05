package model

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"github.com/ungerik/go-gravatar"
	"golang.org/x/crypto/bcrypt"
)

// Users is simply a collection of user structs.
type Users []*User

// User represents a user model definition.
type User struct {
	ID        int64     `json:"id" xorm:"pk autoincr"`
	Slug      string    `json:"slug" xorm:"unique not null"`
	Username  string    `json:"username" xorm:"unique not null"`
	Email     string    `json:"email" xorm:"unique not null"`
	Hash      string    `json:"-" xorm:"unique not null"`
	Password  string    `json:"password,omitempty" xorm:"-"`
	Hashword  string    `json:"-" xorm:"not null"`
	Avatar    string    `json:"avatar,omitempty" xorm:"-"`
	Active    bool      `json:"active" xorm:"default false not null"`
	Admin     bool      `json:"admin" xorm:"default false not null"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	Teams     Teams     `json:"teams,omitempty" xorm:"-"`
	TeamUsers TeamUsers `json:"team_users,omitempty" xorm:"-"`
	Mods      Mods      `json:"mods,omitempty" xorm:"-"`
	UserMods  UserMods  `json:"user_mods,omitempty" xorm:"-"`
	Packs     Packs     `json:"packs,omitempty" xorm:"-"`
	UserPacks UserPacks `json:"user_packs,omitempty" xorm:"-"`
}

// AfterFind invokes required after loading a record from the database.
func (u *User) AfterFind(db *gorm.DB) {
	u.Avatar = gravatar.SecureUrlDefault(u.Email, gravatar.Retro)
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

	if u.Hash == "" {
		u.Hash = base32.StdEncoding.EncodeToString(
			securecookie.GenerateRandomKey(32),
		)
	}

	return nil
}

// BeforeDelete invokes required actions before deletion.
func (u *User) BeforeDelete(tx *gorm.DB) error {
	if err := tx.Model(u).Association("Teams").Clear().Error; err != nil {
		return err
	}

	// TODO(tboerger): Prevent delete if user is last owner
	if err := tx.Model(u).Association("Mods").Clear().Error; err != nil {
		return err
	}

	// TODO(tboerger): Prevent delete if user is last owner
	if err := tx.Model(u).Association("Packs").Clear().Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
// func (u *User) Validate(db *gorm.DB) {
// 	if !govalidator.StringLength(u.Username, "2", "255") {
// 		db.AddError(fmt.Errorf("Username should be longer than 2 and shorter than 255"))
// 	}

// 	if u.Username != "" {
// 		notFound := db.Where(
// 			"username = ?",
// 			u.Username,
// 		).Not(
// 			"id",
// 			u.ID,
// 		).First(
// 			&User{},
// 		).RecordNotFound()

// 		if !notFound {
// 			db.AddError(fmt.Errorf("Username is already present"))
// 		}
// 	}

// 	if u.Hash != "" {
// 		notFound := db.Where(
// 			"hash = ?",
// 			u.Hash,
// 		).Not(
// 			"id",
// 			u.ID,
// 		).First(
// 			&User{},
// 		).RecordNotFound()

// 		if !notFound {
// 			db.AddError(fmt.Errorf("Hash is already present"))
// 		}
// 	}

// 	if !govalidator.IsEmail(u.Email) {
// 		db.AddError(fmt.Errorf(
// 			"Email must be a valid email address",
// 		))
// 	}

// 	if u.Email != "" {
// 		normalized, _ := govalidator.NormalizeEmail(
// 			u.Email,
// 		)

// 		notFound := db.Where(
// 			"email = ?",
// 			normalized,
// 		).Not(
// 			"id",
// 			u.ID,
// 		).First(
// 			&User{},
// 		).RecordNotFound()

// 		if !notFound {
// 			db.AddError(fmt.Errorf("Email is already present"))
// 		}
// 	}

// 	if db.NewRecord(u) {
// 		if !govalidator.StringLength(u.Password, "5", "255") {
// 			db.AddError(fmt.Errorf("Password should be longer than 5 and shorter than 255"))
// 		}
// 	}
// }

// MatchPassword checks if the provided password matches.
func (u *User) MatchPassword(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.Hashword),
		[]byte(password),
	)
}
