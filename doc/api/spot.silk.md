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
    },
    {
      "id": 6,
      "name": "Brooklyn Bridge Park - Pier 1",
      "photo_url": ""
    },
    {
      "id": 7,
      "name": "Brookyn Bridge Park Greenway",
      "photo_url": ""
    },
    {
      "id": 8,
      "name": "Pierrepont Playground",
      "photo_url": ""
    },
    {
      "id": 9,
      "name": "Brooklyn Heights Promenade Garden 2",
      "photo_url": ""
    },
    {
      "id": 10,
      "name": "Hillside Dog Park",
      "photo_url": ""
    },
    {
      "id": 11,
      "name": "Jack the Horse Tavern",
      "photo_url": ""
    },
    {
      "id": 12,
      "name": "Bargemusic",
      "photo_url": ""
    },
    {
      "id": 13,
      "name": "Shake Shack",
      "photo_url": ""
    },
    {
      "id": 14,
      "name": "Under The Brooklyn Bridge",
      "photo_url": ""
    },
    {
      "id": 15,
      "name": "Brooklyn Ice Cream Factory",
      "photo_url": ""
    },
    {
      "id": 16,
      "name": "Brooklyn Bridge Park - Pier 5",
      "photo_url": ""
    },
    {
      "id": 17,
      "name": "Pier 5 Soccer Fields",
      "photo_url": ""
    },
    {
      "id": 18,
      "name": "Juliana's Pizza",
      "photo_url": ""
    },
    {
      "id": 19,
      "name": "The River Café",
      "photo_url": ""
    },
    {
      "id": 20,
      "name": "Ample Hills Creamery",
      "photo_url": ""
    },
    {
      "id": 21,
      "name": "Vineapple Cafe",
      "photo_url": ""
    },
    {
      "id": 22,
      "name": "Sushi Gallery",
      "photo_url": ""
    },
    {
      "id": 23,
      "name": "Celebrate Brooklyn - Pier 1",
      "photo_url": ""
    },
    {
      "id": 24,
      "name": "Pier 15",
      "photo_url": ""
    },
    {
      "id": 25,
      "name": "Dellarocco's",
      "photo_url": ""
    },
    {
      "id": 26,
      "name": "Tutt Cafe",
      "photo_url": ""
    },
    {
      "id": 27,
      "name": "Brooklyn Bridge Park Boathouse",
      "photo_url": ""
    },
    {
      "id": 28,
      "name": "Luke's Lobster",
      "photo_url": ""
    },
    {
      "id": 29,
      "name": "Gran Eléctrica",
      "photo_url": ""
    },
    {
      "id": 30,
      "name": "South Street Seaport",
      "photo_url": ""
    }
  ]
}
```

## `GET /api/spot`

Gets tour's spots.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"
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