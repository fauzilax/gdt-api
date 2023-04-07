package services

import "gdt-api/features/job"

type jobUseCase struct {
	qry job.JobData
}

func New(jb job.JobData) job.JobService {
	return &jobUseCase{
		qry: jb,
	}
}

// GetJobList implements job.JobService
func (*jobUseCase) GetJobList() (job.Core, error) {
	panic("unimplemented")
}

// GetJobDetail implements job.JobService
func (*jobUseCase) GetJobDetail() (job.Core, error) {
	panic("unimplemented")
}
