package VariantStore

import (
	"../../datastore/ModelStore"
	"../../entity"
	"database/sql"
	"errors"
)
type VariantStore struct {
	db *sql.DB
}

func New(db *sql.DB) VariantStore {
	return VariantStore{db: db}
}

func (VarStr VariantStore) GetById(id int) (entity.Variant,error) {

	db := VarStr.db
	if db == nil {
		return entity.Variant{},errors.New("DB not configured properly")
	}

	//removing usage of joins.
	viewQuery, err := db.Query("select variant_id,model_id,name,disp,peak_power,peak_torque from variant where variant_id=?",id)
	if err != nil {
		return entity.Variant{},err
	}
	var variant entity.Variant
	if viewQuery.Next() {
		var variant_id,model_id int
		var disp,peak_power,peak_torque float64
		var variant_name string
		err = viewQuery.Scan(&variant_id,&model_id,&variant_name,&disp,&peak_power,&peak_torque)
		if err != nil {
			return entity.Variant{},err
		}
		variant = entity.Variant{Variant_id: variant_id,Model_id: model_id,Variant_name: variant_name,Peak_torque: peak_torque,Peak_power: peak_power,Disp: disp}

		modelStr := ModelStore.New(VarStr.db)
		model, err := modelStr.GetById(variant.Model_id)
		if err != nil {
			return entity.Variant{},err
		}

		variant.Brand_name = model.Brand_name
		variant.Model_name = model.Model_name
		variant.Brand_Id = model.Brand_id

	} else{
		err := errors.New("Record does not exist")
		return entity.Variant{},err
	}
	return variant, nil

}

func (VarStr VariantStore) Read(filters map[string]string) ([]entity.Variant,error){

	queryStr := "select variant_id,model_id,name,disp,peak_power,peak_torque from variant where "
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
	db := VarStr.db
	if db == nil {
		return []entity.Variant{},errors.New("DB not configured properly")
	}

	//removing usage of joins.
	viewQuery, err := db.Query(queryStr)
	if err != nil {
		return []entity.Variant{},err
	}
	var variants []entity.Variant
	flag := 0
	for viewQuery.Next() {
		flag = 1
		var variant entity.Variant
		var variant_id,model_id int
		var disp,peak_power,peak_torque float64
		var variant_name string
		err = viewQuery.Scan(&variant_id,&model_id,&variant_name,&disp,&peak_power,&peak_torque)
		if err != nil {
			return []entity.Variant{},err
		}
		variant = entity.Variant{Variant_id: variant_id,Model_id: model_id,Variant_name: variant_name,Peak_torque: peak_torque,Peak_power: peak_power,Disp: disp}

		modelStr := ModelStore.New(VarStr.db)
		model, err := modelStr.GetById(variant.Model_id)
		if err != nil {
			return []entity.Variant{},err
		}
		variant.Brand_name = model.Brand_name
		variant.Model_name = model.Model_name
		variant.Brand_Id = model.Brand_id
		variants = append(variants,variant)
	}
	if flag == 0{
		err := errors.New("Record does not exist")
		return []entity.Variant{},err
	}
	return variants, nil



}

func (VarStr VariantStore) Create(variant entity.Variant) (entity.Variant,error){
	db := VarStr.db

	if db == nil {
		return entity.Variant{}, errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO variant(model_id,name,disp,peak_power,peak_torque) values(?,?,?,?,?)")
	if err != nil {
		return entity.Variant{},err
	}
	res, error := insertQuery.Exec(variant.Model_id,variant.Variant_name,variant.Disp,variant.Peak_power,variant.Peak_torque)
	if error != nil {
		return entity.Variant{},errors.New("Foreign key does not exist")
	}
	rowsAffect, err := res.RowsAffected()
	if rowsAffect == 0 {
		return entity.Variant{},errors.New("No rows inserted")
	}
	return variant,nil
}

func (VarStr VariantStore) Update(variant entity.Variant) (entity.Variant,error){
	db := VarStr.db
	if db == nil {
		return entity.Variant{},errors.New("DB not configured properly")
	}
	updateQuery, err := db.Prepare("UPDATE variant SET model_id=?,name=?,disp=?,peak_power=?,peak_torque=? where variant_id = ?")
	if err != nil {
		return entity.Variant{}, err
	}
	res1,err := updateQuery.Exec(variant.Model_id,variant.Variant_name,variant.Disp,variant.Peak_power,variant.Peak_torque,variant.Variant_id)
	if err != nil {
		error := errors.New("Error occured")
		return entity.Variant{},error
	}
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		return entity.Variant{},errors.New("The id does not exist or no modifications were done")
	}
	return variant,nil
}

func (VarStr VariantStore) Delete(variantId int) (entity.Variant,error){

	db := VarStr.db
	if db == nil {
		return entity.Variant{},errors.New("DB not configured properly")
	}
	variant, err := VarStr.GetById(variantId)

	if err != nil {
		return entity.Variant{},err
	}

	deleteStatement,err := db.Prepare("DELETE from variant where variant_id=?")
	if err != nil {
		return entity.Variant{}, err
	}

	result,err := deleteStatement.Exec(variantId)
	if err != nil {
		return entity.Variant{},err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("Record does not exist")
		return entity.Variant{},err
	}
	return variant,nil
}

