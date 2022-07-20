package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Route holds the schema definition for the Route entity.
type Route struct {
	ent.Schema
}

// Fields of the Route.
func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain").NotEmpty(),
		field.String("path").Nillable().Optional(),
	}
}

// Edges of the Route.
func (Route) Edges() []ent.Edge {
	return nil
}
