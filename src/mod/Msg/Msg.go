/*
 BSD 3-Clause License

 Copyright (c) 2017 - 2021, © Badassops LLC
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

 * Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

 * Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package Msg

// Msg makes it possible to associate an `error` with a HTTP status code while
// allowing packages to create their own unique errors that can be used by
// other packages without referencing.
// Go error semantics are not changed.
// Example
//
//     package Foo
//     var (
//         // Notice eFoo is not exported
//         eFoo = Msg.New(http.StatusBadRequest, "Invalid foo value")
//     )
//     func DoSomething(int a) error {
//         if a < 10 {
//     	        return eFoo
//     	    }
//     	    return nil
//     }
//
//     package Bar
//     func HandleBar(w http.ResponseWrite, r *http.Request) {
//         err := Foo.DoSomething(20)
//         if err != nil {
//             Logs.Error(err)
//             //...
//             return
//         }
//         //...
//     }
// Logging (see the Logs module), basically:
//
//     To log a HTTP 200 response:
//         Logs.Response(request, nil)
//
//     To log a HTTP error response:
//         Logs.Response(request, err)
//
//         See also: Util.SendResponseError()
//
// Note that this package does not, and must not, depend on any other package
// except standard Go packages.
// This package was originally written by people @ badassops.com.

import (
	"errors"
)

// List of all codes.
var msgs = map[error]int{}

// New creates a new `error` and associate it with a HTTP status.
func New(status int, msg string) error {
	err := errors.New(msg)
	msgs[err] = status
	return err
}

// List of common errors.
// Remember, the message is sent to clients so you may want to be discreet.
var (
	// Client errors
	ERequest      = New(400, "Bad request")
	EPath         = New(400, "Invalid path")
	EPayload      = New(400, "Invalid payload")
	EQuery        = New(400, "Invalid query")
	EUnauthorized = New(401, "Unauthorized")
	EForbidden    = New(403, "Forbidden")
	ENotFound     = New(404, "Not found")
	EMethod       = New(405, "Invalid method")
	EConflict     = New(409, "Conflict")
	EMedia        = New(415, "Unsupported media type")
	ETeapot       = New(418, "I'm a teapot")
	ELegal        = New(451, "Unavailable for legal reasons")

	// Server errors
	EServer   = New(500, "Internal server error")
	EDatabase = New(500, "Internal server error")
)

// GetStatus returns the HTTP status associated with an `error`.
// If the `error` is not associated with a HTTP status, then returns 500.
func GetStatus(err error) int {
	status, exists := msgs[err]
	if exists {
		return status
	}
	return 500
}
