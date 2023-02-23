package helper

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Page  int         `query:"page" json:"page"`
	Limit int         `query:"limit" json:"limit"`
	Data  interface{} `json:"data"`
}

func GetPagination(c *fiber.Ctx) (*Pagination, error) {
	pagination := new(Pagination)
	if err := c.QueryParser(pagination); err != nil {
		if strings.Contains(err.Error(), "page") {
			return nil, errors.New("page value is invalid")
		}

		if strings.Contains(err.Error(), "limit") {
			return nil, errors.New("limit value is invalid")
		}
	}

	return pagination, nil
}

func CountOffset(page int) int {
	return (page - 1) * 10
}
