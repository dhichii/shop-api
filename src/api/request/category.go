package request

import "shop-api/src/model"

type Category struct {
	NamaCategory string `json:"nama_category"`
}

func (r *Category) MapRequest() *model.Category {
	return &model.Category{
		NamaCategory: r.NamaCategory,
	}
}
