# Tour API

## `POST /api/tour`

Creates tour.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"

```
{
  "name": "tour",
  "photo_url": "",
  "spot_ids": [1,2,3,4,5]
}
```

===

* Status: `201`

* Content-Type: `application/json; charset=utf-8`

```
null
```

## `GET /api/tour`

Gets user's tours.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"
* ?user_id=1
* ?lat=40.66&lng=-74

===

* Status: `200`

* Content-Type: `application/json; charset=utf-8`

```
{
  "instances": [
    {
      "id": 1,
      "name": "tour",
      "photo_url": ""
    }
  ]
}
```
