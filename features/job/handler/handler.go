package handler

import (
	"encoding/json"
	"gdt-api/features/job"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jobController struct {
	srv job.JobService
}

func New(js job.JobService) job.JobHandler {
	return &jobController{
		srv: js,
	}
}

// GetJobList implements job.JobHandler
func (*jobController) GetJobList() echo.HandlerFunc {
	return func(c echo.Context) error {

		resp, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to get job data")
		}
		log.Println(resp)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		// log.Println(string(body))

		var positions []Position
		err = json.Unmarshal(body, &positions)
		if err != nil {
			log.Println(err)
		}
		log.Println(positions)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			// "data":    positions,
		})
	}
}

// GetJobDetail implements job.JobHandler
func (*jobController) GetJobDetail() echo.HandlerFunc {
	panic("unimplemented")
}
