package funding_source

import "backend-app/pkg/pagination"

type FundingSourceFilterRequest struct {
	pagination.BaseRequest
}

type CreateFundingSourceRequest struct {
	FundingSource string `json:"funding_source" binding:"required,max=100"`
	ExternalCode  string `json:"external_code" binding:"omitempty,max=20"`
}

type UpdateFundingSourceRequest struct {
	FundingSource string `json:"funding_source" binding:"required,max=100"`
	ExternalCode  string `json:"external_code" binding:"omitempty,max=20"`
	IsActive      *bool  `json:"is_active"`
}
