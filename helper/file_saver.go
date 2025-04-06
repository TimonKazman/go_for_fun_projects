package helper

import (
	"fmt"
	"os"
)

type Documents struct {
	Customer string
	Value    []byte
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Method of the document struct
func (p *Documents) Save() error {
	filename := p.Customer + "_balance.txt"
	folder := "balance_paper"
	if pathExists(folder) {
		path := folder + "/" + filename
		return os.WriteFile(path, p.Value, 0644)
	}
	return nil
}

func SaveFile(customer *string, value []byte) {

	docs := &Documents{
		Customer: *customer,
		Value:    value,
	}

	err := docs.Save()
	fmt.Println("file saved")

	if err != nil {
		panic(fmt.Sprintf("saving the file not possible: %v", err))
	}

}
