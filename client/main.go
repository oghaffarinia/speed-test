package main

import (
	"fmt"
	"github.com/omidplus/arvan/config"
	"github.com/omidplus/arvan/protocol"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <path-to-config-file>", os.Args[0])
	}
	config, err := config.LoadClient(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// FIXME: Use the same connection for both tests
	testUpload(config)
	testDownload(config)
}

func testUpload(c *config.Client) {
	tester, err := protocol.NewTester(c.Protocol, c.Address, protocol.TestUpload, c.Timeout)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tester.Close()
	speed := tester.Test(c.UploadTime)
	fmt.Printf("Upload speed: %v/s\n", speed)
}

func testDownload(c *config.Client) {
	tester, err := protocol.NewTester(c.Protocol, c.Address, protocol.TestDownload, c.Timeout)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tester.Close()
	speed := tester.Test(c.DownloadTime)
	fmt.Printf("Download speed: %v/s\n", speed)
}
