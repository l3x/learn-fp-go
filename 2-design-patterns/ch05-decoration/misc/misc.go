package main

import (
	"io"
	"os"
	"time"
)

type Hash interface {
	io.Writer
	Sum(b []byte) []byte
	Reset()
	Size() int
	BlockSize() int
}

//func ReadAll(r io.Reader) ([]byte, error)
//
//func LoggingReader(r io.Reader) io.Reader
//func LimitingReader(r io.Reader, n int64) io.Reader
//func ErrorInjector(r io.Reader) io.Reader

// The File interface is implemented by os.File.  App specific
// implementations may add concurrency, caching, stats, fuzzing, etc.
type File interface {
	io.ReaderAt
	io.WriterAt
	io.Closer
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
}

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type File interface {
	io.ReadSeeker
	io.Closer
	Name() string
}
// File is an interface to access the file part of a multipart message.
// Its contents may be either stored in memory or on disk.
// If stored on disk, the File's underlying concrete type will be an *os.File.
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

/*
Pick a Product Type:
(1) Appliance
(2) Book
(3) Clothing
3
Pick a Clothing Type:
(1) Men
(2) Women
(3) Children
2



$ go run main.go --help

Usage of main:

 -proxyPort int

    Server Port (default 8080)

 -serverPort int

    Server Port (default 3000)



INFO  : 13:46:19 Metrics server listening on 127.0.0.1:3000
INFO  : 13:46:20 Proxy listening on 127.0.0.1:8080
DEBUG : 2017/05/17 13:46:30 requester.go:114: makeRequest:
client: 13:46:30 GET http://127.0.0.1:3000
DEBUG : 2017/05/17 13:46:30 metrics.go:66: - randInt: 3081
DEBUG : 2017/05/17 13:46:31 decorator.go:107: backing off...
client: 13:46:31 GET http://web02:3000
DEBUG : 2017/05/17 13:46:31 metrics.go:66: - randInt: 2887
DEBUG : 2017/05/17 13:46:32 decorator.go:107: backing off...
client: 13:46:33 GET http://web03:3000
DEBUG : 2017/05/17 13:46:33 metrics.go:66: - randInt: 1847
DEBUG : 2017/05/17 13:46:34 decorator.go:107: backing off...
INFO  : 13:46:36 FAILURE!


DEBUG : 2017/05/17 13:47:30 requester.go:114: makeRequest:
client: 13:47:30 GET http://web03:3000
DEBUG : 2017/05/17 13:47:30 metrics.go:66: - randInt: 1445
DEBUG : 2017/05/17 13:47:31 decorator.go:107: backing off...
client: 13:47:31 GET http://web01:3000
DEBUG : 2017/05/17 13:47:31 metrics.go:66: - randInt: 3237
DEBUG : 2017/05/17 13:47:32 decorator.go:107: backing off...
client: 13:47:33 GET http://web02:3000
DEBUG : 2017/05/17 13:47:33 metrics.go:66: - randInt: 4106
DEBUG : 2017/05/17 13:47:34 decorator.go:107: backing off...
INFO  : 13:47:36 FAILURE!
DEBUG : 2017/05/17 13:47:36 requester.go:65: > 7 requests done.
DEBUG : 2017/05/17 13:47:40 requester.go:114: makeRequest:
client: 13:47:40 GET http://web03:3000
DEBUG : 2017/05/17 13:47:40 metrics.go:66: - randInt: 495
INFO  : 13:47:41 SUCCESS!
DEBUG : 2017/05/17 13:47:41 requester.go:65: > 8 requests done.


*/

func main() {

}
