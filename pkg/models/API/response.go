package API

type M map[string]interface{}

type ResponseMeta struct {
	Code    int
	Status  string
	Message string
}

type Response struct {
	Meta ResponseMeta
	Data interface{}
}
