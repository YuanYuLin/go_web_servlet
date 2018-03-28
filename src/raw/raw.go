package raw

import "encoding/json"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "ops_log"
import "ops_uds"
import "io/ioutil"

type json_msg_t struct {
    Status		int	`json:"status"`
    Version		int	`json:"version"`
    Data		interface{} `json:"data"`
}

func Init() {
}

func responseWithJsonV1(w http.ResponseWriter, code int,  data interface{}) {
    json_msg := json_msg_t { Status:1, Version:1, Data:data }
    response, _ := json.Marshal(json_msg)
    ops_log.Debug(0x01, string(response))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func Get(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _	:= params["api_version"]
    switch api {
    case "v1":
	bodyBytes := []byte{0, 0}
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(0), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func Put(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _      := params["api_version"]
    switch api {
    case "v1":
        bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(len(bodyBytes)), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
	//_, _, msg_status, _, res_bytes := ops_uds.SendAndRecvByMsg(fn, cmd, uint16(len(bodyBytes)), bodyBytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _	:= params["api_version"]
    switch api {
    case "v1":
        bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(len(bodyBytes)), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func Post(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _	:= params["api_version"]
    //decoder := json.NewDecoder(r.Body)
    switch api {
    case "v1":
        bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(len(bodyBytes)), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func Patch(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _	:= params["api_version"]
    switch api {
    case "v1":
        bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(len(bodyBytes)), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

func Head(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn, _	:= (strconv.Atoi(params["fn"]))
    cmd, _	:= (strconv.Atoi(params["cmd"]))
    api, _	:= params["api_version"]
    switch api {
    case "v1":
        bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	_, _, _, _, res_bytes := ops_uds.SendAndRecvByMsg(uint8(fn), uint8(cmd), uint16(len(bodyBytes)), bodyBytes)
        responseWithJsonV1(w, http.StatusOK, res_bytes)
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

