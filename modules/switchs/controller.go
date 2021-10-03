package switchs

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/n4sunday/go-fiber-mongo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllSwitch(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("switch")
	ctx := c.Context()

	// var result []SwtichsResult = make([]SwtichsResult, 0)

	query := bson.M{}
	findOptions := options.Find()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := collection.CountDocuments(ctx, query)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{"age", -1}})

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "brand"}, {"localField", "brand_id"}, {"foreignField", "_id"}, {"as", "brand"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$brand"}, {"preserveNullAndEmptyArrays", true}}}}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})

	// cursor, err := collection.Find(ctx, query, findOptions)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var showsLoaded []bson.M
	if err = cursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}
	log.Println(showsLoaded)
	j, _ := json.MarshalIndent(showsLoaded, "", "  ")
	log.Println(string(j))
	// if err := cursor.All(ctx, &result); err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }

	// last := math.Ceil(float64(total / limit))
	// if last < 1 && total > 0 {
	// 	last = 1
	// }

	return c.Status(200).JSON(fiber.Map{
		"data":      showsLoaded,
		"total":     total,
		"page":      page,
		"last_page": 0,
		"limit":     limit,
	})
}
