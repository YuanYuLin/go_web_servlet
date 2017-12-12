package gpio

import "encoding/json"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "ops_log"

type json_msg_t struct {
    Status		int	`json:"status"`
    Version		int	`json:"version"`
    Data		interface{} `json:"data"`
}

func Init() {
	V1_init()
}

func responseWithJsonV1(w http.ResponseWriter, code int,  data interface{}) {
    json_msg := json_msg_t { Status:1, Version:1, Data:data }
    response, _ := json.Marshal(json_msg)
    ops_log.Debug(0x01, string(response))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func GetIoGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    port, _	:= strconv.Atoi(params["port"])
    pin, _	:= strconv.Atoi(params["pin"])
    api, _	:= params["api_version"]
    switch api {
    case "v1":
        responseWithJsonV1(w, http.StatusOK, V1_GetIoGpio(port, pin))
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func GetIoListGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    //port, _       := strconv.Atoi(params["port"])
    //pin, _        := strconv.Atoi(params["pin"])
    api, _        := params["api_version"]
    switch api {
    case "v1":
        responseWithJsonV1(w, http.StatusOK, V1_GetIoListGpio())
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func PutIoGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    port, _	:= strconv.Atoi(params["port"])
    pin, _	:= strconv.Atoi(params["pin"])
    api, _      := params["api_version"]
    switch api {
    case "v1":
        var gpio io_gpio_v1_t
        _ = json.NewDecoder(r.Body).Decode(&gpio)
        responseWithJsonV1(w, http.StatusOK, V1_PutIoGpio(port, pin, gpio))
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

