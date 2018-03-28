package gpio

import "encoding/json"
import "ops_log"
import "ops_uds"

type io_gpio_v1_t struct {
    Io_type             string  `json:"type"`
    Port                int     `json:"port"`
    Pin                 int     `json:"pin"`
    Value               int     `json:"value"`
    Direction           string  `json:"direction"`
    Comment             string  `json:"comment"`
    Name                string  `json:"name"`
    Unix_timestamp      string  `json:"unix_timestamp"`
}

//var io_list_gpio []io_gpio_v1_t

func V1_init() {
/*
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:0, Pin:7, Value:0, Direction:"out", Comment:"", Name:"RY2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:0, Pin:26, Value:0, Direction:"out", Comment:"", Name:"DO_Y1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:0, Pin:27, Value:0, Direction:"out", Comment:"", Name:"LED2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:12, Value:0, Direction:"out", Comment:"", Name:"DO_Y0", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:13, Value:0, Direction:"out", Comment:"", Name:"LED3", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:14, Value:0, Direction:"out", Comment:"", Name:"DO_Y2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:15, Value:0, Direction:"out", Comment:"", Name:"LED1(PWD)", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:16, Value:0, Direction:"in", Comment:"", Name:"DI_X1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:17, Value:0, Direction:"in", Comment:"", Name:"DI_X2", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:28, Value:0, Direction:"in", Comment:"", Name:"DI_X0", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:1, Pin:29, Value:0, Direction:"out", Comment:"", Name:"LED4", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:2, Pin:1, Value:0, Direction:"out", Comment:"", Name:"DO_Y3", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:3, Pin:19, Value:0, Direction:"out", Comment:"", Name:"RY1", Unix_timestamp:"0"})
    io_list_gpio = append(io_list_gpio, io_gpio_v1_t{Io_type:"gpio", Port:3, Pin:21, Value:0, Direction:"in", Comment:"", Name:"DI_X3", Unix_timestamp:"0"})
*/
}

type req_get_gpio_t struct {
    Io_type	string	`json:"type"`
    Port	int	`json:"port"`
    Pin		int	`json:"pin"`
}

/*
 * Io_type : 
 *    1 : GPIO
 */
 /*
type res_get_gpio_t {
    Io_type             string	`json:"type"`
    Port                int   `json:"port"`
    Pin                 int   `json:"pin"`
    Value               int   `json:"value"`
    Direction           string  `json:"direction"`
    Comment             string  `json:"comment"`
    Name                string  `json:"name"`
    Unix_timestamp      string  `json:"unix_timestamp"`
}
*/

func V1_GetIoGpio(port int, pin int) (io_gpio_v1_t){
    var req req_get_gpio_t
    var res io_gpio_v1_t
    req.Io_type = "gpio"
    req.Port = port
    req.Pin = pin
    req_bytes, req_err := json.Marshal(req)
    if req_err != nil {
	    ops_log.Error(0x01, req_err.Error())
    }
    _, _, msg_status, res_len, res_bytes := ops_uds.SendAndRecvByMsg(1, 1, uint16(len(req_bytes)), req_bytes)
    if msg_status != 0 {
	    ops_log.Error(0x01, "message status %d", msg_status)
    } else {
	    ops_log.Debug(0x01, "%s", string(res_bytes))
	    for i:=0;i<int(res_len);i++ {
		    ops_log.Debug(0x01, "%x,", res_bytes[i])
	    }
	    res_err := json.Unmarshal(res_bytes, &res)
	    if res_err != nil {
		    ops_log.Error(0x01, res_err.Error())
	    }
    }

    return res
}

func V1_GetIoListGpio() ([]io_gpio_v1_t){
    var req req_get_gpio_t
    var res []io_gpio_v1_t
    req.Io_type = "gpio"
    req_bytes, req_err := json.Marshal(req)
    if req_err != nil {
	    ops_log.Error(0x01, req_err.Error())
    }
    _, _, msg_status, _, res_bytes := ops_uds.SendAndRecvByMsg(1, 3, uint16(len(req_bytes)), req_bytes)
    if msg_status != 0 {
	    ops_log.Error(0x01, "message status %d", msg_status)
    } else {
	    res_err := json.Unmarshal(res_bytes, &res)
	    if res_err != nil {
		    ops_log.Error(0x01, res_err.Error())
	    }
    }
    return res
}

func V1_PutIoGpio(port int, pin int, gpio io_gpio_v1_t) (io_gpio_v1_t){
    //var req io_gpio_v1_t
    var res io_gpio_v1_t
    req_bytes, req_err := json.Marshal(gpio)
    if req_err != nil {
	    ops_log.Error(0x01, req_err.Error())
    }
    _, _, msg_status, _, res_bytes := ops_uds.SendAndRecvByMsg(1, 2, uint16(len(req_bytes)), req_bytes)
    if msg_status != 0 {
	    ops_log.Error(0x01, "message status %d", msg_status)
    } else {
	    res_err := json.Unmarshal(res_bytes, &res)
	    if res_err != nil {
		    ops_log.Error(0x01, res_err.Error())
	    }
    }
    return res
}

