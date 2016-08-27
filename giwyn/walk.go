package giwyn

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func isGitRepository(dirName *os.FileInfo) bool {
	return ((*dirName).IsDir() && (*dirName).Name() == ".git")
}

func GetGitObject(pathname string) error {

	files, err := ioutil.ReadDir(pathname)

	if err != nil {
		return err
	}

	for _, file := range files {
		if isGitRepository(&file) {
			addGiwynConfigurationFile(pathname, false)
			break
		}
	}

	return nil

}

/*
	FindGitObjects is a function that find git object paths from a pathname given as parameter.
	This function returns a slice of strings, which each string corresponds to a git path.
*/
func FindGitObjects(pathname string) []string {

	var listOfGitPaths []string

	if err := filepath.Walk(pathname, findGitPaths(&listOfGitPaths)); err != nil {
		ErrorTracer.Panicln(err)
	}

	return listOfGitPaths

}

/*
	Function that walk from the pathname given as parameter.
	Each time that the function find a ".git" repository, the function add the pathfile to a data structure.
*/
func findGitPaths(listOfGitPaths *[]string) filepath.WalkFunc {

	return func(pathname string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == ".git") {
			*listOfGitPaths = append(*listOfGitPaths, pathname)
			addGiwynConfigurationFile(filepath.Dir(pathname), true)
		}

		return nil
	}

}
