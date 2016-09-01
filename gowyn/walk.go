package gowyn

import (
	"os"
	"path/filepath"
)

func isGitRepositoryExists(pathname string) bool {
	_, err := os.Stat(filepath.Join(pathname, GIT_NAME_DIR))
	/*
		If the directory exists, so err = nil, and err == nil!
	*/
	return err == nil
}

/*
	FindGitObjects is a function that find git object paths from a pathname given as parameter.
	This function returns a slice of strings, which each string corresponds to a git path.
*/
func FindGitObjects(pathname string, group *string) {
	findGitPaths(group)
}

/*
	Function that walk from the pathname given as parameter.
	Each time that the function find a ".git" repository, the function add the pathfile to a data structure.
*/
func findGitPaths(group *string) filepath.WalkFunc {

	return func(pathname string, info os.FileInfo, err error) error {

		if err != nil {
			ErrorTracer.Fatal(err)
		}

		if info.IsDir() && (info.Name() == ".git") {
			if err := addGowynObjectFile(filepath.Dir(pathname), *group, true); err != nil {
				ErrorTracer.Fatal(err)
			}
		}

		return nil

	}

}
