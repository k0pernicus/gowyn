# GOwyn
Giwyn, in Go

## Tutorial

1.	Install Gowyn...

```
git clone https://github.com/k0pernicus/GOwyn
cd GOwyn/
go build
go install
```

2.	Clone 2 "fake" repositories...

```
git clone https://github.com/k0pernicus/ocaml_exercism
git clone https://github.com/k0pernicus/fdcrawler
```

3.	Ok, now, let's go(wyn)!
		Go in `ocaml_exercism` and save it in group `Ocaml`: `gowyn add --group=Ocaml`

4.	Ok, now, you can list `ocaml_exercism` and you will see that the repository contains a new file: `.gowyn` that you can commit and push - this file contains the global configuration of this repository.  
		The local configuration file (`~/.gowyn_conf`) does not have to be commit and push ;-)  
		List saved git objects and you will see like:
```
gowyn list
You saved 1 repositories
	* 0: ".../ocaml_exercism"
You saved 1 groups
	 Ocaml => [".../ocaml_exercism"]
``` 

5.	You can now add Gowyn to the list of prefered repositories, in the gowyn repository: `gowyn add`

6. 	OH NO!  
		I forgot to tell you to add the group `Golang` for Gowyn :-/  
		Ok so, stay in the Gowyn repository and just add `gowyn group --add=Golang` and...

```
gowyn list
You saved 2 repositories
	* 0: ".../ocaml_exercism"
	* 1: ".../gowyn"
You saved 2 groups
	 Golang => [".../gowyn"]
	 Ocaml => [".../ocaml_exercism"]
```

7.	Finaly, you can add `fdcrawler` in the group Golang too using: `gowyn add --path=my/path/to/fdcrawler --group=Golang`

8.	Now, you can list the 3 objects using `gowyn status`, and you have 1 untrack file for each project, which is a gowyn statement to ignore it (in a .gitignore file).  
This is a feature that will be remove in the next version of Gowyn...

9.	Imagine that you don't want to work with Golang... You can just remove the entire group `Golang` and all associated repositories using `gowyn group --rm=Golang`, while keeping those git repositories in your hard drive.  
		Now, if you list your saved git objects:

```
> gowyn list
You saved 1 repositories
	* 0: ".../ocaml_exercism"
You saved 1 groups
	 Ocaml => [".../ocaml_exercism"]
```

10.	Have fun and PLEASE, hack Gowyn!!!!!

### macOS requirements
In order to use [git2go](https://github.com/libgit2/git2go), please to install the following packages: `libgit2`, `cmake` and `pkg-config`.   
Ex: `brew install libgit2 cmake pkg-config`

### TODO List
*	Add a structure as the git repository in the configuration file
*	Remove the repo in the group when removing a repo from the config file
*	Get the status of a group ONLY
* Ignored .gowyn or not?
*	Online tutorial
