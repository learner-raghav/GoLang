package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../entity"
)


func ConnectToDB(config entity.MySQLConfig) (*sql.DB,error){
	dbDriver := "mysql"
	dbUser := "raghav"
	dbPass := "raghav@123M"
	dbName := "testDB"

	db,err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)

	if err != nil{
		return nil,err
	}
	return db,nil
}

