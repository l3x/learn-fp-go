package main

import (
	. "utils"
	. "interfaces"
	"os"
	"io/ioutil"
	"infrastructure"
	"github.com/pkg/errors"
	"net/http"
)

const defaultFileName = "eventset1.jsonl"
var	fileName	string
var wsh WebserviceHandler

func init() {
	GetOptions()
	if Config.LogDebugInfo {
		InitLog("trace-debug-log.txt", os.Stdout, os.Stdout, os.Stderr)
	} else {
		InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)
	}
	fileName = os.Getenv("TEST_FILENAME")
	if len(fileName) == 0 {
		fileName = defaultFileName
	}
	Debug.Printf("ProjectRoot: %s", PadRight(Config.ProjectRoot, " ", 20))
	Debug.Printf("AppEnv: %s", PadRight(Config.AppEnv, " ", 20))
	Debug.Printf("GcpSourceKeyFile: %s", PadRight(Config.GcpSourceKeyFile, " ", 20))
	Debug.Printf("GcpSinkKeyFile: %s", PadRight(Config.GcpSinkKeyFile, " ", 20))
	Debug.Printf("LogDebugInfo: %v", Config.LogDebugInfo)
	HandlePanic(os.Chdir(Config.ProjectRoot))
}

type endpoint struct {
	Api
	uriExample	 string
}

func printApiExample(url, uriExample string) {
	if len(uriExample) == 0 {
		Info.Printf("http://localhost:%s%s", Config.ApiPort, url)
	} else {
		Info.Printf("http://localhost:%s%s?%s", Config.ApiPort, url, uriExample)
	}
}

func main() {
	gcpi, err := infrastructure.GetGcpInteractor()
	HandlePanic(errors.Wrap(err, "unable to get gcp interactor"))
	li, err := infrastructure.GetLocalInteractor()
	HandlePanic(errors.Wrap(err, "unable to get local interactor"))

	wsh = WebserviceHandler{}
	wsh.GcpInteractor = gcpi
	wsh.LocalInteractor = li

	var endpoints = []endpoint{
		{Api{wsh.Health, "/health"}, ""},
		{Api{wsh.ListSourceBuckets, "/list-source-buckets"}, "projectId="+ Config.GcpSourceProjectId},
		{Api{wsh.ListSinkBuckets, "/list-sink-buckets"}, "projectId="+ Config.GcpSinkProjectId},
		{Api{wsh.SourceFileExists, "/source-file-exists"}, "fileName="+fileName},
		{Api{wsh.DownloadFile, "/download-file"}, "fileName="+fileName},
		{Api{wsh.UploadFile, "/upload-file"}, "fileName="+fileName},
		{Api{wsh.LocalFileExists, "/local-file-exists"}, "fileName="+fileName},
	}
	Info.Println("Example API endpoints:")
	{
		for _, ep := range endpoints {
			http.HandleFunc(ep.Api.Url, ep.Api.Handler)
			printApiExample(ep.Api.Url, ep.uriExample)
		}
	}
	http.ListenAndServe(":"+Config.ApiPort, nil)
}
