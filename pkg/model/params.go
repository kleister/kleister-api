package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int
	Offset int
}
