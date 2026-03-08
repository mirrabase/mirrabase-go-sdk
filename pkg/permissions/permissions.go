package permissions

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) Get(role, projectID string) (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) Update(role string, permissions []map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}
