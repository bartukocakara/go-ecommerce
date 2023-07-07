package repository

import "gorm.io/gorm"

func paginateQuery(query *gorm.DB, offset, limit int) *gorm.DB {
	if offset >= 0 {
		query = query.Offset(offset)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	return query
}
