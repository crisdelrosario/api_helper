# api_helper
API Helper

### Example 1:
```
api.New("http://your-host/path-to-api",api.ContentTypeJSON)
api.Get("/client",nil)

auth := api.SetBasicAuth("user","password")
api.Get("/client",auth)
```

### Example 2:

```
type Person struct {
  Name string `json:"name"`
  Age int `json:"age"`
}

person := Person {
  Name: "John Doe",
  Age: 100,
}

response := api.Post("/client",person,auth)
```
