package ModelService


import (
	"../../entity"
	"../../service"
)
type ModelService struct {
	modelStore service.ModelStoreHandler
}

func New(modelStore service.ModelStoreHandler) ModelService  {
	return ModelService{modelStore: modelStore}
}

func (modelService ModelService) GetById(id int) (entity.Model,error) {

	model,err := modelService.modelStore.GetById(id)
	return model,err

}

func (modelService ModelService) Read(filter map[string]string) ([]entity.Model,error){
	models,err := modelService.modelStore.Read(filter)
	return models,err
}

func (modelService ModelService) Create(model entity.Model) (entity.Model,error){

	model,err := modelService.modelStore.Create(model)
	return model,err

}

func (modelService ModelService) Delete(modelId int) (entity.Model,error){
	model,err := modelService.modelStore.Delete(modelId)
	return model,err
}

func (modelService ModelService) Update(model entity.Model) (entity.Model,error){
	model,err := modelService.modelStore.Update(model)
	return model,err
}



