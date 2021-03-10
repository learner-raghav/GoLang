package BrandService


import (
	"../../entity"
	"../../service"
)
type BrandService struct {
	brandStore service.BrandStoreHandler
}

func New(brandStore service.BrandStoreHandler) BrandService  {
	return BrandService{brandStore: brandStore}
}

func (brandService BrandService) GetById(id int) (entity.Brand,error) {

	brand,err := brandService.brandStore.GetById(id)
	return brand,err

}

func (brandService BrandService) Create(brand entity.Brand) (entity.Brand,error){
	
	brand,err := brandService.brandStore.Create(brand)
	return brand,err

}

func (brandService BrandService) Delete(brandId int) (entity.Brand,error){
	brand,err := brandService.brandStore.Delete(brandId)
	return brand,err
}

func (brandService BrandService) Update(brand entity.Brand) (entity.Brand,error){
	brand,err := brandService.brandStore.Update(brand)
	return brand,err
}


