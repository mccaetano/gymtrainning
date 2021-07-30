package models

import (
	"encoding/json"
	"log"
)

type GymProfile struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Address    string `json:"address,omitempty"`
	Phone      string `json:"phone,omitempty"`
	CelPhone   string `json:"celPhone,omitempty"`
	IsWhatsapp bool   `json:"isWhatsApp,omitempty"`
}

func (c *GymProfile) Create() error {
	log.Println("(GymProfile - Create) Init")
	log.Println("(GymProfile - Create) body:", c)

	stmt, err := DB.Prepare("insert into gym.tbGymProfile (name, address, phone, celphone, iswhatsapp) values ($1, $2, $4, $5) RETURNING ID")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Name, c.Address, c.Phone, c.CelPhone, c.IsWhatsapp).Scan(&c.Id)
	if err != nil {
		log.Println("Error: Create error", err.Error())
		return err
	}
	log.Println("(GymProfile - Create) id:", c.Id)

	log.Println("(GymProfile - Create) Finish")
	return nil
}

func (c *GymProfile) Update(id int) error {
	log.Println("(GymProfile - Update) Init")
	log.Println("(GymProfile - Update) body:", c)

	stmt, err := DB.Prepare("update from gym.tbGymProfile set name = $2, address = $3, phone = $4, celphone = $5, iswhatsapp = $6 where id = $1")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Id, c.Name, c.Address, c.Phone, c.CelPhone, c.IsWhatsapp)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	log.Println("(GymProfile - Update) rows affected:", rows)

	log.Println("(GymProfile - Update) Finish")
	return nil
}

func (c *GymProfile) GetById(id int) error {
	log.Println("(GymProfile - GetById) Init")
	log.Println("(GymProfile - GetById) body:", c)

	stmt, err := DB.Prepare("select id, name, address, phone, celphone, iswhatsapp from gym.tbGymProfile where id $1")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Name, &c.Address, &c.Phone, &c.CelPhone, &c.IsWhatsapp)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(GymProfile - Update) body:", string(jsonPrint))

	log.Println("(GymProfile - GetById) Finish")
	return nil
}

func (c *GymProfile) GetAll(name string) ([]GymProfile, error) {
	log.Println("(GymProfile - GetAll) Init")
	log.Println("(GymProfile - GetAll) body:", c)

	stmt, err := DB.Prepare("select id, name, address, phone, celphone, iswhatsapp  from gym.tbGymProfile where lower(name) like '%' || lower($1) || '%'")
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(c.Name)
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer rows.Close()
	gymProfiles := []GymProfile{}
	for rows.Next() {
		gymProfile := GymProfile{}
		rows.Scan(&gymProfile.Id, &gymProfile.Name, &gymProfile.Address, &gymProfile.Phone, &gymProfile.CelPhone, &gymProfile.IsWhatsapp)
		gymProfiles = append(gymProfiles, gymProfile)
	}
	log.Println("(GymProfile - GetAll) rows:", len(gymProfiles))
	jsonPrint, _ := json.Marshal(gymProfiles)
	log.Println("(GymProfile - GetAll) body:", string(jsonPrint))

	log.Println("(GymProfile - GetAll) Finish")
	return gymProfiles, nil
}
