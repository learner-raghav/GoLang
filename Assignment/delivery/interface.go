package delivery
import "Assignment/entity"

type BrandServicer interface {
	GetById(id int) (entity.Brand,error)
	Create(emp entity.Brand) (entity.Brand,error)
	Update(emp entity.Brand) (entity.Brand,error)
	Delete(employeeId int) (entity.Brand,error)
}

type ModelServicer interface {
	Read(filter map[string]string) ([]entity.Brand,error)
	GetById(id int) (entity.Model,error)
	Create(emp entity.Model) (entity.Model,error)
	Update(emp entity.Model) (entity.Model,error)
	Delete(employeeId int) (entity.Model,error)
}

type VariantServicer interface {
	Read(filter map[string]string) ([]entity.Variant,error)
	GetById(id int) (entity.Variant,error)
	Create(emp entity.Variant) (entity.Variant,error)
	Update(emp entity.Variant) (entity.Variant,error)
	Delete(employeeId int) (entity.Variant,error)
}
