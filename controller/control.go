package controller

import (
	"NETSCANNER/hosts"
	"fmt"
	"log"
	"time"
)

const timeout = 1000

var Run = true

func mainCycle() {
	for Run {
		for _, address := range hosts.Adresses {

			if address.IsError() {
				continue
			}
			err := address.Check()
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(time.Millisecond * timeout)
	}
}

func Process() {
	err := hosts.ReadAdresses(hosts.AddrFilename)
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
	}
	//hosts.InitAdress()

	fmt.Println("Запуск цикла")
	mainCycle()
}
