package main
import ( 
	"fmt"
	"net/http"
	"log"
  "encoding/json"
  "github.com/gorilla/mux"
)
//App struct
type App struct {
  ID        string   `json:"id,omitempty"`
  Appname   string   `json:"Appname,omitempty"`
  Disabled  bool     `json:"disabled,omitempty"`
  Versions  *Versions `json:"versions,omitempty"`
}

//Versions struct
type Versions struct {
  Version  string `json:"version,omitempty"`
  Disabled bool   `json:"Disabled,omitempty"`
}
var apps []App

func main() {
  fmt.Println("App served on http://localhost:8000")
  //populate the array
  apps = append(apps, App{ID: "1", Appname: "MDC", Disabled: false, Versions: &Versions{Version: "1.1.1", Disabled: false}})
  apps = append(apps, App{ID: "2", Appname: "Integreatly", Disabled: false, Versions: &Versions{Version: "1.0.1", Disabled: false}})
  apps = append(apps, App{ID: "3", Appname: "RHMAP", Disabled: true, Versions: &Versions{Version: "4.6.2", Disabled: true}})

  // File server for public directroy
  http.Handle("/", http.FileServer(http.Dir("./public")))
  
  router := mux.NewRouter()
  router.HandleFunc("/apps", GetAllApps).Methods("GET")
  router.HandleFunc("/apps/{id}", GetApp).Methods("GET")
  router.HandleFunc("/apps/{id}", UpdateApp).Methods("PUT")
  
  log.Fatal(http.ListenAndServe("localhost:8000", router))
  
}

//GetAllApps get
func GetAllApps(w http.ResponseWriter, r *http.Request) {
  fmt.Println("http://localhost:8000/apps")
  json.NewEncoder(w).Encode(apps)
}

//GetApp get
func GetApp(w http.ResponseWriter, r *http.Request) {
  fmt.Println("http://localhost:8000/apps/{id} GET")
  params := mux.Vars(r)
  for _, item := range apps {
      if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
      }
  }
  json.NewEncoder(w).Encode(&App{})
}

//UpdateApp put
func UpdateApp(w http.ResponseWriter, r *http.Request) {
  fmt.Println("http://localhost:8000/apps/{id} PUT")
}
