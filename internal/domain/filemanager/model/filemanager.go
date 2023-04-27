package model

import "io"

type UploadRequest struct {
	Folder   string
	Filename string
	File     io.Reader
}
