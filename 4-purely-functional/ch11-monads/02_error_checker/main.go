package main

import (
    "os"
    "log"
    "io"
)

func logInfo(numBytesRead int, byteSlice []byte) {
    log.Printf("Number of bytes read: %d\n", numBytesRead)
    log.Printf("Data read: %s\n", byteSlice)

}

type twoByteReader struct {
    err      error
    reader io.Reader
}

func (tbr *twoByteReader) read() (numBytesRead int, byteSlice []byte)  {
    if tbr.err != nil {
        return
    }
    byteSlice = make([]byte, 2)
    numBytesRead, tbr.err = io.ReadFull(tbr.reader, byteSlice)
    logInfo(numBytesRead, byteSlice)
    return
}

func main() {
    file, err := os.Open("alphabet.txt")
    if err != nil {
        log.Fatal(err)
    }

    byteSlice := make([]byte, 2)
    numBytesRead, err := io.ReadFull(file, byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    logInfo(numBytesRead, byteSlice)

    byteSlice = make([]byte, 2)
    numBytesRead, err = io.ReadFull(file, byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    logInfo(numBytesRead, byteSlice)

    byteSlice = make([]byte, 2)
    numBytesRead, err = io.ReadFull(file, byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    logInfo(numBytesRead, byteSlice)

    tbr := &twoByteReader{reader: file}
    byteSlice = make([]byte, 2)
    tbr.read()
    tbr.read()
    tbr.read()
}
