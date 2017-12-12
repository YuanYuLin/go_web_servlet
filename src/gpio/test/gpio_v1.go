package gpio

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

func ConvertListToV1(io_list []io_gpio_t) ([]io_gpio_v1_t){
    var io_list_v1 []io_gpio_v1_t
    for _, item := range io_list {
        io_list_v1 = append(io_list_v1, io_gpio_v1_t{Io_type:item.Io_type,Port:item.Port,Pin:item.Pin,Value:item.Value,Direction:item.Direction,Comment:item.Comment,Name:item.Name,Unix_timestamp:item.Unix_timestamp})
    }
    return io_list_v1
}

func V1_GetIoGpio(msg json_msg_t, io_list []io_gpio_v1_t, port int, pin int) (json_msg_t){
    msg.Status = 1
    msg.Data = nil
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

