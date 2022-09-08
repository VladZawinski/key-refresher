package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ApiKey holds the schema definition for the ApiKey entity.
type ApiKey struct {
	ent.Schema
}

// Fields of the ApiKey.
func (ApiKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int("remaining_credit"),
		field.Enum("status").Values("in-use", "expired", "available").Default("available"),
		field.String("key"),
	}
}

// Edges of the ApiKey.
func (ApiKey) Edges() []ent.Edge {
	return nil
}
