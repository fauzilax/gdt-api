package job

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
}

type JobHandler interface {
	GetJobList() echo.HandlerFunc
	GetJobDetail() echo.HandlerFunc
}
type JobService interface {
	GetJobList() (Core, error)
	GetJobDetail() (Core, error)
}
type JobData interface {
	GetJobList() (Core, error)
	GetJobDetail() (Core, error)
}
