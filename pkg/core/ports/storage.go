package ports

import (
	"io"
	"os"
)

type Storage interface {
	Add(storagePath string, image io.Reader) error
	GetByStoragePath(storagePath string) (*os.File, error)
}
