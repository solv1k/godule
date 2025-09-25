package query

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm/utils"
)

// Query parameters
type Params struct {
	Page    PageParams
	Sort    []SortField
	Filters map[string]string
}

// Pagination parameters
type PageParams struct {
	Number int
	Size   int
}

// Sort field
type SortField struct {
	Field string
	Order string // "asc", "desc"
}

// Parse query parameters
func Parse(c *fiber.Ctx) Params {
	params := Params{
		Filters: make(map[string]string),
	}

	// Extract page number and page size
	pageNumberStr := futils.CopyString(c.Query("page[number]", "1"))
	pageSizeStr := futils.CopyString(c.Query("page[size]", "10"))

	params.Page.Number, _ = strconv.Atoi(pageNumberStr)
	params.Page.Size, _ = strconv.Atoi(pageSizeStr)

	if params.Page.Number <= 0 {
		params.Page.Number = 1
	}

	if params.Page.Size <= 0 {
		params.Page.Size = 10
	}

	if params.Page.Size > 100 {
		params.Page.Size = 100
	}

	// Extract sort fields
	if sort := futils.CopyString(c.Query("sort")); sort != "" {
		for _, field := range strings.Split(sort, ",") {
			field = strings.TrimSpace(field)
			if strings.HasPrefix(field, "-") {
				params.Sort = append(params.Sort, SortField{
					Field: strings.TrimPrefix(field, "-"),
					Order: "desc",
				})
			} else {
				params.Sort = append(params.Sort, SortField{
					Field: field,
					Order: "asc",
				})
			}
		}
	}

	// Extract filters
	queries := c.Queries()
	for key, value := range queries {
		safeKey := futils.CopyString(key)
		safeValue := futils.CopyString(value)

		if strings.HasPrefix(safeKey, "filter[") && strings.HasSuffix(safeKey, "]") {
			filterKey := safeKey[7 : len(safeKey)-1] // remove "filter[" and "]"
			params.Filters[filterKey] = safeValue
		}
	}

	return params
}

// Returns SQL string for sorting, e.g. "field1 asc, field2 desc, ..."
func (p Params) BuildSortSQL(avaliableSorts []string) string {
	var clauses []string
	for _, sort := range p.Sort {
		if !utils.Contains(avaliableSorts, sort.Field) {
			continue
		}
		clauses = append(clauses, sort.Field+" "+sort.Order)
	}
	return strings.Join(clauses, ", ")
}

// Returns offset calculated by page number and page size
func (p Params) Offset() int {
	return (p.Page.Number - 1) * p.Page.Size
}

// Returns limit calculated by page size
func (p Params) Limit() int {
	return p.Page.Size
}
