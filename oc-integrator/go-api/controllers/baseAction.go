package controllers

type Resp struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
}
