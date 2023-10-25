package models

type JSONResponse struct {
	Error bool        `json:"error"`
	Msg   string      `json:"message"`
	Data  interface{} `json:"data,omitempty"`
}
