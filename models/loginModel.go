package models

import (
	"errors"
	"log"

	"github.com/mccaetano/gymtranning/utils"
)

type UserLogin struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *UserLogin) Login() (tokenString string, err error) {

	stmt, err := DB.Prepare(`select id from gym.tbuserprofile where email = $1 and password = $2`)
	if err != nil {
		log.Println("Erro Validar Usuario, err:", err.Error())
		return
	}
	defer stmt.Close()

	stmt.QueryRow(c.Username, c.Password).Scan(&c.Id)
	if c.Id == 0 {
		err = errors.New("invalid user or password")
		log.Println("Erro ao Gerar o Token, err:", err.Error())
		return
	}

	if tokenString, err = utils.CreateToken(uint64(c.Id)); err != nil {
		log.Println("Erro ao Gerar o Token, err:", err.Error())
		return
	}

	return
}
