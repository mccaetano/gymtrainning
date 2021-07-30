package models

import (
	"encoding/json"
	"log"
)

type UserProfile struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	CelPhone string `json:"celPhone,omitempty"`
	Address  string `json:"address,omitempty"`
}

func (c *UserProfile) Create() error {
	log.Println("(UserProfile - Create) Init")
	log.Println("(UserProfile - Create) body:", c)

	stmt, err := DB.Prepare(`insert into gym.tbUserProfile (name,
		email,
		phone,
		celphone,
		address) values ($1, $2, $3, $4, $5) RETURNING ID`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Name, c.Email, c.Phone, c.CelPhone, c.Address).Scan(&c.Id)
	if err != nil {
		log.Println("Error: Create error", err.Error())
		return err
	}
	log.Println("(UserProfile - Create) id:", c.Id)

	log.Println("(UserProfile - Create) Finish")
	return nil
}

func (c *UserProfile) Update(id int) error {
	log.Println("(UserProfile - Update) Init")
	log.Println("(UserProfile - Update) body:", c)

	stmt, err := DB.Prepare(`update from gym.tbUserProfile set name = $2,
		email = $3,
		phone = $4,
		celphone = $5,
		address = $1 where id = $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Id, c.Name, c.Email, c.Phone, c.CelPhone, c.Address)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	log.Println("(UserProfile - Update) rows affected:", rows)

	log.Println("(UserProfile - Update) Finish")
	return nil
}

func (c *UserProfile) GetById(id int) error {
	log.Println("(UserProfile - GetById) Init")
	log.Println("(UserProfile - GetById) body:", c)

	stmt, err := DB.Prepare(`select id, name,
	email,
	phone,
	celphone,
	address from gym.tbUserProfile where id $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Name, &c.Email, &c.Phone, &c.CelPhone, &c.Address)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(UserProfile - Update) body:", string(jsonPrint))

	log.Println("(UserProfile - GetById) Finish")
	return nil
}

func (c *UserProfile) GetAll(name string) ([]UserProfile, error) {
	log.Println("(UserProfile - GetAll) Init")
	log.Println("(UserProfile - GetAll) body:", c)

	stmt, err := DB.Prepare(`select id, name,
	email,
	phone,
	celphone,
	address from gym.tbUserProfile where lower(name) like '%' || lower($1) || '%'`)
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
	userProfiles := []UserProfile{}
	for rows.Next() {
		userProfile := UserProfile{}
		rows.Scan(&userProfile.Id, &userProfile.Name, &userProfile.Email,
			&userProfile.Phone, &userProfile.CelPhone, &userProfile.Address)
		userProfiles = append(userProfiles, userProfile)
	}
	log.Println("(UserProfile - GetAll) rows:", len(userProfiles))
	jsonPrint, _ := json.Marshal(userProfiles)
	log.Println("(UserProfile - GetAll) body:", string(jsonPrint))

	log.Println("(UserProfile - GetAll) Finish")
	return userProfiles, nil
}
