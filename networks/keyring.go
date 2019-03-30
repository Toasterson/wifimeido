package networks

import (
	"encoding/json"
	"github.com/zalando/go-keyring"
	"os/user"
)

type KeyringStorage struct {
}

func (k *KeyringStorage) SaveNetwork(net Network) error {
	currentUser, err := user.Current()

	if err != nil {
		return err
	}

	netJson, err := json.Marshal(net)

	if err != nil {
		return err
	}

	keyring.Set("wifimeido", currentUser.Name, string(netJson))

	return nil
}
