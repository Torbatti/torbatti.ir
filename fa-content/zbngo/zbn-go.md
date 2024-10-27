---
title: زبان برنامه نویسی گو
---

<!-- ```go  
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Salam Mosafer!"))
    })
    http.ListenAndServe(":3000", r)
}
``` -->

- [Go By Example](https://gobyexample.com/)



# کلوچه ها | Cookies
## Create
## Set
## Update
## Delete



# محیط | Env
## تابع : os.Setenv(key,value)
## تابع : os.Getenv(key)
## فایل : .env
```bash
# نصب | install
go get github.com/joho/godotenv
```
```go
// File: main.go

// Load the .env file in the current directory
err := godotenv.Load()
// or
err := godotenv.Load(".env")
```




# MiddleWares

# JWT

# Form
```go
func send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	msg := &Message{
		Email:	 r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
    }
}
```

# Json
## Marshal
## Unmarshal

# html/template

# Templ
- [a-h/temple](https://templ.guide)

# Sqlc
- [Sqlc Github](https://github.com/sqlc-dev/sqlc)

# Sqlite

# Redka
- [Redka Github](https://github.com/nalgeon/redka)

# Chi
- [Chi Router](https://go-chi.io)

# web frameworks
## Fiber
## Gin
