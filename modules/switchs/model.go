package switchs

import (
	"time"

	"github.com/n4sunday/go-fiber-mongo/modules/brand"
)

type Swtichs struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name"`
	Force     int       `json:"force"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type SwtichsResult struct {
	ID        string      `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string      `json:"name"`
	Force     int         `json:"force"`
	Type      string      `json:"type"`
	Brand     brand.Brand `json:"brand,omitempty"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
}
