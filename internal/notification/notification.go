package notification

type Message struct {
	Method   []string // Push, Email
	Subject  string
	Message  string
	Priority int // priority -2 no alert, -1 quiet notification, 1 high-priority, 2 require confirmation
}

func (m *Message) Send() {
	SendMail(m)
}
