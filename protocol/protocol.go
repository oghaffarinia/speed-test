package protocol

import (
	"io"
	"io/ioutil"
	"log"
)

func Handle(conn io.ReadWriteCloser) {
	defer conn.Close()
	header := make([]byte, 1)
	_, err := conn.Read(header)
	if err != nil {
		log.Printf("Error reading header: %v", err)
		return
	}
	switch TestType(header[0]) {
	case TestUpload: // Discard all incoming data
		io.Copy(ioutil.Discard, conn)

	case TestDownload: // Write dummy data
		io.Copy(conn, dummy(0))
	}
}

type dummy int

func (dummy) Read(p []byte) (n int, err error) {
	return len(p), nil
}
