package gowyn

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func isGowynObjectFileExists(pathname string) bool {
	_, err := os.Stat(filepath.Join(pathname, GOWYN_NAME_FILE))
	/*
		Same for isGitRepositoryExists()
	*/
	return err == nil
}

/*
	addGowynObjectFile is a function that add comments and informations about a git repository, in a GIWYN_NAME_FILE file.
*/
func addGowynObjectFile(pathname string, crawlBehaviour bool) error {

	InfoTracer.Printf(" found \"%s\"\n", pathname)

	if crawlBehaviour && !askForConfirmation(fmt.Sprintf("Would you like to follow this repository \"%s\"?", pathname)) {
		return nil
	}

	file, err := os.OpenFile(filepath.Join(pathname, GOWYN_NAME_FILE), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := io.WriteString(file, UPDATED_S+" "+time.Now().String()+"\n"); err != nil {
		return err
	}

	addEntryInConfigFile(pathname)

	return nil

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
		rmEntryInConfigFile(pathname)
		return nil
	} else {
		return errors.New("No gowyn configuration file in the current directory.")
	}

}
