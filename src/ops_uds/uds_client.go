package ops_uds

import "fmt"
import "os"
import "bytes"
import "encoding/binary"
import "net"
import "encoding/json"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "io"
import "ops_log"

type json_msg_t struct {
    Status		int	`json:"status"`
    Version		int	`json:"version"`
    Data		interface{} `json:"data"`
}

type msg_hdr_t struct {
	Data_size	uint16	`json:"size"`
	Fn		uint8	`json:"fn"`
	Cmd		uint8	`json:"cmd"`
	Status		uint8	`json:"status"`
	Crc32		uint32	`json:"crc32"`
}

type msg_t struct {
	Hdr		msg_hdr_t	`json:"hdr"`
	Data		[]byte `json:"data"`
}

func WriteBytes(dst io.Writer, src interface{}) (error) {
	we := binary.Write(dst, binary.LittleEndian, src)
	if we != nil {
		ops_log.Error(0x01, "error: ")
		ops_log.Error(0x01, we.Error())
	}
	return we
}

func ReadBytes(src io.Reader, dst interface{}) (error) {
	re := binary.Read(src, binary.LittleEndian, dst)
	if re != nil {
		ops_log.Error(0x01, "error: ")
		ops_log.Error(0x01, re.Error())
	}
	return re
}

func SendAndRecvByMsg(fn uint8, cmd uint8, data_size uint16, data []byte) (uint8, uint8, uint8, uint16, []byte){
	req_msg := msg_t{}
	req_msg.Hdr.Data_size = data_size
	req_msg.Hdr.Fn = fn
	req_msg.Hdr.Cmd = cmd
	req_msg.Hdr.Status = 0
	req_msg.Hdr.Crc32 = 0
	req_msg.Data = data
	req_buf := new(bytes.Buffer)
	we := WriteBytes(req_buf, req_msg.Hdr.Data_size)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	we = WriteBytes(req_buf, req_msg.Hdr.Fn)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	we = WriteBytes(req_buf, req_msg.Hdr.Cmd)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	we = WriteBytes(req_buf, req_msg.Hdr.Status)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	we = WriteBytes(req_buf, req_msg.Hdr.Crc32)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	we = WriteBytes(req_buf, req_msg.Data)
	if we != nil {
		ops_log.Error(0x01, we.Error())
	}
	req_bytes := req_buf.Bytes()
	res_bytes := SendAndRecvByBytes(req_bytes)
	res_buf := bytes.NewReader(res_bytes)
	for i:=0;i<100;i++ {
		ops_log.Debug(0x01, "%x,", res_bytes[i])
	}
	res_msg := msg_t{}
	re := ReadBytes(res_buf, &res_msg.Hdr.Data_size) // idx : 0, len : 2
	if re != nil {
		ops_log.Error(0x01, re.Error())
	}
	re = ReadBytes(res_buf, &res_msg.Hdr.Fn) // idx : 2, len : 1
	if re != nil {
		ops_log.Error(0x01, re.Error())
	}
	re = ReadBytes(res_buf, &res_msg.Hdr.Cmd) // idx : 3, len : 1
	if re != nil {
		ops_log.Error(0x01, re.Error())
	}
	re = ReadBytes(res_buf, &res_msg.Hdr.Status) // idx : 4, len : 1
	if re != nil {
		ops_log.Error(0x01, re.Error())
	}
	re = ReadBytes(res_buf, &res_msg.Hdr.Crc32) // idx : 5, len : 4
	if re != nil {
		ops_log.Error(0x01, re.Error())
	}
	data_index := int(9) // message header size
	ops_log.Debug(0x01, "data index = %d\n", data_index, len(res_bytes))
	res_msg.Data = make([]byte, int(res_msg.Hdr.Data_size))

	ops_log.Debug(0x01, "fn: %x", int(res_msg.Hdr.Fn))
	ops_log.Debug(0x01, "cmd: %x", int(res_msg.Hdr.Cmd))
	ops_log.Debug(0x01, "status: %x", int(res_msg.Hdr.Status))
	ops_log.Debug(0x01, "data size: %d", int(res_msg.Hdr.Data_size))
	ops_log.Debug(0x01, "Crc32: %d", int(res_msg.Hdr.Crc32))
	for i:=0;i<int(res_msg.Hdr.Data_size);i++ {
		res_msg.Data[i] = res_bytes[data_index + i]
		ops_log.Debug(0x01, "i:%x\n", int(i), byte(res_msg.Data[i]))
	}

	return res_msg.Hdr.Fn, res_msg.Hdr.Cmd, res_msg.Hdr.Status, res_msg.Hdr.Data_size, res_msg.Data
}

