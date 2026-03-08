package syllabus

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Service handles syllabus-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetSyllabus returns a syllabus.
func (s *Service) GetSyllabus(ctx context.Context, syllabusID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateSyllabus creates a new syllabus.
func (s *Service) CreateSyllabus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/syllabus", body, &result)
	return result, err
}

// UpdateSyllabus updates a syllabus.
func (s *Service) UpdateSyllabus(ctx context.Context, syllabusID int, body interface{}) error {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteSyllabus deletes a syllabus.
func (s *Service) DeleteSyllabus(ctx context.Context, syllabusID int) error {
	return s.DeleteSyllabusWithOptions(ctx, syllabusID, nil)
}

// DeleteSyllabusWithOptions deletes a syllabus with optional query parameters.
func (s *Service) DeleteSyllabusWithOptions(ctx context.Context, syllabusID int, opts *DeleteSyllabusOptions) error {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	if opts != nil && opts.DeleteRelatedActivity {
		u += "?delete_related_activity=true"
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ResortSyllabus resorts syllabuses.
func (s *Service) ResortSyllabus(ctx context.Context, body interface{}) error {
	_, err := s.client.Put(ctx, "/api/syllabus/resort", body, nil)
	return err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}
