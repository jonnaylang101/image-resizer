package ports

import (
	"io"
	"os"
)

const (
	ErrInvalidStoragePath   = "storage error: invalid storagePath"
	ErrInvalidImageFile     = "storage error: invalid image file"
	ErrImageFileDuplication = "storage error: file already exists"
	ErrFileNotFound         = "storage error: file not found"
)

//go:generate mockgen -destination=./../../actors/repositories/storage/mockstorage/mock_store.go -package=mockstorage -source=storage.go Storage
type Storage interface {
	Add(storagePath string, image io.Reader) error
	GetByStoragePath(storagePath string) (*os.File, error)
}
