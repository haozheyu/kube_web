package initial

import (
	"kube_web/bus/newbus"
)

func InitBus() {
	var err error
	// bus.DefaultBus, err = bus.NewBus(beego.AppConfig.String("BusRabbitMQURL"))
	if newbus.UBus == nil {
		newbus.UBus, err = newbus.NewBus()
		if err != nil {
			panic(err)
		}
	}
}
