package model

import (
  "gopkg.in/go-playground/validator.v8"
)

var (
  validate *validator.Validate
)

func init() {
  validate = validator.New(
    &validator.Config{
      TagName: "validate",
    },
  )
}
