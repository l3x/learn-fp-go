package main

import (
	"fmt"
	"net"
)

//import "theirpkg"

//func MyFunction(t *theirpkg.AType)
//
//func MyOtherFunction(i theirpkg.AnInterface)

type errorBehavior interface {
	Retryable() bool
}


func IsRetryable(err error) bool {
	eb, ok := err.(errorBehavior)
	return ok && eb.Retryable()
}

//type writeFlusher interface {
//     io.Writer
//     http.Flusher
//}

type Dividend struct {
	Val int
}
func (n Dividend) Divide(divisor int) int {
	return n.Val/divisor
}

type BytesReadConn struct {
	net.Conn
	BytesRead uint64
}

func (brc *BytesReadConn) Read(p []byte) (int, error) {
	n, err := brc.Conn.Read(p)
	brc.BytesRead += uint64(n)
	return n, err
}

func main() {
	d := Dividend{2}
	fmt.Printf("%d", d.Divide(0))
}
