package VariantService

import (
	"../../entity"
	"../../service"
)
type VariantService struct {
	variantStore service.VariantStoreHandler
}

func New(variantStore service.VariantStoreHandler) VariantService  {
	return VariantService{variantStore: variantStore}
}

func (variantService VariantService) GetById(id int) (entity.Variant,error) {

	variant,err := variantService.variantStore.GetById(id)
	return variant,err

}

func (variantService VariantService) Read(filter map[string]string) ([]entity.Variant,error){

	variants,err := variantService.variantStore.Read(filter)
	return variants,err
}

func (variantService VariantService) Create(variant entity.Variant) (entity.Variant,error){

	variant,err := variantService.variantStore.Create(variant)
	return variant,err
}

func (variantService VariantService) Delete(variantId int) (entity.Variant,error){
	variant,err := variantService.variantStore.Delete(variantId)
	return variant,err
}

func (variantService VariantService) Update(variant entity.Variant) (entity.Variant,error){
	variant,err := variantService.variantStore.Update(variant)
	return variant,err
}



