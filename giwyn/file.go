package giwyn

import (
	"fmt"
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

	if _, err := io.WriteString(file, UPDATED_S+" "+time.Now().String()+"\n"); err != nil {
		ErrorTracer.Println(err)
		return
	}

	/*
		Procedure to confirm if the user wants to follow the current path project or not
	*/
	confirmationToFollow := STATUS_IGNORING
	if askForConfirmation(fmt.Sprintf("Would you like to follow %s?", pathname)) {
		confirmationToFollow = STATUS_FOLLOWING
	}
	if _, err := io.WriteString(file, STATUS+" "+confirmationToFollow+"\n"); err != nil {
		ErrorTracer.Println(err)
		return
	}

	file.Close()
}
