package infrastructure

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var MinioClient *minio.Client

type MinioConfig struct {
	BaseUrl string `mapstructure:"baseUrl"`
	ID      string `mapstructure:"id"`
	Secret  string `mapstructure:"secret"`
	Token   string `mapstructure:"token"`
}

func InitStorage() {
	var config MinioConfig
	if err := viper.UnmarshalKey("minio", &config); err != nil {
		panic(fmt.Errorf("failed to load up minio config: %v", err))
	}

	client, err := minio.New(config.BaseUrl, &minio.Options{
		Creds: credentials.NewStaticV4(config.ID, config.Secret, config.Token),
	})
	if err != nil {
		panic(fmt.Errorf("failed to make minio client: %v", err))
	}

	MinioClient = client
}
