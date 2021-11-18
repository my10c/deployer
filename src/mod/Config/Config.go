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
package Config

import (
	"encoding/json"
	"log"
	"os"
)

// ServerT is the server configuration.
type ServerT struct {
	Port int    `json:"port"` // server will listen to this port
	Ip   string `json:"ip"`   // server IP address
}

// LogsT is the Logs module configuration.
type LogsT struct {
	FilePath       string `json:"logfile"`        // Absolute path of the log file.
	FileMaxSize    int    `json:"maxsize"`        // Maximum size of the log file before rotated (MB).
	FileMaxAge     int    `json:"maxage"`         // Maximum file age before rotated (days).
	FileMaxBackups int    `json:"maxbackups"`     // Maximum number of backups.
	FileCompress   bool   `json:"compress"`       // Whether to compress rotated files.
	FileLocalTime  bool   `json:"localtime"`      // Whether to use server localtime (false means UTC).
	Debug          bool   `json:"debug"`          // Whether to log debug messages.
	Info           bool   `json:"info"`           // Whether to log informational messages.
	ResponseOK     bool   `json:"response-ok"`    // Whether to log request OK responses.
	ResponseError  bool   `json:"response-error"` // Whether to log request error responses.
}

// ApiT is the Api module configuration.
type ApiT struct {
	Prefix string            `json:"prefix"` // api URL prefix, e.g. "/api/v1/"
	Auth   string            `json:"auth"`   // authentication token, e.g. "QBiaxVGbYqofjAQVmk7qAmiI3JlrC1cOFWgpJVwj"
	Acl    []string          `json:"acl"`    // allowed remoted IP addresses
	Cmds   map[string]string `json:"cmds"`   // list of commands, e.g. "blue":"/usr/local/sbin/api/blue"
}

var (
	// Server configuration.
	Server ServerT

	// Logs configuration.
	Logs LogsT

	// Api configuration.
	Api ApiT
)

func Init(path string) {
	var conf struct {
		Server ServerT `json:"server"`
		Logs   LogsT   `json:"logs"`
		Api    ApiT    `json:"api"`
	}

	f, err := os.Open(path)
	if err != nil {
		log.Printf("Config.Init %s: %s", path, err.Error())
		panic(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("Config.Init %s: %s", path, err.Error())
		panic(err)
	}
	Server = conf.Server
	Logs = conf.Logs
	Api = conf.Api
}
