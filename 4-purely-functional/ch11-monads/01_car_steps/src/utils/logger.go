package utils

import (
	"io"
	log "github.com/sirupsen/logrus"
	"os"
	"fmt"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	InfoHandler io.Writer
	ErrorHandler io.Writer
)

const BasicTimeStampFormat = "2006-01-02 15:04:05"
var LevelDescriptions = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
var DebugChars = ">>"

func InitLog (
	traceFileName string,
	debugHandler io.Writer,
	infoHandler io.Writer,
	errorHandler io.Writer,
) {
	if len(traceFileName) > 0 {
		_ = os.Remove(traceFileName)
		file, err := os.OpenFile(traceFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to create log file: %s", traceFileName)
		}
		debugHandler = io.MultiWriter(file, debugHandler)
		infoHandler = io.MultiWriter(file, infoHandler)
		errorHandler = io.MultiWriter(file, errorHandler)
	}

	InfoHandler = infoHandler
	ErrorHandler = errorHandler

	plainFormatter := new(PlainFormatter)

	basicFormatter := new(BasicFormatter)
	basicFormatter.TimestampFormat = BasicTimeStampFormat
	basicFormatter.LevelDesc = LevelDescriptions

	basicDebugFormatter := new(BasicDebugFormatter)

	plusVFormatter := new(PlusVFormatter)
	plusVFormatter.TimestampFormat = BasicTimeStampFormat
	plusVFormatter.LevelDesc = LevelDescriptions
	plusVFormatter.FullTimestamp = true

	logLevel, err := log.ParseLevel(Config.LogLevel)
	if err != nil {
		println("ERROR: " + err.Error())
		logLevel = log.InfoLevel
	}

	Debug = log.New()
	Debug.Out = debugHandler
	//Debug.Formatter = new(log.TextFormatter) //new(log.JSONFormatter)
	if Config.LogVerbose {
		Debug.Formatter = basicFormatter
	} else {
		Debug.Formatter = basicDebugFormatter
	}
	//Debug.Formatter = plainFormatter
	Debug.Hooks= make(log.LevelHooks)
	Debug.Level = logLevel


	Info = log.New()
	Info.Out = infoHandler
	//Info.Formatter = customFormatter
	if Config.LogVerbose {
		Info.Formatter = basicFormatter
	} else {
		Info.Formatter = plainFormatter
	}
	Info.Hooks= make(log.LevelHooks)
	Info.Level = logLevel

	Error = log.New()
	Error.Out = errorHandler
	//Error.Formatter = plusVFormatter
	if Config.LogVerbose {
		Error.Formatter = basicFormatter
	} else {
		Error.Formatter = plainFormatter
	}
	Error.Hooks= make(log.LevelHooks)
	Error.Level = logLevel

	DebugChars = Config.LogDebugChars
}


type PlainFormatter struct {}
func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

type BasicDebugFormatter struct {
	TimestampFormat string
	LevelDesc []string
}
func (f *BasicDebugFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s\n", DebugChars, entry.Message)), nil
}

type BasicFormatter struct {
	TimestampFormat string
	LevelDesc []string
}
func (f *BasicFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], timestamp, entry.Message)), nil
}

type PlusVFormatter struct {
	TimestampFormat string
	LevelDesc []string
	FullTimestamp bool
}
func (f *PlusVFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	//TODO: Find bug in logrus that prevents entry.Level from returning correct value
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[Error.Level], timestamp, entry.Message)), nil
}