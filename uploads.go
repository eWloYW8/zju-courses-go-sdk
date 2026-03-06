package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// UploadsService handles file upload and resource-related API operations.
type UploadsService struct {
	client *Client
}

// --- Response Types ---

type UploadsListResponse struct {
	Uploads []*model.Upload `json:"uploads"`
	model.Pagination
}

// --- Upload Operations ---

// GetUpload returns information about an uploaded file.
func (s *UploadsService) GetUpload(ctx context.Context, uploadID int) (*model.Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	result := new(model.Upload)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// UploadFile uploads a file from disk.
func (s *UploadsService) UploadFile(ctx context.Context, filePath string, params map[string]string) (*model.Upload, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return s.UploadReader(ctx, file, filepath.Base(filePath), params)
}

// UploadReader uploads a file from a reader.
func (s *UploadsService) UploadReader(ctx context.Context, reader io.Reader, filename string, params map[string]string) (*model.Upload, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		part, err := writer.CreateFormFile("file", filename)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, reader); err != nil {
			pw.CloseWithError(err)
			return
		}
		for k, v := range params {
			if err := writer.WriteField(k, v); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		writer.Close()
	}()

	urlStr := "/api/uploads"
	reqURL, err := s.client.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), pr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	result := new(model.Upload)
	_, err = s.client.do(req, result)
	return result, err
}

// DeleteUpload deletes an uploaded file.
func (s *UploadsService) DeleteUpload(ctx context.Context, uploadID int) error {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// UpdateUpload updates upload metadata.
func (s *UploadsService) UpdateUpload(ctx context.Context, uploadID int, body interface{}) (*model.Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	result := new(model.Upload)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// --- Upload References ---

// UploadReference uploads a reference file for an activity.
func (s *UploadsService) UploadReference(ctx context.Context, filePath string, activityID int) (*model.Upload, error) {
	return s.UploadFile(ctx, filePath, map[string]string{
		"parent_id":   fmt.Sprintf("%d", activityID),
		"parent_type": "materialactivity",
	})
}

// CreateReference creates a reference link.
func (s *UploadsService) CreateReference(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/uploads/reference", body, &result)
	return result, err
}

// DeleteReference deletes a reference.
func (s *UploadsService) DeleteReference(ctx context.Context, referenceID int) error {
	u := fmt.Sprintf("/api/uploads/references/%d", referenceID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- Share & Document ---

// ShareToCourses shares uploads to courses.
func (s *UploadsService) ShareToCourses(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/uploads/share-to-courses", body, nil)
	return err
}

// GetDocumentPreviewURL returns a document preview URL.
func (s *UploadsService) GetDocumentPreviewURL(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/document/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// UploadScreenShot uploads a screenshot.
func (s *UploadsService) UploadScreenShot(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/uploads/screen-shot", body, &result)
	return result, err
}

// ShareScreenShot shares a screenshot.
func (s *UploadsService) ShareScreenShot(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/uploads/screen-shot/share-to", body, nil)
	return err
}

// BatchBlobUpload uploads multiple blobs.
func (s *UploadsService) BatchBlobUpload(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/uploads/batch/blob", body, &result)
	return result, err
}

// QueryUploadDetails queries details for multiple uploads.
func (s *UploadsService) QueryUploadDetails(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/uploads/details/query", body, &result)
	return result, err
}

// GetUptoken returns an upload token.
func (s *UploadsService) GetUptoken(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uptoken?id=%d", uploadID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Audio Upload ---

// UploadAudio uploads an audio file.
func (s *UploadsService) UploadAudio(ctx context.Context, filePath string, params map[string]string) (*model.Upload, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, file); err != nil {
			pw.CloseWithError(err)
			return
		}
		for k, v := range params {
			if err := writer.WriteField(k, v); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		writer.Close()
	}()

	reqURL, err := s.client.baseURL.Parse("/api/uploads/audio")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), pr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	result := new(model.Upload)
	_, err = s.client.do(req, result)
	return result, err
}

// --- H5 Courseware ---

// GetH5Courseware returns H5 courseware information.
func (s *UploadsService) GetH5Courseware(ctx context.Context, coursewareID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/h5-courseware/%d", coursewareID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// UploadH5Courseware uploads H5 courseware.
func (s *UploadsService) UploadH5Courseware(ctx context.Context, uploadID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/h5-courseware/upload/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// --- SCORM ---

// UploadSCORM uploads SCORM content.
func (s *UploadsService) UploadSCORM(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/scorm/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, nil, &result)
	return result, err
}

// --- Online Videos ---

// GetOnlineVideo returns online video information.
func (s *UploadsService) GetOnlineVideo(ctx context.Context, videoID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/online-videos/%d", videoID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Media Captions ---

// GetMediaCaptionTemplate returns a media caption template.
func (s *UploadsService) GetMediaCaptionTemplate(ctx context.Context, lang string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/media-captions/template?lang=%s", lang)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetMediaCaptionProgress returns media caption task progress.
func (s *UploadsService) GetMediaCaptionProgress(ctx context.Context, mediaIDs string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/media/media-caption-tasks/progress?media_ids=%s", mediaIDs)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Duplicate Detection ---

// CheckDuplicate checks a file for duplicate content.
func (s *UploadsService) CheckDuplicate(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/duplicate-detect/file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// DownloadDuplicateReport downloads a duplicate detection report.
func (s *UploadsService) DownloadDuplicateReport(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/duplicate-detect/report/download", &result)
	return result, err
}
