package interfaces_test

import (
	. "interfaces"
	. "utils"
	"infrastructure"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const failure = "\u2717"
const defaultFileName = "eventset1.jsonl"

var	fileName               string
var	wsh                    WebserviceHandler

func init() {
	GetOptions()
	if Config.LogDebugInfoForTests {
		InitLog("trace-debug-log.txt", os.Stdout, os.Stdout, os.Stderr)
	} else {
		InitLog("trace-debug-log.txt", ioutil.Discard, os.Stdout, os.Stderr)
	}
	HandlePanic(os.Chdir(Config.ProjectRoot))
	Debug.Printf("Config: %+v\n", Config)
	// use a filename in a downloads subdirectory
	fileName = os.Getenv("TEST_FILENAME")
	if len(fileName) == 0 {
		fileName = defaultFileName
	}
	// instantiate interactors
	gcpi, err := infrastructure.GetGcpInteractor()
	HandlePanic(errors.Wrap(err, "unable to get gcp interactor"))
	li, err := infrastructure.GetLocalInteractor()
	HandlePanic(errors.Wrap(err, "unable to get local interactor"))
	// wire up interactors to webservice handler
	wsh = WebserviceHandler{}
	wsh.GcpInteractor = gcpi
	wsh.LocalInteractor = li
}

type endpoint struct {
	Api
	expectedBody	string
}

func TestEndpoints(t *testing.T) {
	Debug.Printf("fileName: %s", fileName)

	var endpoints = []endpoint{
		{Api{wsh.Health,
			"/health"},
			`{"alive": true}`},
		{Api{wsh.ListSourceBuckets,
			"/list-source-buckets?projectId="+ Config.GcpSourceProjectId},
			`{"buckets":[{"name":"lexttc3-my-backup-bucket"},{"name":"lexttc3-my-source-bucket"}]}`},
		{Api{wsh.ListSinkBuckets,
			"/list-sink-buckets?projectId="+ Config.GcpSinkProjectId},
			`{"buckets":[{"name":"lexttc3-my-backup-bucket"},{"name":"lexttc3-my-source-bucket"}]}`},
		{Api{wsh.UploadFile,
			"/upload-file?fileName="+fileName},
			`{"success":true}`},
		//{Api{wsh.DownloadFile,
		//	"/download-file?fileName="+fileName},
		//	`{"success":true}`},
		{Api{wsh.SourceFileExists,
			 "/source-file-exists?fileName="+fileName},
			`{"exists":true}`},
		{Api{wsh.LocalFileExists,
			"/local-file-exists?fileName="+fileName},
			`{"exists":true}`},
	}

	t.Log("Testing API endpoints...")
	{
		for _, ep := range endpoints {
			{
				req, err := http.NewRequest("GET", ep.Api.Url, nil)
				if err != nil {
					t.Fatal(err)
				}
				// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(ep.Api.Handler)
				// Our handlers implement http.Handler, so we can call their ServeHTTP method directly
				handler.ServeHTTP(rr, req)
				t.Logf("\tChecking \"%s\" for status code \"%d\"",
					ep.Api.Url, http.StatusOK)
				if status := rr.Code; status != http.StatusOK {
					t.Errorf("\t\t%v handler returned wrong status code: got %v want %v",
						failure, status, http.StatusOK)
				}
				t.Logf("\tChecking \"%s\" for expected body", ep.Api.Url)
				Debug.Println("rr.Body.String(): ", rr.Body.String())
				if strings.TrimSpace(rr.Body.String()) != ep.expectedBody {
					t.Errorf("\t\t%v handler returned unexpected body: got %v want %v",
						failure, rr.Body.String(), ep.expectedBody)
				}
			}
		}
	}
}
