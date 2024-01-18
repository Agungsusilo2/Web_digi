package dto

type Heroes struct {
	Id           string `json:"Id"`
	NamaLengkap  string `json:"NamaLengkap"`
	Email        string `json:"Email"`
	NomorTelepon string `json:"NomorTelepon"`
	Umur         string `json:"Umur"`
	Pekerjaan    string `json:"Pekerjaan"`
}
