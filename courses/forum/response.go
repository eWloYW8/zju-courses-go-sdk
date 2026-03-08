package forum

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type TopicsResponse struct {
	Topics []*Topic `json:"topics"`
	model.Pagination
}

type LatestTopicsResponse struct {
	Topics []*Topic `json:"topics"`
}

type ForumScoreResponse struct {
	ForumScore *ForumScore `json:"forum_score"`
}

type TopicCategoriesResponse struct {
	TopicCategories []*TopicCategory `json:"topic_categories"`
}
