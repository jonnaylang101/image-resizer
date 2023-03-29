package ports

import "github.com/jonnaylang101/image-resizer/pkg/core/domain"

type Service interface {
	Resize(width int32, height int32, filenameSuffix string, sourceFileStoragePaths ...string) (domain.ResizeResponse, error)
}
