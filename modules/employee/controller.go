package employee

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

func GetEmployee(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("employees")

	idParam := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	var result Employee
	collection.FindOne(c.Context(), bson.M{"_id": employeeID}).Decode(&result)
	return c.JSON(result)
}

func GetAllEmployee(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("employees")
	ctx := c.Context()

	var employees []Employee = make([]Employee, 0)

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

	if err := cursor.All(ctx, &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(200).JSON(fiber.Map{
		"data":      employees,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}

func CreateEmployee(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("employees")

	// New Employee struct
	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// force MongoDB to always set its own generated ObjectIDs
	employee.ID = ""
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	insertionResult, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Employee
	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)

	// return the created Employee in JSON format
	return c.Status(201).JSON(createdEmployee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("employees")

	idParam := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	employee := new(Employee)
	// Parse body into struct
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Find the employee and update its data
	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
				{Key: "updated_at", Value: time.Now()},
			},
		},
	}
	err = collection.FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated employee
	employee.ID = idParam
	return c.Status(200).JSON(response.UpdateSuccess())
}

func DeleteEmployee(c *fiber.Ctx) error {
	db := database.DB
	collection := db.Db.Collection("employees")

	employeeID, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(response.ErrorInvalidID())
	}

	// find and delete the employee with the given ID
	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := collection.DeleteOne(c.Context(), &query)

	if err != nil {
		return c.Status(500).JSON(response.ErrorDeletion())
	}

	// the employee might not exist
	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(response.DeleteSuccess())
}
