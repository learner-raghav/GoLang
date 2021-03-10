package service

import 	"../entity"

type BrandStoreHandler interface {
	GetById(id int) (entity.Brand,error)
	Create(brand entity.Brand) (entity.Brand,error)
	Update(brand entity.Brand) (entity.Brand,error)
	Delete(id int) (entity.Brand,error)
}

type ModelStoreHandler interface {
	Read(filter map[string]string) ([]entity.Model,error)
	GetById(id int) (entity.Model,error)
	Create(model entity.Model) (entity.Model,error)
	Update(model entity.Model) (entity.Model,error)
	Delete(id int) (entity.Model,error)
}

type VariantStoreHandler interface {
	Read(filter map[string]string) ([]entity.Variant,error)
	GetById(id int) (entity.Variant,error)
	Create(variant entity.Variant) (entity.Variant,error)
	Update(variant entity.Variant) (entity.Variant,error)
	Delete(id int) (entity.Variant,error)
}

