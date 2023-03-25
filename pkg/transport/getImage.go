package transport

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetImage(r *http.Request) (io.Reader, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		logrus.WithContext(r.Context()).Error(err)
		return nil, nil, err
	}
	file, header, err := r.FormFile("File")

	if err != nil {
		logrus.Infof("Error Retrieving the File %v", err)
		return nil, nil, err
	}

	buf := streamToByte(file)

	contentType := http.DetectContentType(buf)
	if contentType != "image/png" && contentType != "image/jpeg" {
		logrus.Infof("Unsupported image format: %v", contentType)
		return nil, nil, newUnsupportedContentType(contentType)
	}

	return bytes.NewReader(buf), header, nil
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		logrus.Error(err)
	}
	return buf.Bytes()
}
