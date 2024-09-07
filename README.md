
# TaskHub

Task management API built with go, you can make project as base of your activity, put some task inside it. 

Using gRPC for the user authentication microservice

This project utilize several packages, including

- #### go-fiber v2 

- #### viper

- #### GORM

- #### protobuffer

- #### redis

- #### go-air



## Installation

Run the docker containers with docker-compose

```bash
  docker-compose up --build
```
    
## Features

- User Management microservice
- Organize tasks with project grouping



## API Reference

#### Get all projects by user

```http
  GET /api/{userid}/projects
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userid` | `string` | **Required**. ID for the user logged in |

#### Get a project by user

```http
  GET /api/{userid}/projects/{projectid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `projectid`   | `string` | **Required**. ID for the user logged in and project id |

#### Create project

```http
  POST /api/{userid}/projects
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userid` | `string` | **Required**. ID for the user logged in |

```json
{
   "user_id":1,
   "title":"this is title",
   "description": "this is description"
}
```

#### Update project

```http
  PUT /api/{userid}/projects/{projectid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `projectid`   | `string` | **Required**. ID for the user logged in and project id |

```json
{
   "title":"this is title revised",
   "description": "this is description"
}
```
#### Delete project

```http
  DELETE /api/{userid}/projects/{projectid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `projectid`   | `string` | **Required**. ID for the user logged in and project id |



#### Get all tasks by user

```http
  GET /api/{userid}/tasks
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userid` | `string` | **Required**. ID for the user logged in |


#### Get a task by user

```http
  GET /api/{userid}/task/{taskid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `taskid`   | `string` | **Required**. ID for the user logged in and task id |

#### Create task

```http
  POST /api/{userid}/tasks
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userid` | `string` | **Required**. ID for the user logged in |

```json
{
  "user_id":1,
  "project_id":1,
  "title":"this is title",
  "description":"this is description",
  "status":"not_completed",
  "priority":"normal"
}
```

#### Update task

```http
  PUT /api/{userid}/task/{taskid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `taskid`   | `string` | **Required**. ID for the user logged in and task id |


```json
{
  "user_id":1,
  "project_id":1,
  "title":"this is title revised",
  "description":"this is description",
  "status":"not_completed",
  "priority":"normal"
}
```

#### Delete task

```http
  Delete /api/{userid}/task/{taskid}
```

| Parameter               | Type     | Description                |
| :---------------------- | :------- | :------------------------- |
| `userid`, `taskid`   | `string` | **Required**. ID for the user logged in and task id |



## Contributing

Contributions are always welcome!

Pull requests are aso welcome. For major changes, please open an issue first to discuss what you would like to change.


Made by revandpratama with ❤️


## License

[MIT](https://choosealicense.com/licenses/mit/)

