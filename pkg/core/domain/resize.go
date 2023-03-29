package domain

type ResizeRequestConfig struct {
	SourceImageStoragePath string
	Width                  int32
	Height                 int32
}

type ResizeResponse struct {
	ResizedImagesStoragePaths []string
}
