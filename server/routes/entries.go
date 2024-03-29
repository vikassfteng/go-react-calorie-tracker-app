package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var entryCollection *mongo.Collection = openCollection(Client, "calories"))
func AddEntry(c *gin.Context) {
	// AddEntry is a function that adds an entry to the database
}

func GetEntries(c *gin.Context) {
	// GetEntries is a function that gets all entries from the database

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	} 
	cursor.All(ctx, &entries); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntriesByIngridient(c *gin.Context) {
	// GetEntryByID is a function that gets an entry by ID from the database
}

func GetEntryByID(c *gin.Context) {
	// GetEntryByID is a function that gets an entry by ID from the database
}

func UpdateEntry(c *gin.Context) {
	// UpdateEntry is a function that updates an entry in the database
}

func UpdateIngredient(c *gin.Context) {
	// UpdateIngredient is a function that updates an ingredient in the database
}

func DeleteEntry(c *gin.Context) {
	// DeleteEntry is a function that deletes an entry from the database

	entryID := c.Param.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)


	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.,JSON(http.StatusOK, result.DeletedCount)})
}	

