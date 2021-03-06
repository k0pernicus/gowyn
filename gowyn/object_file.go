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

func AddGroupForAnExistingPath(pathname string, group string) {

	if isGowynObjectFileExists(pathname) {
		addGroupInConfigFile(pathname, group)
	} else {
		ErrorTracer.Printf("Cannot found an existing Gowyn repository for %s, please to run \"gowyn add --group %s\" instead\n", pathname, group)
	}

}

/*
	addGowynObjectFile is a function that add comments and informations about a git repository, in a GIWYN_NAME_FILE file.
*/
func addGowynObjectFile(pathname string, groupname string) error {

	InfoTracer.Printf(" found \"%s\"\n", pathname)
	if groupname != "" {
		InfoTracer.Printf("==> Belongs to group \"%s\"\n", groupname)
	}

	file, err := os.OpenFile(filepath.Join(pathname, GOWYN_NAME_FILE), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	if groupname != "" {
		if _, err := io.WriteString(file, GROUP_S+groupname+"\n"); err != nil {
			return err
		}
		addGroupInConfigFile(pathname, groupname)
	}

	if _, err := io.WriteString(file, UPDATED_S+time.Now().String()+"\n"); err != nil {
		return err
	}

	/*
			Add GOWYN_NAME_FILE to the .gitignore file

		if gitignoreFile, err := os.OpenFile(filepath.Join(pathname, GITIGNORE_NAME_FILE), os.O_RDWR|os.O_APPEND, 0644); err == nil {
			defer gitignoreFile.Close()
			gitignoreFile.WriteString(GOWYN_NAME_FILE)
		}
	*/

	addEntryInConfigFile(pathname)

	return nil

}

func GetGitObject(pathname string, group *string, crawlBehaviour bool) error {

	if isGitRepositoryExists(pathname) {
		if !isGowynObjectFileExists(pathname) {
			if crawlBehaviour && !askForConfirmation(fmt.Sprintf("Would you like to follow this repository \"%s\"?", pathname)) {
				return nil
			}
			addGowynObjectFile(pathname, *group)
			return nil
		} else {
			if indexOf(pathname) == -1 {
				if *group != "" {
					addGroupInConfigFile(pathname, *group)
				}
				addEntryInConfigFile(pathname)
			} else {
				return errors.New("Gowyn configuration file already exists.")
			}
		}
	} else {
		return errors.New("The pathname does not point to a git repository.")
	}

	return nil

}

func RmGitObject(pathname string, hardRm bool) error {

	if isGowynObjectFileExists(pathname) {
		rmEntryInConfigFile(pathname)
		if hardRm {
			if err := os.Remove(filepath.Join(pathname, GOWYN_NAME_FILE)); err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("No gowyn configuration file in the current directory.")
	}

}
