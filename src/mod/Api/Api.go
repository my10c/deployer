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
package Api

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"deployer.badassops.com/Config"
	"deployer.badassops.com/Logs"
	"deployer.badassops.com/Msg"
)

func getIp(r *http.Request) string {
	a := r.RemoteAddr
	host, _, err := net.SplitHostPort(a)
	if err == nil {
		return host
	}
	return a
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Make sure the method is GET.
	if r.Method != "GET" {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Always authenticate, i.e. the request must contain an Auth header
	// containing the token as set in the configuration.
	auth := r.Header.Get("Auth")
	if auth != Config.Api.Auth {
		err := Msg.EUnauthorized
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Check whether the remote host is allowed access.
	if len(Config.Api.Acl) > 0 {
		ip := net.ParseIP(getIp(r))
		ok := false
		for _, v := range Config.Api.Acl {
			_, ipv4Net, err := net.ParseCIDR(v)
			if err != nil {
				continue
			}
			if ipv4Net.Contains(ip) {
				ok = true
				break
			}
		}
		if !ok {
			err := Msg.EUnauthorized
			if Logs.Response != nil {
				Logs.Response(r, err)
			}
			w.WriteHeader(Msg.GetStatus(err))
			return
		}
	}

	// The path must start with the prefix as set in the configuration.
	path := r.URL.Path
	if !strings.HasPrefix(path, Config.Api.Prefix) {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// There must be at least one argument to pass to the script.
	if len(r.URL.Query()) < 1 {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Since we only need the name, then strip it to get the command (script).
	path = strings.TrimPrefix(path, Config.Api.Prefix)
	cmd, exists := Config.Api.Cmds[path]
	if !exists {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// WARNING We pass the arguments as is.
	// It is up to the script to do whatever it has to do with it.
	exe := exec.Command(cmd, r.URL.RawQuery)

	// If we're debugging, it's nice to print out the result of the script.
	// That is, if it does print out something.
	if Logs.Debug != nil {
		exe.Stdout = os.Stdout
	}

	err := exe.Run()
	if err != nil {
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}
	if Logs.Response != nil {
		Logs.Response(r, nil, fmt.Sprintf("%s %s", cmd, r.URL.RawQuery))
	}
}

func Init() {
	Logs.AddPrefix("Api.")

	// This handles every URL.
	http.HandleFunc("/", handle)
}
