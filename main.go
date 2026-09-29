package main

import (
	"bridge/adapter"
	"bridge/adapter/legacy"
	"bridge/bridger"
	"fmt"
	"log"
	"os"
)

func main() {
	registry := bridger.NewChannelRegistry()
	registry.Register("email", &bridger.EmailChannel{})
	registry.Register("sms", &bridger.SMSChannel{})
	registry.Register("pager", adapter.NewPagerAdapter(&legacy.Pager{Online: true}))

	channel, err := registry.Selection(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "regular":
		if len(os.Args) != 5 {
			os.Exit(2)
		}
		
		regular := bridger.NewRegularAlert(channel, os.Args[3], os.Args[4])
		if err := regular.Send(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Regular Alert System: {To: %q, Message: %q, Channel: %T}",
		regular.CollectionsTo, regular.Message, channel)
	case "critical":
		if len(os.Args) != 6 {
			os.Exit(2)
		}

		critical := bridger.NewCriticalAlert(channel, os.Args[3], os.Args[4], os.Args[5])
		if err := critical.Send(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Critical Alert System: {To Maintenance: %q, To Facilities: %q, Message: %q, Channel: %T}",
		critical.CollectionsTo, critical.FacilitiesTo, critical.Message, channel)
	default:
		log.Fatal("alert type is either regular or critical")
	}

}