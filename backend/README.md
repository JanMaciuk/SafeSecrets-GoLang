## Backend

### Uruchamianie
`go run main.go`

### Push na heroku
`git subtree push --prefix backend heroku master`

### Scieżki API

### Pobieranie - `GET /?key={id}:{key}`

* Zwraca
  ```
    {
      "secret": {
        "id": 7,
        "content": "TRESC",
        "usagesLeft": 5
      }
    }
  ```

### Dodawanie - `POST /`
* Body - `{"content": "TRESC", "uses": 0}`
* Headers - Content-Type: 'application/json'
* Zwraca
    ```
  {
      "key": "KLUCZ",
      "removalKey": "KLUCZ DO USUNIĘCIA",
      "secret": {
        "id": 9,
        "content": "ZAKODOWANA TRESC (nie obchodzi frontend)",
        "usagesLeft": 5
      }
    }
  ```
  
### Usuwanie - `DELETE /?key={id}:{removalKey}`