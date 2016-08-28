package gowyn

import (
	"io"
	"log"
)

/*
  error is Logger object to output logs for error messages
  info is Logger object to output logs for information messages
  warning is Logger object to output warning messages
*/
var (
	ErrorTracer   *log.Logger
	InfoTracer    *log.Logger
	WarningTracer *log.Logger
)

/*
InitTraces is a function that initialize Loggers
*/
func InitTraces(errorHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer) {

	/*
	  Initialize the error field
	*/
	ErrorTracer = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	/*
	  Initialize the info field
	*/
	InfoTracer = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	/*
	  Initialize the warning field
	*/
	WarningTracer = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

}
