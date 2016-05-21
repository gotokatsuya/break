# Spot API

## `POST /api/spot`

Creates spot.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"

```
{
  "spot_logs": [
      {
          "visit_time": 123456789,
          "lat": 40.7,
          "lng": -74
      }
  ]
}
```

===

* Status: `201`

* Content-Type: `application/json; charset=utf-8`

```
null
}
```

## `GET /api/spot`

Gets user's spots.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"
* ?start_visit_time=123456788&end_visit_time=123456790
* ?tour_id=1
===

* Status: `200`

* Content-Type: `application/json; charset=utf-8`

```
{
  "instances": [
    {
      "id": 1,
      "name": "Brooklyn Bridge Park - Pier 2",
      "photo_url": ""
    },
    {
      "id": 2,
      "name": "Brooklyn Bridge Park - Pier 4",
      "photo_url": ""
    },
    {
      "id": 3,
      "name": "Brooklyn Heights Promenade",
      "photo_url": ""
    },
    {
      "id": 4,
      "name": "Brooklyn Bridge Park",
      "photo_url": ""
    },
    {
      "id": 5,
      "name": "Blue Marble Ice Cream",
      "photo_url": ""
    }
  ]
}
```
