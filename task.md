Act as a Staff Software Engineer. I have a large seeder.go file for a Hospital Information System (HIS) that needs to be refactored into a modular structure to avoid circular dependencies and foreign key violations.

Your Mission:
Refactor the provided code into a domain-based directory structure inside internal/database/seeders/.

1. Directory Structure:

seeder.go: The orchestrator/entry point.

helper.go: Contains createBaseModel and pointer helpers.

location_seeder.go: Logic for Provinces, Cities, Subdistricts, and Villages.

profile_seeder.go: Logic for Profiles and ProfileDetails.

master_seeder.go: Logic for Religions, Genders, Job Categories, Titles, etc.

employee_seeder.go: Logic for Employees and EmployeeDetails.

auth_seeder.go: Logic for Roles, Modules, Menus, Users, and RBAC Pivots.

2. Strict Execution Order (The Orchestrator Logic):
   You must ensure the SeedAll function executes in this exact sequence to satisfy Database Constraints:

Locations (Provinces -> Cities -> Districts -> Villages)

Profiles (Needs Location IDs)

Master References (Needs Profile ID)

Employees (Needs Master IDs & Profile ID. Note: First employee 'Winarno' created_by is nil)

Auth & RBAC (Needs Employee ID & Profile ID)

3. State Passing Strategy:

Each seeder function should return a struct or map containing the IDs created (e.g., locationIDs, masterIDs) so they can be passed to the next function in the chain.

Use GORM Transactions (db.Transaction) in the main orchestrator.

4. Code Rules:

Use 'github.com/brianvoe/gofakeit/v6' for all dummy data.

Ensure all structs follow the previously defined GORM tags and BaseModel embedding.

Maintain the _m table naming convention via TableName() methods if necessary.
