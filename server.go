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
  DoorStop int `json:DoorStop`
  Lights Lights `json:Lights`
}

type Lights struct {
  LightExt int `json:LightExt`
  LightFL  int `json:LightFL`
  LightFR  int `json:LightFR`
  LightML  int `json:LightML`
  LightMM  int `json:LightMM`
  LightMR  int `json:LightMR`
  LightRL  int `json:LightRL`
  LightRM  int `json:LightRM`
  LightRR  int `json:LightRR`
}

func defaultConnection(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to the default Page!")
  fmt.Println("Default page requested.")
}

func getLightStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Lights Data requested.")
  LightLabel := r.Header.Get("Light")
  fmt.Println("Lights : " + LightLabel)
  val := GetLightSpecificValue(LightLabel)
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(val)
}

// Lights value based on label
func GetLightSpecificValue(label string) int {
  switch label {
  case "Light_Ext":
    return data.Lights.LightExt
  case "Light_F_L":
    return data.Lights.LightFL
  case "Light_F_R":
    return data.Lights.LightFR
  case "Light_M_L":
    return data.Lights.LightML
  case "Light_M_M":
    return data.Lights.LightMM
  case "Light_M_R":
    return data.Lights.LightMR
  case "Light_R_L":
    return data.Lights.LightRL
  case "Light_R_M":
    return data.Lights.LightRM
  case "Light_R_R":
    return data.Lights.LightRR
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
  case "Light_F_L":
    data.Lights.LightFL = value
    if (data.Lights.LightFL == value) {
      return true
    }
    return false
  case "Light_F_R":
    data.Lights.LightFR = value
    if (data.Lights.LightFR == value) {
      return true
    }
    return false
  case "Light_M_L":
    data.Lights.LightML = value
    if (data.Lights.LightML == value) {
      return true
    }
    return false
  case "Light_M_M":
    data.Lights.LightMM = value
    if (data.Lights.LightMM == value) {
      return true
    }
    return false
  case "Light_M_R":
    data.Lights.LightMR = value
    if (data.Lights.LightMR == value) {
      return true
    }
    return false
  case "Light_R_L":
    data.Lights.LightRL = value
    if (data.Lights.LightRL == value) {
      return true
    }
    return false
  case "Light_R_M":
    data.Lights.LightRM = value
    if (data.Lights.LightRM == value) {
      return true
    }
    return false
  case "Light_R_R":
    data.Lights.LightRR = value
    if (data.Lights.LightRR == value) {
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
  writeDB()
}


func getDoorStatus(w http.ResponseWriter, r *http.Request) {
  val := data.Door
  fmt.Println("Door Data requested : ", val)
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(val)
}

// Command = OPEN | CLOSE
func openCloseDoor(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Door Status Update.")
  Command := r.Header.Get("Command")
  if Command == "OPEN" {
    if (data.DoorStop == 1) {
      data.DoorStop = 0
      fmt.Fprint(w, "Door Lock removed.")
    }
    if data.Door != 1 {
      data.Door = 1
      fmt.Fprint(w, "Opening Door.")
    } else {
      fmt.Fprint(w, "Door Already Open.")
    }
  } else if Command == "CLOSE" {
    if (data.DoorStop == 1) {
      data.DoorStop = 0
      fmt.Fprint(w, "Door Lock removed.")
    }
    if data.Door != 0 {
      data.Door = 0
      fmt.Fprint(w, "Closing Door.")
    } else {
      fmt.Fprint(w, "Door Already Closed.")
    }
  } else {
    fmt.Fprint(w, "Unknown Command requested.")
  }
  writeDB()
}

// Status = LOCK | UNLOCK
func stopDoor(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Locking Door Position.")
  Status := r.Header.Get("Status")
  if Status == "LOCK" {
    data.DoorStop = 1
    fmt.Fprint(w, "LOCK Command requested.")
  } else if Status == "UNLOCK" {
    data.DoorStop = 0
    fmt.Fprint(w, "UNLOCK Command requested.")
  }
  writeDB()
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

func writeDB() {
  file, _ := json.MarshalIndent(data, "", " ")
  _ = ioutil.WriteFile("db.json", file, 0644)
}

func handleRequests() {
  router := mux.NewRouter()

  router.HandleFunc("/Lights", getLightStatus).Methods("GET")
  router.HandleFunc("/Lights", setLightStatus).Methods("PUT")
  router.HandleFunc("/Door", getDoorStatus).Methods("GET")
  router.HandleFunc("/Door", openCloseDoor).Methods("PUT")
  router.HandleFunc("/DoorStop", stopDoor).Methods("PUT")

  http.HandleFunc("/", defaultConnection)
  log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
  readDb()
  fmt.Println("ASE Service running.")
  handleRequests()
}
