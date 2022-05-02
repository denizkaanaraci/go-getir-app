package model

type GetResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
