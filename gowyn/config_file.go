package gowyn

import (
	"errors"
	"os"

	"github.com/Jeffail/gabs"
)

var globalContainer gabs.Container

/*
	parseConfigurationFile is a function that allows to parse a gowyn object file, which is a simple JSON file
*/
func parseGowynConfigurationFile(pathname string) error {

	if isGowynConfigurationFile(pathname) {

		if container, err := gabs.ParseJSONFile(pathname); err != nil {
			return err
		} else {
			globalContainer = *container
		}
	} else {
		return errors.New("Gowyn configuration file not found")
	}

	return nil

}

/*
	saveGowynConfigurationFile is a function that allows to store the data structure (bytes) that represents git repositories, in the configuration file of Gowyn
*/
func saveGowynConfigurationFile(pathname string) error {

	if isGowynConfigurationFile(pathname) {
		file, err := os.OpenFile(pathname, os.O_WRONLY, 0666)

		if err != nil {
			return err
		}

		defer file.Close()

		if _, err := file.Write(globalContainer.Bytes()); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("Gowyn configuration file not found.")
	}

}

func addEntryInConfigFile(filepath string) error {

	return nil

}
