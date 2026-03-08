package members

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) List(projectID string) ([]map[string]any, error) { return nil, errNotImplemented }

func (s *Service) UpdateRole(userID, role, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Remove(userID, projectID string) error { return errNotImplemented }
