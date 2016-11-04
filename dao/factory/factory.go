package factory

import (
	"github.com/alexyslozada/daopattern/dao/interfaces"
	"github.com/alexyslozada/daopattern/dao/psql"
	"github.com/alexyslozada/daopattern/dao/mysql"
	"log"
)

func FactoryDAO(e string) interfaces.UserDAO {
	var i interfaces.UserDAO
	switch e {
	case "postgres":
		i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.UserImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", e)
		return nil
	}
	return i
}
