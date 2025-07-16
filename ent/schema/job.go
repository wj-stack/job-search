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
		field.String("description").
			Comment("岗位描述"),
		field.String("requirement").
			Comment("岗位要求"),
		field.String("job_category").
			Comment("岗位类别"),
		field.String("city_info").
			Comment("城市信息"),
		field.String("recruit_type").
			Comment("招聘类型"),
		field.Time("publish_time").
			Comment("发布时间"),
		field.String("code").
			Comment("代码"),
		field.Strings("city_list").
			Optional().
			Comment("城市列表"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
