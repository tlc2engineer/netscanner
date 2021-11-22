package hosts

import (
	"encoding/csv"
	"fmt"
	"os"
)

const AddrFilename = "addr.conf"

func ReadAdresses(fname string) error {
	Adresses = make([]*Address, 0)
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, record := range records {
		addr := NewAddress(record[0], record[1], 4, record[2], record[3] == "true")
		Adresses = append(Adresses, addr)
	}
	return nil

}

func RemoveAddress(name, ip string) error {
	file, err := os.Open(AddrFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	num := -1
	for i, record := range records {
		if record[0] == name && record[1] == ip {
			num = i
			break
		}
	}
	if num == -1 {
		return fmt.Errorf("Запись не найдена")
	}
	records = append(records[:num], records[num+1:]...)
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return nil

}

func AddAddress(name, ip, hosttype string, on bool) error {
	file, err := os.Open(AddrFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	strOn := "false"
	if on {
		strOn = "true"
	}
	record := []string{name, ip, hosttype, strOn}
	records = append(records, record)
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}
