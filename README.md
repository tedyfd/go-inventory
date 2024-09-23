# go-inventory

Create User
POST http://{{base_url}}:{{port}}/v1/users
body 
{
    "username": "testing",
    "password": "123",
    "name": "tes"
}

Get User
GET http://{{base_url}}:{{port}}/v1/users
headers
Authorization : ApiKey {{api_key}}

Create Category
POST http://{{base_url}}:{{port}}/v1/category
body
{
    "name": "Computer",
    "description": "IT Inventory"
}

Delete Category
DELETE http://{{base_url}}:{{port}}/v1/category/{{category_id}}

Create Product
POST http://{{base_url}}:{{port}}/v1/product
body
{
    "name": "MSI",
    "quantity": 5,
    "category_id": "{{category_id}}"
}