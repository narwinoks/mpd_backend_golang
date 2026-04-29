# Role: Senior Golang Developer

## Context
1. Make scope pagination
2. Goal: Avoid making 2 separate queries (count and find) in `role_repository`. Retrieve both the paginated data and the total count using a single query via a GORM scope.

## Implementation Detail: Single Query Pagination (Window Function)

To achieve pagination and total count in a single query with PostgreSQL and GORM, we utilize the `COUNT(*) OVER()` window function.

### 1. Update Pagination Scope (`pkg/pagination/base_request.go`)

We modify the existing `PaginateScope` to include the `COUNT(*) OVER() AS total_count` selection.

```go
func PaginateScope(req Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.Page <= 0 {
			req.Page = 1
		}
		if req.Paginate <= 0 {
			req.Paginate = 10
		}

		offset := (req.Page - 1) * req.Paginate
		
		return db.Select("*, COUNT(*) OVER() AS total_count").
			Offset(offset).
			Limit(req.Paginate)
	}
}
```

### 2. Implementation in Repository (`internal/modules/master/repository/role/role_repository.go`)

We use a struct to capture the additional `total_count` column.

```go
type RoleWithCount struct {
	role.Role
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *roleRepositoryImpl) FindAll(req pagination.Request) ([]role.Role, int64, error) {
	var results []RoleWithCount
	var roles []role.Role
	var total int64 = 0

	err := r.db.Model(&role.Role{}).
		Scopes(pagination.PaginateScope(req)).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			roles = append(roles, res.Role)
		}
	}

	return roles, total, nil
}
```
