package main

import (
	"github.com/alexyslozada/daopattern/utilities"
	"log"
	"github.com/alexyslozada/daopattern/dao/factory"
	"github.com/alexyslozada/daopattern/models"
	"fmt"
)

func main() {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	userDAO := factory.FactoryDAO(config.Engine)
	user := models.User{}

	fmt.Print("Nombre: ")
	fmt.Scan(&user.FirstName)
	fmt.Print("Apellido: ")
	fmt.Scan(&user.LastName)
	fmt.Print("Correo: ")
	fmt.Scan(&user.Email)

	err = userDAO.Create(&user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
}
