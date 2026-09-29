package adapter

import (
	"bridge/adapter/legacy"
	"errors"
	"strconv"
)

var (
	ErrChannelUnavailiable = errors.New("channel not availiable")
	ErrInvalidRecipient = errors.New("recipient invalid")
	ErrDeliveryFailure = errors.New("delivery failed")
)
type PagerAdapter struct {
	pager *legacy.Pager
}

func NewPagerAdapter(pager *legacy.Pager) *PagerAdapter {
	return &PagerAdapter{pager: pager}
}

func (p *PagerAdapter) Deliver(to, message string) error {
	if p == nil || p.pager == nil {
		return ErrChannelUnavailiable
	}
	id, err := strconv.Atoi(to)
	if err != nil {
		return ErrInvalidRecipient
	}

	status := p.pager.Page(id, []byte(message))

	switch status {
	case legacy.StatusOK:
		return nil
	case legacy.StatusBadID:
		return ErrInvalidRecipient
	case legacy.StatusOffline:
		return ErrChannelUnavailiable
	default:
		return ErrDeliveryFailure
	}
}