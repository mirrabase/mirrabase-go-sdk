package auth

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) Signup(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) VerifyEmail(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) ResendConfirmation(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Login(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) Logout() error { return errNotImplemented }

func (s *Service) GetUser() (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) LoginAndSetSession(payload map[string]any) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) SetSession(token string) error { return errNotImplemented }

func (s *Service) ClearSession() {}

func (s *Service) GetSession() string { return "" }
