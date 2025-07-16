package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Comment("岗位 ID"),
		field.String("title").
			MaxLen(100).
			Comment("岗位名称"),
		field.String("company").
			MaxLen(100).
			Comment("公司名称"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
