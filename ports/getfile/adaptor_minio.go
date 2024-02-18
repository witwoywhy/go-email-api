package getfile

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type adaptorMinio struct {
	client *minio.Client
	ctx    context.Context
}

func NewAdaptorMinio(client *minio.Client) Port {
	return &adaptorMinio{
		client: client,
		ctx:    context.Background(),
	}
}

func (a *adaptorMinio) Execute(request Request) (*Response, error) {
	obj, err := a.client.GetObject(a.ctx, request.BucketName, request.FileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &Response{
		Object: obj,
	}, nil
}
