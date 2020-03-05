package protocol

//TODO
type ResponseOperation string

const (
	RESPONSE_MESSAGE ResponseOperation = "RES_MSG"
	ERROR            ResponseOperation = "ERROR"
)

type Response struct {
	Op   ResponseOperation `json:"op"`
	Data map[string]string `json:"data"`
}
