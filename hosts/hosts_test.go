package hosts

import "testing"

var testAdresses=[]*Address{
	&Address{Name: "Шлюз",IP: "192.168.51.1"},
	&Address{Name: "Свой",IP: "192.168.51.48"},
	&Address{Name: "Боровлев",IP: "192.168.51.48"},
}
func TestRead(t *testing.T) {
	err := ReadAdresses("testaddr.conf")
	if err != nil {
		t.Errorf("Ошибка чтения файла конфигурации %v", err)
	}
	if len(testAdresses)!=len(Adresses){
		t.Errorf("Записано %d адресов, прочиталось %d адресов",len(testAdresses),len(Adresses))
	}
	for i:=range testAdresses{
		if testAdresses[i].Name!=Adresses[i].Name || testAdresses[i].IP!=Adresses[i].IP{
			t.Errorf("Ошибка конфигурации: необходимо имя %s получено имя %s  ",testAdresses[i].Name,Adresses[i].Name)
		}
		if  testAdresses[i].IP!=Adresses[i].IP{
			t.Errorf("Ошибка конфигурации:  необходимо ip %s получено ip %s ",testAdresses[i].IP,Adresses[i].IP)
		}

	}
}
