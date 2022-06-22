## Usage

- `make build`
- `make tests`
- `make run`
- `make clean`

## DOC API

- `Get User`

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

- `Get User By Id`

  URL: localhost:8000/user-show?id=1

  Method: GET

  Response :
  {
  "id": "1",
  "name": "Eka"
  }

- `Store User`

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

- `Update User`

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

- `Delete User`

  URL: localhost:8000/user-delete?id=1

  Method: DELETE

  Response :
  {
  "message": "berhasil hapus data"
  }
