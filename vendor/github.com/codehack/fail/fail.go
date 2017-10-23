// Copyright 2017 Codehack. All rights reserved.
// For mobile and web development visit http://codehack.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fail

import (
	"fmt"
	"runtime"

	"net/http"
	"strconv"
	"strings"
)

const (
	messageOK         = "OK"
	messageNotFound   = "object not found"
	messageUnexpected = "an unexpected error has occurred"
)

// ErrUnspecified is an error with no cause or reason.
var ErrUnspecified = fmt.Errorf("unspecified error")

/*
Fail is an error that can be used in an HTTP response.

	- Status: the HTTP Status code of the response (400-4XX, 500-5XX)
	- Message: friendly error message (for clients)
	- Details: slice of error details. e.g., form validation errors.

*/
type Fail struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
	prev    error
	file    string
	line    int
}

// Error implements the error interface.
// This function should be used for logs and not sent to clients.
// It will unwrap all the stacked previous errors.
func (f *Fail) Error() string {
	return fmt.Sprintf("%s:%d: %s", f.file, f.line, f.prev.Error())
}

// String implements the fmt.Stringer interface, to make fails errors print nicely.
func (f *Fail) String() string {
	return f.Message
}

/*
Format implements the fmt.Formatter interface. This allows a fail to have
fmt.Sprintf verbs for its values. This is handy for sending to logs.

	Verb  Description
	----  ---------------------------------------------------
	%%    Percent sign
	%d    All fail details separated with commas (``Fail.Details``)
	%e    The original error (``error.Error``)
	%f    File name where the fail was called, minus the path.
	%l    Line of the file for the fail
	%m    The message of the fail (``Fail.Message``)
	%s    HTTP Status code (``Fail.Status``)

Example:

	// Print file, line, and original error.
	// Note: we use index [1] to reuse `f` argument.
	f := fail.Cause(err)
	fmt.Printf("%[1]f:%[1]l %[1]e", f)
	// Output:
	// alerts.go:123 missing argument to vars

*/
func (f *Fail) Format(s fmt.State, c rune) {
	var str string

	p, pok := s.Precision()
	if !pok {
		p = -1
	}

	switch c {
	case 'd':
		str = strings.Join(f.Details, ", ")
	case 'e':
		if f.prev == nil {
			str = ErrUnspecified.Error()
		} else {
			str = f.prev.Error()
		}
	case 'f':
		str = f.file
	case 'l':
		str = strconv.Itoa(f.line)
	case 'm':
		str = f.Message
	case 's':
		str = strconv.Itoa(f.Status)
	}
	if pok {
		str = str[:p]
	}
	s.Write([]byte(str))

}

// Cause wraps an error into a fail so it can be used in a response.
func Cause(prev error) *Fail {
	f := &Fail{
		prev: prev,
	}
	f.Caller(1)
	return f
}

/*
Because returns the previous error if it's a fail but keeps the current context
of where it was called. If the error err is not a fail, it will return an
Unexpected fail. Use this function when you know that err is a fail.

Example:

	func Authorize() error {
		return fail.Unauthorized("you dont have access")
	}

	func main() {
		if err := Authorize(); !fail.IsUnknown(err) {
			// we know this error is a fail.
			log.Println("cmd:", fail.Because(err))
		}
	}

*/
func Because(err error) error {
	if e, ok := err.(*Fail); ok {
		f := &Fail{
			Status:  e.Status,
			Message: e.Message,
			Details: e.Details,
			prev:    e.prev,
		}
		f.Caller(1)
		return f
	}
	return Cause(err).Unexpected()
}

// Caller finds the file and line where the failure happened.
// 'skip' is the number of calls to skip, not including this call.
// If you use this from a point(s) which is not the error location, then that
// call must be skipped.
func (f *Fail) Caller(skip int) {
	// check if this cause is internal and show the proper file:line.
	if f.prev == ErrUnspecified {
		skip++
	}
	_, file, line, _ := runtime.Caller(skip + 1)
	f.file = file[strings.LastIndex(file, "/")+1:]
	f.line = line
}

