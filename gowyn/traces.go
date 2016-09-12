package gowyn

import (
	"io"
	"log"
)

/*
	DebugTracer is a Logger object to output logs to debug the program
  ErrorTracer is a Logger object to output logs for runtime errors
  InfoTracer is a Logger object to output basic informations about the program
  WarningTracer is a Logger object to output runtime warnings - warnings are informative messages that are not considered as errors
*/
var (
	DebugTracer   *log.Logger
	ErrorTracer   *log.Logger
	InfoTracer    *log.Logger
	WarningTracer *log.Logger
)

/*
InitTraces is a function that initialize Loggers
*/
func InitTraces(debugHandle io.Writer, errorHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer) {

	/*
		Initialize the debug field
	*/
	DebugTracer = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

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
