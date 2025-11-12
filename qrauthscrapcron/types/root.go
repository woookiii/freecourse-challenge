package types

type AddReq struct {
	URL           string   `json:"url" binding:"required"` //startWith is also binding condition, gin have various binding conditions
	CardSelector  string   `json:"cardSelector" binding:"required"`
	InnerSelector string   `json:"innerSelector" binding:"required"`
	Tag           []string `json:"tag" binding:"required"`
}

type ViewReq struct {
	URL string `form:"url" binding:"required"`
}

type UpdateReq struct {
	URL           string   `json:"url" binding:"required"`
	CardSelector  string   `json:"cardSelector" binding:"required"`
	InnerSelector string   `json:"innerSelector" binding:"required"`
	Tag           []string `json:"tag" binding:"required"`
}

type DeleteReq struct {
	URL string `form:"url" binding:"required"`
}