// BadRequest changes the error to a "Bad Request" fail.
// 'm' is the reason why this is a bad request.
// 'details' is an optional slice of details to explain the fail.
func (f *Fail) BadRequest(m string, details ...string) error {
	f.Status = http.StatusBadRequest
	f.Message = m
	f.Details = details
	return f
}

// BadRequest is a convenience function to return a BadRequest fail when there's
// no error.
func BadRequest(m string, fields ...string) error {
	return Cause(ErrUnspecified).BadRequest(m, fields...)
}

// Conflict changes the error to a "Conflict" fail.
// 'm' is the reason why this is a conflict.
// 'details' is an optional slice of details to explain the fail.
func (f *Fail) Conflict(m string, details ...string) error {
	f.Status = http.StatusConflict
	f.Message = m
	f.Details = details
	return f
}

// Conflict is a convenience function to return a Conflict fail when there's
// no error.
func Conflict(m string, fields ...string) error {
	return Cause(ErrUnspecified).Conflict(m, fields...)
}

// Forbidden changes an error to a "Forbidden" fail.
// 'm' is the reason why this action is forbidden.
func (f *Fail) Forbidden(m string) error {
	f.Status = http.StatusForbidden
	f.Message = m
	return f
}

// Forbidden is a convenience function to return a Forbidden fail when there's
// no error.
func Forbidden(m string) error {
	return Cause(ErrUnspecified).Forbidden(m)
}

// NotFound changes the error to an "Not Found" fail.
func (f *Fail) NotFound(m ...string) error {
	f.Status = http.StatusNotFound
	if m != nil {
		f.Message = m[0]
	} else {
		f.Message = messageNotFound
	}
	return f
}

// NotFound is a convenience function to return a Not Found fail when there's
// no error.
func NotFound(m ...string) error {
	return Cause(ErrUnspecified).NotFound(m...)
}

// Unauthorized changes the error to an "Unauthorized" fail.
func (f *Fail) Unauthorized(m string) error {
	f.Status = http.StatusUnauthorized
	f.Message = m
	return f
}

// Unauthorized is a convenience function to return an Unauthorized fail when there's
// no Go error.
func Unauthorized(m string) error {
	return Cause(ErrUnspecified).Unauthorized(m)
}

// Unexpected morphs the error into an "Internal Server Error" fail.
func (f *Fail) Unexpected() error {
	f.Status = http.StatusInternalServerError
	f.Message = messageUnexpected
	return f
}

// Unexpected is a convenience function to return an Internal Server Error fail
// when there's no error.
func Unexpected() error {
	return Cause(ErrUnspecified).Unexpected()
}

// Say returns the HTTP status and message response of a fail.
// If the error is nil, then there's no error -- say everything is OK.
// If the error is not a handled fail, then convert it to an unexpected fail.
func Say(err error) (int, string) {
	switch e := err.(type) {
	case nil:
		return http.StatusOK, messageOK
	case *Fail:
		return e.Status, e.Message
	}
	return http.StatusInternalServerError, messageUnexpected
}

// Error complements `http.Error` by sending a fail response.
func Error(w http.ResponseWriter, err error) {
	status, m := Say(err)
	http.Error(w, m, status)
}

// IsBadRequest returns true if fail is a BadRequest fail, false otherwise.
func IsBadRequest(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusBadRequest
}

// IsConflict returns true if fail is a Conflict fail, false otherwise.
func IsConflict(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusConflict
}

// IsUnauthorized returns true if fail is a Unauthorized fail, false otherwise.
func IsUnauthorized(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusUnauthorized
}

// IsForbidden returns true if fail is a Forbidden fail, false otherwise.
func IsForbidden(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusForbidden
}

// IsNotFound returns true if fail is a NotFound fail, false otherwise.
func IsNotFound(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusNotFound
}

// IsUnexpected returns true if fail is an Unexpected fail, false otherwise.
func IsUnexpected(err error) bool {
	e, ok := err.(*Fail)
	return ok && e.Status == http.StatusInternalServerError
}

// IsUnknown returns true if err is not handled as a fail, false otheriwse.
func IsUnknown(err error) bool {
	_, ok := err.(*Fail)
	return !ok
}
