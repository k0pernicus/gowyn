# GOwyn
Giwyn, in Go

## Tutorial

* Please to install [Golang](http://golang.org)

*	Install Gowyn...
```
git clone https://github.com/k0pernicus/gowyn
cd gowyn/
go build
go install
```  

*	Clone 2 "fake" repositories...
```
git clone https://github.com/k0pernicus/ocaml_exercism
git clone https://github.com/k0pernicus/fdcrawler
```  

*	Ok, now, let's go(wyn)!
	Go in `ocaml_exercism` and save it in group `Ocaml`: `gowyn add --group=Ocaml`

*	Ok, now, you can list `ocaml_exercism` and you will see that the repository contains a new file: `.gowyn` that you can commit and push - this file contains the global configuration of this repository.  
	The local configuration file (`~/.gowyn_conf`) does not have to be commit and push ;-)  
	List saved git objects and you will see like:
```
gowyn list
You saved 1 repositories
	* 0: ".../ocaml_exercism"
You saved 1 groups
	 Ocaml => [".../ocaml_exercism"]
``` 

*	You can now add Gowyn to the list of prefered repositories, in the gowyn repository: `gowyn add`

*	OH NO!  
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

*	Finaly, you can add `fdcrawler` in the group Golang too using: `gowyn add --path=my/path/to/fdcrawler --group=Golang`

*	Now, you can list the 3 objects using `gowyn status`, or `gowyn status --full` if you want to get all informations for each saved git object.

*	Imagine that you don't want to work with Golang... You can just remove the entire group `Golang` and all associated repositories using `gowyn group --ignore=Golang`, while keeping those git repositories in your hard drive and gowyn object files.  
	If you want ALSO to delete gowyn object files for those repositories, use `--rm` instead of `--ignore`.  
	Now, if you list your saved git objects:

```
> gowyn list
You saved 1 repositories
	* 0: ".../ocaml_exercism"
You saved 1 groups
	 Ocaml => [".../ocaml_exercism"]
```

*	Have fun and PLEASE, hack Gowyn!!!!!

## FAQ

*	What is the difference between `ignore` and `rm`?  
	The main difference between those is that `rm` will delete your local object file (`.gowyn`), not `ignore`. This difference is really important if you want to push your local object file in the git server. 

*	There is no much informations about my git object files...  
	I know :-/  
	Gowyn has to be improve **a lot**, and do not hesitate to push some updates if you wanna contribute to this project :-)

### macOS requirements
In order to use [git2go](https://github.com/libgit2/git2go), please to install the following packages: `libgit2`, `cmake` and `pkg-config`.   
Ex: `brew install libgit2 cmake pkg-config`

### TODO List
*	Add `retrieve` command
*	Add a structure as the git repository in the configuration file
