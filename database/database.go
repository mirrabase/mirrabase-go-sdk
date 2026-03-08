package database

import "errors"

var errNotImplemented = errors.New("not implemented")

type Service struct{}

type TableClient struct {
	table string
}

func New() *Service { return &Service{} }

func (s *Service) ListTables(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) DropTable(name, projectID string) error { return errNotImplemented }

func (s *Service) Query(sql, projectID string) (map[string]any, error) { return nil, errNotImplemented }

func (s *Service) ListBackups(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) CreateBackup(projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (s *Service) FromTable(table string) *TableClient { return &TableClient{table: table} }

func (t *TableClient) Select(projectID string) ([]map[string]any, error) {
	return nil, errNotImplemented
}

func (t *TableClient) SelectByID(id, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (t *TableClient) Insert(payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (t *TableClient) Update(id string, payload map[string]any, projectID string) (map[string]any, error) {
	return nil, errNotImplemented
}

func (t *TableClient) Delete(id, projectID string) error { return errNotImplemented }
