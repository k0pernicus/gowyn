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
	/*
		app
		===
		Field that contains commands and fields about GOwyn
	*/
	app = kingpin.New("GOwin", "A command-line app to follow your git repositories")
	/*
		add
		===
		Command that add a git object in the list of git objects, which is the list of objects to be inform about modifications
		Subcommands
		-----------
		*	crawl: boolean to enable the behaviour to add all git objects contained in the current repository
		*	group: string that contains the group to add the git object contained in the current repository
		*	path:	string that contains the path of the git object to add in the list of git objects, and not in the current repository
	*/
	add       = app.Command("add", "Add the current directory to the list of git repositories")
	add_crawl = add.Flag("crawl", "Crawl to add git repositories found since the current directory").Bool()
	add_group = add.Flag("group", "Add the/all git repository/ies in a group").String()
	add_path  = add.Flag("path", "Give as parameter the path of the git repository").String()
	/*
		debug
		=====
		Simple flag to display logs in order to debug the program
	*/
	debug = app.Flag("debug", "Run the app with debug traces").Bool()
	/*
		group
		=====
		Manage groups for the git repository contained in the current path
		Subcommands
		-----------
		* add: Add a new group to the git object
		* ignore: Ignore the given group, AND ALL associated git objects
		*	rm: Remove/delete the given group, AND ALL associated git objects
	*/
	group        = app.Command("group", "Add group to a followed git repository")
	group_add    = group.Flag("add", "Add group to a followed git repository").String()
	group_ignore = group.Flag("ignore", "Ignore an existing group and associated repositories - gowyn configuration files WILL NOT BE erased!").String()
	group_rm     = group.Flag("rm", "Remove an existing group and associated repositories - gowyn configuration files WILL BE erased!").String()
	/*
		ignore
		======
		Ignore the current path - the current path will be unlisted, but the gowyn configuration file will not be erased (which is different with the `rm` command)
	*/
	ignore = app.Command("ignore", "Ignore the current path")
	/*
		list
		====
		List appreciated/saved git objects
	*/
	list = app.Command("list", "List followed git repositories")
	/*
		retrieve
		========
		Retrieve git objects in your hard drive - really interesting if you have some repositories with created .gowyn files inside
		Subcommands
		-----------
		* group: Retrieve git objects IF it belong to the given group
	*/
	retrieve       = app.Command("retrieve", "Retrieve a gowyn configuration, based on existing gowyn objects")
	retrieve_group = retrieve.Flag("group", "Retrieve only gowyn objects that corresponding to the specified group").String()
	/*
		rm
		==
		Remove the git object from the list of appreciated/saved git objects
	*/
	rm = app.Command("rm", "Remove the current directory to the list followed git repositories")
	/*
		status
		======
		Get the status of each appreciated/saved git objects
		Subcommands
		-----------
		* group: Get the status of a given group ONLY
	*/
	status       = app.Command("status", "Get the status of each listed git repositories")
	status_group = status.Flag("group", "Get the status of the given group ONLY").String()
	/*
		update
		======
		Update the global configuration file (~/.gowyn_conf) - usefull to delete useless/removed git objects
		Subcommands
		-----------
		* group: Update only the given group
	*/
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
			gowyn.FindGitObjects(pwd, add_group, false)
		} else {
			if err := gowyn.GetGitObject(pwd, add_group, false); err != nil {
				panic(fmt.Sprintf("ERROR: Canno't get the git object from %s, due to \"%s\"", pwd, err))
			}
		}

	case group.FullCommand():
		if *group_add != "" {
			gowyn.AddGroupForAnExistingPath(pwd, *group_add)
		} else if *group_rm != "" {
			gowyn.RmGroupInConfigFile(*group_rm, true)
		} else if *group_ignore != "" {
			gowyn.RmGroupInConfigFile(*group_ignore, false)
		}

	case ignore.FullCommand():
		if err := gowyn.RmGitObject(pwd, false); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't ignore the git object from %s, due to \"%s\"", pwd, err))
		}

	case list.FullCommand():
		gowyn.ListGitObjects()

	case retrieve.FullCommand():
		if *retrieve_group != "" {
			gowyn.FindGitObjects(pwd, retrieve_group, true)
		} else {
			gowyn.FindGitObjects(pwd, nil, true)
		}

	case rm.FullCommand():
		if err := gowyn.RmGitObject(pwd, true); err != nil {
			panic(fmt.Sprintf("ERROR: Canno't remove the git object from %s, due to \"%s\"", pwd, err))
		}

	case status.FullCommand():
		gowyn.CheckStateOfGitObjects(status_group)

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
