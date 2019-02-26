package restCRUD

type Student struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Email string `json:"email"`
}

type Students []Student