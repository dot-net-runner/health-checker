package notification

type (
	Nitifier interface {
		Notify(m Message) error
	}

	Message struct {
		Title string
		Body  string
	}
)

func NewMessage(title string, msg string) *Message {
	return &Message{
		Title: title,
		Body:  msg,
	}
}
