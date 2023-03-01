package global

import "net/http"

var (
	GLO_REQ_CH  chan []string
	GLO_RESP_CH chan http.Response
)
