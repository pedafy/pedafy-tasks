# API Routes

> Here will be defined all the API's routes using the Richardson Maturity Model, more info [here](https://martinfowler.com/articles/richardsonMaturityModel.html).

### Tasks

| Action | Method | Route | Response code |
| ------ | ------ | ----- | ------------- |
| Get all tasks | `GET` | `/tasks` | `200`, `204`, `403`, `503` |
| Get all tasks by order | `GET` | `/tasks?sort=status|new|due` | `200`, `204`, `403`, `503` |
| Get all by id kind | `GET` | `/tasks/{id_kind}/{id:[0-9]+}` | `200`, `204`, `403`, `503` |
| Create a task | `PUT` | `/task` | `201`, `403`, `422`, `503` |
| Modify a task | `POST` | `/task/{id:[0-9]+}` | `200`,`403`, `410`, `422`, `503` |
| Archive a task | `POST` | `/task/archive/{id:[0-9]+}` | `200`, `403`, `410`, `503` |

### Status
| Action | Method | Route | Resp code |
| ------ | ------ | ----- | --------- |
| Get all status | `GET` | `/status` | `200`, `204`, `403`, `503` |
| Get status by ID | `GET` | `/status/{id:[0-9]+}` | `200`, `403`, `410`, `503` |
| Get status by name | `GET` | `/status/{name}` | `200`, `403`, `410`, `503` |
