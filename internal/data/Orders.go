package data

type Order struct{
 Id string `bson:"id"`
 Items []Product `bson:"items"`
}