# Web Applications in Go

## 3 Essential Components
1) Handler
2) Servemux (router)
3) Server

## Servemux
- `"/"` catches all paths
- local scope the servemux for security
- longer URL >>> shorter URL; meaning order of route declaration doesn't matter
- servemux is NOT RESTful??
- Go cannot detect difference between JSON and txt, set header -> `w.Header().Set("Content-Type", "application/json")`
- get url params `id` with `id, err := strconv.Atoi(r.URL.Query().Get("id"))`
