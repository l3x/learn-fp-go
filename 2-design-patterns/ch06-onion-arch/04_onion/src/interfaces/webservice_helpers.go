package interfaces

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"fmt"
	"domain"
	. "utils"
)

var ErrorResponse = []byte("Error")
var sf = fmt.Sprintf

func getFormat(r *http.Request) (format string) {
	//format = r.URL.Query()["format"][0]
	// Hard code json for now
	format = "json"
	return
}

func setFormat(format string, data interface{}) ([]byte, error) {
	var apiOutput []byte
	if format == "json" {
		output, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal data to json")
		}
		apiOutput = output
	} else {
		Error.Printf("invalid data format encountered")
		apiOutput = ErrorResponse
	}
	return apiOutput, nil
}

func handleSuccess(debugMsg, msg string, req *http.Request, res http.ResponseWriter, err error, success bool) {
	Debug.Printf(debugMsg)
	response := domain.Outcome{}
	response.Success = success
	if err != nil {
		Error.Printf("Failed to %s. %v", msg, err)
	}
	output, err := setFormat(getFormat(req), response)
	if err != nil {
		output = ErrorResponse
		Error.Printf("Failed to setFormat. %v",  err)
	}
	Debug.Printf("string(output): %s", string(output))
	fmt.Fprintln(res, string(output))
}


func handleExists(debugMsg, msg string, req *http.Request, res http.ResponseWriter, err error, exists bool) {
	Debug.Printf(debugMsg)
	response := domain.Existence{}
	response.Exists = exists
	if err != nil {
		Error.Printf("Failed to %s. %v", msg, err)
	}
	output, err := setFormat(getFormat(req), response)
	if err != nil {
		output = ErrorResponse
		Error.Printf("Failed to setFormat. %v",  err)
	}
	Debug.Printf("string(output): %s", string(output))
	fmt.Fprintln(res, string(output))
}

func handleBuckets(debugMsg, msg string, req *http.Request, res http.ResponseWriter, err error, bucketNames []domain.Bucket) {
	Debug.Printf(debugMsg)
	response := domain.Buckets{}
	for _, bucketName := range bucketNames {
		Debug.Printf("bucketName: %s", bucketName)
		response.Buckets = append(response.Buckets, bucketName)
	}
	if err != nil {
		Error.Printf("Failed to %s. %v", msg, err)
	}
	output, err := setFormat(getFormat(req), response)
	if err != nil {
		output = ErrorResponse
		Error.Printf("Failed to setFormat. %v",  err)
		//return
	}
	Debug.Printf("string(output): %s", string(output))
	fmt.Fprintln(res, string(output))
}






