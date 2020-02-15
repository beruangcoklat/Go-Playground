package entity

type RunCodeRequest struct {
	Code string `json:"code"`
}

type RunCodeResponse struct {
	Data string `json:"data"`
}
