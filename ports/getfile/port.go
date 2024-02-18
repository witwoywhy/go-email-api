package getfile

import "github.com/minio/minio-go/v7"

type Port interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	BucketName string
	FileName   string
}

type Response struct {
	Object *minio.Object
}
