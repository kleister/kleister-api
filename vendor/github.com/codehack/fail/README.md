# Fail [![GoDoc](https://godoc.org/github.com/codehack/fail?status.svg)](https://godoc.org/github.com/codehack/fail) [![ghit.me](https://ghit.me/badge.svg?repo=codehack/fail)](https://ghit.me/repo/codehack/fail) [![Go Report Card](https://goreportcard.com/badge/github.com/codehack/fail)](https://goreportcard.com/report/github.com/codehack/fail)

*Manage and handle [Go](http://golang.org) errors with nice and correct HTTP responses.*

*Fail* allows to wrap errors and describe them as HTTP responses, such as, "Not Found" and "Bad Request" and their status code. Also, help you mask internal errors with friendly messages that are appropriate for clients.

The goal of this package is to handle the behavior of errors to minimize static errors and type assertions. If all errors are wrapped with ``fail.Cause`` we can inspect them better and give clients proper responses.

This package is inspired by Dave Cheney's excellent blog post "[Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)".

## Quick Start

Install using "go get":

	go get github.com/codehack/fail

Then import from your source:

	import "github.com/codehack/fail"

## Documentation

The documentation at GoDoc:

[http://godoc.org/github.com/codehack/fail](http://godoc.org/github.com/codehack/fail)

## Example

```go
package main

import "github.com/codehack/fail"

// Create a new user, using JSON values.
http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	// we want a POST request, this will fail.
	if r.Method != "POST" {
		fail.Error(w, fail.BadRequest("request not POST"))
		return
	}

	var payload struct {
		Name string        `json:"name"`
		Created *time.Time `json:"created"`
	}

	// JSON is hard, this will fail.
	if err := json.Decode(r.Request.Body, &payload); err != nil {
		var (
			m string
			status int
		)
		switch {
		case err == json.SyntaxError:
			fail.Error(w, fail.Cause(err).Unexpected())
		default:
			fail.Error(w, fail.Cause(err).BadRequest("your payload is terrible"))
		}
		return
	}

	// if we manage to get this far, this will fail.
	if err := saveUserDB(&payload); err != nil {
		fail.Error(w, fail.Cause(err).Unexpected())
		return
	}

	// Hah hah, this will fail.
	fail.Error(w, fail.Forbidden("resistance is futile."))
})
```

## Credits

**Fail** is Copyright (c) 2017 [Codehack](http://codehack.com).
Published under an [MIT License](https://raw.githubusercontent.com/codehack/fail/master/LICENSE)

