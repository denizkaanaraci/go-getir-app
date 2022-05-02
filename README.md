# Getir Case Challenge

https://go-getir-app.herokuapp.com/

### Requirements

go 1.18

### Configuration
    Database and server configs can be changed in config.yml file
    

### Build:
        go build -o go-getir ./cmd/go-getir
    
### run: 
        ./go-getir

### docker

    build:
        docker build -t go-getir .

    run: 
        docker run -d -p 8080:8080 --name go-getir go-getir
        
    interactive run: 
        docker run -it --rm -p 8080:8080 --name go-getir go-getir

### heroku
    docker build -t go-getir .
    
    heroku container:push web -a go-getir-app

    heroku container:release web -a go-getir-app

## Endpoints

### Get Records from MongoDB

**Endpoint:** `/record`

**Method:** `POST`

**Example URL:** https://go-getir-app.herokuapp.com/record

**Request Example:** 

```json
{
  "startDate": "2016-01-26",
  "endDate": "2018-02-02",
  "minCount": 2700,
  "maxCount": 3000

}
```
**Response Example:**
```json
{
  "code":0,
  "msg":"Success",
  "records": [
    {
      "key":"TAKwGc6Jr4i8Z487",
      "createdAt":"2017-01-28T01:22:14.398Z",
      "totalCount":2800
    },
    {
      "key":"NAeQ8eX7e5TEg7oH",
      "createdAt":"2017-01-27T08:19:14.135Z",
      "totalCount":2900
    }
  ]
}
```
### Set to In-Memory Storage

**Endpoint:** `/in-memory`

**Method:** `POST`

**Example URL:** https://go-getir-app.herokuapp.com/in-memory

**Request Example:**

```json
{
  "key": "active-tabs",
  "value": "getir"
}
```
**Response Example:**

200 OK
```json
{
  "key": "active-tabs",
  "value": "getir"
}
```

404 Not Found
```json
{
  "key": "active-tabs",
  "value": "getir"
}
```

### Get from In-Memory Storage

**Endpoint:** `/in-memory?key={key}`

**Method:** `GET`

**Params:** `key: key`


**Example URL:** https://go-getir-app.herokuapp.com/in-memory?key=active-tabs

**Response Example:**

200 OK
```json
{
  "key": "active-tabs",
  "value": "getir"
}
```

404 Not Found
```json
{
  "error": "Key not found"
}
```
