package BrandDelivery

import (
	"../../entity"
	"../../delivery"
	"encoding/json"
	"net/http"
	"strconv"
)

type BrandHandler struct {
	brandServicer delivery.BrandServicer
}


type Response struct{
	Msg string
	Brand entity.Brand
}

func New(brandServicer delivery.BrandServicer) BrandHandler {
	return BrandHandler{brandServicer: brandServicer}
}

func (brandHandler BrandHandler) Handle(res http.ResponseWriter,req *http.Request){
	switch req.Method {
	case http.MethodGet:
		brandHandler.getById(res,req)
	case http.MethodPost:
		brandHandler.create(res,req)
	case http.MethodPut:
		brandHandler.update(res,req)
	case http.MethodDelete:
		brandHandler.delete(res,req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (brandHandler BrandHandler) getById(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	brandId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	brand,err := brandHandler.brandServicer.GetById(brandId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Brand: brand}
	json.NewEncoder(res).Encode(r)
}

func (brandHandler BrandHandler) create(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	brand := entity.Brand{}
	err := json.NewDecoder(req.Body).Decode(&brand)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Msg": "Error UnMarshalling the Body"}`))
		return
	}

	brand_res,err := brandHandler.brandServicer.Create(brand)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Brand: brand_res}
	json.NewEncoder(res).Encode(r)
}

func (brandHandler BrandHandler) update(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	brandId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}

	var brand entity.Brand
	json.NewDecoder(req.Body).Decode(&brand)
	brand.Brand_id = brandId //Assigning the id to the object.
	brand_res,err := brandHandler.brandServicer.Update(brand)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Brand: brand_res}
	json.NewEncoder(res).Encode(r)
}

func (brandHandler BrandHandler) delete(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","application/json")
	brandId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: "Invalid Id Format",Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	brand_res,err := brandHandler.brandServicer.Delete(brandId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r := Response{Msg: err.Error(),Brand: entity.Brand{}}
		json.NewEncoder(res).Encode(r)
		return
	}
	res.WriteHeader(http.StatusOK)
	r := Response{Msg: "Success",Brand: brand_res}
	json.NewEncoder(res).Encode(r)

}