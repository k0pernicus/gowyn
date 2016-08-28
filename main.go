package main

import (
	"fmt"
	"os"
	/*"os/user"*/

	"github.com/k0pernicus/gowyn/gowyn"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("GOwin", "A command-line app to follow your git repositories")
	/*ADD*/
	add   = app.Command("add", "Add the current directory to the list of git repositories")
	crawl = add.Flag("crawl", "Crawl to add git repositories found since the current directory").Bool()
	/*RM*/
	rm = app.Command("rm", "Remove the current directory to the list followed git repositories")
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

	/*
		userHomeDirectory, err := user.Current()
		if err != nil {
			panic(fmt.Sprintf("ERROR: Canno't retrieve the user home directory, due to \"%s\"", err))
		}
	*/

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case add.FullCommand():
		if *crawl {
			if err := gowyn.FindGitObjects(pwd); err != nil {
				panic(fmt.Sprintf("ERROR: Crawl to find git paths failed, due to \"%s\"", err))
			}
		} else {
			if err := gowyn.GetGitObject(pwd); err != nil {
				panic(fmt.Sprintf("ERROR: Canno't get the git object from %s, due to \"%s\"", pwd, err))
			}
		}

	case rm.FullCommand():
		if err := gowyn.RmGitObject(pwd); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't remove the git object from %s, due to \"%s\"", pwd, err))
		}
	}

}
