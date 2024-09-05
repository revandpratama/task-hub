package dto

type TaskAttResponse struct {
	TaskID   int    `json:"task_id"`
	UserID   int    `json:"user_id"`
	FilePath string `json:"filepath"`
	FileType string `json:"filetype"`
}
type TaskAttRequest struct {
	TaskID   int    `json:"task_id"`
	UserID   int    `json:"user_id"`
	FilePath string `json:"filepath"`
	FileType string `json:"filetype"`
}
