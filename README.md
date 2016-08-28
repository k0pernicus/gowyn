# GOwyn
Giwyn, in Go

## Procedure

1. Install `GOwyn` using `go install`
2. Go in the source dir of a git repository you want to follow
3. `gowyn add` will add the current git repository in the list of git repositories you want to follow
4. You can see if the path has been added using `cat ~/.gowyn_conf`
5. If you don't want to follow the repository, return to the source dir of the repository and type `gowyn rm`
6. List the `.gowyn_conf` and you will see that the repository doesn't more exists...
7. If you want to add all git repositories INSIDE a specific directory, go to the source dir and type `git add --crawl` 

### macOS requirements
In order to use [git2go](https://github.com/libgit2/git2go), please to install the following packages: `libgit2`, `cmake` and `pkg-config`.   
Ex: `brew install libgit2 cmake pkg-config`
