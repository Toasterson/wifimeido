package networks

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Storage interface {
	SaveNetwork(net Network)
}

type FileStorage struct {
	path string
}

func NewFileStorage(path string) FileStorage {
	return FileStorage{path}
}

func (f FileStorage) SaveNetwork(net RegularNetwork) error {
	data, err := json.Marshal(net)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.path+net.SSID(), data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (f FileStorage) OpenNetwork(ssid string, n Network) error {
	file, err := os.Open(f.path + ssid)

	if err != nil {
		return nil
	}

	data, _ := ioutil.ReadAll(file)

	file.Close()

	json.Unmarshal(data, n)

	return nil
}
