# API Endpoints for Timeslots

<- [Back to All endpoints](../../../README.md)

## createHour

*/api/hour* - POST

> **User Must be logged in as an Admin**
>
>startTime :string
>
>endTime :string
>
>dayOfWeek :int [0 - 6]

```json
{
    "id": 0,
    "startTime": "7:00",
    "endTime": "10:00",
    "dayOfWeek": 0
} 
```

---

### getHour

*/api/hour/{id}* - GET

> id is the id of the hour object that is being fetched

```json
{
    "id": 1,
    "startTime": "7:00",
    "endTime": "10:00",
    "dayOfWeek": 0
}
```

---

### getHours

*/api/hours* - GET

```json
[
    {
        "id": 1,
        "startTime": "7:00",
        "endTime": "10:00",
        "dayOfWeek": 0
    },
    {
        "id": 2,
        "startTime": "7:09",
        "endTime": "5:00",
        "dayOfWeek": 2
    }
]
```

---

### getHoursByDay

*/api/hour/day/{dayOfWeek}* - GET

> dayOfWeek is an integer between 0-6

```json
[
    {
        "id": 1,
        "startTime": "7:00",
        "endTime": "10:00",
        "dayOfWeek": 0
    }
]
```

---

### deleteHour

*/api/hour/{id}* - DELETE

> id is the id of the hour object that is being deleted

```json
{
    "result": "success"
}
```

---

### getUsersByHour

*/api/hour/availability/{id}* - GET

> id is the id of the hour object

```json

```
