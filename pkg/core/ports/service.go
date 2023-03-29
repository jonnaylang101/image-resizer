package ports

import "github.com/jonnaylang101/image-resizer/pkg/core/domain"

type Service interface {
	Resize(domain.ResizeRequestConfig) (domain.ResizeResponse, error)
}
