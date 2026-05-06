package funding_source

import "time"

type FundingSourceResponse struct {
	ID            string    `json:"id"`
	FundingSource string    `json:"funding_source"`
	ExternalCode  string    `json:"external_code"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
