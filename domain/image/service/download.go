package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func (s Image) Download(ctx context.Context, request model.DownloadRequest) (io.Reader, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s%s/%s", s.Destination, request.ImageFolder, request.ImageName), os.O_RDONLY, 0600)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logrus.WithContext(ctx).Errorf("File not found %v", err)
			return nil, utils.NotFoundError{
				Item: "image",
				Id:   request.ImageName,
			}
		}
		logrus.WithContext(ctx).Errorf("Error opening local file %v", err)
		return nil, err
	}
	return f, nil
}
