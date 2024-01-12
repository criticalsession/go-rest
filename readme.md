# GO REST API

## REST API written in GO.

Bare-bones REST API written in GO with a sqlite database and JWT Authentication. Allows signup, login, CRUD operations on events, registering and unregistering.

## Install/Run

1. Download the code
2. Make sure you have GO installed on your system
3. Run `go run .`

## User Routes

### Create User

**Request**

```
POST http://localhost:8080/signup
Content-Type: application/json

{
    "email": "my@email.com",
    "password": "password"
}
```

**Response**

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:06:28 GMT
Content-Length: 87
Connection: close

{
  "message": "User created",
  "user": {
    "Id": 1,
    "Email": "my@email.com",
    "Password": "password"
  }
}
```

### Login

**Request**

```
POST http://localhost:8080/login
Content-Type: "application/json"
    
    {
        "email": "my@email.com",
        "password": "password"
    }
```

**Response**

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:07:19 GMT
Content-Length: 192
Connection: close

{
  "message": "Login successful",
  "token": "<token>"
}
```

## Event Routes

### Get Events

**Request**

`GET http://localhost:8080/events`

**Response**

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:10:44 GMT
Content-Length: 183
Connection: close

[
  {
    "Id": 1,
    "Name": "Test event",
    "Description": "A test event",
    "Location": "A test location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserId": 1,
    "Registrations": [
      {
        "Id": 1,
        "UserId": 1,
        "EventId": 1
      }
    ]
  }
]
```

### Get Event by Id

**Request**

`GET http://localhost:8080/events/1`

**Response**

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:10:33 GMT
Content-Length: 181
Connection: close

{
  "Id": 1,
  "Name": "Test event",
  "Description": "A test event",
  "Location": "A test location",
  "DateTime": "2025-01-01T15:30:00Z",
  "UserId": 1,
  "Registrations": [
    {
      "Id": 1,
      "UserId": 1,
      "EventId": 1
    }
  ]
}
```

### Create Event

**Request**

```
POST http://localhost:8080/events
Content-Type: application/json
Authorization: <token>

{
    "name" : "Test event",
    "description" : "A test event",
    "location": "A test location",
    "datetime" : "2025-01-01T15:30:00.000Z"
} 
```

**Response**

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:08:59 GMT
Content-Length: 188
Connection: close

{
  "event": {
    "Id": 1,
    "Name": "Test event",
    "Description": "A test event",
    "Location": "A test location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserId": 1,
    "Registrations": null
  },
  "message": "Event created"
}
```

### Update Event

Note: Only the user who created the event can update it.

**Request**

```
PUT http://localhost:8080/events/1
content-type: application/json
Authorization: <token>

{
    "name": "Updated test event",
    "description": "Updated test event description",
    "datetime": "2025-01-01T15:30:00Z",
    "location": "Updated test event location"
}
```

**Response**

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:14:18 GMT
Content-Length: 226
Connection: close

{
  "event": {
    "Id": 1,
    "Name": "Updated test event",
    "Description": "Updated test event description",
    "Location": "Updated test event location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserId": 0,
    "Registrations": null
  },
  "message": "Event updated"
}
```

### Delete Event

Note: Only the user who created the event can update it.

**Request**

```
DELETE http://localhost:8080/events/1
Authorization: <token>
```

**Response**

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:19:17 GMT
Content-Length: 39
Connection: close

{
  "eventId": 1,
  "message": "Event deleted"
}
```

### Register for Event

**Request**

```
POST http://localhost:8080/events/1/register
Authorization: <token>
```

**Response**

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:17:15 GMT
Content-Length: 277
Connection: close

{
  "event": {
    "Id": 1,
    "Name": "Updated test event",
    "Description": "Updated test event description",
    "Location": "Updated test event location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserId": 1,
    "Registrations": [
      {
        "Id": 0,
        "UserId": 1,
        "EventId": 1
      }
    ]
  },
  "message": "User registered to event",
  "userId": 1
}
```

### Unregister from Event

**Request**

```
DELETE http://localhost:8080/events/1/register
Authorization: <token>
```

**Response**

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 06 Jan 2024 15:17:41 GMT
Content-Length: 250
Connection: close

{
  "event": {
    "Id": 1,
    "Name": "Updated test event",
    "Description": "Updated test event description",
    "Location": "Updated test event location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserId": 1,
    "Registrations": []
  },
  "message": "User unregistered from event",
  "userId": 1
}
```

## To-Do

- [ ] Testing