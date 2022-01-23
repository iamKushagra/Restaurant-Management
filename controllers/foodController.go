package controller

import(
	"context"
	"fmt"
	"golang-restaurant-mangement/database"
	"golang-restaurant-mangement/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection * mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c * gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)

		recordPerPage, err: = strvconv.Atoi(c.Query("record_per_page"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err: = strvconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex: = (page - 1) * recordPerPage
		startIndex, err = strvconv.Atoi(c.Query("start_index"))

		matchStage := bson.D{
			{ "$match", bson.D{
				{"category", c.Query("category")},
			}}}

		groupStage := bson.D{
			{ "$group", bson.D{
				{"_id",bson.D{{"_id","null"}}},
				{"total_count",bson.D{{"$sum",1}}},
				{"data",bson.D{{"$push","$$ROOT"}}},
			}}}

		projectStage := bson.D{
			{ "$project", bson.D{
				{"_id",0},
				{"total_count",1},
				{"food_items",bson.D{{"$slice",[]interface{}{"$data",startIndex,recordPerPage}}}},
			}}}

		foodCollection.Aggregate(ctx,mongo.Pipeline{
			matchStage, groupStage, projectStage
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c * gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
		foodId: = c.Param("food_id")
		var food models.Food

		err: = foodCollection.FindOne(ctx, bson.M {
			"food_id": foodId
		}).Decode( & food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": "Error occured while fetching food"
			})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c * gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
		var menu models.Menu
		var food models.Food

		if err: = c.BindJSON( & food);
		err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error()
			})
			return
		}

		validationErr: = validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": validationErr.Error()
			})
			return
		}
		err: = foodCollection.FindOne(ctx, bson.M {
			"food_id": food.Menu_id
		}).Decode( & food)
		defer cancel()
		if err != nil {
			msg: = fmt.Sprintf("Menu was not found")
			c.JSON(http.StatusNotFound, gin.H {
				"error": msg
			})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed( * food.Price, 2)
		food.Price = & num

		result, insertErr: = foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg: = fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": msg
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, result)
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc {
	return func(c * gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		var food models.Menu

		if err: = c.BindJSON( & food);
		err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error()
			})
			return
		}

		foodId: = c.Param("food_id")
		filter: = bson.M {
			"food_id": foodId
		}

		var updatedObj primitive.D

		if food.Start_Date != nil && food.End_Date != nil {
			if !inTimeSpan( * food.Start_Date, * food.End_Date, time.Now()) {
				msg: = fmt.Sprintf("Menu is not in the time span")
				c.JSON(http.StatusBadRequest, gin.H {
					"error": msg
				})
				defer cancel()
				return
			}
			updatedObj = append(updatedObj, bson.E {
				"start_date",
				food.Start_Date
			})
			updatedObj = append(updatedObj, bson.E {
				"end_date",
				food.End_Date
			})

			if food.Name != "" {
				updatedObj = append(updatedObj, bson.E {
					"name",
					food.Name
				})
			}
			if food.Category != "" {
				updatedObj = append(updatedObj, bson.E {
					"category",
					food.Category
				})
			}
			food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updatedObj = append(updatedObj, bson.E {
				"updated_at",
				food.Updated_at
			})

			upsert: = true
			opt: = options.UpdateOptions {
				Upsert: & upsert
			}
			result, err: = food.Collection.UpdateOne(ctx, filter, bson.D {
				{
					"$set",
					updatedObj
				}
			}, & opt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"error": "Error occured while updating food"
				})
			}
			defer cancel()
			c.JSON(http.StatusOK, result)
		}
	}
}