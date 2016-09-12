package gowyn

import (
	"fmt"
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

func UpdateConfigFile() {

	var notFound = 0

	var indexToRemove []int

	var gitObjects []*gabs.Container

	gitObjects, err := globalContainer.S(FILENAME_PATH).Children()

	if err != nil {
		ErrorTracer.Printf("Canno't get the data structure %s from %s!\n", FILENAME_PATH, GOWYN_NAME_CONF)
		return
	}

	for index, gitObject := range gitObjects {
		pathValue, ok := gitObject.Data().(string)
		if !(ok) {
			ErrorTracer.Printf("The data structure %s in %s contains a non-string value!\n", FILENAME_PATH, GOWYN_NAME_CONF)
			continue
		}
		var currentPath = pathValue + "/"
		if _, err := os.Stat(currentPath); err != nil && os.IsNotExist(err) {
			notFound += 1
			indexToRemove = append(indexToRemove, index)
			WarningTracer.Printf("%s not found...\n", currentPath)
			rmEntryInConfigFile(currentPath)
		}
	}

	if notFound != 0 && askForConfirmation("We found some (re)moved repositories, would you like to delete those?") {
		for _, indexOfFilepath := range indexToRemove {
			_ = globalContainer.ArrayRemoveP(indexOfFilepath, FILENAME_PATH)
		}
	}

}

func UpdateConfigFileByGroup(group string) {

	var notFound = 0

	var indexToRemove []int
	var pathsToRemove []string

	var gitObjects []*gabs.Container

	gitObjects, err := globalContainer.S(GROUPS_PATH, group).Children()

	if err != nil {
		ErrorTracer.Printf("Canno't get the group %s from %s, in %s!\n", group, GROUPS_PATH, GOWYN_NAME_CONF)
		return
	}

	for index, gitObject := range gitObjects {
		pathValue, ok := gitObject.Data().(string)
		if !(ok) {
			ErrorTracer.Printf("The group %s in %s contains a non-string value!\n", group, GOWYN_NAME_CONF)
			continue
		}
		var currentPath = pathValue
		if _, err := os.Stat(currentPath); err != nil && os.IsNotExist(err) {
			notFound += 1
			pathsToRemove = append(pathsToRemove, currentPath)
			indexToRemove = append(indexToRemove, index)
			WarningTracer.Printf("%s not found...\n", currentPath)
			rmEntryInConfigFile(currentPath)
		}
	}

	if notFound != 0 && askForConfirmation("We found some (re)moved repositories, would you like to delete those?") {
		for _, indexOfFilepath := range indexToRemove {
			_ = globalContainer.ArrayRemove(indexOfFilepath, GROUPS_PATH, group)
			_ = globalContainer.ArrayRemove(indexOf(pathsToRemove[indexOfFilepath]), FILENAME_PATH)
		}
	}

}

/*
	parseConfigurationFile is a function that allows to parse a gowyn object file, which is a simple JSON file
*/
func parseGowynConfigurationFile() error {

	container, err := gabs.ParseJSONFile(configPath)
	if err != nil {
		return err
	}
	if !container.ExistsP(FILENAME_PATH) {
		InfoTracer.Println("Does not exists a basic JSON structure to store filenames...")
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

func addGroupInConfigFile(filepath string, group string) {

	if !globalContainer.Exists(GROUPS_PATH, group) {
		if _, err := globalContainer.Array(GROUPS_PATH, group); err != nil {
			ErrorTracer.Printf("Canno't create the array %s, due to %s\n", GROUPS_PATH, err)
		}
	}

	if globalContainer.Exists(GROUPS_PATH, group, filepath) {
		WarningTracer.Printf("The entry %s exists in the group %s\n", filepath, group)
		return
	}

	if err := globalContainer.ArrayAppend(filepath, GROUPS_PATH, group); err != nil {
		ErrorTracer.Printf("Canno't add the current entry (%s) in the group %s, due to \"%s\"\n", filepath, group, err)
	} else {
		InfoTracer.Printf("The current entry (%s) has been added in the group %s\n", filepath, group)
	}

}

func RmGroupInConfigFile(group string, hardRm bool) {

	if globalContainer.Exists(GROUPS_PATH, group) {
		/* Delete paths in the group */
		pathsToRemove, err := globalContainer.S(GROUPS_PATH, group).Children()
		if err != nil {
			ErrorTracer.Fatalf("Canno't get the group %s from %s to delete it!\n", group, GROUPS_PATH)
		}
		for _, path := range pathsToRemove {
			data, _ := path.Data().(string)
			/*
				if err := globalContainer.ArrayRemove(indexOf(data), FILENAME_PATH); err != nil {
					ErrorTracer.Printf("Canno't delete %s from %s, included in group %s!\n", data, FILENAME_PATH, group)
				}
			*/
			if hardRm {
				if err := RmGitObject(data, true); err != nil {
					ErrorTracer.Printf("Canno't delete the gowyn configuration file from %s!\n", data)
				}
			}
		}
		/* Delete the group */
		if err := globalContainer.Delete(GROUPS_PATH, group); err != nil {
			ErrorTracer.Printf("The group %s has not been deleted from %s, due to \"%s\"\n", group, GROUPS_PATH, err)
		}
	} else {
		WarningTracer.Printf("The group %s does not exists and canno't be deleted\n", group)
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

func ListGitObjects() {

	if gitObjects, err := globalContainer.S(FILENAME_PATH).Children(); err == nil {
		fmt.Printf("You saved %d repositories\n", len(gitObjects))
		for index, gitObject := range gitObjects {
			fmt.Printf("\t* %d: %s\n", index, gitObject)
		}
	}

	if groups, err := globalContainer.S(GROUPS_PATH).ChildrenMap(); err == nil {
		fmt.Printf("You saved %d groups\n", len(groups))
		for group, gitObjects := range groups {
			fmt.Printf("\t %s => %s\n", group, gitObjects.String())
		}
	}

}
