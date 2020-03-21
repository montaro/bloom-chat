package protocol

type Envelop interface {
	String() string
}

type MessageKind string

const (
	MessageText          MessageKind = "MSG_TXT"
	MessagePhoto         MessageKind = "MSG_PHOTO"
	MessagePhotoAnimated MessageKind = "MSG_PHOTO_ANIM"
)

type ClientDevice string

const (
	ClientWeb    ClientDevice = "WEB"
	ClientMobile ClientDevice = "Mobile"
)

type ReplyTo struct {
	Id int64 `json:"id"`
}

type Permission string

const (
	PermissionCanDelete Permission = "CAN_DELETE"
)

type Sender struct {
	Id     int64        `json:"id"`
	Name   string       `json:"name"`
	Status string       `json:"status"`
	Client ClientDevice `json:"client"`
}

type ImageKind string

type ImageSize struct {
	URl    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Message struct {
	Id               int64         `json:"id"`
	Kind             MessageKind   `json:"kind"`
	Content          string        `json:"content"`
	FormattedContent string        `json:"formatted_content"`
	Timestamp        int64         `json:"timestamp"`
	Status           string        `json:"status"`
	Sender           Sender        `json:"sender"`
	ReplyTo          ReplyTo       `json:"reply_to"`
	Permissions      []*Permission `json:"permissions"`
	Sizes            []*ImageSize  `json:"sizes"`
}
