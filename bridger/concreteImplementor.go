package bridger

type Email struct {
	To string
	Message string
}

type EmailChannel struct {
	Message []Email
}

func (e *EmailChannel) Deliver(to, message string) error {
	e.Message = append(e.Message, Email{To: to, Message: message})
	return nil
}

type SMS struct {
	To string
	Message string
}

type SMSChannel struct {
	Message []SMS
}

func (sms *SMSChannel) Deliver(to, message string) error {
	sms.Message = append(sms.Message, SMS{To: to, Message: message})
	return nil
}
