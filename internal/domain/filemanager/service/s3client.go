package service

import (
	"context"

	"miniWiki/internal/config"
	"miniWiki/internal/domain/filemanager/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct {
	uploader *s3manager.Uploader
}

func NewS3Client(config config.S3Config) (*S3Client, error) {
	s3Config := &aws.Config{
		Endpoint:    aws.String(config.Endpoint),
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.KeyID, config.SecretKey, config.Token),
	}
	s3Session, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	return &S3Client{
		uploader: s3manager.NewUploader(s3Session),
	}, nil
}

func (c *S3Client) Upload(ctx context.Context, req model.UploadRequest) error {
	_, err := c.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Body:   req.File,
		Key:    aws.String(req.Filename),
		Bucket: aws.String(req.Folder),
	})
	return err
}
