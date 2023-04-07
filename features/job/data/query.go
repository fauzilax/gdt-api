package data

import (
	"gdt-api/features/job"

	"gorm.io/gorm"
)

type jobQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) job.JobData {
	return &jobQuery{
		db: db,
	}
}

// GetJobDetail implements job.JobData
func (*jobQuery) GetJobDetail() (job.Core, error) {
	panic("unimplemented")
}

// GetJobList implements job.JobData
func (*jobQuery) GetJobList() (job.Core, error) {
	panic("unimplemented")
}
