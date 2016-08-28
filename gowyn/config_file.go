package gowyn

import (
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
)

var (
	globalContainer gabs.Container
	configPath      string
)

func InitGowynConfigurationFile(pathname string) error {
	configPath = pathname
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if _, err := os.Create(configPath); err != nil {
			return err
		} else {
			InfoTracer.Printf("Created file %s\n", pathname)
			ioutil.WriteFile(configPath, []byte("{}"), 0666)
		}
	}
	return parseGowynConfigurationFile()
}

func SaveCurrentConfiguration() error {
	return saveGowynConfigurationFile()
}

/*
	parseConfigurationFile is a function that allows to parse a gowyn object file, which is a simple JSON file
*/
func parseGowynConfigurationFile() error {

	container, err := gabs.ParseJSONFile(configPath)
	if err != nil {

	}
	if !container.ExistsP(FILENAME_PATH) {
		InfoTracer.Println("Does not exists a basic JSON structure inside the configuration file...")
		if _, err := container.Array(FILENAME_PATH); err != nil {
			return err
		}
	}
	globalContainer = *container
	return nil

}

/*
	saveGowynConfigurationFile is a function that allows to store the data structure (bytes) that represents git repositories, in the configuration file of Gowyn
*/
func saveGowynConfigurationFile() error {

	if err := ioutil.WriteFile(configPath, globalContainer.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func addEntryInConfigFile(filepath string) {

	if err := globalContainer.ArrayAppend(filepath, FILENAME_PATH); err != nil {
		ErrorTracer.Printf("Canno't add the current entry (%s) in the list of git repositories to follow, due to \"%s\"", filepath, err)
	} else {
		InfoTracer.Printf("The current entry (%s) has been added in the list of git repositories to follow", filepath)
	}

}

func indexOf(element string) int {

	children, _ := globalContainer.S(FILENAME_PATH).Children()

	for index, child := range children {
		if child.Data().(string) == element {
			return index
		}
	}

	return -1

}

func rmEntryInConfigFile(filepath string) {

	indexOfFilepath := indexOf(filepath)

	if indexOfFilepath == -1 {
		ErrorTracer.Printf("Element %s not found\n", filepath)
		return
	}

	if err := globalContainer.ArrayRemoveP(indexOfFilepath, FILENAME_PATH); err != nil {
		ErrorTracer.Printf("Canno't delete the current entry (%s) from the list of git repositories, due to error %s\n", filepath, err)
	} else {
		InfoTracer.Printf("The current entry (%s) has been deleted from the list of git repositories to follow\n", filepath)
	}

}
