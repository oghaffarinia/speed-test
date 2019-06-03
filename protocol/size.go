package protocol

import "fmt"

// Source: https://golang.org/doc/effective_go.html#constants

type Size float64

const (
	_       = iota // ignore first value by assigning to blank identifier
	KB Size = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (size Size) String() string {
	switch {
	case size >= YB:
		return fmt.Sprintf("%.2f YB", size/YB)
	case size >= ZB:
		return fmt.Sprintf("%.2f ZB", size/ZB)
	case size >= EB:
		return fmt.Sprintf("%.2f EB", size/EB)
	case size >= PB:
		return fmt.Sprintf("%.2f PB", size/PB)
	case size >= TB:
		return fmt.Sprintf("%.2f TB", size/TB)
	case size >= GB:
		return fmt.Sprintf("%.2f GB", size/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", size/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", size/KB)
	}
	return fmt.Sprintf("%.2f B", size)
}
