package domain

const (
	GoogleCloudBucket HostProvider = iota
	SourceFlow FlowType = iota
	SinkFlow
)

type (
	HostProvider int
	FlowType  int
)

type CloudStorage struct {
	HostProvider HostProvider //Host location for log files, e.g., google cloud bucket
	ProjectId    string       //Project Id for this GCP storage account
	FlowType     FlowType     //source or sink
}

type LocalRepository interface {
	FileExists(fileName string) (fileExists bool, err error)
}

type BucketRepository interface {
	List(projectId string) (buckets []Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type FileRepository interface {
	Store(file File)
	FindById(id int) File
}

type Bucket struct {
	Name    string `json:"name"`
}
type Buckets struct {
	Buckets []Bucket `json:"buckets"`
}
