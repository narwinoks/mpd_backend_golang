role
senior golang development

context
add uuid in response all response api 

task 1
install package
https://github.com/gofrs/uuid

task 2
set middleware ResponseIDMiddleware

task 3 
example response
{
"rc": "500",
"message": "Internal Server Error",
"errors": "dial tcp: i/o timeout",
"request_id": "c1f72e34-9c8a-4b9f-8d2e-1a2b3c4d5e6f"
}