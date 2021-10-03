package employee

import (
	"time"

	"github.com/n4sunday/go-fiber-mongo/modules/position"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Salary    float64            `json:"salary,omitempty"`
	Age       float64            `json:"age,omitempty"`
	Position  primitive.ObjectID `json:"position,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type EmployeeResult struct {
	ID        string            `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Salary    float64           `json:"salary,omitempty"`
	Age       float64           `json:"age,omitempty"`
	Positions position.Position `json:"positions,omitempty"`
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}
