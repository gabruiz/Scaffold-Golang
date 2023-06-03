package converter

type converter interface {
	ModelToView()
	ViewToModel()
}
