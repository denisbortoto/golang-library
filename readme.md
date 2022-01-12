# 📝 About

The library api is a project utilized to practice the content studied about the Golang programming language.

## 🚀 Technologies

This project was developed using the following technologies:

- Golang
- Postgresql
- Docker

## How to download and run:
```bash
    $git clone https://github.com/denisbortoto/golang-library.git
    $cd .\library\cmd
```
- Install [docker](https://docs.docker.com/get-docker/) and open it
- Run **docker-compose up**
- Apply migrations with: $goose postgres "user=postgres dbname=library sslmode=disable password=postgres" up


- Navigate to .\library\cmd and run $go run server.go