package pagination

import "gorm.io/gorm"

type Pagination struct {
	Limit  *int
	Offset *int
}

func Apply(db *gorm.DB, p Pagination) *gorm.DB {
	limit := 100
	offset := 0

	if p.Limit != nil && *p.Limit > 0 {
		limit = *p.Limit
	}

	if p.Offset != nil && *p.Offset > 0 {
		offset = (*p.Offset - 1) * limit
	}
	db = db.Limit(limit).Offset(offset)
	return db
}
