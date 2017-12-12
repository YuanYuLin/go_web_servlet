package gpio

import "encoding/json"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "uds"

type json_msg_t struct {
    Status		int	`json:"status"`
    Version		int	`json:"version"`
    Data		interface{} `json:"data"`
}

type io_gpio_t struct {
    Io_type             string  `json:"type"`
    Port                int     `json:"port"`
    Pin                 int     `json:"pin"`
    Value               int     `json:"value"`
    Direction           string  `json:"direction"`
    Comment             string  `json:"comment"`
    Name                string  `json:"name"`
    Unix_timestamp      string  `json:"unix_timestamp"`
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func GetIoGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    port, _	:= strconv.Atoi(params["port"])
    pin, _	:= strconv.Atoi(params["pin"])
    api, _	:= params["api_version"]
    msg := json_msg_t { 1, 0, nil }
    switch api {
    case "v1":
	msg.Version = 1
	msg.Data = V1_GetIoGpio(port, pin)
	responseWithJson(w, http.StatusOK, msg)
    default:
        responseWithJson(w, http.StatusOK, msg)
    }
}

func GetIoListGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    //port, _       := strconv.Atoi(params["port"])
    //pin, _        := strconv.Atoi(params["pin"])
    api, _        := params["api_version"]
    msg := json_msg_t { 1, 0, nil }
    switch api {
    case "v1":
	msg.Version = 1
        responseWithJson(w, http.StatusOK, V1_GetIoListGpio(msg, ConvertListToV1(io_list_gpio)))
    default:
        responseWithJson(w, http.StatusOK, msg)
    }
}

func PutIoGpio(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    port, _	:= strconv.Atoi(params["port"])
    pin, _	:= strconv.Atoi(params["pin"])
    api, _      := params["api_version"]
    msg := json_msg_t { 1, 0, nil }
    switch api {
    case "v1":
	msg.Version = 1
        var io_gpio_v1 io_gpio_v1_t
        _ = json.NewDecoder(r.Body).Decode(&io_gpio_v1)
        responseWithJson(w, http.StatusOK, V1_PutIoGpio(msg, ConvertListToV1(io_list_gpio), port, pin, io_gpio_v1))
    default:
        responseWithJson(w, http.StatusOK, msg)
    }
}

