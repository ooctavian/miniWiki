package model

import "io"

type UploadRequest struct {
	ImageFolder string
	ImageName   string
	Image       io.Reader
}

type DownloadRequest struct {
	ImageFolder string
	ImageName   string
}
