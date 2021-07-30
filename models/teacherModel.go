package models

import (
	"encoding/json"
	"log"
)

type Teacher struct {
	Id   int64
	Name string
}

func (c *Teacher) Create() error {
	log.Println("(Teacher - Create) Init")
	log.Println("(Teacher - Create) body:", c)

	stmt, err := DB.Prepare(`insert into gym.tbteacher 
		(name) 
		values ($1) RETURNING ID`)
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
	log.Println("(Teacher - Create) id:", c.Id)

	log.Println("(Teacher - Create) Finish")
	return nil
}

func (c *Teacher) Update(id int) error {
	log.Println("(Teacher - Update) Init")
	log.Println("(Teacher - Update) body:", c)

	stmt, err := DB.Prepare(`update from gym.tbteacher set 
		name = $2
		where id = $1`)
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
	log.Println("(Teacher - Update) rows affected:", rows)

	log.Println("(Teacher - Update) Finish")
	return nil
}

func (c *Teacher) GetById(id int) error {
	log.Println("(Teacher - GetById) Init")
	log.Println("(Teacher - GetById) body:", c)

	stmt, err := DB.Prepare(`select t.id, t.name
		from gym.tbteacher where t.id $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Id, &c.Name)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(Teacher - Update) body:", string(jsonPrint))

	log.Println("(Teacher - GetById) Finish")
	return nil
}

func (c *Teacher) GetAll(name string) ([]Teacher, error) {
	log.Println("(Teacher - GetAll) Init")
	log.Println("(Teacher - GetAll) body:", c)

	stmt, err := DB.Prepare(`select t.id, t.name
		from gym.tbteacher t where t.name like '%' || $1 || '%' `)
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer rows.Close()
	trainnings := []Teacher{}
	for rows.Next() {
		trainning := Teacher{}
		rows.Scan(&trainning.Id, &trainning.Name)
		trainnings = append(trainnings, trainning)
	}
	log.Println("(Teacher - GetAll) rows:", len(trainnings))
	jsonPrint, _ := json.Marshal(trainnings)
	log.Println("(Teacher - GetAll) body:", string(jsonPrint))

	log.Println("(Teacher - GetAll) Finish")
	return trainnings, nil
}
