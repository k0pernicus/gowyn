package main

import (
	"errors"
	"fmt"
	"github.com/k0pernicus/giwyn/giwyn"
	"os"
	"os/user"
)

func main() {

	/*
		Init traces (Error, Info, Warning)
	*/
	giwyn.InitTraces(os.Stderr, os.Stdout, os.Stdout)

	/*
		Get current user, in order to retrieve his home directory path
	*/
	userHomeDirectory, err := user.Current()
	if err != nil {
		fmt.Panicln("ERROR: Canno't retrieve the user home directory, due to: ", err)
	}

	/*
		Get each git path - currently (demo), we don't use them...
	*/
	_ := giwyn.FindGitObjects(userHomeDirectory.HomeDir)

}
