package inmem

import (
	"errors"
	"io"
	"net/url"
	"os"

	"github.com/jonnaylang101/image-resizer/pkg/core/ports"
)

const (
	ErrInvalidStoragePath   = "storage error: invalid storagePath"
	ErrInvalidImageFile     = "storage error: invalid image file"
	ErrImageFileDuplication = "storage error: file already exists"
	ErrFileNotFound         = "storage error: file not found"
)

type storage struct {
	memPath string
}

func NewStore(memPath string) (ports.Storage, error) {
	if err := os.MkdirAll(memPath, os.ModePerm); err != nil {
		return nil, err
	}

	return &storage{
		memPath: memPath,
	}, nil
}

func (s *storage) Add(storagePath string, image io.Reader) error {
	if storagePath == "" {
		return errors.New(ErrInvalidStoragePath)
	}
	if image == nil {
		return errors.New(ErrInvalidImageFile)
	}

	filepath := s.memPath + "/" + url.PathEscape(storagePath)

	if _, err := os.Stat(filepath); !errors.Is(err, os.ErrNotExist) {
		return errors.New(ErrImageFileDuplication)
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, image); err != nil {
		return err
	}

	return nil
}

func (s *storage) GetByStoragePath(storagePath string) (*os.File, error) {
	if storagePath == "" {
		return nil, errors.New(ErrInvalidStoragePath)
	}

	filepath := s.memPath + "/" + url.PathEscape(storagePath)
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New(ErrFileNotFound)
	}

	return os.Open(filepath)
}
