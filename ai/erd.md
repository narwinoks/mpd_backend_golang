make table  erd result must be Mermaid to draw io

context
default table field is

id int 32,
uuid string 36,
is_active bool,
profile_id fk from profile table,
external_code string 20,
created_at timestamp,
updated_at timestamp,
deleted_at timestamp,
created_by fk for employee table,
updated_by fk for employee table,
deleted_by fk for employee table,

make table

1.religions_m
field is
religion string 100,

2.profiles_m
province_id fk,
city_id,
subdistrict_id,
village_id,
postal_code,
email,
name,
profile,
government_name,
phone,
telp
full_address text,

3.provinces_m
province

4.cities_m
city
province_id -> fk for provinces_m

5.subdistrict_m
subdistrict
city_id fk for cities_m
provinces_id fk for provinces_m

6.genders_m
gender

7.job_categories_m
code
job_category

8.job_titles_m
code
job_title
job_category_id fk  to job_categories_m

9.employment_statuses_m
code
employee_status

10.marital_status_m 
 material_status

11.employee_group_m
employee_group

12.positions_m
position

13.employees_m
religion_id fk for religion_m
gender_id fk for gender_m
job_title_id  fk  for job_titles_m
employment_status_id fk for employment_statuses_m
full_name
identity_number
nip
npwp
birth_place varchar
birth_date date


14.employee_details_m
employee_id uniques for employee id
marital_status_id bigint [fk]
functional_position_id bigint [fk]  positions_m
structural_position_id bigint [fk]  positions_m
join_date date
resign_date date
retirement_date date

15.roles_m
role

16.users_m

username
email
password
role_id fk roles_m
employee_id fk to employees_m

17.app_modules_m
code
name 
category
sort_order

18.app_menus_m
app_module_id fk app_modules_m
parent_id fk app_menus_m
code
name
path
description
icon
sort order

19.role_modules_m
role_id fk from roles
modules_id fk from modules

20.user_modules_m
user_id fk from users_m
modules_id fk from modules_m

21.app_permission_m
permission

22.role_permission_m
role_id fk from roles
permission_id fk from permission

23.user_permission_m
user_id fk from users_m
permission_id fk from permission

24.profile_detail_m

profile_id
website
longitude
latitude
registration_date
moto

25.subdistrict_villages_m
province_id fk  provinces_m
city_id fk cities_m
district_id fk
subdistrict_village
