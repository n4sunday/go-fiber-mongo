package position

import (
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/n4sunday/go-fiber-mongo/database"
	"github.com/n4sunday/go-fiber-mongo/utils/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "positions"

func GetPosition(c *fiber.Ctx) error {
	collection := database.DB.Db.Collection(CollectionName)

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	var result Position
	collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&result)
	return c.JSON(result)
}

func GetAllPosition(c *fiber.Ctx) error {
	collection := database.DB.Db.Collection(CollectionName)
	ctx := c.Context()

	var results []Position = make([]Position, 0)

	query := bson.D{{}}
	findOptions := options.Find()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := collection.CountDocuments(ctx, query)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{"age", -1}})

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if err := cursor.All(ctx, &results); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(200).JSON(fiber.Map{
		"data":      results,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}

func CreatePosition(c *fiber.Ctx) error {
	collection := database.DB.Db.Collection(CollectionName)

	data := new(Position)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	data.ID = ""
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	insertionResult, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdResult := &Position{}
	createdRecord.Decode(createdResult)

	return c.Status(201).JSON(createdResult)
}

func UpdatePosition(c *fiber.Ctx) error {
	collection := database.DB.Db.Collection(CollectionName)

	idParam := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	data := new(Position)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "level", Value: data.Level},
				{Key: "updated_at", Value: time.Now()},
			},
		},
	}
	err = collection.FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	data.ID = idParam
	return c.Status(200).JSON(response.UpdateSuccess())
}

func DeletePosition(c *fiber.Ctx) error {
	collection := database.DB.Db.Collection(CollectionName)

	id, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	query := bson.D{{Key: "_id", Value: id}}
	result, err := collection.DeleteOne(c.Context(), &query)

	if err != nil {
		return c.Status(500).JSON(response.ErrorDeletion())
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(response.DeleteSuccess())
}
