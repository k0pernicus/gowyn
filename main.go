package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/k0pernicus/giwyn/giwyn"
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
		panic(fmt.Sprintf("ERROR: Canno't retrieve the user home directory, due to %s", err))
	}

	/*
		Get each git path - currently (demo), we don't use them...
	*/
	giwyn.FindGitObjects(userHomeDirectory.HomeDir)

}
