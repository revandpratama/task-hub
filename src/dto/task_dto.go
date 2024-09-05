package dto

type TaskRequest struct {
	UserID      int    `json:"user_id"`
	ProjectID   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

type TaskResponse struct {
	ID             int               `json:"id"`
	UserID         int               `json:"user_id"`
	ProjectID      int               `json:"project_id"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	Status         string            `json:"status"`
	Priority       string            `json:"priority"`
	TaskAttachment []TaskAttResponse `json:"task_attachment"`
}
