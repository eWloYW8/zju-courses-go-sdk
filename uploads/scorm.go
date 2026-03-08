package uploads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// SCORM-related methods are added to Service since SCORM
// is a content packaging standard related to uploads/content.

// --- SCORM CMI (Computer Managed Instruction) ---

// GetSCORMCMI returns SCORM CMI data for an activity and SCO.
func (s *Service) GetSCORMCMI(ctx context.Context, activityID int, scoID string) (*SCORMCMIData, error) {
	u := fmt.Sprintf("/api/activity/%d/scorm/%s/cmi", activityID, scoID)
	result := new(SCORMCMIData)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateSCORMCMI updates SCORM CMI data for an activity and SCO.
func (s *Service) UpdateSCORMCMI(ctx context.Context, activityID int, scoID string, data *SCORMCMIData) error {
	u := fmt.Sprintf("/api/activity/%d/scorm/%s/cmi", activityID, scoID)
	_, err := s.client.Post(ctx, u, data, nil)
	return err
}

// ReuploadFile re-uploads a file.
func (s *Service) ReuploadFile(ctx context.Context, uploadID int, body interface{}) (*model.Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d/reupload", uploadID)
	result := new(model.Upload)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// BuildSCORMPreviewURL builds the frontend preview URL for uploaded SCORM content.
func (s *Service) BuildSCORMPreviewURL(uploadID int, sco string, parameters string) string {
	u := fmt.Sprintf("/api/uploads/scorm/%d?sco=%s&preview=true", uploadID, url.QueryEscape(sco))
	if parameters != "" {
		u += "&para=" + url.QueryEscape(parameters)
	}
	return u
}

// GetUploadURL returns the file access URL.
func (s *Service) GetUploadURL(ctx context.Context, uploadID int) (*UploadURLResponse, error) {
	u := fmt.Sprintf("/api/uploads/%d/url", uploadID)
	result := new(UploadURLResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetUploadBlob returns the file blob/preview.
func (s *Service) GetUploadBlob(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/%d/blob", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}
