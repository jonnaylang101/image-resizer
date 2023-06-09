package inmem

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jonnaylang101/image-resizer/pkg/core/ports"
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
		return errors.New(ports.ErrInvalidStoragePath)
	}
	if image == nil {
		return errors.New(ports.ErrInvalidImageFile)
	}

	fp := filepath.Clean(filepath.Join(s.memPath, url.PathEscape(storagePath)))

	if _, err := os.Stat(fp); !errors.Is(err, os.ErrNotExist) {
		return errors.New(ports.ErrImageFileDuplication)
	}

	f, err := os.Create(fp)
	if f != nil {
		defer func() { _ = f.Close() }()
	}
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
		return nil, errors.New(ports.ErrInvalidStoragePath)
	}

	fp := filepath.Clean(filepath.Join(s.memPath, url.PathEscape(storagePath)))

	if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New(ports.ErrFileNotFound)
	}

	return os.Open(fp)
}
