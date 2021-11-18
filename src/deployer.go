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
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"deployer.badassops.com/Api"
	"deployer.badassops.com/Config"
	"deployer.badassops.com/Initialize"
	"deployer.badassops.com/Logs"
	"deployer.badassops.com/Utils"
	"deployer.badassops.com/Variables"
)

func main() {
	// get given argument overwrite the values in the configuration file
	Initialize.InitArgs()

	if Variables.AsRoot == true {
		// must run as root
		Utils.IsRoot()
	}

	// get value from the configuration file, default or overwritten
	Config.Init(Variables.GivenValues["configFile"])

	// initialize the logger system
	Logs.Init()

	// initialize the aoi system
	Api.Init()

	// install a signale handler so we capture issue if the application dies
	Utils.SignalHandler()

	// start the api server
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", Config.Server.Ip, Config.Server.Port), nil)
	if err != nil {
		e := fmt.Sprintf("http.ListenAndServer failed `%s`, panic follows...", err.Error())
		Logs.Error(errors.New(e))
		panic(err)
	}

	// should never be reached
	os.Exit(0)
}
