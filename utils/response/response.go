package response

import "github.com/gofiber/fiber/v2"

func ErrorInvalidID() *fiber.Map {
	return &fiber.Map{
		"message": "invalid id",
	}
}

// UPDATE
func UpdateSuccess() *fiber.Map {
	return &fiber.Map{
		"message": "update success",
	}
}

// DELETE
func DeleteSuccess() *fiber.Map {
	return &fiber.Map{
		"message": "delete success",
	}
}

func ErrorDeletion() *fiber.Map {
	return &fiber.Map{
		"message": "deletion error",
	}
}
