package api

type APIResponse struct {
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}
