package projects

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) List() ([]map[string]any, error) { return nil, errNotImplemented }

func (s *Service) Create(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Get(projectID string) (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) Update(projectID string, payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Remove(projectID string) error { return errNotImplemented }

func (s *Service) SetActive(projectID string) {}
