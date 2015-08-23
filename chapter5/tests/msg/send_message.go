package msg

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error {
	// ...
	return nil
}

type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}
