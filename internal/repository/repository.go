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

// Paginate applies pagination to the query and returns the paginated results.
func Paginate(db *gorm.DB, offset, limit int) *gorm.DB {
	return db.Offset((offset - 1) * limit).Limit(limit)
}

// CountTotal returns the total count of records for the given query.
func CountTotal(db *gorm.DB, model interface{}) (int64, error) {
	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
