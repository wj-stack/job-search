package dao

import (
	"context"

	"github.com/wj-stack/job-search/ent"
	"github.com/wj-stack/job-search/ent/job"
	"github.com/wj-stack/job-search/internal/domain/model"
	"github.com/wj-stack/job-search/internal/domain/repository"
)

var _ repository.JobRepository = (*JobDAO)(nil)

type JobDAO struct {
	client *ent.Client
}

func NewJobDAO(client *ent.Client) *JobDAO {
	return &JobDAO{
		client: client,
	}
}

// CreateJob 创建新的岗位
func (r *JobDAO) CreateJob(ctx context.Context, title, company string) (*model.Job, error) {
	entJob, err := r.client.Job.Create().
		SetTitle(title).
		SetCompany(company).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:      entJob.ID,
		Title:   entJob.Title,
		Company: entJob.Company,
	}, nil
}

// GetJobByID 根据 ID 获取岗位
func (r *JobDAO) GetJobByID(ctx context.Context, id int) (*model.Job, error) {
	entJob, err := r.client.Job.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:      entJob.ID,
		Title:   entJob.Title,
		Company: entJob.Company,
	}, nil
}

// ListJobs 列出所有岗位
func (r *JobDAO) ListJobs(ctx context.Context) ([]*model.Job, error) {
	entJobs, err := r.client.Job.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	jobs := make([]*model.Job, len(entJobs))
	for i, entJob := range entJobs {
		jobs[i] = &model.Job{
			ID:      entJob.ID,
			Title:   entJob.Title,
			Company: entJob.Company,
		}
	}
	return jobs, nil
}

// UpdateJob 更新岗位信息
func (r *JobDAO) UpdateJob(ctx context.Context, id int, title, company string) (*model.Job, error) {
	entJob, err := r.client.Job.UpdateOneID(id).
		SetTitle(title).
		SetCompany(company).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:      entJob.ID,
		Title:   entJob.Title,
		Company: entJob.Company,
	}, nil
}

// DeleteJob 删除岗位
func (r *JobDAO) DeleteJob(ctx context.Context, id int) error {
	return r.client.Job.DeleteOneID(id).Exec(ctx)
}

// ListJobsWithPagination implements repository.JobRepository.
func (r *JobDAO) ListJobsWithPagination(ctx context.Context, query *repository.JobPaginationQuery) ([]*model.Job, int, error) {
	queryBuilder := r.client.Job.Query()
	if query.Query != "" {
		queryBuilder.Where(job.Or(
			job.TitleContainsFold(query.Query),
			job.CompanyContainsFold(query.Query),
			job.DescriptionContainsFold(query.Query),
			job.RequirementContainsFold(query.Query),
		))
	}
	if query.JobCategory != "" {
		queryBuilder.Where(job.JobCategory(query.JobCategory))
	}
	if query.CityInfo != "" {
		queryBuilder.Where(job.CityInfo(query.CityInfo))
	}
	if len(query.CityList) > 0 {
		queryBuilder.Where(job.CityInfoIn(query.CityList...))
	}
	total, err := queryBuilder.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	entJobs, err := queryBuilder.
		Limit(query.PageSize).
		Offset((query.Page - 1) * query.PageSize).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return r.entJobsToModelJobs(entJobs), total, nil
}

func (r *JobDAO) entToModel(entJob *ent.Job) *model.Job {
	return &model.Job{
		ID:          entJob.ID,
		Title:       entJob.Title,
		Company:     entJob.Company,
		Description: entJob.Description,
		Requirement: entJob.Requirement,
		JobCategory: entJob.JobCategory,
		CityInfo:    entJob.CityInfo,
		RecruitType: entJob.RecruitType,
		PublishTime: entJob.PublishTime,
		Code:        entJob.Code,
		CityList:    entJob.CityList,
	}
}

func (r *JobDAO) entJobsToModelJobs(entJobs []*ent.Job) []*model.Job {
	jobs := make([]*model.Job, len(entJobs))
	for i, entJob := range entJobs {
		jobs[i] = r.entToModel(entJob)
	}
	return jobs
}
