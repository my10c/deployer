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
package Initialize

import (
	"fmt"
	"os"
	"strconv"

	"deployer.badassops.com/Help"
	"deployer.badassops.com/Utils"
	"deployer.badassops.com/Variables"

	"github.com/akamensky/argparse"
)

// Function to process the given args
func InitArgs() {
	parser := argparse.NewParser(Variables.MyProgname, Variables.MyTask)

	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:     "Path to the configuration file",
			Default:  Variables.DefaultConfigFile,
		})

	asRoot := parser.Flag("R", "runAsRoot",
		&argparse.Options{
			Required: false,
			Help:     "Do run as root",
			Default:  false,
		})

	showSetup := parser.Flag("S", "showconfig",
		&argparse.Options{
			Required: false,
			Help:     "Show configuration example",
			Default:  false,
		})

	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     "Show version",
			Default:  false,
		})

	showInfo := parser.Flag("i", "info",
		&argparse.Options{
			Required: false,
			Help:     "Show information",
			Default:  false,
		})

	logFile := parser.String("l", "logFile",
		&argparse.Options{
			Required: false,
			Help:     "Path to the log file",
			Default:  Variables.DefaultLog,
		})

	logMaxSize := parser.String("M", "logMaxSize",
		&argparse.Options{
			Required: false,
			Help:     "Max size of the log file (MB). Default: " + strconv.Itoa(Variables.DefaultLogMaxSize),
		})

	logMaxBackups := parser.String("B", "logMaxBackups",
		&argparse.Options{
			Required: false,
			Help:     "Max log file count. Default: " + strconv.Itoa(Variables.DefaultLogMaxBackups),
		})

	logMaxAge := parser.String("A", "logMaxAge",
		&argparse.Options{
			Required: false,
			Help:     "Max days to keep a log file. Default: " + strconv.Itoa(Variables.DefaultLogMaxAge),
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *asRoot {
		Variables.AsRoot = true
	}

	if *showSetup {
		Help.HelpSetup()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s\n", Variables.MyVersion)
		os.Exit(0)
	}

	if *showInfo {
		fmt.Printf("%s\n", Variables.MyInfo)
		fmt.Printf("%s\n", Variables.MyTask)
		os.Exit(0)
	}

	if configFile != nil {
		Variables.GivenValues["configFile"] = fmt.Sprintf("%s", *configFile)
	}

	Variables.GivenValues["logfile"] = fmt.Sprintf("%s", *logFile)

	if *logMaxSize != "" {
		Variables.GivenValues["maxsize"] = fmt.Sprintf("%s", *logMaxSize)
	}

	if *logMaxBackups != "" {
		Variables.GivenValues["maxbackups"] = fmt.Sprintf("%s", *logMaxBackups)
	}

	if *logMaxAge != "" {
		Variables.GivenValues["maxage"] = fmt.Sprintf("%s", *logMaxAge)
	}

	if !Utils.CheckFileExist(Variables.GivenValues["configFile"]) {
		fmt.Printf("Configuration file %s, does not exist\n", Variables.GivenValues["configFile"])
		os.Exit(1)
	}

	if !Utils.CheckFileExist(Variables.GivenValues["logfile"]) {
		os.Exit(1)
	}
}
