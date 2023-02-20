package request

import "shop-api/src/model"

type CategoryRequest struct {
	NamaCategory string `json:"nama_category"`
}

func (r *CategoryRequest) MapRequest() *model.Category {
	return &model.Category{
		NamaCategory: r.NamaCategory,
	}
}
