package hosts

import "time"

var Hosts []*Host

const Debug = true

const (
	Computer = iota + 1
	Controller
	Server
	Unknown
)

type Host struct {
	Name     string
	IP       string
	Alive    bool
	HostType HostType
	AlarmOn  bool
	LastScan time.Time
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

func NewAddress(name string, IP string, hostType string, alarmOn bool) *Host {
	return &Host{Name: name, IP: IP, HostType: NewHost(hostType), AlarmOn: alarmOn}
}

func (addr *Host) IsAlive() bool {
	return addr.Alive
}
