package gpio

import "uds"

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

type uds_msg_gpio_t struct {
    Hdr		uds.msg_hdr_t
    Data	io_gpio_v1_t
}

func msg_to_bytes(data interface{}, data_len int) {
	buf = make([]byte, data_len)
	for i:= range data {
		buf[i] = byte(data[i])
	}
	return buf
}

func V1_GetIoGpio(port int, pin int) (io_gpio_v1_t) {
    data := io_gpio_v1_t{}
    data.Io_type = "GPIO"
    data.Port = port
    data.Pin = pin
    data.Value = 0
    data.Direction = "IN"
    data.Comment = ""
    data.Name = "N/A"
    data.Unix_timestamp = ""

    msg := uds_msg_gpio_t{}
    msg.Hdr.Data_size = len(io_gpio_v1_t)
    msg.Hdr.Fn = 1
    msg.Hdr.Cmd = 1
    msg.Hdr.Crc32 = 0
    msg.Data = gpio

    req_buf := msg_to_bytes(msg)

    res_buf := uds.SendAndRecvByBytes(req_buf)
    for _, item := range io_list {
        if item.Port == port &&
            item.Pin == pin {
	    msg.Status = 0
	    msg.Data = item
	    return msg
        }
    }
    return msg
}

func V1_GetIoListGpio(msg json_msg_t, io_list []io_gpio_v1_t) (json_msg_t){
    msg.Status = 0
    msg.Data = io_list
    return msg
}

func V1_PutIoGpio(msg json_msg_t, io_list []io_gpio_v1_t, port int, pin int, obj io_gpio_v1_t) (json_msg_t){
    msg.Status = 1
    msg.Data = obj
    obj.Port = port
    obj.Pin = pin
    for _, item := range io_list_gpio {
        if item.Port == port &&
           item.Pin == pin {
            //fmt.Printf("Port:%d, Pin:%d, value:%d\n", io_gpio.Port, io_gpio.Pin, io_gpio.Value)
            item.Value = obj.Value
            item.Comment = obj.Comment
            item.Unix_timestamp = obj.Unix_timestamp
	    msg.Status = 0
	    msg.Data = obj
	    return msg
        }
    }
    return msg
}

