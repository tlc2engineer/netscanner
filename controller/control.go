package controller

import (
	"netscan/hosts"
	"netscan/lowlevel"
	"time"
)

var Alarm bool
var Run = true

func mainCycle() {
	for Run {
		alarm := false
		for _, host := range hosts.Hosts {
			time.Sleep(time.Second)
			_, _, err := lowlevel.Ping(host.IP)
			host.LastScan = time.Now()
			if err == nil {
				host.Alive = true
				continue
			}
			if !host.AlarmOn || (host.AlarmOn && !host.Alive) {
				if host.AlarmOn {
					alarm = true
				}
				host.Alive = false
				continue
			}
			var i int
			for i = 0; i < 3; i++ {
				_, _, err := lowlevel.Ping(host.IP)
				if err == nil {
					break
				}
			}
			if i == 3 {
				host.Alive = false
				alarm = true
			}

		}
		Alarm = alarm

	}
}

func Process() {

	hosts.JsonReadConf()

	mainCycle()
}
