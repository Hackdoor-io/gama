package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "github.com/speps/go-hashids"
  "log"
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

  if os.Getenv("ENV") == "development" {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
  }

  port := os.Getenv("PORT")
  hackdoor := os.Getenv("HACKDOOR_ENDPOINT")

  http.HandleFunc("/{target:articles|tags|authors|patrons}/{id:[0-9]+}/{path:(\\w+?_?-?\\d?)+}", func (w http.ResponseWriter, r *http.Request) {
    pathVars := mux.Vars(r)
    target := pathVars["target"]
    id := pathVars["id"]
    path := pathVars["path"]
    hashId := makeHashId(id, target)

    redirectPath := hackdoor + "/" + target + "/" + hashId + "/" + path

    http.Redirect(w, r, redirectPath, http.StatusPermanentRedirect)
  })

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("[ROUTER] -> %s\n", r.URL.Path)
    http.Redirect(w, r, hackdoor + "?source=gama", http.StatusPermanentRedirect)
  })

  fmt.Printf("Running Gama at http://localhost:%s\n", port)
  http.ListenAndServe(":" + port, nil)
}
