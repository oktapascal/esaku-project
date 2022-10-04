package storages

import (
	"context"
	"esaku-project/configs"
	"esaku-project/helpers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewSessionAws(configuration configs.Config) *s3.Client {
	accessKey := configuration.Get("AWS_ACCESS_KEY_ID")
	accessSecretKey := configuration.Get("AWS_SECRET_ACCESS_KEY")
	region := configuration.Get("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: accessSecretKey,
			},
		}),
	)

	helpers.PanicIfError(err)

	client := s3.NewFromConfig(cfg)

	return client
}
