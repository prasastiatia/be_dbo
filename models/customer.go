package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model

	Name         string `json:"name"`
	TanggalLahir string `json:"tanggal_lahir"`
	Mobile       string `json:"mobile" `
	Alamat       string `json:"alamat"`
}

type ShowCustomer struct {
	Name         string
	TanggalLahir string
	Mobile       string
	Alamat       string
}
