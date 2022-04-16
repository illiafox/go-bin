# GoBIN - pastebin copy written in pure go with memcached and postgreSQL
![Model](https://user-images.githubusercontent.com/61962654/163591990-4daf2f1f-8eea-493d-94f7-5e40040b8a9f.png)
![Screenshot](https://user-images.githubusercontent.com/61962654/163588613-e1fc5cd1-023c-44b2-a331-5f35b80d871c.png)

---
## Requirements
* **Go:** `1.18`
* **PostgreSQL:** `14.2`
* **Memcached:** `1.6.15`

## docker-compose
No additional settings are required, just start and use
```shell
docker-compose up
```
### Dockerfile
```shell
docker build . --tag=gobin
docker run gobin
```
---

## Building 

```shell
git clone https://github.com/illiafox/go-bin.git
cd go-bin/cmd/server
go build -o gobin
```
```shell
./gobin
```
---
## Endpoints:
### `:8080` main page
### `:5433` postgreSQL

---

### Custom config path
```shell
./gobin -config conf/app/config.toml
```

### HTTP mode
```shell
./gobin -http
```
### Update config from environment variables
```shell
./gobin -env
```
view `.env` template


### Skip config reading, environment only
```shell
./gobin -noread
```
---
## SQL

### Table

```sql
key  char(16) primary key,
data json 
```

### Migrations
docker-compose do this automatically
```sql
SOURCE migrate-up.sql
```
```sql
SOURCE migrate-down.sql
```



