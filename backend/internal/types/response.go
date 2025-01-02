package types

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`   // omitempty to omit the field if it's nil
	Errors  map[string]string      `json:"errors,omitempty"` // omitempty to omit the field if it's nil
	Success bool                   `json:"success"`
}
