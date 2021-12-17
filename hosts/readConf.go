package hosts

import (
	"encoding/json"
	"fmt"
	"os"
)

const AddrFilename = "conf.json"

func JsonReadConf() error {
	bts, err := os.ReadFile(AddrFilename)
	if err != nil {
		return err
	}
	Hosts = make([]*Host, 0)
	err = json.Unmarshal(bts, &Hosts)
	if err != nil {
		return err
	}
	return nil
}

func RemoveHost(ip string) error {
	for i, host := range Hosts {
		if host.IP == ip {
			Hosts = append(Hosts[:i], Hosts[i+1:]...)
			err := JsonWriteConf()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("нет такого IP")

}

func AddHost(host *Host) error {
	Hosts = append(Hosts, host)
	err := JsonWriteConf()
	if err != nil {
		return err
	}
	return nil
}

func UpdateHost(host *Host) error {
	ip := host.IP
	for i, h := range Hosts {
		if h.IP == ip {
			Hosts[i] = host
			err := JsonWriteConf()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("нет такого IP")
}

func JsonWriteConf() error {
	bts, err := json.Marshal(Hosts)
	if err != nil {
		return err
	}
	err = os.WriteFile("conf.json", bts, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
