package BrandStore

import (
	"database/sql"
	"errors"
)
import "../../entity"

type BrandStore struct {
	db *sql.DB
}

func New(db *sql.DB) BrandStore {
	return BrandStore{db: db}
}

func (brndStr BrandStore) GetById(id int) (entity.Brand,error) {

	db := brndStr.db
	if db == nil {
		return entity.Brand{},errors.New("DB not configured properly")
	}

	viewQuery, err := db.Query("SELECT * FROM brand where brand_id = ?",id)
	if err != nil {
		return entity.Brand{},err
	}
	var brand entity.Brand
	if viewQuery.Next() {
		var brand_id int
		var brand_name string
		err = viewQuery.Scan(&brand_id,&brand_name)
		if err != nil {
			return entity.Brand{},err
		}
		brand = entity.Brand{Brand_id: brand_id,Brand_name: brand_name}
	} else{
		err := errors.New("BrandStore record does not exist")
		return entity.Brand{},err
	}
	return brand, nil

}

func (brndStr BrandStore) Create(brand entity.Brand) (entity.Brand,error){
	db := brndStr.db

	if db == nil {
		return entity.Brand{}, errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO brand(brand_name) values(?)")
	if err != nil {
		return entity.Brand{},err
	}
	res, error := insertQuery.Exec(brand.Brand_name)
	rowsAffect,err := res.RowsAffected()
	if rowsAffect == 0 {
		return entity.Brand{},errors.New("The id does not exist or no modifications were done")
	}
	if error != nil {
		return entity.Brand{},error
	}
	return brand,nil
}

func (brndStr BrandStore) Update(brand entity.Brand) (entity.Brand,error){
	db := brndStr.db
	if db == nil {
		return entity.Brand{},errors.New("DB not configured properly")
	}
	updateQuery, err := db.Prepare("UPDATE brand SET brand_name=? where brand_id = ?")
	if err != nil {
		return entity.Brand{}, err
	}
	res1,err := updateQuery.Exec(brand.Brand_name,brand.Brand_id)
	if err != nil {
		error := errors.New("Error occured")
		return entity.Brand{},error
	}
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		return entity.Brand{},errors.New("The id does not exist or no modifications were done")
	}
	return brand,nil
}

func (brndStr BrandStore) Delete(brandId int) (entity.Brand,error){

	db := brndStr.db
	brand, err := brndStr.GetById(brandId)

	if err != nil {
		return entity.Brand{},err
	}
	if db == nil {
		return entity.Brand{},errors.New("DB not configured properly")
	}

	deleteStatement,err := db.Prepare("DELETE from brand where brand_id=?")
	if err != nil {
		return entity.Brand{}, err
	}

	result,err := deleteStatement.Exec(brand.Brand_id)
	if err != nil {
		return entity.Brand{},err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("The id does not exist")
		return entity.Brand{},err
	}
	return brand,nil
}