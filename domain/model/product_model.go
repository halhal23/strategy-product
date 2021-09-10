package model

type ProductModel struct {
	ID int
	Name string
	Price int
	UserId int
}

func NewProductModel(id int, name string, price int, userId int) *ProductModel {
	return &ProductModel{
		ID: id,	
		Name: name,
		Price: price,
		UserId: userId,
	}
}