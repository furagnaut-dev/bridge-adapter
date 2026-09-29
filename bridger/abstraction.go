package bridger

type Alert struct {
	channel DeliveryChannel
}

func (a Alert) deliver(to, message string) error {
	return a.channel.Deliver(to, message)
}
