package infrastructure

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"domain"
	"google.golang.org/api/iterator"
	. "utils"
	"io/ioutil"
	"os"
	"io"
	"interfaces"
	"usecases"
	"github.com/pkg/errors"
)

type GcpHandler struct {
	Client *storage.Client
}

var GcpInteractor *usecases.GcpInteractor

func NewGcpHandler(keyFile string) (gcpHandler *GcpHandler, err error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(keyFile))
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new storage client")
	}
	gcpHandler = new(GcpHandler)
	gcpHandler.Client = client
	return
}

func (handler *GcpHandler) ListBuckets(flowType domain.FlowType, projectId string) (buckets []domain.Bucket, err error) {
	Debug.Printf("Running: ListBuckets(%v, %v)", flowType, projectId)
	client := handler.Client
	ctx := context.Background()
	it := client.Buckets(ctx, projectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "bucket iterator error")
		}
		buckets = append(buckets, domain.Bucket{battrs.Name})
	}
	return
}

func (handler *GcpHandler) FileExists(fileName string) (fileExists bool, err error) {
	ctx := context.Background()
	bucketName := Config.SourceBucketName
	newFile := domain.NewFile(fileName)
	fullPath := newFile.FullHostPath(Config.GcpSourceDir)
	Debug.Printf("fullPath: %s", fullPath)
	br, err := handler.Client.Bucket(bucketName).Object(fullPath).NewReader(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "bucket reader error for %s", fullPath)
	} else {
		data, err := ioutil.ReadAll(br)
		defer br.Close()
		if err != nil {
			return false, errors.Wrapf(err, "ioutil.ReadAll error for %s", fullPath)
		} else if len(data) == 0 {
			return false, errors.Wrapf(err, "File size must be greater than 0 for %s", fullPath)
		}
	}
	return true, err
}

func (handler *GcpHandler) GetBucketObject(flowType domain.FlowType, projectId string, bucketName string, fileName string) storage.ObjectHandle {
	client := handler.Client
	fileObject := client.Bucket(bucketName).Object(fileName)
	return *fileObject
}

func (handler *GcpHandler) DownloadFile(fileName string) (success bool, err error) {
	newFile := domain.NewFile(fileName)
	fullFilePath := newFile.FullHostPath(Config.GcpSourceDir)
	Debug.Printf("fullFilePath: %s", fullFilePath)
	ctx := context.Background()

	Debug.Printf("Config.GcpSourceProjectId: %s", Config.GcpSourceProjectId)
	Debug.Printf("Config.SourceBucketName: %s", Config.SourceBucketName)
	Debug.Printf("fullFilePath: %s", fullFilePath)


	bucketObject := handler.GetBucketObject(domain.SourceFlow, Config.GcpSourceProjectId, Config.SourceBucketName, fullFilePath)
	fr, err := bucketObject.NewReader(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "unable to get file (%s) from bucket(%s)", fullFilePath, Config.SourceBucketName)
	}
	defer fr.Close()
	fileBytes, err := ioutil.ReadAll(fr)
	if err != nil {
		return false, errors.Wrap(err, "ioutil.ReadAll failed")
	}
	logFiles, err := newFile.Parse(fileBytes)
	if err != nil {
		return false, errors.Wrap(err, "newFile.Parse failed")
	}
	success = true
	var logFileName string
	var cachedLogFiles []string
	for i, logFile := range *logFiles {
		logFileName = newFile.FullParsedFileName(i)
		Info.Println("Encoding, caching and saving logFileName: "+logFileName)
		logFileJson, err := logFile.ToJson()
		if err != nil {
			Error.Printf("Unable to encode logFileName (%s) - ERROR: %v", logFileName, err)
			break
		}
		cachedLogFiles = append(cachedLogFiles, logFileJson)
		logFile.Write(logFileName, logFileJson)
	}
	return
}

func (handler *GcpHandler) UploadFile(fileName string) (success bool, err error) {
	ctx := context.Background()
	newFile := domain.NewFile(fileName)
	newFullPath := newFile.FullLocalPath()
	f, err := os.Open(newFullPath)
	if err != nil {
		return false, errors.Wrapf(err, "unable to open local file: %s", newFullPath)
	}
	defer f.Close()
	bucketObject := handler.GetBucketObject(domain.SinkFlow, Config.GcpSinkProjectId, Config.SinkBucketName, newFile.FullHostPath(Config.GcpSinkDir))
	wc := bucketObject.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return false, errors.Wrapf(err, "io.Copy failed for %s", newFullPath)
	}
	if err := wc.Close(); err != nil {
		return false, errors.Wrapf(err, "io.Close failed for %s", newFullPath)
	}
	success = true
	return
}

func GetGcpInteractor() (gcpInteractor *usecases.GcpInteractor, err error) {
	if GcpInteractor == nil {
		sourceHandler, err := NewGcpHandler(Config.GcpSourceKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new source gcp handler")
		}
		sinkHandler, err := NewGcpHandler(Config.GcpSinkKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new sink gcp handler")
		}
		handlers := make(map[string] interfaces.GcpHandler)
		handlers["SourceBucketRepo"] = sourceHandler
		handlers["SinkBucketRepo"] = sinkHandler
		gcpInteractor = new(usecases.GcpInteractor)
		gcpInteractor.SourceBucketRepository = interfaces.NewSourceBucketRepo(handlers)
		gcpInteractor.SinkBucketRepository = interfaces.NewSinkBucketRepo(handlers)
		GcpInteractor = gcpInteractor
	}
	return GcpInteractor, err
}



type List interface {
	Cons(val interface{}) List
	ForEach(f func(interface{}))
	Head() interface{}
	IsNil() bool
	Reverse() List
	Size() int
	Tail() List
}

type Map interface {
	Delete(key string) Map
	ForEach(f func(key string, val interface{}))
	IsNil() bool
	Keys() []string
	Lookup(key string) (interface{}, bool)
	Set(key string, value interface{}) Map
	Size() int
	String() string
	UnsafeMutableSet(key string, value interface{}) Map
}
