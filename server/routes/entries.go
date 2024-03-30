package routes

import (
	"log"
    "context"
    "fmt"
    "net/http"
    "time"

	"github.com/vikassfteng/go-react-calorie-tracker-app/models"

	"github.com/go-playground/validator/v10"
	
    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	// AddEntry is a function that adds an entry to the database

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entry models.Entry
	var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(entry)
	log.Println("validating entry")

	if validationErr != nil {
		log.Println("Inserting entry to the database and validationErr: ", validationErr)

		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	log.Println("Inserting entry to the database")

	entry.ID = primitive.NewObjectID()
	// for inserting entry to the database
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		msg := fmt.Sprintf("Order Item was not created: %v", insertErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetEntries(c *gin.Context) {
	// GetEntries is a function that gets all entries from the database

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	err = cursor.All(ctx, &entries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntriesByIngredient(c *gin.Context) {
	// GetEntryByID is a function that gets an entry by ID from the database
	ingredient := c.Param("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntryByID(c *gin.Context) {
	// GetEntryByID is a function that gets an entry by ID from the database

	EntryID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry bson.M
	if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)
}

func UpdateIngredient(c *gin.Context) {
	// UpdateIngredient is a function that updates an ingredient in the database
	entryID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient
	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if ingredient.Ingredients == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ingredients cannot be nil"})
		return
	}
	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdateEntry(c *gin.Context) {
	// UpdateIngredient is a function that updates an ingredient in the database

	entryID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	docID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	result, err := entryCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"dish":        entry.Dish,
			"ingredients": entry.Ingredients,
			"fat":         entry.Fat,
			"calories":    entry.Calories,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("Error replacing entry: %v\n", err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteEntry(c *gin.Context) {
	// DeleteEntry is a function that deletes an entry from the database

	entryID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
