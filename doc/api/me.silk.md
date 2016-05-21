# Me API

## `POST /api/me/login`

Logins user.

* Content-Type: "application/json"

```
{
  "name": "katsuya",
  "email": "katsuya.goto@eure.jp",
  "photo_url": "https://avatars1.githubusercontent.com/u/5179593?v=3&s=460"
}
```

===

* Status: `200`

* Content-Type: `application/json; charset=utf-8`

```
{
  "instance": {
    "id": 1,
    "name": "katsuya",
    "email": "katsuya.goto@eure.jp",
    "photo_url": "https://avatars1.githubusercontent.com/u/5179593?v=3&s=460",
    "token": "test-1"
  }
}
```

## `GET /api/me`

Gets logined user.

* Content-Type: "application/json"
* Authorization: "AccessToken test-1"

===

* Status: `200`

* Content-Type: `application/json; charset=utf-8`

```
{
  "instance": {
    "id": 1,
    "name": "katsuya",
    "email": "katsuya.goto@eure.jp",
    "photo_url": "https://avatars1.githubusercontent.com/u/5179593?v=3&s=460",
    "token": "test-1"
  }
}
```