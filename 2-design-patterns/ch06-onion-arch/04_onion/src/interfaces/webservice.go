package interfaces

import (
	"io"
	"net/http"
	"domain"
)

type Api struct {
	Handler		 func(res http.ResponseWriter, req *http.Request)
	Url          string
}

type LocalInteractor interface {
	LocalFileExists(fileName string) (fileExists bool, err error)
}

type GcpInteractor interface {
	ListSourceBuckets(projectId string) (buckets []domain.Bucket, err error)
	ListSinkBuckets(projectId string) (buckets []domain.Bucket, err error)
	SourceFileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type WebserviceHandler struct {
	LocalInteractor LocalInteractor
	GcpInteractor   GcpInteractor
}

func (handler WebserviceHandler) Health(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"alive": true}`)
}

func (handler WebserviceHandler) LocalFileExists(res http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("fileName")
	exists, err := handler.LocalInteractor.LocalFileExists(fileName)
	handleExists(sf("Running LocalFileExists for fileName: %s...", fileName), "find file", req, res, err, exists)
}

func (handler WebserviceHandler) ListSourceBuckets(res http.ResponseWriter, req *http.Request) {
	projectId := req.FormValue("projectId")
	bucketNames, err := handler.GcpInteractor.ListSourceBuckets(projectId)
	handleBuckets(sf("Running ListSourceBuckets for projectId: %s...", projectId), "list source buckets", req, res, err, bucketNames)
}

func (handler WebserviceHandler) ListSinkBuckets(res http.ResponseWriter, req *http.Request) {
	projectId := req.FormValue("projectId")
	bucketNames, err := handler.GcpInteractor.ListSinkBuckets(projectId)
	handleBuckets(sf("Running ListSinkBuckets for projectId: %s...", projectId), "list sink buckets", req, res, err, bucketNames)
}

func (handler WebserviceHandler) SourceFileExists(res http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("fileName")
	exists, err := handler.GcpInteractor.SourceFileExists(fileName)
	handleExists(sf("Running SourceFileExists for fileName: %s...", fileName), "find file", req, res, err, exists)
}

func (handler WebserviceHandler) DownloadFile(res http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("fileName")
	success, err := handler.GcpInteractor.DownloadFile(fileName)
	handleSuccess(sf("Running DownloadFile for fileName: %s...", fileName), "upload file", req, res, err, success)
}

func (handler WebserviceHandler) UploadFile(res http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("fileName")
	success, err := handler.GcpInteractor.UploadFile(fileName)
	handleSuccess(sf("Running UploadFile for fileName: %s...", fileName), "upload file", req, res, err, success)
}
