package data

type Adress struct{
  Location string `bson:"location"`
  State string `bson:"state"`
  Pin string `bson:"pin"`
}