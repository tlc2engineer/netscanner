package hosts

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

var Adresses []*Address

const Debug = false

type Address struct {
	Name      string
	IP        string
	Alive     bool
	HostType  string
	AlarmOn   bool
	countPing int
	pinger    *ping.Pinger
	error     bool
	errString string
	LastScan  time.Time
}

func NewAddress(name string, IP string, countPing int, hostType string, alarmOn bool) *Address {
	return &Address{Name: name, IP: IP, countPing: countPing, HostType: hostType, AlarmOn: alarmOn}
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
		addr.error = true
		return err
	}
	err = addr.pinger.Run() // Blocks until finished.
	if err != nil {
		addr.Alive = false
		return err
	}
	stats := addr.pinger.Statistics() // get send/receive/duplicate/rtt stat
	if Debug {

		log.Println(addr.Name, addr.IP, stats.PacketLoss)
	}
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
	return addr.error
}
