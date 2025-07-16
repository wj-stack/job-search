package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Comment("投递记录 ID"),
		field.Int("user_id").
			Comment("用户 ID"),
		field.Int("job_id").
			Comment("岗位 ID"),
		field.String("resume_url").
			MaxLen(200).
			Comment("简历文件地址"),
		field.String("status").
			MaxLen(20).
			Comment("投递状态"),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return nil
}
