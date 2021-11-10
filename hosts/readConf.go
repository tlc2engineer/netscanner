package hosts

import (
	"encoding/csv"
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
