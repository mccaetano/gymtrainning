package models

import (
	"encoding/json"
	"time"

	"log"
)

type TrainningStep struct {
	Id         int64     `json:"id,omitempty"`
	Exercise   Exercises `json:"exercise,omitempty"`
	Series     int64     `json:"series,omitempty"`
	Fat        int64     `json:"fat,omitempty"`
	Time       time.Time `json:"time,omitempty"`
	Tranningid int64     `json:"tranningid,omitempty"`
}

func (c *TrainningStep) Create() error {
	log.Println("(TrainningStep - Create) Init")
	log.Println("(TrainningStep - Create) body:", c)

	stmt, err := DB.Prepare(`insert into gym.tbtrainningstep 
		(exerciseid, series, fat, time, tranningid) 
		values ($1, $2, $3, $4, $5) RETURNING ID`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Exercise.Id, c.Series, c.Fat, c.Time, c.Tranningid).Scan(&c.Id)
	if err != nil {
		log.Println("Error: Create error", err.Error())
		return err
	}
	log.Println("(TrainningStep - Create) id:", c.Id)

	log.Println("(TrainningStep - Create) Finish")
	return nil
}

func (c *TrainningStep) Update(id int) error {
	log.Println("(TrainningStep - Update) Init")
	log.Println("(TrainningStep - Update) body:", c)

	stmt, err := DB.Prepare(`update from gym.tbtrainningstep set 
		exerciseid = $2, series = $3, fat = $4, time = $5, 
		tranningid  = $6
		where id = $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Id, c.Exercise.Id, c.Series, c.Fat, c.Time, c.Tranningid)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}

	rows, _ := res.RowsAffected()
	log.Println("(TrainningStep - Update) rows affected:", rows)

	log.Println("(TrainningStep - Update) Finish")
	return nil
}

func (c *TrainningStep) GetById(id int) error {
	log.Println("(TrainningStep - GetById) Init")
	log.Println("(TrainningStep - GetById) body:", c)

	stmt, err := DB.Prepare(`select t.id, t.exerciseid, t2."name" exercisename, t.series, t.fat, 
		to_char(t.time,'HH24:MI') as time, t.tranningid
		from gym.tbtrainningsteps t inner join gym.tbexercises t2 on t2.id = t.exerciseid where t.id $1`)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&c.Id, &c.Exercise.Id, &c.Exercise.Name, &c.Series, &c.Fat, &c.Time, &c.Tranningid)
	jsonPrint, _ := json.Marshal(c)
	log.Println("(TrainningStep - Update) body:", string(jsonPrint))

	log.Println("(TrainningStep - GetById) Finish")
	return nil
}

func (c *TrainningStep) GetAll(name string) ([]TrainningStep, error) {
	log.Println("(TrainningStep - GetAll) Init")
	log.Println("(TrainningStep - GetAll) body:", c)

	stmt, err := DB.Prepare(`select t.id, t.exerciseid, t2."name" exercisename, t.series, t.fat, 
		to_char(t.time,'HH24:MI') as time, t.tranningid
		from gym.tbtrainningsteps t inner join gym.tbexercises t2 on t2.id = t.exerciseid `)
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error:", err.Error())
		return nil, err
	}
	defer rows.Close()
	trainning := []TrainningStep{}
	for rows.Next() {
		exercise := TrainningStep{}
		rows.Scan(&exercise.Id, &exercise.Exercise.Id, &exercise.Exercise.Name,
			&exercise.Series, &exercise.Fat, &exercise.Time, &exercise.Tranningid)
		trainning = append(trainning, exercise)
	}
	log.Println("(TrainningStep - GetAll) rows:", len(trainning))
	jsonPrint, _ := json.Marshal(trainning)
	log.Println("(TrainningStep - GetAll) body:", string(jsonPrint))

	log.Println("(TrainningStep - GetAll) Finish")
	return trainning, nil
}
