# Middleware Test Developer

This Backend service is an example of dockerized fibergo service with testing. This service is not intended to be use for production use.

## How to Run

 1. Make sure docker is installed
 2. Copy `.env.example`	and paste it in repository directory. Rename it to `.env` and fill in or replace the value declared
 3. Run
	 ```
	 make docker.run
	 ```
4. Wait for docker to finish building and running
5. The service api is ready to be used
6. To stop the service run
	```
	make docker.stop 
	```

## How to Run Test

 1. Make sure docker is installed
 2. Copy `.env.example`	and paste it in repository directory. Rename it to `.env` and fill in or replace the value declared.
 3. Run
	 ```
	 make docker.test
	 ```
4. Wait for docker to finish building and running the test
5. The app should close itself after testing

## API Docs
Base URL : `/api`

### GET `/employee`
Fetch all employee in database
#### Response Body Example
**200 OK**
```
{
	"data":  [
		{
			"id":  1,
			"firstName":  "gabriel1",
			"lastName":  "gabriel2",
			"email":  "gabriel@gmail.com",
			"hireDate":  "2024-04-23T00:00:00Z"
		},
		{
			"id":  2,
			"firstName":  "gabriel1",
			"lastName":  "gabriel2",
			"email":  "gabriel@gmail.com",
			"hireDate":  "2024-04-23T00:00:00Z"
		},
	],
	"message":  "Successfully fetch all employee data"
}
```

### GET `/employee/{id}`
Fetch existing employee data based on `{id}` provided
#### Path Variable
| Path | Description | Accepted Values |
|--|--| -- |
| id | Specify employee id that needed to be fetch | number |

#### Response Body Example
**200 OK**
`/employee/1`
```
{
	"data":  {
			"id":  1,
			"firstName":  "gabriel1",
			"lastName":  "gabriel2",
			"email":  "gabriel@gmail.com",
			"hireDate":  "2024-04-23T00:00:00Z"
		},
	"message":  "Successfully fetched employee data"
}
```
### POST`/employee`
Create New Employee
#### Request Body
| Body Data | Description | Required? | Datatype | Example
|--|--| -- | -- | -- |
| firstName | First Name of the employee | ✅ |  string | "John"
| lastName | Last Name of the employee | ✅ |  string | "Doe"
| email| Email of the employee. Must be an email format | ✅ |  string | "john@email.com"
| hireDate | Hiring Date of the employee. Time will be ignored by database | ✅ |  datetime | "2024-04-23T15:00:00.000Z"
```
{
	"firstName"  :  "John",
	"lastName"  :  "Doe",
	"email"  :  "john@email.com",
	"hireDate"  :  "2024-04-23T15:00:00.000Z"
}
```

#### Response Body Example
**200 OK**
```
{
	"data":  {
			"id":  2,
			"firstName"  :  "John",
			"lastName"  :  "Doe",
			"email"  :  "john@email.com",
			"hireDate"  :  "2024-04-23T15:00:00.000Z"
		},
	"message":  "Successfully created new employee"
}
```
### PUT`/employee/{id}`
Update existing employee based on `{id}` provided
#### Path Variable
| Path | Description | Accepted Values |
|--|--| -- |
| id | Specify employee id that needed to be updated| number |
#### Request Body
| Body Data | Description | Required? | Datatype | Example
|--|--| -- | -- | -- |
| firstName | First Name of the employee | ❌ |  string | "John"
| lastName | Last Name of the employee | ❌ |  string | "Doe"
| email| Email of the employee. Must be an email format | ❌ |  string | "john@email.com"
| hireDate | Hiring Date of the employee. Time will be ignored by database | ❌ |  datetime | "2024-04-23T15:00:00.000Z"
```
{
	"firstName"  :  "Doe",
	"lastName"  :  "John",
	"email"  :  "die@email.com",
	"hireDate"  :  "2024-05-23T15:00:00.000Z"
}
```

#### Response Body Example
**200 OK**
`/employee/2`
```
{
	"data":  {
			"id":  2,
			"firstName"  :  "Doe",
			"lastName"  :  "John",
			"email"  :  "die@email.com",
			"hireDate"  :  "2024-05-23T15:00:00.000Z"
		},
	"message":  "Successfully updated employee data"
}
```
### Delete`/employee/{id}`
Update existing employee based on `{id}` provided
#### Path Variable
| Path | Description | Accepted Values |
|--|--| -- |
| id | Specify employee id that needed to be fetch | number |

#### Response Body Example
**200 OK**
`/employee/1`
```
{
	"data":  {
			"id":  1,
			"firstName":  "gabriel1",
			"lastName":  "gabriel2",
			"email":  "gabriel@gmail.com",
			"hireDate":  "2024-04-23T00:00:00Z"
		},
	"message":  "Successfully deleted employee data"
}
```
