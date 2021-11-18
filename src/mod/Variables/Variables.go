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
package Variables

import (
	"os"
	"path"
	"strconv"
	"time"
)

var (
	// varaibles start with a kapital letter are global
	now = time.Now()

	MyVersion   = "0.3"
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Marc Krisnanto and Luc Suryo"
	myCopyright = "Copyright 2017 - " + strconv.Itoa(now.Year()) + " © Badassops LLC"
	myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
	myEmail     = "<marc@badassops.com> and <luc@badassops.com>"
	MyInfo      = MyProgname + " " + MyVersion + "\n" +
		myCopyright + "\nLicense: " + myLicense +
		"\nWritten by " + myAuthor + "\n" + 
        "Authors email: " + myEmail + "\n"
	MyTask = "A blue - green deployment API server"

	// defaults
	GivenValues map[string]string

	DefaultHome          = "/etc/deployer"
	DefaultConfigFile    = DefaultHome + "/config.json"
	DefaultLog           = "/var/log/deployer.log"
	DefaultLogMaxSize    = 512 // MB
	DefaultLogMaxBackups = 28  // count
	DefaultLogMaxAge     = 30  // days
	AsRoot bool
)

func init() {
	// setup the default value, these are hardcoded.
	GivenValues = make(map[string]string)
	GivenValues["configFile"] = DefaultConfigFile
	GivenValues["logfile"] = DefaultLog
	AsRoot = true
}
