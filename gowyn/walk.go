package gowyn

import (
	"errors"
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

func isGowynObjectFileExists(pathname string) bool {
	_, err := os.Stat(filepath.Join(pathname, GOWYN_NAME_FILE))
	/*
		Same for isGitRepositoryExists()
	*/
	return err == nil
}

func isGowynConfigurationFile(pathname string) bool {
	_, err := os.Stat(pathname)
	return err == nil
}

func GetGitObject(pathname string) error {

	if isGitRepositoryExists(pathname) {
		if !isGowynObjectFileExists(pathname) {
			addGowynObjectFile(pathname, false)
			return nil
		} else {
			return errors.New("Gowyn configuration file already exists.")
		}
	} else {
		return errors.New("The pathname does not point to a git repository.")
	}

}

func RmGitObject(pathname string) error {

	if isGowynObjectFileExists(pathname) {
		if err := os.Remove(filepath.Join(pathname, GOWYN_NAME_FILE)); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("No gowyn configuration file in the current directory.")
	}

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
			addGowynObjectFile(filepath.Dir(pathname), true)
		}

		return nil
	}

}
