package main

import "github.com/gorilla/mux"
import "log"
import "net/http"
import "os"
import "gpio"
import "raw"
import "ops_uds"

type rest_api_function_t func(w http.ResponseWriter, r *http.Request)

type rest_api_t struct {
    Url		string
    Method	string
    Function	rest_api_function_t
}

var api_url_prefix = "/api/{api_version}"

var rest_api_list = []rest_api_t {
    {api_url_prefix + "/gpio",			"GET",	gpio.GetIoListGpio},
    {api_url_prefix + "/gpio/{port}/{pin}",	"GET",	gpio.GetIoGpio},
    {api_url_prefix + "/gpio/{port}/{pin}",	"PUT",	gpio.PutIoGpio},

    {api_url_prefix + "/uds/{fn}/{cmd}/{data}",	"GET",	ops_uds.Test},

    {api_url_prefix + "/raw/{fn}/{cmd}",	"GET",	raw.Get},
    {api_url_prefix + "/raw/{fn}/{cmd}",	"PUT",	raw.Put},
    {api_url_prefix + "/raw/{fn}/{cmd}",	"DELETE",raw.Delete},
    {api_url_prefix + "/raw/{fn}/{cmd}",	"POST",	raw.Post},
    {api_url_prefix + "/raw/{fn}/{cmd}",	"PATCH",raw.Patch},
    {api_url_prefix + "/raw/{fn}/{cmd}",	"HEAD",	raw.Head},
}

func main() {
    if len(os.Args) <= 1 {
        log.Print("prog <www_dir>")
        os.Exit(-1)
    }
    www_path := os.Args[1]
    log.Print(www_path)
    router := mux.NewRouter()

    gpio.Init()
    //api_url_prefix := "/api/{api_version}"
    for _, rest := range rest_api_list {
        //rest_url := api_url_prefix + rest.Url
	log.Printf("%s:%s\n", rest.Method, rest.Url)
        router.HandleFunc(rest.Url, rest.Function).Methods(rest.Method)
    }
    log.Printf("WWW(80) starting ...\n")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir(www_path)))
    log.Fatal(http.ListenAndServe(":80", router))
}

