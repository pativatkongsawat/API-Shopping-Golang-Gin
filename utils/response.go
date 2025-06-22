package utils

import "go_gin/models/category"

type ResponseMessage struct {
	Status  interface{} `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type CategoryResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Result  []category.Category `json:"result"`
}
