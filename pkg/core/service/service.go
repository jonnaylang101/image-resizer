package service

import (
	"errors"

	"github.com/jonnaylang101/image-resizer/pkg/core/domain"
	"github.com/jonnaylang101/image-resizer/pkg/core/ports"
)

type service struct{}

func New() ports.Service {
	return &service{}
}

func (s *service) Resize(cfg domain.ResizeRequestConfig) (domain.ResizeResponse, error) {
	res := domain.ResizeResponse{}
	return res, errors.New("not implemented")
}
