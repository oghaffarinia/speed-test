package protocol

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

type Tester interface {
	Test(duration time.Duration) Size
	io.Closer
}

type TestType byte

const (
	TestUpload TestType = iota + 1
	TestDownload
)

func NewTester(protocol, address string, t TestType, timeout time.Duration) (Tester, error) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	_, err = conn.Write([]byte{byte(t)})
	if err != nil {
		conn.Close()
		return nil, err
	}
	packet := NewPacket(32 * 1024)
	switch t {
	case TestUpload:
		return &tester{
			ReadWriteCloser: conn,
			packet:          packet,
			job: func(ctx context.Context, p Packet, conn io.ReadWriteCloser) (int64, error) {
				return p.WriteTo(ctx, conn)
			},
		}, nil
	case TestDownload:
		return &tester{
			ReadWriteCloser: conn,
			packet:          packet,
			job: func(ctx context.Context, p Packet, conn io.ReadWriteCloser) (int64, error) {
				return p.ReadFrom(ctx, conn)
			},
		}, nil
	default:
		return nil, fmt.Errorf("Invalid test type %d", t)
	}
}

type tester struct {
	io.ReadWriteCloser
	packet Packet
	job    func(ctx context.Context, p Packet, conn io.ReadWriteCloser) (int64, error)
}

func (t *tester) Test(duration time.Duration) Size {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	startTime := time.Now()
	length, _ := t.job(ctx, t.packet, t.ReadWriteCloser)
	return Size(float64(length) / time.Since(startTime).Seconds())

}
