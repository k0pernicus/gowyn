package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/k0pernicus/gowyn/gowyn"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("GOwin", "A command-line app to follow your git repositories")
	/*ADD*/
	add   = app.Command("add", "Add the current directory to the list of git repositories")
	crawl = add.Flag("crawl", "Crawl to add git repositories found since the current directory").Bool()
	/*CONFIG*/
	config   = app.Command("config", "Add informations about the Gowyn configuration")
	computer = config.Arg("computer", "Add a computer name to your configuration").String()
	/*LIST*/
	list = app.Command("list", "List followed git repositories")
	/*RM*/
	rm = app.Command("rm", "Remove the current directory to the list followed git repositories")
	/*STATUS*/
	status = app.Command("status", "Get the status of each listed git repositories")
	/*UPDATE*/
	update = app.Command("update", "Update the configuration file removing useless links")
)

func main() {

	/*
		Init traces (Error, Info, Warning)
	*/
	gowyn.InitTraces(os.Stderr, os.Stdout, os.Stdout)

	userHomeDirectory, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("ERROR: Canno't retrieve the user home directory, due to \"%s\"", err))
	}

	if err := gowyn.InitGowynConfigurationFile(filepath.Join(userHomeDirectory.HomeDir, gowyn.GOWYN_NAME_CONF)); err != nil {
		panic(fmt.Sprintf("ERROR: Canno't retrieve the configuration file, due to \"%s\"", err))
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("ERROR: Canno't retrieve the user current path, due to \"%s\"", err))
	}

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

	case config.FullCommand():
		/*TODO*/

	case list.FullCommand():
		gowyn.ListGitObjects()

	case rm.FullCommand():
		if err := gowyn.RmGitObject(pwd); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't remove the git object from %s, due to \"%s\"", pwd, err))
		}

	case status.FullCommand():
		gowyn.CheckStateOfGitObjects()

	case update.FullCommand():
		/*TODO*/
	}

	if err := gowyn.SaveCurrentConfiguration(); err != nil {
		panic(fmt.Sprintf("ERROR: Canno't save current configuration in configuration file, due to \"%s\"", err))
	}
}
