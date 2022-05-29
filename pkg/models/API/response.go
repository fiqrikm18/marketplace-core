package API

type M map[string]interface{}

type ResponseMeta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}
