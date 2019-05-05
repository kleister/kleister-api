package user

import (
	"github.com/go-kit/kit/log"
)

type LoggingOptions struct {
	Service Service
	Logger  log.Logger
}

func NewLogging(opts LoggingOptions) Service {
	return &logging{
		service: opts.Service,
		logger:  opts.Logger,
	}
}

type logging struct {
	service Service
	logger  log.Logger
}

func (l *logging) ListUser() {

}

func (l *logging) CreateUser() {

}

func (l *logging) UpdateUser() {

}

func (l *logging) DeleteUser() {

}

func (l *logging) ShowUser() {

}
