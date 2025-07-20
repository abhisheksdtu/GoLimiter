package api

// CheckLimitRequest is the request body for the /check-limit endpoint.
type CheckLimitRequest struct {
	Key string `json:"key" binding:"required" example:"my-api-key"`
}

// CheckLimitResponse is the response body for the /check-limit endpoint.
type CheckLimitResponse struct {
	Allowed bool `json:"allowed" example:"true"`
}
