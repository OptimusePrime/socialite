package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}). /*.Default(uuid.New)*/ StorageKey("old"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("name"),
		field.String("password"),
		field.Time("birthDate").Optional(),
		field.String("avatar").Optional(),
		field.String("biography").Optional(),
		field.String("gender").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
