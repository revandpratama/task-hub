package util

import "github.com/revandpratama/task-hub/dto"

func NewResponse(rp dto.ResponseParam) any {
	var response any
	var status string

	if rp.StatusCode >= 200 && rp.StatusCode < 400 {
		status = "success"
	} else {
		status = "failed"
	}

	if rp.Data != nil {
		response = &dto.ResponseWithData{
			Status:  status,
			Message: rp.Message,
			Data:    rp.Data,
		}
	} else {
		response = &dto.ResponseWithoutData{
			Status:  status,
			Message: rp.Message,
		}
	}

	return response
}
