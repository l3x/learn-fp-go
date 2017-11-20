package usecases

import (
	"domain"
)

type LocalInteractor struct {
	LocalRepository domain.LocalRepository
}

func (interactor *LocalInteractor) LocalFileExists(fileName string) (fileExists bool, err error) {
	return interactor.LocalRepository.FileExists(fileName)
}

type GcpInteractor struct {
	SourceBucketRepository domain.BucketRepository
	SinkBucketRepository domain.BucketRepository
}

func (interactor *GcpInteractor) ListSourceBuckets(projectId string) (buckets []domain.Bucket, err error) {
	return interactor.SourceBucketRepository.List(projectId)
}

func (interactor *GcpInteractor) ListSinkBuckets(projectId string) (buckets []domain.Bucket, err error) {
	return interactor.SinkBucketRepository.List(projectId)
}

func (interactor *GcpInteractor) SourceFileExists(fileName string) (fileExists bool, err error) {
	return interactor.SourceBucketRepository.FileExists(fileName)
}

func (interactor *GcpInteractor) DownloadFile(fileName string) (success bool, err error) {
	return interactor.SourceBucketRepository.DownloadFile(fileName)
}

func (interactor *GcpInteractor) UploadFile(fileName string) (success bool, err error) {
	return interactor.SinkBucketRepository.UploadFile(fileName)
}
