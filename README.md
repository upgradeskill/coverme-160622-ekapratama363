[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=7994602&assignment_repo_type=AssignmentRepo)

# golang-starter

![CI](https://github.com/lumochift/golang-starter/workflows/CI/badge.svg)

Basic starting point for a new golang project.

## Creating a new project

- Rename `cmd/golang-starter` folder
- Rename `APP_NAME` in `Makefile`
- Rename `go.mod` with proper username and repository names
- Update imports accordingly

## Usage

- `make build`
- `make tests`
- `make run`
- `make clean`

## DOC API

- Get User

  URL: localhost:8000/user

  Method: GET

  Response :
  [
  {
  "id": "1",
  "name": "Eka"
  },
  {
  "id": "2",
  "name": "eka1"
  }
  ]

- Get User By Id

  URL: localhost:8000/user-show?id=1

  Method: GET

  Response :
  {
  "id": "1",
  "name": "Eka"
  }

- Store User

  URL: localhost:8000/user-add

  Method: POST

  Body:
  {
  "id": "2",
  "name": "eka1"
  }

  Response :
  {
  "message": "success create data"
  }

- Update User

  URL: localhost:8000/user-update?id=1

  Method: PUT

  Body:
  {
  "name": "eka update"
  }

  Response :
  {
  "id": "1",
  "name": "eka update"
  }

- Delete User

  URL: localhost:8000/user-delete?id=1

  Method: DELETE

  Response :
  {
  "message": "berhasil hapus data"
  }
