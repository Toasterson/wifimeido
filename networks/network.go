package networks

import (
	"encoding/json"
)

type Network interface {
	SSID() string
}

type RegularNetwork struct {
	ssid string
	psk  string
}

type regularNetworkJSON struct {
	SSID string `json:"ssid"`
	PSK  string `json:"psk"`
}

func NewRegularNetwork(ssid string, psk string) RegularNetwork {
	return RegularNetwork{ssid, psk}
}

func (n RegularNetwork) SSID() string {
	return n.ssid
}

func (n RegularNetwork) SetSSID(ssid string) {
	n.ssid = ssid
}

func (n RegularNetwork) PSK() string {
	return n.psk
}

func (n RegularNetwork) SetPSK(psk string) {
	n.psk = psk
}

func (n *RegularNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(regularNetworkJSON{
		n.ssid,
		n.psk,
	})
}

func (n *RegularNetwork) UnmarshalJSON(b []byte) error {
	temp := &regularNetworkJSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	n.ssid = temp.SSID
	n.psk = temp.PSK

	return nil
}

type EapNetwork struct {
	ssid     string
	username string
	password string
}

type eapNetworkJSON struct {
	SSID     string `json:"ssid"`
	Username string `json:"user"`
	Password string `json:"password"`
}

func (n EapNetwork) SSID() string {
	return n.ssid
}

func (n EapNetwork) Username() string {
	return n.username
}

func (n EapNetwork) Password() string {
	return n.password
}

func (n *EapNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(eapNetworkJSON{
		n.ssid,
		n.username,
		n.password,
	})
}

func (n *EapNetwork) UnmarshalJSON(b []byte) error {
	temp := &eapNetworkJSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	n.ssid = temp.SSID
	n.username = temp.Username
	n.password = temp.Password

	return nil
}
