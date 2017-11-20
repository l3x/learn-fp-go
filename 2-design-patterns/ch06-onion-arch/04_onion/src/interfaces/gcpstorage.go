package interfaces

import (
	"domain"
)

type GcpHandler interface {
	ListBuckets(flowType domain.FlowType, projectId string) (buckets []domain.Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type GcpRepo struct {
	gcpHandlers map[string]GcpHandler
	gcpHandler  GcpHandler
}

type SourceBucketRepo GcpRepo
type SinkBucketRepo GcpRepo

func NewSourceBucketRepo(gcpHandlers map[string]GcpHandler) *SourceBucketRepo {
	return &SourceBucketRepo{
		gcpHandlers: gcpHandlers,
		gcpHandler: gcpHandlers["SourceBucketRepo"],
		}
}

func (repo *SourceBucketRepo) List(projectId string) (buckets []domain.Bucket, err error) {
	return repo.gcpHandler.ListBuckets(domain.SourceFlow, projectId)
}

func (repo *SourceBucketRepo) FileExists(fileName string) (fileExists bool, err error) {
	return repo.gcpHandler.FileExists(fileName)
}

func (repo *SourceBucketRepo) DownloadFile(fileName string) (success bool, err error) {
	return repo.gcpHandler.DownloadFile(fileName)
}

func (repo *SourceBucketRepo) UploadFile(fileName string) (success bool, err error) {
	return false, nil
}

func NewSinkBucketRepo(gcpHandlers map[string]GcpHandler) *SinkBucketRepo {
	return &SinkBucketRepo{
		gcpHandlers: gcpHandlers,
		gcpHandler: gcpHandlers["SinkBucketRepo"],
	}
}

func (repo *SinkBucketRepo) List(projectId string) (buckets []domain.Bucket, err error) {
	return repo.gcpHandler.ListBuckets(domain.SinkFlow, projectId)
}

func (repo *SinkBucketRepo) FileExists(fileName string) (fileExists bool, err error) {
	return repo.gcpHandler.FileExists(fileName)
}

func (repo *SinkBucketRepo) DownloadFile(fileName string) (success bool, err error) {
	return false, nil
}

func (repo *SinkBucketRepo) UploadFile(fileName string) (success bool, err error) {
	return repo.gcpHandler.UploadFile(fileName)
}

func (repo *SinkBucketRepo) ListFileNamesToFetch(fileName string) (cloudFiles domain.CloudFiles, err error) {
	return cloudFiles, err
}


