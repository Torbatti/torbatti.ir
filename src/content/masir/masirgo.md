---
title: مسیر برنامه نویسی زبان گو
---

```go  
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
```

# سلام
# سلام
# سلام
# سلام
# سلام
# سلام