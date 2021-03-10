package ModelStore

import (
	"database/sql"
	"errors"
	"fmt"
)
import "../../entity"
import "../../datastore/BrandStore"
type ModelStore struct {
	db *sql.DB
}

func New(db *sql.DB) ModelStore {
	return ModelStore{db: db}
}

func (modelStr ModelStore) GetById(id int) (entity.Model,error) {

	db := modelStr.db
	if db == nil {
		return entity.Model{},errors.New("DB not configured properly")
	}

	//Should not use Joins, as in real life we might have different databases.
	viewQuery, err := db.Query("select model_id,model_name,brand_id from model where model.model_id= ?",id)
	if err != nil {
		return entity.Model{},err
	}
	var model entity.Model
	if viewQuery.Next() {
		var model_id int
		var model_name string
		var brand_id int
		err = viewQuery.Scan(&model_id,&model_name,&brand_id)
		if err != nil {
			return entity.Model{},err
		}
		model = entity.Model{Model_id: model_id,Model_name: model_name,Brand_id: brand_id}
		brandStr := BrandStore.New(modelStr.db)
		brand, err := brandStr.GetById(model.Brand_id)
		if err != nil {
			return entity.Model{},err
		}
		model.Brand_name = brand.Brand_name
	} else{
		err := errors.New("Record does not exist")
		return entity.Model{},err
	}
	return model, nil
}

func (modelStr ModelStore) Read(filters map[string]string) ([]entity.Model,error){
	queryStr := "Select * from model where "
	n := len(filters)
	i := 1
	for k,v := range filters {
		var temp string
		if i < n {
			temp = (k + "=\"" + v + "\" and ")
		} else {
			temp = (k + "=\"" + v+"\"")
		}
		queryStr += temp
		i += 1
	}
	db := modelStr.db
	if db == nil {
		return []entity.Model{},errors.New("DB not configured properly")
	}
	//Should not use Joins, as in real life we might have different databases.
	viewQuery, err := db.Query(queryStr)
	if err != nil {
		return []entity.Model{},err
	}
	var models []entity.Model
	flag := 0
	for viewQuery.Next() {
		flag = 1
		var model entity.Model
		var model_id int
		var model_name string
		var brand_id int
		err = viewQuery.Scan(&model_id,&model_name,&brand_id)
		if err != nil {
			return []entity.Model{},err
		}
		model = entity.Model{Model_id: model_id,Model_name: model_name,Brand_id: brand_id}
		brandStr := BrandStore.New(modelStr.db)
		brand, err := brandStr.GetById(model.Brand_id)
		if err != nil {
			return []entity.Model{},err
		}
		model.Brand_name = brand.Brand_name
		models = append(models,model)
	}
	fmt.Println(models)
	if flag == 0 {
		err := errors.New("Record does not exist")
		return []entity.Model{},err
	}
	return models, nil
}

func (modelStr ModelStore) Create(model entity.Model) (entity.Model,error){
	db := modelStr.db

	if db == nil {
		return entity.Model{}, errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO model(model_name,brand_id) values(?,?)")
	if err != nil {
		return entity.Model{},err
	}
	res, error := insertQuery.Exec(model.Model_name,model.Brand_id)
	if error != nil {
		return entity.Model{},errors.New("Foreign key does not exist")
	}
	rowsAffect, err := res.RowsAffected()
	if rowsAffect == 0 {
		return entity.Model{},errors.New("No rows inserted")
	}
	return model,nil
}

func (modelStr ModelStore) Update(model entity.Model) (entity.Model,error){
	db := modelStr.db
	if db == nil {
		return entity.Model{},errors.New("DB not configured properly")
	}
	updateQuery, err := db.Prepare("UPDATE model SET model_name=?, brand_id=? where model_id = ?")
	if err != nil {
		return entity.Model{}, err
	}
	res1,err := updateQuery.Exec(model.Model_name,model.Brand_id,model.Model_id)
	if err != nil {
		error := errors.New("Error occured")
		return entity.Model{},error
	}
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		return entity.Model{},errors.New("The id does not exist or no modifications were done")
	}
	return model,nil
}

func (modelStr ModelStore) Delete(modelId int) (entity.Model,error){

	db := modelStr.db
	if db == nil {
		return entity.Model{},errors.New("DB not configured properly")
	}
	model, err := modelStr.GetById(modelId)

	if err != nil {
		return entity.Model{},err
	}

	deleteStatement,err := db.Prepare("DELETE from model where model_id=?")
	if err != nil {
		return entity.Model{}, err
	}

	result,err := deleteStatement.Exec(modelId)
	if err != nil {
		return entity.Model{},err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("ModelStore record does not exist")
		return entity.Model{},err
	}
	return model,nil
}
