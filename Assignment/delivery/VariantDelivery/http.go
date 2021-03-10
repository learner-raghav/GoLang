package VariantDelivery

import (
	"../../delivery"
	"../../entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type VariantHandler struct {
	variantServicer delivery.VariantServicer
}


type Response struct{
	Msg string
	Variant entity.Variant
}

func New(variantServicer delivery.VariantServicer) VariantHandler {
	return VariantHandler{variantServicer: variantServicer}
}

func (variantHandler VariantHandler) Handle(res http.ResponseWriter,req *http.Request){
	switch req.Method {
	case http.MethodGet:
		variantHandler.getById(res,req)
	case http.MethodPost:
		variantHandler.create(res,req)
	case http.MethodPut:
		variantHandler.update(res,req)
	case http.MethodDelete:
		variantHandler.delete(res,req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}


func (varHandler VariantHandler) handleGet(res http.ResponseWriter,req *http.Request){
	id := req.URL.Query().Get("id")
	if id != "" {
		varHandler.getById(res,req)
	} else {
		varHandler.Read(res,req)
	}
}

func (varHandler VariantHandler) Read(res http.ResponseWriter,req *http.Request){
	params := req.URL.Query()
	fmt.Println(params)
	res.Write([]byte("Got it!!"))
}

func (variantHandler VariantHandler) getById(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	variantId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	variant,err := variantHandler.variantServicer.GetById(variantId)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Variant: variant}
	json.NewEncoder(res).Encode(r)
}

func (variantHandler VariantHandler) create(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	variant := entity.Variant{}
	err := json.NewDecoder(req.Body).Decode(&variant)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Msg": "Error UnMarshalling the Body"}`))
		return
	}

	variant_res,err := variantHandler.variantServicer.Create(variant)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Variant: variant_res}
	json.NewEncoder(res).Encode(r)
}

func (variantHandler VariantHandler) update(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	variantId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	var variant entity.Variant
	json.NewDecoder(req.Body).Decode(&variant)

	variant.Variant_id = variantId
	variant_res,err := variantHandler.variantServicer.Update(variant)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Variant: variant_res}
	json.NewEncoder(res).Encode(r)
}

func (variantHandler VariantHandler) delete(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	variantId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	variant_res,err := variantHandler.variantServicer.Delete(variantId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Variant: entity.Variant{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Variant: variant_res}
	json.NewEncoder(res).Encode(r)

}