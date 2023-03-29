package domain

// TODO: this is a slightly smaller problem space that might not require a domain struct of any kind.
// ResizeResponse may find a better home in the ports

type ResizeResponse struct {
	ResizedImagesStoragePaths []string
}
