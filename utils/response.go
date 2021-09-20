package response

import "github.com/gofiber/fiber/v2"

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

func ErrorInvalidID() *fiber.Map {
	return &fiber.Map{
		"message": "invalid id",
	}
}

func ErrorDeletion() *fiber.Map {
	return &fiber.Map{
		"message": "deletion error",
	}
}
