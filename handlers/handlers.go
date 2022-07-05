package handlers

import (
	"context"
	"fmt"
	"ginweb/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipesHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRecipesHandler(ctx context.Context, collection *mongo.Collection) *RecipesHandler {
	return &RecipesHandler{
		collection: collection,
		ctx:        ctx,
	}
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//      description: Successful operation
func (handler *RecipesHandler) ListRecipeHandler(c *gin.Context) {

	curr, err := handler.collection.Find(handler.ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer curr.Close(handler.ctx)
	recipes := make([]models.Recipe, 0)
	for curr.Next(handler.ctx) {
		var recipe models.Recipe
		curr.Decode(&recipe)
		recipes = append(recipes, recipe)
	}
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation GET /recipes/{id} recipes GetOneRecipe
// Search recipes based on id
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: recipe id
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
func (handler *RecipesHandler) GetOneRecipeHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	curr := handler.collection.FindOne(handler.ctx, bson.M{
		"_id": objectId,
	})
	var recipe models.Recipe
	err := curr.Decode(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// swagger:operation PUT /recipes/{id} recipes UpdateRecipes
// Update an Excisting recipe
// ---
// parameters:
// - name: id
//   in: path
//   desciption: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//       description: Successfull operation StatusOk
//     '400':
//       description: invalid input StatusBadRequest
//     '404':
//       description: invalid recipe id StatusNotFound
func (handler *RecipesHandler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	objectId, _ := primitive.ObjectIDFromHex(id)

	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{
		"_id": objectId,
	}, bson.D{{
		"$set", bson.D{
			{
				"name", recipe.Name,
			},
			{
				"instructions", recipe.Instructions,
			},
			{
				"incredients", recipe.Ingredients,
			},
			{
				"tags", recipe.Tags,
			},
		}}})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe Has Been Updated",
	})
}

// swagger:operation DELETE /recipes/{id} recipes DeleteRecipes
// Deletes an existing recipe
// ---
// parameters:
// - name: id
//   required: true
//   in: path
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//       description: Successfull Operation StatusOk
//     '400':
//       description: invalid input StatusBadRequest
func (handler *RecipesHandler) DeleteRecipeHandler(c *gin.Context) {

	id := c.Params.ByName("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	_, err := handler.collection.DeleteOne(handler.ctx, bson.M{
		"_id": objectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe Has Been Successfully Deleted",
	})

}

// swagger:operation POST /recipes recipes AddNewRecipes
// Adds a new recipe from post request
// ---
// consumes:
// - application/json
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {

	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}
	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now()
	_, err := handler.collection.InsertOne(handler.ctx, recipe)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error While inserting a new recipe",
		})
		return
	}
	c.JSON(http.StatusOK, recipe)

}
