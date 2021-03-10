package ModelDelivery

import (
	"../../entity"
	"../../delivery"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ModelHandler struct {
	modelServicer delivery.ModelServicer
}


type Response struct{
	Msg string
	Model entity.Model
}

func New(modelServicer delivery.ModelServicer) ModelHandler {
	return ModelHandler{modelServicer: modelServicer}
}

func (modelHandler ModelHandler) Handle(res http.ResponseWriter,req *http.Request){
	switch req.Method {
	case http.MethodGet:
		modelHandler.handleGet(res,req)
	case http.MethodPost:
		modelHandler.create(res,req)
	case http.MethodPut:
		modelHandler.update(res,req)
	case http.MethodDelete:
		modelHandler.delete(res,req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (modelHandler ModelHandler) handleGet(res http.ResponseWriter,req *http.Request){
	id := req.URL.Query().Get("id")
	if id != "" {
		modelHandler.getById(res,req)
	} else {
		modelHandler.Read(res,req)
	}
}

func (modelHandler ModelHandler) Read(res http.ResponseWriter,req *http.Request){
	params := req.URL.Query()
	fmt.Println(params)
	res.Write([]byte("Got it!!"))
}

func (modelHandler ModelHandler) getById(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	modelId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	model,err := modelHandler.modelServicer.GetById(modelId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Model: model}
	json.NewEncoder(res).Encode(r)
}

func (modelHandler ModelHandler) create(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	model := entity.Model{}
	err := json.NewDecoder(req.Body).Decode(&model)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Msg": "Error UnMarshalling the Body"}`))
		return
	}

	model_res,err := modelHandler.modelServicer.Create(model)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Model: model_res}
	json.NewEncoder(res).Encode(r)
}

func (modelHandler ModelHandler) update(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	modelId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	var model entity.Model
	json.NewDecoder(req.Body).Decode(&model)
	model.Model_id = modelId //Assigning the id to the object.
	model_res,err := modelHandler.modelServicer.Update(model)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Model: model_res}
	json.NewEncoder(res).Encode(r)
}

func (modelHandler ModelHandler) delete(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	modelId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	model_res,err := modelHandler.modelServicer.Delete(modelId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Model: entity.Model{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Model: model_res}
	json.NewEncoder(res).Encode(r)

}