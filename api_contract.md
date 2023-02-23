# API CONTRACT 
## USER 

- Login 

Url :
```sh
{{url}}/login
```

Request :
```sh
{
    "email"     : "ammar@gmail.com",
	"password"  : "12345",
}
```
Response Success :
```sh
{
    "code": 200,
    "message": "Login user succeed",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NH0.njwxZfCyfpSfQH9Xk5v2rcvjBSkibWVtkjlXixh2V_o",
        "user": {
            "id": 4,
            "email": "ammar@gmail.com",
            "name": "ammar",
            "role": "admin",
            "CreatedAt": "2023-02-22T15:56:16.591+07:00",
            "UpdatedAt": "2023-02-22T15:56:16.591+07:00"
        }
    }
}
```

Response Not Valid Password :
```sh
{
    "code": 401,
    "message": "Input data error",
    "data": {
        "errors": "password invalid"
    }
}
```
Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Register 

Url :
```sh
{{url}}/admin/register
{{url}}/user/register
```

Request :
```sh
{
    "name"      : "ammar",
    "email"     : "ammar@gmail.com",
	"password"  : "12345",
}
```

Response Success :
```sh
{
    "code": 201,
    "message": "Create new user succeed",
    "data": {
        "id": 4,
            "email": "ammar@gmail.com",
            "name": "ammar",
            "role": "admin",
            "CreatedAt": "2023-02-22T15:56:16.591+07:00",
            "UpdatedAt": "2023-02-22T15:56:16.591+07:00"
    }
}
```

Response Not Valid Password :
```sh
{
    "code": 400,
    "message": "Input data required",
    "data": {
        "errors": "invalid data"
    }
}
```
Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Create Monster 


Url :
```sh
{{url}}/create-monster
```

Request Header :

| Key Header  | Acceptable Value
| ------------- | ------------- |
| Authorization | Headers |

Request :
```sh
   {
    "name":"lugia",
    "typeId":["1", "3"],
    "description":"Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
    "statistics":{
        "hp":100,
        "def":77,
        "attack":78,
        "speed":88
    },
    "kind":"Lizard"
   }
```

Response Success :
```sh
{
    "code": 201,
    "message": "Create new user succeed",
    "data": "success to create monster"
}
```

Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Update Monster 

Url :
```sh
{{url}}/create-monster
```

Request Header :

| Key Header  | Acceptable Value
| ------------- | ------------- |
| Authorization | Headers |

Request :
```sh
   {
    "name":"lugia",
    "description":"Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
    "statistics":{
        "hp":100,
        "def":77,
        "attack":78,
        "speed":88
    },
    "kind":"Lizard",
    "typeId":"1",
    "typeIdOld":"2"
   }
```

Response Success :
```sh
{
    "code": 200,
    "message": "success updated monster",
    "data": "succes to update"
}
```

Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Delete Monster 


Url :
```sh
{{url}}/create-monster
```

Request Header :

| Key Header  | Acceptable Value
| ------------- | ------------- |
| Authorization | Headers |

Request Query Param :

| Key Query  | 
| ------------- |
| monster_id     | 
| type_id      | 

Request Query Param :

| Key Query  | 
| ------------- |
| monster_id     | 
| type_id      | 

Response Success :
```sh
{
    "code": 200,
    "message": "success deleted monster",
    "data": "success to delete"
}
```

Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Get All Monster 


Url :
```sh
{{url}}/getall-monsters
```

Request Query Param :

| Key Query  | 
| ------------- |
| sort     | 
| order      | 
| id      | 
| types (array query param)      | 


Response Success :
```sh
{
    "code": 200,
    "message": "Success",
    "data": [
        {
            "id": 15,
            "name": "munila",
            "description": "Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
            "statistics": {
                "hp": 99,
                "def": 77,
                "speed": 88,
                "attack": 78
            },
            "kind": "Leaf",
            "type_id": 2,
            "type_name": "grass",
            "CreatedAt": "2023-02-24T01:15:02.921+07:00",
            "UpdatedAt": "2023-02-24T01:15:02.921+07:00"
        },
        {
            "id": 15,
            "name": "munila",
            "description": "Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
            "statistics": {
                "hp": 99,
                "def": 77,
                "speed": 88,
                "attack": 78
            },
            "kind": "Leaf",
            "type_id": 3,
            "type_name": "iron",
            "CreatedAt": "2023-02-24T01:15:02.921+07:00",
            "UpdatedAt": "2023-02-24T01:15:02.921+07:00"
        },
        {
            "id": 14,
            "name": "lugia",
            "description": "Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
            "statistics": {
                "hp": 100,
                "def": 77,
                "speed": 88,
                "attack": 78
            },
            "kind": "Lizard",
            "type_id": 1,
            "type_name": "fire",
            "CreatedAt": "2023-02-24T01:14:33.212+07:00",
            "UpdatedAt": "2023-02-24T01:14:33.212+07:00"
        },
        {
            "id": 14,
            "name": "lugia",
            "description": "Sed quia soluta voluptatem animi qui. Qui rem cumque dolore et autem error et voluptatum. Sunt eius nulla perferendis in enim non similique accusantium. Voluptatem dolore iusto sequi labore et dignissimos. Corporis suscipit fuga tempora perspiciatis eius libero.",
            "statistics": {
                "hp": 100,
                "def": 77,
                "speed": 88,
                "attack": 78
            },
            "kind": "Lizard",
            "type_id": 3,
            "type_name": "iron",
            "CreatedAt": "2023-02-24T01:14:33.212+07:00",
            "UpdatedAt": "2023-02-24T01:14:33.212+07:00"
        },
    ]
}
```

Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```

- Get All Types 


Url :
```sh
{{url}}/getall-types
```

Response Success :
```sh
{
    "code": 200,
    "message": "Success",
    "data": [
        {
            "id": 1,
            "name": "fire"
        },
        {
            "id": 2,
            "name": "grass"
        },
    ]
}
```

Response Internal Server Error :
```sh
{
    "code": 500,
    "message": "Internal Server Error",
    "data": {
        "errors": "internal server error"
    }
}
```





