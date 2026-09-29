package bridger

import "errors"

type RegularAlert struct {
    Alert
    CollectionsTo string
    Message       string
}

func NewRegularAlert(channel DeliveryChannel, to, message string) RegularAlert {
    return RegularAlert{
        Alert:   Alert{channel: channel},
        CollectionsTo:      to,
        Message: message,
    }
}

func (r RegularAlert) Send() error {
    return r.deliver(r.CollectionsTo, r.Message)
}

type CriticalAlert struct {
    Alert
    CollectionsTo string
    FacilitiesTo  string
    Message       string
}

func NewCriticalAlert(channel DeliveryChannel, primary, backup, message string) CriticalAlert {
    return CriticalAlert{
        Alert:   Alert{channel: channel},
        CollectionsTo: primary,
        FacilitiesTo:  backup,
        Message: message,
    }
}

func (c CriticalAlert) Send() error {
    message := "CRITICAL: " + c.Message

    collectionsErr := c.deliver(c.CollectionsTo, message)
    facilitiesErr := c.deliver(c.FacilitiesTo, message)

    return errors.Join(collectionsErr, facilitiesErr)
}