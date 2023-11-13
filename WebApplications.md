# Web Applications in Go

## 3 Essential Components
1) Handler
2) Servemux (router)
3) Server

## Handler
- must implement ServeHTTP()
- servemux implements the Handler interface
- Handler requests are handled concurrently

## Servemux
- `"/"` catches all paths
- local scope the servemux for security
- longer URL >>> shorter URL; meaning order of route declaration doesn't matter
- Go cannot detect difference between JSON and txt, set header -> `w.Header().Set("Content-Type", "application/json")`
- get url params `id` with `id, err := strconv.Atoi(r.URL.Query().Get("id"))`
