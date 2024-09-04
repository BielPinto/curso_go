

## Running the serve.go

```
server# go run cmd/server/main.go 
```




## db Sqlite3
```
To access the sqlite3 database follow the command
server# sqlite3 ./data/database.db
```
List Records
select * from quotation;

```
Exit sqlite3
.exit
---
Table structure
PRAGMA table_info(quotation);
