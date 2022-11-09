package main

import (
  "os"
  "fmt"
  "log"
  "strconv"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
)

var data DB

type DB struct {
  Door int     `json:Door`
  Co   int     `json:Co`
  Lights Lights `json:Lights`
}

type Lights struct {
  LightExt int `json:LightExt`
  LightL  int `json:LightL`
  LightM  int `json:LightM`
  LightR  int `json:LightR`
}

func defaultConnection(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to the Smart Garage Welcome Page!")
  fmt.Println("Default page requested.")
}

func getLightStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Lights Data requested.")
  LightLabel := r.Header.Get("Light")
  fmt.Println("Lights : " + LightLabel)
  //val := GetLightSpecificValue(LightLabel)
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(data.Lights)
}

// Lights value based on label
func GetLightSpecificValue(label string) int {
  switch label {
  case "Light_Ext":
    return data.Lights.LightExt
  case "Light_L":
    return data.Lights.LightL
  case "Light_M":
    return data.Lights.LightM
  case "Light_R":
    return data.Lights.LightR
  default:
    return -1
  }
}

// Lights value based on label
func SetLightSpecificValue(label string, value int) bool{
  switch label {
  case "Light_Ext":
    data.Lights.LightExt = value
    if (data.Lights.LightExt == value) {
      return true
    }
    return false
  case "Light_L":
    data.Lights.LightL = value
    if (data.Lights.LightL == value) {
      return true
    }
    return false
  case "Light_M":
    data.Lights.LightM = value
    if (data.Lights.LightM == value) {
      return true
    }
    return false
  case "Light_R":
    data.Lights.LightR = value
    if (data.Lights.LightR == value) {
      return true
    }
    return false
  default:
    data.Lights.LightExt = value
    if (data.Lights.LightExt == value) {
      return true
    }
    return false
  }
  return false
}

// Light stat = 0 (Off) | 1 (On)
func setLightStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Lights Data Updating.")
  LightLabel := r.Header.Get("Light")
  Value,_ := strconv.Atoi(r.Header.Get("Value"))
  res := SetLightSpecificValue(LightLabel, Value)
  if (res) {
    fmt.Println("Operation complete.")
    fmt.Fprint(w, "Light status changed for ", LightLabel, " to ", Value)
  } else {
    fmt.Println("Operation failed.")
    fmt.Fprint(w, "Error ", LightLabel, " to ", Value)
  }
  go writeDB()
}

// Command = OPEN | STOP | CLOSE
func setDoorStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Door Status Update.")
  Command := r.Header.Get("Command")
  if Command == "OPEN" {
    if data.Door != 1 {
      data.Door = 1
      fmt.Fprint(w, "Opening Door.")
    } else {
      fmt.Fprint(w, "Door Already Open.")
    }
  } else if Command == "CLOSE" {
    if data.Door != -1 {
      data.Door = -1
      fmt.Fprint(w, "Closing Door.")
    } else {
      fmt.Fprint(w, "Door Already Closed.")
    }
  } else if Command == "STOP" {
    if data.Door != 0 {
      data.Door = 0
      fmt.Fprint(w, "Stopping Door.")
    } else {
      fmt.Fprint(w, "Door Already Stopped.")
    }
  } else {
    fmt.Fprint(w, "Unknown command.")
  }
  go writeDB()
}

// Get Door status
func getDoorStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Door Data requested : ", data.Door)
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(data.Door)
}

// Get Current stored Co sensor Value (int) 1 to 1000ppm 25 safe
// sensor ESSINIE MICS-6814 Gas Sensor Module Air Quality Detection Sensor Carbon Monoxide CO Nitrogen Dioxide NO2 Ammonia NH3 Gas Detector Module for Arduino
func getCoStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Co Data requested : ", data.Co)
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(data.Co)
}

// Set the Co Sensor value [Value = int] 1 to 1000ppm 25 safe
func setCoStatus(w http.ResponseWriter, r *http.Request) {
  Value,_ := strconv.Atoi(r.Header.Get("Value"))
  fmt.Println("CO Sensor Data Updating.", Value)
  data.Co = Value
  fmt.Fprint(w, "Sensor Value updated.")
  go writeDB()
}

// Perform On every server start
func readDb() {
  jsonFile, err := os.Open("db.json")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("DB Read complete.... ")
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &data)

  fmt.Println(data)
  defer jsonFile.Close()
}

// Perform on server sever dat change
func writeDB() {
  file, _ := json.MarshalIndent(data, "", " ")
  _ = ioutil.WriteFile("db.json", file, 0644)
}

func handleRequests() {
  router := mux.NewRouter()

  router.HandleFunc("/Lights", getLightStatus).Methods("GET")
  router.HandleFunc("/Lights", setLightStatus).Methods("PUT")
  router.HandleFunc("/Door", getDoorStatus).Methods("GET")
  router.HandleFunc("/Door", setDoorStatus).Methods("PUT")
  router.HandleFunc("/Co", getCoStatus).Methods("GET")
  router.HandleFunc("/Co", setCoStatus).Methods("PUT")

  http.HandleFunc("/", defaultConnection)
  log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
  readDb()
  fmt.Println("ASE Service running.")
  handleRequests()
}
