package protocol

type Request struct {
	Op   string      `json:"op"`
	Data interface{} `json:"data"`
}

type Response struct {
	Op   string      `json:"op"`
	Data interface{} `json:"data"`
}
