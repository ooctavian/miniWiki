package service

import (
	"context"
	"fmt"
	"io"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"
)

func (s Image) Upload(ctx context.Context, request model.UploadRequest) error {
	outputFile, err := os.OpenFile(fmt.Sprintf("%s/%s", request.ImageFolder, request.ImageName), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Error opening local file: %v", err)
		return err
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, request.Image)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Error writing image: %v", err)
	}
	return err
}
