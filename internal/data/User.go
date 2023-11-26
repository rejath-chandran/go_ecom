package data



type User struct{
	Name string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password"`
	Adresss []Adress `bson:"address"`
	Cart  []Product `bson:"cart"`
	Orders []Order `bson:"order"`
}

