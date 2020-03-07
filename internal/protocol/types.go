package protocol

type RequestData interface {
	String() string
}

type RequestI interface {
	decode()
}

type ResponseData interface {
	decode()
}