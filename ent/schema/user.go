package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Comment("用户 ID"),
		field.String("username").
			MaxLen(50).
			NotEmpty().
			Comment("用户名"),
		field.String("password").
			MaxLen(100).
			NotEmpty().
			Sensitive().
			Comment("密码"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
