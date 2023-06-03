package converter

import (
	model "example.com/main-service/model"
	view "example.com/main-service/view"
)

func ModelToView(model model.Example) view.Example {
	var view view.Example

	view.ID = model.ID
	view.Name = model.Name
	view.Password = model.Password

	return view
}

func ViewToModel(view view.Example) model.Example {
	var model model.Example

	model.ID = view.ID
	model.Name = view.Name
	model.Password = view.Password

	return model
}
