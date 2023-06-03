package dao

import "example.com/main-service/model"

func SaveExample(model model.Example) model.Example {
	result := getDatabase().FirstOrCreate(&model)
	checkErrors(result)

	return model
}

func GetExampleById(id int) model.Example {
	var model model.Example
	result := getDatabase().First(&model, id)
	checkErrors(result)

	return model
}
