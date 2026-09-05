package types

type Students struct {
	Id    int
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   string `validate:"required"`
}
