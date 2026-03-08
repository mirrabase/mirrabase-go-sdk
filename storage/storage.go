package storage

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) ListBuckets(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) CreateBucket(payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) DeleteBucket(bucketID, projectID string) error { return errNotImplemented }

func (s *Service) ListFiles(params map[string]string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) UploadObject(bucket, path, filePath, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) DownloadObject(bucket, path, projectID string) ([]byte, error) {
	return nil, errNotImplemented
}
