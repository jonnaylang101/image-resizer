package ports

import (
	"io"
	"os"
)

//go:generate mockgen -destination=./../../actors/repositories/storage/mockstorage/mock_store.go -package=mockstorage -source=storage.go Storage
type Storage interface {
	Add(storagePath string, image io.Reader) error
	GetByStoragePath(storagePath string) (*os.File, error)
}
