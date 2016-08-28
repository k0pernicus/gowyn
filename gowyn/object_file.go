package gowyn

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

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

	return nil

}
