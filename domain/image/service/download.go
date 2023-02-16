package service

import (
	"context"
	"fmt"
	"io"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"
)

func (s Image) Download(ctx context.Context, request model.DownloadRequest) (io.Reader, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", request.ImageFolder, request.ImageName), os.O_RDONLY, 0600)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Error opening local file %v", err)
		return nil, err
	}
	return f, nil
}
