package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/k0pernicus/gowyn/gowyn"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("GOwin", "A command-line app to follow your git repositories")
	add   = app.Command("add", "Add the current directory to the list of git repositories")
	crawl = app.Command("crawl", "Crawl the entire hard drive to find and save git repositories")
	rm    = app.Command("rm", "Remove the current directory to the list followed git repositories")
)

func main() {

	/*
		Init traces (Error, Info, Warning)
	*/
	gowyn.InitTraces(os.Stderr, os.Stdout, os.Stdout)

	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("ERROR: Canno't retrieve the user current path, due to \"%s\"", err))
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case add.FullCommand():
		if err := gowyn.GetGitObject(pwd); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't get the git object from %s, due to \"%s\"", pwd, err))
		}

	case crawl.FullCommand():
		/*
			Get current user, in order to retrieve his home directory path
		*/
		userHomeDirectory, err := user.Current()
		if err != nil {
			panic(fmt.Sprintf("ERROR: Canno't retrieve the user home directory, due to \"%s\"", err))
		}
		/*
			Get each git path - currently (demo), we don't use them...
		*/
		gowyn.FindGitObjects(userHomeDirectory.HomeDir)

	case rm.FullCommand():
		if err := gowyn.RmGitObject(pwd); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't remove the git object from %s, due to \"%s\"", pwd, err))
		}
	}

}
