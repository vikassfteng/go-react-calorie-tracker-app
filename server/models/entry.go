package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID 				primitive.ObjectID 	`bson:"_id`
	Dish 			*string 			`json:"dish"`
	Fat 			*float64 			`json:"fat"`
	Ingriedients	*string 			`json:"ingriedients"`
	Calories 		*string 			`json:"calories"`
}
