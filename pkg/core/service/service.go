package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jonnaylang101/image-resizer/pkg/core/domain"
	"github.com/jonnaylang101/image-resizer/pkg/core/ports"
)

const (
	ErrInvalidWidth                 = "Resize error: invalid width provided via ResizeRequestConfig"
	ErrInvalidHeight                = "Resize error: invalid height provided via ResizeRequestConfig"
	ErrNoProvidedStoragePaths       = "Resize error: no image storage paths provided"
	minDimension              int32 = 1
	defaultSuffixFormat             = "--resized-%d-%d"
)

type service struct {
	store ports.Storage
}

func New(store ports.Storage) ports.Service {
	return &service{
		store: store,
	}
}

// Resize will locate files in storage using the provided sourceStoragePaths, it will resize all these images
// to the width and height provided and resave them to the storage account with the filenameSuffix.
// If any files can't be found using the storagePaths these will be reported after the method has completed working
// on any files it could locate. If no filenameSuffix is passed, the suffix will default to --resized-<width>-<height>
func (s *service) Resize(width, height int32, filenameSuffix string, sourceFileStoragePaths ...string) (domain.ResizeResponse, error) {
	res := domain.ResizeResponse{}
	if len(sourceFileStoragePaths) < 1 {
		return res, errors.New(ErrNoProvidedStoragePaths)
	}
	if width < minDimension {
		return res, errors.New(ErrInvalidWidth)
	}
	if height < minDimension {
		return res, errors.New(ErrInvalidHeight)
	}
	if filenameSuffix == "" {
		filenameSuffix = fmt.Sprintf(defaultSuffixFormat, width, height)
	}

	// TODO: build this without concurrency first then examine the use of concurrent pipelines - benchmark first

	res.ProcessedFileStoragePaths = make([]string, len(sourceFileStoragePaths))
	for i, sp := range sourceFileStoragePaths {
		res.ProcessedFileStoragePaths[i] = addSuffix(sp, filenameSuffix)
	}

	return res, nil
}

func addSuffix(origPath, suffix string) string {
	if origPath == "" || suffix == "" {
		return origPath
	}

	sb := strings.Builder{}
	sb.WriteString(strings.TrimSuffix(origPath, filepath.Ext(origPath)))
	sb.WriteString(filepath.Clean(suffix))
	sb.WriteString(filepath.Ext(origPath))

	return sb.String()
}
