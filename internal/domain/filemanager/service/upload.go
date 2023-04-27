package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"miniWiki/internal/domain/filemanager/model"

	"github.com/sirupsen/logrus"
)

func (s FileManager) Upload(ctx context.Context, request model.UploadRequest) error {
	createDir(request.Folder)
	outputFile, err := os.OpenFile(
		fmt.Sprintf("%s%s/%s",
			s.Destination,
			request.Folder,
			request.Filename),
		os.O_WRONLY|os.O_CREATE,
		0600)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error opening local file: %v", err)
		return err
	}

	_, err = io.Copy(outputFile, request.File)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error writing image: %v", err)
		return err
	}
	err = outputFile.Close()
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error writing image: %v", err)
		return err
	}
	return err
}

func createDir(dirName string) {
	if _, err := os.Stat(dirName); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			logrus.Fatalf("Creating dir error: %v", err)
		}
	}
}
