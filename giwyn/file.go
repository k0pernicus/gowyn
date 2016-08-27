package giwyn

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

/*
	addGiwynConfigurationFile is a function that add comments and informations about a git repository, in a GIWYN_NAME_FILE file.
*/
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
