package bridger

type DeliveryChannel interface {
	Deliver(to, message string) error
}