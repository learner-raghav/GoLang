package EmployeeDelivery

import (
	"../../entity"
	"../../service"
	"encoding/json"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	empService service.EmployeeServicer
}


type Response struct{
	Msg string
	Emp entity.Employee
}

func New(empService service.EmployeeServicer) EmployeeHandler {
	return EmployeeHandler{empService: empService}
}

func (empHandler EmployeeHandler) Handle(res http.ResponseWriter,req *http.Request){
	switch req.Method {
	case http.MethodGet:
		res.WriteHeader(http.StatusOK)
		empHandler.getById(res,req)
	case http.MethodPost:
		res.WriteHeader(http.StatusOK)
		empHandler.create(res,req)
	case http.MethodPut:
		res.WriteHeader(http.StatusOK)
		empHandler.update(res,req)
	case http.MethodDelete:
		res.WriteHeader(http.StatusOK)
		empHandler.delete(res,req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (empHandler EmployeeHandler) getById(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	employeeId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	employee,err := empHandler.empService.GetById(employeeId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Emp: employee}
	json.NewEncoder(res).Encode(r)
}

func (empHandler EmployeeHandler) create(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	emp := entity.Employee{}
	err := json.NewDecoder(req.Body).Decode(&emp)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Msg": "Error UnMarshalling the Body"}`))
		return
	}

	employee,err := empHandler.empService.Create(emp)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Emp: employee}
	json.NewEncoder(res).Encode(r)
}

func (empHandler EmployeeHandler) update(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	employeeId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	var emp entity.Employee
	json.NewDecoder(req.Body).Decode(&emp)
	emp.EmployeeId = employeeId //Assigning the id to the object.
	employee,err := empHandler.empService.Update(emp)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Employee Updated Successfully!!",Emp: employee}
	json.NewEncoder(res).Encode(r)
}

func (empHandler EmployeeHandler) delete(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	employeeId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	emp,err := empHandler.empService.Delete(employeeId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Emp: entity.Employee{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Employee Deleted Successfully",Emp: emp}
	json.NewEncoder(res).Encode(r)

}