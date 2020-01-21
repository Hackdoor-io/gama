package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/speps/go-hashids"
  "net/http"
  "os"
  "strconv"
)

func makeHashId(id string, target string) string {

  strId, _ := strconv.Atoi(id)

  var targetSalt string;

  switch target {
    case "articles":
      targetSalt = "ARTICLES_SALT"
      break
    case "tags":
      targetSalt = "TOPICS_SALT"
      break
    case "authors":
      targetSalt = "AUTHORS_SALT"
      break
    case "patrons":
      targetSalt = "PATRONS_SALT"
      break
    default:
      panic("Unknown target: " + target)
  }

  hashIdLength, _ := strconv.Atoi(os.Getenv("HASHID_LENGTH"))

  hd := hashids.NewData()
  hd.Salt = os.Getenv(targetSalt)
  hd.MinLength = hashIdLength
  h, _ := hashids.NewWithData(hd)
  encodedHashId, _ := h.Encode([]int{strId})

  return encodedHashId
}

func main() {

  port := os.Getenv("PORT")
  hackdoor := os.Getenv("HACKDOOR_ENDPOINT")

  var dynamicRoutes = [4]string{"articles", "tags", "authors", "patrons"}

  for _, route := range dynamicRoutes {

    fmt.Printf("Created route /%s\n", route)

    http.HandleFunc("/" + route + "/{id:[0-9]+}/{path:(\\w+?_?-?\\d?)+}", func (w http.ResponseWriter, r *http.Request) {
      pathVars := mux.Vars(r)
      id := pathVars["id"]
      path := pathVars["path"]
      hashId := makeHashId(id, route)

      redirectPath := hackdoor + "/" + route + "/" + hashId + "/" + path

      http.Redirect(w, r, redirectPath, http.StatusPermanentRedirect)
    })
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("[ROUTER] -> %s\n", r.URL.Path)
    http.Redirect(w, r, hackdoor + "?source=gama", http.StatusPermanentRedirect)
  })

  fmt.Printf("Running Gama at http://localhost:%s\n", port)
  http.ListenAndServe(":" + port, nil)
}
