package config

import (
	"github.com/BurntSushi/toml"
	"strings"
	"time"
)

type Client struct {
	Protocol     string
	Address      string
	Timeout      time.Duration
	UploadTime   time.Duration
	DownloadTime time.Duration
}

type Server struct {
	Protocol string
	Port     int
	Timeout  time.Duration
}

type client struct {
	Server clientConnection
	Test   testInfo
}

type server struct {
	Server serverConnection
}

type clientConnection struct {
	Protocol           string
	Address            string
	Read_write_timeout duration
}

type serverConnection struct {
	Protocol           string
	Port               int
	Read_write_timeout duration
}

type testInfo struct {
	Upload_time   duration
	Download_time duration
}

func LoadClient(path string) (*Client, error) {
	var c client
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return nil, err
	}
	result := &Client{
		Protocol:     c.Server.Protocol,
		Address:      c.Server.Address,
		Timeout:      time.Duration(c.Server.Read_write_timeout),
		UploadTime:   time.Duration(c.Test.Upload_time),
		DownloadTime: time.Duration(c.Test.Download_time),
	}
	if !strings.Contains(result.Address, ":") {
		result.Address = result.Address + ":7777"
	}
	return result, err
}

func LoadServer(path string) (*Server, error) {
	var s server
	_, err := toml.DecodeFile(path, &s)
	if err != nil {
		return nil, err
	}
	return &Server{
		Protocol: s.Server.Protocol,
		Port:     s.Server.Port,
		Timeout:  time.Duration(s.Server.Read_write_timeout),
	}, nil
}

type duration time.Duration

func (d *duration) UnmarshalText(text []byte) error {
	v, err := time.ParseDuration(string(text))
	if err == nil {
		*d = duration(v)
	}
	return err
}
