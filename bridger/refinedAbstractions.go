package bridger

import "errors"

type RegularAlert struct {
    Alert
    To      string
    Message string
}

func NewRegularAlert(channel DeliveryChannel, to, message string) RegularAlert {
    return RegularAlert{
        Alert:   Alert{channel: channel},
        To:      to,
        Message: message,
    }
}

func (r RegularAlert) Send() error {
    return r.deliver(r.To, r.Message)
}

type CriticalAlert struct {
    Alert
    Primary string
    Backup  string
    Message string
}

func NewCriticalAlert(channel DeliveryChannel, primary, backup, message string) CriticalAlert {
    return CriticalAlert{
        Alert:   Alert{channel: channel},
        Primary: primary,
        Backup:  backup,
        Message: message,
    }
}

func (c CriticalAlert) Send() error {
    urgent := "CRITICAL: " + c.Message
    primaryErr := c.deliver(c.Primary, urgent)
    backupErr := c.deliver(c.Backup, urgent)
    return errors.Join(primaryErr, backupErr)
}