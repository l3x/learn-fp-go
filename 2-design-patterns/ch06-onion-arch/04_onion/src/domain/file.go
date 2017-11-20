package domain

import (
	"os"
	"path"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	. "utils"
	"bytes"
	"strings"
	"path/filepath"
	"fmt"
	"github.com/pkg/errors"
)

const Parsed = "parsed"
const NormalMode = 0666

type File struct {
	Id         int
	Name       string `json:"name"`
	ErrorMsg   string `json:"error"`
	Contents   LogFile `json:"logFile"`
	Bytes      []byte  `json:"bytes"`
}

type CloudFile struct {
	Name       string `json:"name"`
}
type CloudFiles struct {
	Names       []CloudFile
}

type CloudPath struct {
	Path       string `json:"path"`
}
type CloudPaths struct {
	Paths	[]CloudPath
}

func (f *File) ToJson() string {
	b, _ := json.MarshalIndent(f, "", "    ")
	return string(b)
}

// NewFile returns a value representing a file. In our case only contains the file name
func NewFile(fileName string) *File {
	fileName = path.Base(fileName)
	return &File{
		Name:       fileName,
	}
}

// NameOnly returns the file name only (no path and no extension)
func (f *File) NameOnly() string {
	fileName := path.Base(f.Name)
	extension := filepath.Ext(fileName)
	nameOnly := fileName[0:len(fileName)-len(extension)]
	return nameOnly
}

func (f *File) FullParsedFileName(i int) string {
	return path.Join(f.LocalParsedPath(), fmt.Sprintf("%d.json", i))
}

// Exists returns true if the file exists locally and is visible to the current user.
func (f *File) Exists() bool {
	_, err := os.Stat(f.FullLocalPath())
	return err == nil
}

// Path returns the file's path (no parent directory)
func (f *File) Path() string {
	return path.Join("")
}

// Path returns the file's full local path
func (f *File) LocalPath() string {
	return path.Join(Config.DownloadDir)
}

// Path returns the file's full host path
func (f *File) HostPath(cloudDir string) string {
	return path.Join(cloudDir)
}

// Path returns the file's full path with filename (less parent directory)
func (f *File) FullPath() string {
	return path.Join(f.Path(), f.Name)
}

// Path returns the file's parsed path (less parent directory)
func (f *File) AllParsedPath() string {
	return path.Join(f.LocalParsedPath(), Parsed)
}

// Path returns the file's full path with filename (less parent directory)
func (f *File) AllFullParsedPath() string {
	return path.Join(f.AllParsedPath(), f.Name)
}

// Path returns the file's full path (less parent directory)
func (f *File) LocalParsedPath() string {
	return path.Join(Config.DownloadDir, f.Path(), f.NameOnly())
}

// Path returns the file's full absolute path with filename.
func (f *File) FullLocalPath() string {
	return path.Join(f.LocalPath(), f.Name)
}

// Path returns the file's full absolute path with filename.
func (f *File) FullHostPath(cloudDir string) string {
	return path.Join(f.HostPath(cloudDir), f.Name)
}

// ContentsJson returns the file's Contents in json
func (f *File) ContentsJson() string {
	b, _ := json.MarshalIndent(f.Contents, "", "    ")
	return string(b)
}

// Write writes out the file's Contents to disk.
func (f *File) Write(bytes []byte) (err error) {
	Debug.Println("Creating file: "+f.FullLocalPath())
	osFile, err := os.Create(f.FullLocalPath())
	if err != nil {
		return errors.Wrapf(err, "unable to open %s", f.FullLocalPath())
	}
	defer osFile.Close()
	_, err = osFile.Write(bytes)
	if err != nil {
		return errors.Wrapf(err, "unable to write to file %s", f.FullLocalPath())
	}
	return
}


// Read the downloaded file into a byte array
func (f *File) Read() (bytes []byte, err error) {
	file, err := os.Open(f.FullLocalPath())
	defer file.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open %s", f.FullLocalPath())
	}
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s", f.FullLocalPath())
	}
	return
}

// Read grabs the parsed data-packet json objects associated with the file
func (f *File) ReadLogFiles() (logFiles []LogFile, multiStatus MultiStatus, err error) {
	// Get parsed (data-packet) files
	var msg string
	Debug.Println("f.LocalParsedPath(): "+f.LocalParsedPath())
	cmd := exec.Command("find", f.LocalParsedPath(), "-type", "f", "-name", "*.json")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err = cmd.Run()
	if err != nil {
		msg = fmt.Sprintf("error running find %s -type f -name *.json", f.LocalParsedPath())
		Error.Printf(msg)
		return nil, MultiStatus{}, errors.Wrap(err, msg)
	}
	findParsedFilesResponse := string(cmdOutput.Bytes())
	if len(findParsedFilesResponse) == 0 {
		msg = fmt.Sprintf("no results from running find %s -type f -name *.json", f.LocalParsedPath())
		Error.Printf(msg)
		return nil, MultiStatus{}, errors.Wrap(err, msg)
	}
	parsedFilePaths := strings.Split(findParsedFilesResponse, "\n")
	parsedFilePaths = parsedFilePaths[:len(parsedFilePaths)-1] // find adds newline that adds empty item at end
	Debug.Printf("parsedFilePaths: %v", parsedFilePaths)
	// Populate logFiles
	multiStatus = MultiStatus{}
	outcomeAndMsgs := []OutcomeAndMsg{}
	var logFileSlice []LogFile
	for i, parsedFilePath := range parsedFilePaths {
		Debug.Printf("%d - encoding %s", i, parsedFilePath)
		file, err := os.Open(parsedFilePath)
		if err != nil {
			msg = "unable to encode parsedFilePath: "+parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success:false, Message:msg})
			break
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			msg = "ioutil.ReadAll failed for "+parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success:false, Message:msg})
			break
		}
		logFile, err := NewLogFile(string(bytes))
		if err != nil {
			msg = "failed to parse "+parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success:false, Message:msg})
			break
		}
		logFileSlice = append(logFileSlice, *logFile)
	}
	logFiles = logFileSlice
	Debug.Printf("logFiles: %+v", logFiles)
	multiStatus.OutcomeAndMsgs = outcomeAndMsgs
	return
}

// FormatJson creates parsed directory and creates a proper .json file from the contents of the .jsonl file
func (f *File) FormatJson() (newContents string, err error) {
	fileName := f.Name
	parsedFullPath := f.AllFullParsedPath() // AllFullParsedDirectory previously created
	read, err := ioutil.ReadFile(f.FullLocalPath())
	if err != nil {
		return "", errors.Wrapf(err, "File.FormatJson: Unable to read (%s)", fileName)
	}
	newContents = "["+strings.Replace(string(read), "}{\"brandId", "},{\"brandId", -1)+"]"
	Debug.Printf("newContents: %v", newContents)
	err = ioutil.WriteFile(parsedFullPath, []byte(newContents), NormalMode)
	if err != nil {
		return "", errors.Wrapf(err, "File.FormatJson: Unable to write newContents (%s)", fileName)
	}
	return
}

// Parse parses the file Contents of JSONL (w/o new lines) into JSON
func (f *File) Parse(fileBytes []byte) (logFiles *[]LogFile,  err error) {
	logFilesJson, err := f.FormatJson()
	if err != nil {
		return nil, errors.Wrap(err, "unable to FormatJson")
	}
	logFiles = &[]LogFile{}
	err = json.Unmarshal([]byte(logFilesJson), logFiles)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshall results")
	}
	return
}
