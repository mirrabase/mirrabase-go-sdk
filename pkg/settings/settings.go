package settings

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) GetGeneral(projectID string) (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) UpdateGeneral(payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) ListDomains(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Pause(projectID string) (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) Delete(projectID string) error { return errNotImplemented }
