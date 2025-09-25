package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
)

func PaginationMeta(total int64, params query.Params) map[string]interface{} {
	return fiber.Map{
		"total": total,
		"page":  params.Page.Number,
		"size":  params.Page.Size,
		"pages": (total + int64(params.Page.Size) - 1) / int64(params.Page.Size),
	}
}
