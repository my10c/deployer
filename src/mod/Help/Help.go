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
package Help

import (
	"fmt"
	"os"

	"deployer.badassops.com/Variables"
)

func HelpSetup() {
	fmt.Printf("%s", Variables.MyInfo)
	fmt.Print(`
Configuration file must be valid json files with this structure:
{
	"server": {
		"port": 9091,        The port to listen to.
		"ip": "192.168.6.6"  IP to bind to.
	},
	"logs": {
		"logfile": "/var/log/deployer.log",  Absolute path of the log file.
		"maxsize": 512,                      Maximum size of the log file before rotated (MB).
		"maxage": 30,                        Maximum file age before rotated (days).
		"maxbackups": 28                     Maximum number of backups.
		"debug": false,                      Whether to log debug messages.
		"info": true,                        Whether to log informational messages.
		"response-ok": true,                 Whether to log successful requests.
		"response-error": true               Whether to log error requests.
	},
	"api": {
		"prefix": "/api/v1/",                        URI prefix for the APIs.
		"auth": "zYAwesomeAutTokeSK"                 Authorization token.
		"acl": ["127.0.0.0/16", "172.16.240.0/24"],  IPs allowed to connect in CIDR notation.
		"cmds": {                                    List of API name and script to execute.
			"blue": "/usr/local/lib/deployer/blue.sh",
			"artifact": "/usr/local/lib/deployer/artifact.sh",
			"green": "/usr/local/lib/deployer/green.sh"
		}
	}
}

Example on how to access the server:

	curl --header \"Auth: zYAwesomeAutTokeSK\" http://127.0.0.1:9091/api/v1/blue?foo:bar

`)
	os.Exit(0)
}
