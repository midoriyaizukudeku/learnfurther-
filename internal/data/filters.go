package data

import "main/internal/validator"

type Filters struct {
	Page         int
	Page_size    int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "value must be reater than zero")
	v.Check(f.Page <= 100, "page", "value must be smaller than 100")
	v.Check(f.Page_size > 0, "page_size", "value must be greater than 0")
	v.Check(f.Page_size < 20, "page_size", "value must be smaller than 20")
	v.Check(validator.PermittedValues(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
