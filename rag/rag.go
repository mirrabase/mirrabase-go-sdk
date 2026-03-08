package rag

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) ListCollections(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) CreateCollection(name, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) DeleteCollection(collectionID, projectID string) error { return errNotImplemented }

func (s *Service) Query(payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Inference(collection, query, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}
