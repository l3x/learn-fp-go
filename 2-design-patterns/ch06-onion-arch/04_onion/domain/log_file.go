package domain

import (
	"os"
	"io"
	"encoding/json"
	"github.com/pkg/errors"
)

type User struct {
	UserId int `json:"userId"`
	Country string `json:"country"`
	DeviceType string `json:"deviceType"`
	IP string `json:"ip"`
	SrcPort int `json:"srcPort"`
}

type LogFile struct {
	EventId     int `json:"eventId"`
	Timestamp   int64 `json:"timestamp"`
	Description string `json:"description"`
	User
}

func NewLogFile(logfileJson string) (logFile *LogFile, err error) {
	err = json.Unmarshal([]byte(logfileJson), &logFile)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal json")
	}
	return
}

func (lf *LogFile) ToJson() (logFileJson string, err error) {
	logFileBytes, err := json.Marshal(lf)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal json")
	}
	logFileJson = string(logFileBytes)
	return
}

func (lf *LogFile) Write(logFilename, contents string) (err error) {
	overwrite := true
	flag := os.O_WRONLY | os.O_CREATE
	if overwrite {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_EXCL
	}
	osFile, err := os.OpenFile(logFilename, flag, 0666)
	if err != nil {
		return errors.Wrapf(err, "unable to open %s", logFilename)
	}
	bytes := []byte(contents)
	n, err := osFile.Write(bytes)
	if err == nil && n < len(bytes) {
		err = io.ErrShortWrite
		return errors.Wrapf(io.ErrShortWrite, "not all bytes written for %s", logFilename)
	}
	if err1 := osFile.Close(); err1 != nil {
		return errors.Wrapf(err, "unable to close %s", logFilename)
	}
	return
}
