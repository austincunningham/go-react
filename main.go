package main
import ( 
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)
type App struct {
  ID        string   `json:"id,omitempty"`
  Appname   string   `json:"Appname,omitempty"`
  Disabled  bool     `json:"disabled,omitempty"`
  Versions  *Versions `json:"versions,omitempty"`
}
type Versions struct {
  Version  string `json:"version,omitempty"`
  Disabled bool   `json:"Disabled,omitempty"`
}
var apps []App

func main() {
  fmt.Println("http://localhost:8000")
  //populate the array
  apps = append(apps, App{ID: "1", Appname: "MDC", Disabled: false, Versions: &Versions{Version: "1.1.1", Disabled: false}})
  apps = append(apps, App{ID: "2", Appname: "Integreatly", Disabled: false, Versions: &Versions{Version: "1.0.1", Disabled: false}})
  apps = append(apps, App{ID: "3", Appname: "RHMAP", Disabled: true, Versions: &Versions{Version: "4.6.2", Disabled: true}})


  http.Handle("/", http.FileServer(http.Dir("./public")))
  
  //router.HandleFunc("/apps", GetAllApps).Methods("GET")
  http.HandleFunc("/apps", GetAllApps)
  http.HandleFunc("/apps/{id}", GetApp)
  http.HandleFunc("/apps/{id}", UpdateApp)
  
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
  
}

func GetAllApps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(apps)
}

func GetApp(w http.ResponseWriter, r *http.Request) {}
func UpdateApp(w http.ResponseWriter, r *http.Request) {}
