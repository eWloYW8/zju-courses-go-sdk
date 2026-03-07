package exams

import "github.com/eWloYW8/zju-courses-go-sdk/model"

type SubjectLib struct {
	ID            int           `json:"id"`
	Title         string        `json:"title,omitempty"`
	ParentID      int           `json:"parent_id,omitempty"`
	IsFolder      bool          `json:"is_folder,omitempty"`
	Children      []*SubjectLib `json:"children,omitempty"`
	SubjectsCount int           `json:"subjects_count,omitempty"`
}

type SubjectLibsResponse struct {
	SubjectLibs []*SubjectLib `json:"subject_libs"`
}

type RubricsResponse struct {
	Rubrics []*model.Rubric `json:"rubrics"`
}

type CoursewareQuizSettings struct {
	QuizCountLimit int `json:"quiz_count_limit,omitempty"`
}

type CoursewareQuizSettingsResponse struct {
	Setting *CoursewareQuizSettings `json:"setting"`
}

type CoursewareQuizSubjectsResponse struct {
	QuizID   int                  `json:"quiz_id,omitempty"`
	Subjects []*model.ExamSubject `json:"subjects,omitempty"`
}

type CoursewareQuizUpdateResponse struct {
	QuizID int `json:"quiz_id"`
}
