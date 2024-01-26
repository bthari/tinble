package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type JWTTokenResponse struct {
	Token string `json:"token"`
}
