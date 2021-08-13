# Go, Gin, Gorm, and MySQL

a simple crud with go and mysql

## Routes

- http://localhost:8080/ || (homepage)
- http://localhost:8080/mahasiswa || (method POST: add mahasiswa)
- http://localhost:8080/mahasiswa || (method GET: get all mahasiswas)
- http://localhost:8080/mahasiswa/{nim} || (method GET: get detail of mahasiswa)
- http://localhost:8080/mahasiswa/{nim} || (method PUT: update mahasiswa)
- http://localhost:8080/mahasiswa/{nim} || (method DELETE: delete mahasiswa)

## Data JSON

```javascript
{
     "id": 1,
     "nim": "181910305",
     "name": "Muhamad Zidan Indratama"
}
```
