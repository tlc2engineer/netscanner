package hosts

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

var Adresses []*Address

const Debug = false

const (
	Computer = iota + 1
	Controller
	Server
	Unknown
)

type Address struct {
	Name         string
	IP           string
	Alive        bool
	HostType     HostType
	AlarmOn      bool
	countPing    int
	pinger       *ping.Pinger
	hostError    bool
	errString    string
	LastScan     time.Time
	PingError    bool
	PingErrorStr string
	PacketLoss   float64
}

type HostType int

func (ht HostType) String() string {
	switch ht {
	case 1:
		return "Компьютер"
	case 2:
		return "Сервер"
	case 3:
		return "Контроллер"
	}
	return "Неизвестный тип"
}

func NewHost(host string) HostType {
	switch host {
	case "Компьютер":
		return Computer
	case "Контроллер":
		return Controller
	case "Сервер":
		return Server
	}
	return Unknown
}

func NewAddress(name string, IP string, countPing int, hostType string, alarmOn bool) *Address {
	return &Address{Name: name, IP: IP, countPing: countPing, HostType: NewHost(hostType), AlarmOn: alarmOn}
}

func (addr *Address) getPinger() error {
	pinger, err := ping.NewPinger(addr.IP)
	if err != nil {
		return err
	}

	pinger.Count = addr.countPing
	pinger.SetPrivileged(true)
	addr.pinger = pinger
	return nil
}
func (addr *Address) IsAlive() bool {
	return addr.Alive
}

func (addr *Address) Check() error {
	err := addr.getPinger()
	defer func() {
		addr.LastScan = time.Now()
	}()

	if err != nil {
		addr.errString = err.Error()
		addr.hostError = true
		return err
	}
	err = addr.pinger.Run() // Blocks until finished.
	if err != nil {
		addr.Alive = false
		addr.PingError = true
		addr.PingErrorStr = err.Error()
		return err
	}
	addr.PingError = false
	addr.PingErrorStr = ""
	stats := addr.pinger.Statistics() // get send/receive/duplicate/rtt stat
	if Debug {

		log.Println(addr.Name, addr.IP, stats.PacketLoss)
	}
	addr.PacketLoss = stats.PacketLoss
	if stats.PacketLoss == 100.0 {
		addr.Alive = false
		return nil
	}
	addr.Alive = true
	return nil
}

func (addr *Address) Error() string {
	return addr.errString
}

func (addr *Address) IsError() bool {
	return addr.hostError
}
