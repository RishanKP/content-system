# Endpoints

## User Service

This service returns information about the users.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| List users | `GET` | `/users/` |
| Get user by Id | `GET` | `/users/{id}` |
| Insert user | `POST` | `/users/` |
| Delete user | `DELETE` | `/users/{id}` |

## Content Service

Used to manage content actions.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| List stories | `GET` | `/contents/` |
| Get story by Id | `GET` | `/contents/{id}` |
| Get latest stories | `GET` | `/contents/new` |
| Insert story | `POST` | `/contents/` |
| Update story | `PUT` | `/contents/{id}` |
| Delete story | `DELETE` | `/contents/{id}` |
| Get most viewed stories | `GET` | `/contents/top?param=reads&limit=` |
| Get most liked stories | `GET` | `/contents/top?param=likes&limit=` |
| Bulk upload stories by CSV| `POST` | `/contents/uploadByCSV` |

## Interactions Service

Used to manage interaction. eg: like,read actions

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Save interaction | `POST` | `/interactions/` |
| Remove an interaction | `POST` | `/interactions?action=remove` |
| Get interactions by story ID | `GET` | `/interactions/{storyID}` |
| Get most viewed stories | `GET` | `/interactions/analytics/top?param=reads&limit=` |
| Get most liked stories | `GET` | `/interactions/analytics/top?param=likes&limit=` |
