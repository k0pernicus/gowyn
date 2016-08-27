package giwyn

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

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

func addGiwynConfigurationFile(pathname string) {

	file, err := os.OpenFile(filepath.Join(pathname, GIWYN_NAME_FILE), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		ErrorTracer.Println(err)
		return
	}

	if _, err := io.WriteString(file, UPDATED_S+time.Now().String()); err != nil {
		ErrorTracer.Println(err)
		return
	}

	file.Close()
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
			InfoTracer.Printf("* Found %s\n", pathname)
			*listOfGitPaths = append(*listOfGitPaths, pathname)
			addGiwynConfigurationFile(filepath.Dir(pathname))
		}

		return nil
	}

}
