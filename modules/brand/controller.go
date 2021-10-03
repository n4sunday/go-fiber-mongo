package brand

import (
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/n4sunday/go-fiber-mongo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateBranch(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("brand")

	// New Employee struct
	data := new(Brand)
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

	createdData := &Brand{}
	createdRecord.Decode(createdData)

	return c.Status(201).JSON(createdData)
}

func GetAllBrand(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("brand")
	ctx := c.Context()

	var result []Brand = make([]Brand, 0)

	query := bson.M{}
	findOptions := options.Find()

	query = bson.M{
		"$or": []bson.M{
			{
				"name": bson.M{
					"$regex": primitive.Regex{
						Pattern: "R2",
						Options: "i",
					},
				},
			},
		},
	}

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

	if err := cursor.All(ctx, &result); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(200).JSON(fiber.Map{
		"data":      result,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}
