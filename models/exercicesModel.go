package models

import (
	"encoding/json"

	"log"
)

type Exercises struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Exercises) Create() error {
	log.Println("(Exercises - Create) Init")
	log.Println("(Exercises - Create) body:", c)

	stmt, err := DB.Prepare("insert into gym.tbexercises (name) values ($1, $2, $3) RETURNING ID")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Name).Scan(&c.Id)
	if err != nil {
		log.Println("Error: Create error", err.Error())
		return err
	}
	log.Println("(Exercises - Create) id:", c.Id)

	log.Println("(Exercises - Create) Finish")
	return nil
}

func (c *Exercises) Update(id int) error {
	log.Println("(Exercises - Update) Init")
	log.Println("(Exercises - Update) body:", c)

	stmt, err := DB.Prepare("update from gym.tbexercises set name = $2 where id = $1")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Id, c.Name)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	log.Println("(Exercises - Update) rows affected:", rows)

	log.Println("(Exercises - Update) Finish")
	return nil
}

func (c *Exercises) GetById(id int) error {
	log.Println("(Exercises - GetById) Init")
	log.Println("(Exercises - GetById) body:", c)

	stmt, err := DB.Prepare("select id, name from gym.tbexercises where id $1")
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Name)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(Exercises - Update) body:", string(jsonPrint))

	log.Println("(Exercises - GetById) Finish")
	return nil
}

func (c *Exercises) GetAll(name string) ([]Exercises, error) {
	log.Println("(Exercises - GetAll) Init")
	log.Println("(Exercises - GetAll) body:", c)

	stmt, err := DB.Prepare("select id, name from gym.tbexercises where lower(name) like '%' || lower($1) || '%'")
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
	exercises := []Exercises{}
	for rows.Next() {
		exercise := Exercises{}
		rows.Scan(&exercise.Id, &exercise.Name)
		exercises = append(exercises, exercise)
	}
	log.Println("(Exercises - GetAll) rows:", len(exercises))
	jsonPrint, _ := json.Marshal(exercises)
	log.Println("(Exercises - GetAll) body:", string(jsonPrint))

	log.Println("(Exercises - GetAll) Finish")
	return exercises, nil
}
