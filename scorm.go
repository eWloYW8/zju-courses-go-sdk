package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// SCORM-related methods are added to UploadsService since SCORM
// is a content packaging standard related to uploads/content.

// --- SCORM CMI (Computer Managed Instruction) ---

// GetSCORMCMI returns SCORM CMI data for an activity and SCO.
func (s *UploadsService) GetSCORMCMI(ctx context.Context, activityID int, scoID string) (*model.SCORMCMIData, error) {
	u := fmt.Sprintf("/api/activity/%d/scorm/%s/cmi", activityID, scoID)
	result := new(model.SCORMCMIData)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// UpdateSCORMCMI updates SCORM CMI data for an activity and SCO.
func (s *UploadsService) UpdateSCORMCMI(ctx context.Context, activityID int, scoID string, data *model.SCORMCMIData) error {
	u := fmt.Sprintf("/api/activity/%d/scorm/%s/cmi", activityID, scoID)
	_, err := s.client.put(ctx, u, data, nil)
	return err
}

// ReuploadFile re-uploads a file.
func (s *UploadsService) ReuploadFile(ctx context.Context, uploadID int, body interface{}) (*model.Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d/reupload", uploadID)
	result := new(model.Upload)
	_, err := s.client.post(ctx, u, body, result)
	return result, err
}

// UploadURLResponse represents the response from GetUploadURL.
type UploadURLResponse struct {
	URL string `json:"url"`
}

// GetUploadURL returns the file access URL.
func (s *UploadsService) GetUploadURL(ctx context.Context, uploadID int) (*UploadURLResponse, error) {
	u := fmt.Sprintf("/api/uploads/%d/url", uploadID)
	result := new(UploadURLResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetUploadBlob returns the file blob/preview.
func (s *UploadsService) GetUploadBlob(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/%d/blob", uploadID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}
