package uploads

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles file upload and resource-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Upload Operations ---

// GetUpload returns information about an uploaded file.
func (s *Service) GetUpload(ctx context.Context, uploadID int) (*Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	result := new(Upload)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UploadFile uploads a file from disk.
func (s *Service) UploadFile(ctx context.Context, filePath string, params map[string]string) (*Upload, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return s.UploadReader(ctx, file, filepath.Base(filePath), params)
}

// UploadReader uploads a file from a reader.
func (s *Service) UploadReader(ctx context.Context, reader io.Reader, filename string, params map[string]string) (*Upload, error) {
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
	reqURL, err := s.client.BaseURL().Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), pr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	result := new(Upload)
	_, err = s.client.Do(req, result)
	return result, err
}

// DeleteUpload deletes an uploaded file.
func (s *Service) DeleteUpload(ctx context.Context, uploadID int) error {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateUpload updates upload metadata.
func (s *Service) UpdateUpload(ctx context.Context, uploadID int, body interface{}) (*Upload, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	result := new(Upload)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// --- Upload References ---

// UploadReference uploads a reference file for an activity.
func (s *Service) UploadReference(ctx context.Context, filePath string, activityID int) (*Upload, error) {
	return s.UploadFile(ctx, filePath, map[string]string{
		"parent_id":   fmt.Sprintf("%d", activityID),
		"parent_type": "materialactivity",
	})
}

// CreateReference creates a reference link.
func (s *Service) CreateReference(ctx context.Context, body *CreateReferenceRequest) (*UploadReference, error) {
	result := new(UploadReference)
	_, err := s.client.Post(ctx, "/api/uploads/reference", body, result)
	return result, err
}

// DeleteReference deletes a reference.
func (s *Service) DeleteReference(ctx context.Context, referenceID int) error {
	u := fmt.Sprintf("/api/uploads/references/%d", referenceID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateReferenceUpload updates the upload bound to a reference.
func (s *Service) UpdateReferenceUpload(ctx context.Context, referenceID int, uploadID int) (*UploadReference, error) {
	u := fmt.Sprintf("/api/uploads/references/%d", referenceID)
	result := new(UploadReference)
	_, err := s.client.Put(ctx, u, &UpdateReferenceUploadRequest{UploadID: uploadID}, result)
	return result, err
}

// DeleteMarkedAttachment deletes an uploaded marked attachment.
func (s *Service) DeleteMarkedAttachment(ctx context.Context, attachmentID int) error {
	u := fmt.Sprintf("/api/uploads/marked_attachment/%d", attachmentID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Share & Document ---

// ShareToCourses shares uploads to courses.
func (s *Service) ShareToCourses(ctx context.Context, body *ShareToCoursesRequest) error {
	_, err := s.client.Post(ctx, "/api/uploads/share-to-courses", body, nil)
	return err
}

// GetDocumentPreviewURL returns a document preview URL.
func (s *Service) GetDocumentPreviewURL(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/document/%d/url?preview=true", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetReferenceDocumentPreviewURL returns a document preview URL for an upload reference.
func (s *Service) GetReferenceDocumentPreviewURL(ctx context.Context, referenceID int) (*UploadURLResponse, error) {
	u := fmt.Sprintf("/api/uploads/reference/document/%d/url?preview=true", referenceID)
	result := new(UploadURLResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetEmbedMaterial returns preview metadata for an embedded material upload.
func (s *Service) GetEmbedMaterial(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/embed-material/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListUploadReferences returns reference details for an upload.
func (s *Service) ListUploadReferences(ctx context.Context, uploadID int, opts *model.ListOptions, conditions string) (*UploadReferencesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/uploads/%d/references", uploadID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(UploadReferencesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetUploadPreview returns preview metadata for an upload.
func (s *Service) GetUploadPreview(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/%d?preview=true", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetUploadPDFInfo returns PDF page information for an upload.
func (s *Service) GetUploadPDFInfo(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/%d/pdf-info?preview=true", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetUploadAudioPreview returns audio preview metadata for an upload.
func (s *Service) GetUploadAudioPreview(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/audio/%d?preview=true", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetUploadRichContent returns upload data for rich content rendering.
func (s *Service) GetUploadRichContent(ctx context.Context, uploadID int, createdAt string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/uploads/%d/in-rich-content", uploadID), map[string]string{"created_at": createdAt})
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UploadScreenShot uploads a screenshot.
func (s *Service) UploadScreenShot(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/uploads/screen-shot", body, &result)
	return result, err
}

// ShareScreenShot shares a screenshot.
func (s *Service) ShareScreenShot(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/uploads/screen-shot/share-to", body, nil)
	return err
}

// BatchBlobUpload uploads multiple blobs.
func (s *Service) BatchBlobUpload(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/uploads/batch/blob", body, &result)
	return result, err
}

// QueryUploadDetails queries details for multiple uploads.
func (s *Service) QueryUploadDetails(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/uploads/details/query", body, &result)
	return result, err
}

// GetUptoken returns an upload token.
func (s *Service) GetUptoken(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uptoken?id=%d", uploadID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Audio Upload ---

// UploadAudio uploads an audio file.
func (s *Service) UploadAudio(ctx context.Context, filePath string, params map[string]string) (*Upload, error) {
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

	reqURL, err := s.client.BaseURL().Parse("/api/uploads/audio")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), pr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	result := new(Upload)
	_, err = s.client.Do(req, result)
	return result, err
}

// ListMoodlePackages lists uploaded Moodle packages.
func (s *Service) ListMoodlePackages(ctx context.Context, opts *model.ListOptions, conditions string) (*MoodlePackagesResponse, error) {
	params := map[string]string{}
	if conditions != "" {
		params["conditions"] = conditions
	}
	u := addListOptions("/api/uploads/moodle-pkg", opts)
	u = addQueryParams(u, params)
	result := new(MoodlePackagesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- H5 Courseware ---

// GetH5Courseware returns H5 courseware information.
func (s *Service) GetH5Courseware(ctx context.Context, coursewareID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/h5-courseware/%d", coursewareID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UploadH5Courseware uploads H5 courseware.
func (s *Service) UploadH5Courseware(ctx context.Context, uploadID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/h5-courseware/upload/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- SCORM ---

// UploadSCORM uploads SCORM content.
func (s *Service) UploadSCORM(ctx context.Context, uploadID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/uploads/scorm/%d", uploadID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// --- Online Videos ---

// GetOnlineVideo returns online video information.
func (s *Service) GetOnlineVideo(ctx context.Context, videoID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/online-videos/%d", videoID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Media Captions ---

// GetMediaCaptionTemplate returns a media caption template.
func (s *Service) GetMediaCaptionTemplate(ctx context.Context, lang string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/media-captions/template?lang=%s", lang)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetMediaCaptionProgress returns media caption task progress.
func (s *Service) GetMediaCaptionProgress(ctx context.Context, mediaIDs string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/media/media-caption-tasks/progress?media_ids=%s", mediaIDs)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Duplicate Detection ---

// CheckDuplicate checks a file for duplicate content.
func (s *Service) CheckDuplicate(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/duplicate-detect/file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DownloadDuplicateReport downloads a duplicate detection report.
func (s *Service) DownloadDuplicateReport(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/duplicate-detect/report/download", &result)
	return result, err
}

// GetFileStatus returns the processing status of an upload file.
func (s *Service) GetFileStatus(ctx context.Context, uploadID int) (string, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	var result struct {
		Status string `json:"status"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.Status, err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
