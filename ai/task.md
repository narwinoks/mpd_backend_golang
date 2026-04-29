# Role: Act as an Expert Golang & GORM Developer.
## Context
I am building an Enterprise application using Clean Architecture. I need to create dynamic database filters (similar to Laravel's when, where, like, and between) using GORM Scopes.

Your Task:
I will provide you with a Filter/Request Struct. Your job is to generate the corresponding GORM Scopes function.

Strict Rules for Generation:

The Laravel "When" Logic: You MUST check if the parameter is not a "Zero Value" (e.g., != "" for string, != 0 for int/float, or != nil for pointers) BEFORE applying the Where clause. If it is empty/zero, simply return db.

ILIKE (Fuzzy Search): If I specify a field should be "ILIKE", wrap the value with % (e.g., "%"+val+"%") and use PostgreSQL ILIKE.

Exact Match: If I specify "Exact", use the standard column = ?.

BETWEEN (Dates/Numbers): If I provide a start and end parameter, check if BOTH are provided, then use db.Where("column BETWEEN ? AND ?", start, end).

Custom Method: If I specify "Custom", I will provide the raw SQL or logic, and you should wrap it in a .Where() or appropriate GORM clause.

Example for Custom Method:
If I say: "Search is Custom (ILIKE on name and mrn)"
You generate:
if filter.Search != "" {
    db = db.Where("name ILIKE ? OR mrn ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
}

Signature: The function must return func(db *gorm.DB) *gorm.DB so it can be chained inside .Scopes().

Here is my Filter Struct:

Go
type Request struct {
Page     int `form:"page,default=1"`
Paginate int `form:"paginate,default=10"`
Search   string `form:"search,default=''"`
}

pagination.go
Please generate the GORM Scope function named FilterPatient(filter PatientFilter) based on the rules above.

