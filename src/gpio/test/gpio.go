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

var io_list_gpio []io_gpio_t

func Init() {
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:0, Pin:7, Value:0, Direction:"out", Comment:"", Name:"RY2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:0, Pin:26, Value:0, Direction:"out", Comment:"", Name:"DO_Y1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:0, Pin:27, Value:0, Direction:"out", Comment:"", Name:"LED2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:12, Value:0, Direction:"out", Comment:"", Name:"DO_Y0", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:13, Value:0, Direction:"out", Comment:"", Name:"LED3", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:14, Value:0, Direction:"out", Comment:"", Name:"DO_Y2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:15, Value:0, Direction:"out", Comment:"", Name:"LED1(PWD)", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:16, Value:0, Direction:"in", Comment:"", Name:"DI_X1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:17, Value:0, Direction:"in", Comment:"", Name:"DI_X2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:28, Value:0, Direction:"in", Comment:"", Name:"DI_X0", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:1, Pin:29, Value:0, Direction:"out", Comment:"", Name:"LED4", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:2, Pin:1, Value:0, Direction:"out", Comment:"", Name:"DO_Y3", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:3, Pin:19, Value:0, Direction:"out", Comment:"", Name:"RY1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_t{Io_type:"gpio", Port:3, Pin:21, Value:0, Direction:"in", Comment:"", Name:"DI_X3", Unix_timestamp:"0"})
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
        responseWithJson(w, http.StatusOK, V1_GetIoGpio(msg, ConvertListToV1(io_list_gpio), port, pin))
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