func SendAndRecvByBytes(req_buf []byte)([]byte) {
	max_client := 5

	res_buf := make([]byte, 0x1000)
	uds_type := "unixgram"
	client_path := ""
	server_path := "/var/run/uds.www"
	for i := 0; i < max_client; i++ {
		client_path = fmt.Sprintf("%s.cli_%x", server_path, i)
		_, err := os.Stat(client_path)
		if os.IsNotExist(err) {
			ops_log.Error(0x01, "not exist..." + client_path)
			break
		} else {
			ops_log.Error(0x01, "exist..." + client_path)
			client_path = ""
		}
	}
	cli_addr := net.UnixAddr{client_path, uds_type}
	ser_addr := net.UnixAddr{server_path, uds_type}
	conn, err_conn := net.DialUnix(uds_type, &cli_addr, &ser_addr)
	if err_conn != nil {
		ops_log.Error(0x01, err_conn.Error())
	}

	ops_log.Debug(0x01, "Server write %ld\n", len(req_buf))
	for i:=0;i<len(req_buf);i++ {
		ops_log.Debug(0x01, "%x,", req_buf[i])
	}
	ops_log.Debug(0x01, "\n")

	num_write, err_write := conn.Write(req_buf)
	ops_log.Debug(0x01, "num write: %d\n", num_write)
	if err_write != nil {
		ops_log.Error(0x01, err_write.Error())
	}
	ops_log.Debug(0x01, "Client Read\n")
	num_read, err_read := conn.Read(res_buf)
	ops_log.Debug(0x01, "num read: %d\n", num_read)
	if err_read != nil {
		ops_log.Error(0x01, err_read.Error())
	}

	conn.Close()
	os.Remove(client_path)
	return res_buf
}

func responseWithJsonV1(w http.ResponseWriter, code int,  data interface{}) {
    json_msg := json_msg_t { Status:1, Version:1, Data:data }
    response, _ := json.Marshal(json_msg)
    ops_log.Debug(0x01, string(response))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

type test_t struct {
	Fn	int	`json:fn`
	Cmd	int	`json:cmd`
	Payload	string	`json:payload`
}

func sendTestV1(fn, cmd uint8, data string) (test_t){
	var req test_t
	var res test_t
	req.Fn = int(fn)
	req.Cmd = int(cmd)
	req.Payload = data
	req_bytes, req_err := json.Marshal(req)
	if req_err != nil {
		ops_log.Error(0x01, req_err.Error())
	}
	_, _, _, res_len, res_bytes := SendAndRecvByMsg(fn, cmd, uint16(len(req_bytes)), req_bytes)
	json.Unmarshal(res_bytes, &res)
	ops_log.Debug(0x01, "len:%ld, %ld\n", res_len, len(res_bytes))
	return res
}

type test_v2_t struct {
	Fn	int	`json:fn`
	Cmd	int	`json:cmd`
	Payload	string	`json:payload`
	Reqpayload	string	`json:req_payload`
}

func sendTestV2(fn, cmd uint8, data string) (test_v2_t){
	var req test_v2_t
	var res test_v2_t
	req.Fn = int(fn)
	req.Cmd = int(cmd)
	req.Payload = data
	req_bytes, req_err := json.Marshal(req)
	if req_err != nil {
		ops_log.Error(0x01, req_err.Error())
	}
	_, _, _, res_len, res_bytes := SendAndRecvByMsg(fn, cmd, uint16(len(req_bytes)), req_bytes)
	json.Unmarshal(res_bytes, &res)
	ops_log.Debug(0x01, "len:%ld, %ld\n", res_len, len(res_bytes))
	return res
}

func Test(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fn_int, _	:= strconv.Atoi(params["fn"])
    cmd_int, _	:= strconv.Atoi(params["cmd"])
    data_str, _	:= params["data"]
    api, _	:= params["api_version"]
    fn := uint8(fn_int)
    cmd := uint8(cmd_int)
    switch api {
    case "v1":
        responseWithJsonV1(w, http.StatusOK, sendTestV1(fn, cmd, data_str))
    case "v2":
	responseWithJsonV1(w, http.StatusOK, sendTestV2(fn, cmd, data_str))
    default:
        responseWithJsonV1(w, http.StatusOK, nil)
    }
}

