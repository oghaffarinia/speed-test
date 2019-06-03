package protocol

import (
	"context"
	"io"
)

type Packet interface {
	ReadFrom(ctx context.Context, r io.Reader) (n int64, err error)
	WriteTo(ctx context.Context, w io.Writer) (n int64, err error)
}

func NewPacket(chunk int) Packet {
	return &packet{chunk: chunk}
}

type packet struct {
	chunk int
}

func (p *packet) ReadFrom(ctx context.Context, r io.Reader) (n int64, err error) {
	data := make([]byte, p.chunk)
	for err == nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			var length int
			length, err = r.Read(data)
			n += int64(length)
		}
	}
	return
}

func (p *packet) WriteTo(ctx context.Context, w io.Writer) (n int64, err error) {
	data := make([]byte, p.chunk)
	for err == nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			var length int
			length, err = w.Write(data)
			n += int64(length)
		}
	}
	return
}
