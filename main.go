package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/k0pernicus/gowyn/gowyn"
	"gopkg.in/alecthomas/kingpin.v2"
)

/*
	PROGRAM ARGUMENTS
	-----------------
	-> add: Add the current directory (where the user is) to the git repositories to watch!
	--> crawl: Crawl subdirectories from where the user is, in order to add these git repositories in the list to watch
	--> group: Tag the repository
	--> path: Add the repository, for a given path
	-----------------
	-> list: Get the list of followed git repositories
	-----------------
	-> rm: Remove the repository where the user is, from the list of git repositories to watch
	--> group: Remove the given group and associated repositories, from your git repositories
	-----------------
	-> status: Get the status of the watched git repositories
	-----------------
	-> update: Update the list of git repositories to watch, by verifying if saved git repositories are still available
	--> group: Update only the given group, if exists
*/
var (
	app          = kingpin.New("GOwin", "A command-line app to follow your git repositories")
	add          = app.Command("add", "Add the current directory to the list of git repositories")
	add_crawl    = add.Flag("crawl", "Crawl to add git repositories found since the current directory").Bool()
	add_group    = add.Flag("group", "Add the/all git repository/ies in a group").String()
	add_path     = add.Flag("path", "Give as parameter the path of the git repository").String()
	list         = app.Command("list", "List followed git repositories")
	rm           = app.Command("rm", "Remove the current directory to the list followed git repositories")
	remove_group = rm.Flag("group", "Remove an existing group and associated repositories, from the configuration file").String()
	status       = app.Command("status", "Get the status of each listed git repositories")
	update       = app.Command("update", "Update the configuration file removing useless links")
	update_group = update.Flag("group", "Update the configuration file only for the specified group").String()
)

func main() {

	/*
		Init traces (Debug, Error, Info, Warning)
		If the user does not want to get debug traces, initialize the field using nil (so, don't print anything)!
	*/
	if !*debug {
		gowyn.InitTraces(nil, os.Stderr, os.Stdout, os.Stdout)
	} else {
		gowyn.InitTraces(os.Stdout, os.Stderr, os.Stdout, os.Stdout)
	}

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
		if *add_path != "" {
			pwd = *add_path
		}
		if *add_crawl {
			gowyn.FindGitObjects(pwd, add_group)
		} else {
			if err := gowyn.GetGitObject(pwd, add_group); err != nil {
				panic(fmt.Sprintf("ERROR: Canno't get the git object from %s, due to \"%s\"", pwd, err))
			}
		}

	case list.FullCommand():
		gowyn.ListGitObjects()

	case rm.FullCommand():
		if *remove_group != "" {
			gowyn.RmGroupInConfigFile(*remove_group)
		} else {
			if err := gowyn.RmGitObject(pwd); err != nil {
				panic(fmt.Sprintf("ERROR: Canno't remove the git object from %s, due to \"%s\"", pwd, err))
			}
		}

	case status.FullCommand():
		gowyn.CheckStateOfGitObjects()

	case update.FullCommand():
		if *update_group != "" {
			gowyn.UpdateConfigFileByGroup(*update_group)
		} else {
			gowyn.UpdateConfigFile()
		}
	}

	if err := gowyn.SaveCurrentConfiguration(); err != nil {
		panic(fmt.Sprintf("ERROR: Canno't save current configuration in configuration file, due to \"%s\"", err))
	}
}
