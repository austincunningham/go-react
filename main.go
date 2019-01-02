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
  GlobalDisableMessage string `json:"globalDisableMessage,omitempty"`
  Versions  *Versions `json:"versions,omitempty"`
}

//Versions struct
type Versions struct {
  Version  string `json:"version,omitempty"`
  Disabled bool   `json:"disabled,omitempty"`
  DisableMessage string `json:"disableMessage,omitempty"`
}
var apps []App

func main() {
  fmt.Println("App served on http://localhost:8000")
  //populate the array hard coded at the moment 
  apps = append(apps, App{
    ID: "1",
    Appname: "MDC", 
    Disabled: false, 
    GlobalDisableMessage: "Disabled", 
    Versions: &Versions{
      Version: "1.1.1", 
      Disabled: false}})
  apps = append(apps, App{
    ID: "2", 
    Appname: "Integreatly", 
    Disabled: false,
    GlobalDisableMessage: "Disabled", 
    Versions: &Versions{
      Version: "1.0.1", 
      Disabled: false,
      DisableMessage: "Disabled by admin"}})
  apps = append(apps, App{
    ID: "3", 
    Appname: "RHMAP", 
    Disabled: true,
    GlobalDisableMessage: "Disabled", 
    Versions: &Versions{
      Version: "4.6.2", 
      Disabled: true,
      DisableMessage: "Disabled by admin"}})

  // File server for public directroy
  http.Handle("/", http.FileServer(http.Dir("./goreact/public")))
  
  // REST End points
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
  params := mux.Vars(r)
  fmt.Println("http://localhost:8000/apps/{id} GET id :", params["id"])
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
  defer r.Body.Close()
  params := mux.Vars(r)
  var app App
  fmt.Println("http://localhost:8000/apps/{id} PUT id :", params["id"])
  //fmt.Println("response.Body ==> ",json.NewDecoder(r.Body).Decode(&app))

  if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
  }
  respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
  for _, item := range apps {
    if item.ID == params["id"] {
      // logic go in here for handeling the update e.g. {2 Integreatly false 0xc00000c320}
       fmt.Println(item.Versions.Disabled)
       fmt.Println(item.Versions.DisableMessage)
       fmt.Println(item.Versions)
       // incoming content
       fmt.Println(app.Versions.DisableMessage)
       // hard coded when the PUT endpoint hit
      //  var disabled = &item.Versions.Disabled 
      //  *disabled = true 
      //  var disabledmessage = &item.Versions.DisableMessage
      //  *disabledmessage = "Disabled by your admin"
      //  fmt.Println(item.Versions.Disabled)
      //  fmt.Println(item.Versions.DisableMessage)
    }
  }
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

