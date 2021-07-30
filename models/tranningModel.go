package models

import (
	"encoding/json"

	"log"
)

type Trainning struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IsSunday      bool   `json:"isSunday,omitempty"`
	IsMonday      bool   `json:"isMonday,omitempty"`
	IsTuesday     bool   `json:"isTuesday,omitempty"`
	IsWednesday   bool   `json:"isWednesday,omitempty"`
	IsThursday    bool   `json:"isThursday,omitempty"`
	IsFriday      bool   `json:"isFriday,omitempty"`
	IsSaturday    bool   `json:"isSaturday,omitempty"`
	LastTrainning string `json:"lastTrainning,omitempty"`
	UserId        int    `json:"UserId,omitempty"`
	GymId         int    `json:"gymId,omitempty"`
}

func (c *Trainning) Create() error {
	log.Println("(Trainning - Create) Init")
	log.Println("(Trainning - Create) body:", c)

	stmt, err := DB.Prepare(`insert into gym.tbtrainning 
		(name, issunday, ismonday, istuesday, iswednesday, isthursday,
			isfriday, issaturday, lasttrainning, userid,
			gymid) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING ID`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Name, c.IsSunday, c.IsMonday, c.IsTuesday, c.IsWednesday,
		c.IsThursday, c.IsFriday, c.IsSaturday, c.LastTrainning, c.UserId, c.GymId).Scan(&c.Id)
	if err != nil {
		log.Println("Error: Create error", err.Error())
		return err
	}
	log.Println("(Trainning - Create) id:", c.Id)

	log.Println("(Trainning - Create) Finish")
	return nil
}

func (c *Trainning) Update(id int) error {
	log.Println("(Trainning - Update) Init")
	log.Println("(Trainning - Update) body:", c)

	stmt, err := DB.Prepare(`update from gym.tbtrainning set 
		name = $2, issunday = $3, ismonday = $4, istuesday = $5, 
		iswednesday = $6, isthursday = $7, isfriday = $8, 
		issaturday = $9, lasttrainning = $10, userid = $11, gymid = $12 
		where id = $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Id, c.Name, c.IsSunday, c.IsMonday, c.IsTuesday, c.IsWednesday,
		c.IsThursday, c.IsFriday, c.IsSaturday, c.LastTrainning, c.UserId, c.GymId)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	log.Println("(Trainning - Update) rows affected:", rows)

	log.Println("(Trainning - Update) Finish")
	return nil
}

func (c *Trainning) GetById(id int) error {
	log.Println("(Trainning - GetById) Init")
	log.Println("(Trainning - GetById) body:", c)

	stmt, err := DB.Prepare(`select id, name, issunday, ismonday, istuesday, 
		iswednesday, isthursday, isfriday, issaturday, lasttrainning, userid,
		gymid from gym.tbtrainning where id $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Name, &c.IsSunday, &c.IsMonday, &c.IsTuesday, &c.IsWednesday,
		&c.IsThursday, &c.IsFriday, &c.IsSaturday, &c.LastTrainning, &c.UserId, &c.GymId)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(Trainning - Update) body:", string(jsonPrint))

	log.Println("(Trainning - GetById) Finish")
	return nil
}

func (c *Trainning) GetAll(name string) ([]Trainning, error) {
	log.Println("(Trainning - GetAll) Init")
	log.Println("(Trainning - GetAll) body:", c)

	stmt, err := DB.Prepare(`select id, name, issunday, ismonday, istuesday, 
		iswednesday, isthursday, isfriday, issaturday, lasttrainning, userid,
		gymid from gym.tbtrainning where lower(name) like '%' || lower($1) || '%'`)
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
	trainning := []Trainning{}
	for rows.Next() {
		exercise := Trainning{}
		rows.Scan(&exercise.Id, &exercise.Name, &exercise.IsSunday, &exercise.IsMonday,
			&exercise.IsTuesday, &exercise.IsWednesday, &exercise.IsThursday,
			&exercise.IsFriday, &exercise.IsSaturday, &exercise.LastTrainning,
			&exercise.UserId, &exercise.GymId)
		trainning = append(trainning, exercise)
	}
	log.Println("(Trainning - GetAll) rows:", len(trainning))
	jsonPrint, _ := json.Marshal(trainning)
	log.Println("(Trainning - GetAll) body:", string(jsonPrint))

	log.Println("(Trainning - GetAll) Finish")
	return trainning, nil
}
