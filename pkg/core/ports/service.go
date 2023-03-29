package ports

import "github.com/jonnaylang101/image-resizer/pkg/core/domain"

type Service interface {
	Resize(int32, int32, string, ...string) (domain.ResizeResponse, error)
}
