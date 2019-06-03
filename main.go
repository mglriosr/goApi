package main

import(
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Usuario struct {
  ID string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct{
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}

var personas []Usuario

func GetPersonasEndPoint(w http.ResponseWriter , req *http.Request ){
  json.NewEncoder(w).Encode(personas)
}

func GetUsuarioEndPoint(w http.ResponseWriter , req *http.Request ){
  params := mux.Vars(req)
  for _, item := range personas {
    if item.ID == params["id"]{
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Usuario{})
}

func CreateUsuarioEndPoint(w http.ResponseWriter , req *http.Request ){
  params := mux.Vars(req)
  var usuario Usuario
  _ = json.NewDecoder(req.Body).Decode(&usuario)
  usuario.ID = params ["id"]
  personas = append(personas, usuario)
  json.NewEncoder(w).Encode(personas)
}

func DeleteUsuarioEndPoint(w http.ResponseWriter , req *http.Request ){
  params := mux.Vars(req)
  for index, item := range personas {
    if item.ID == params ["id"] {
      personas = append(personas[:index], personas[index + 1:]... )
      break
    }
  }
  json.NewEncoder(w).Encode(personas)
}

func main(){
  router := mux.NewRouter()

  personas = append(personas, Usuario{ID: "1", FirstName: "Miguel", LastName: "Rios", Address: &Address{City: "Quito", State: "Pichincha"}})
  personas = append(personas, Usuario{ID: "2", FirstName: "Michelle", LastName: "Perez", Address: &Address{City: "Quito", State: "Pichincha"}})

  // Endpoints
  router.HandleFunc("/personas", GetPersonasEndPoint).Methods("GET")
  router.HandleFunc("/personas/{id}", GetUsuarioEndPoint).Methods("GET")
  router.HandleFunc("/personas/{id}", CreateUsuarioEndPoint).Methods("POST")
  router.HandleFunc("/personas/{id}", DeleteUsuarioEndPoint).Methods("DELETE")


  log.Fatal(http.ListenAndServe(":3000", router))

}