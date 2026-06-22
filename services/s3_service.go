package services

import (
	"context"
	"mime/multipart"
	"path"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const BucketName = "go-image-api-yui-iwamoto"

func UploadToS3(
	file multipart.File,
	fileName string,
) (string, error) {

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
	)

	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	bucket := BucketName

	_, err = client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: &bucket,
			Key:    &fileName,
			Body:   file,
		},
	)

	if err != nil {
		return "", err
	}

	url :=
		"https://" +
			BucketName +
			".s3.ap-northeast-1.amazonaws.com/" +
			fileName

	return url, nil
}

func DeleteFromS3(fileURL string) error {

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
	)

	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	fileName := path.Base(fileURL)

	bucket := BucketName

	_, err = client.DeleteObject(
		context.Background(),
		&s3.DeleteObjectInput{
			Bucket: &bucket,
			Key:    &fileName,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
