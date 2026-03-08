package apikeys

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) List(projectID string) ([]map[string]any, error) { return nil, errNotImplemented }

func (s *Service) Create(payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Revoke(keyID, projectID string) error { return errNotImplemented }

func (s *Service) RateLimits(projectID string) (map[string]any, error) { return nil, errNotImplemented }
