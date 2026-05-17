package models

var (
	modelRegistry []interface{}
)

// RegisterModel giúp một model tự đăng ký bản thân vào danh sách migrate
func RegisterModel(model interface{}) {
	modelRegistry = append(modelRegistry, model)
}

// GetModels trả về danh sách tất cả các model đã đăng ký
func GetModels() []interface{} {
	return modelRegistry
}
