package bridger

import (
	"errors"
)

var UnknowNChannel = errors.New("unknown chanel")
type ChannelRegistry struct {
	channels map[string]DeliveryChannel
}

func NewChannelRegistry() *ChannelRegistry {
	return &ChannelRegistry{channels: make(map[string]DeliveryChannel)}
}

func (r *ChannelRegistry) Register(name string, channel DeliveryChannel) {
	r.channels[name] = channel
}

func (r *ChannelRegistry) Selection(name string) (DeliveryChannel, error) {
	channel, statusOK := r.channels[name]
	if !statusOK {
		return nil, UnknowNChannel
	}

	return channel, nil
}