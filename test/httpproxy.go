// httpproxy
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var SETTING_FILE = "setting.conf"
var SETTING_INIT = `{
	"DistHost":"chelappfntews13",
	"Port":"57608",
	"UrlMapping":{
		"/accountmgmt/v1/extendedconfigs/services/51ae7a6e86d08c1051acdc6f/envs/1111":"serviceclient.json"
	}
}`
var COUNTER = 1

type Config struct {
	DistHost   string
	Port       string
	UrlMapping map[string]string
}

type httpFun func(w http.ResponseWriter, r *http.Request)

func fileHandler(file string) httpFun {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%03d] URI:%s\nResponse:%s\n", COUNTER, r.RequestURI, file)
		COUNTER++
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			//http.NotFound(w, r)
			w.WriteHeader(404)
			w.Write([]byte("{}"))
			return
		}
		w.Write(bytes)
	}
}
func defaultHandler(host, port string) httpFun {

	return func(w http.ResponseWriter, r *http.Request) {
		newURL := fmt.Sprintf("http://%s:%s%s", host, port, r.RequestURI)
		fmt.Printf("[%03d] URI:%s\nRedirect To:%s\n", COUNTER, r.RequestURI, newURL)
		COUNTER++
		client := &http.Client{}
		req, err := http.NewRequest(r.Method, newURL, nil)
		if err != nil {
			panic(err)
		}
		for k, v := range r.Header {
			for _, vv := range v {
				req.Header.Set(k, vv)
			}
		}

		resp, err := client.Do(req)
		defer resp.Body.Close()

		for _, c := range resp.Cookies() {
			w.Header().Add("Set-Cookie", c.Raw)
		}

		w.WriteHeader(resp.StatusCode)
		var buffer bytes.Buffer

		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(resp.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				buffer.Write(buf[0:n])
			}
		default:
			bodyByte, _ := ioutil.ReadAll(resp.Body)
			buffer.Write(bodyByte)
		}

		if err != nil {
			panic(err)
		}
		buffer.WriteTo(w)
	}
}

func main() {
	jsonBytes, err := ioutil.ReadFile(SETTING_FILE)
	if err != nil {
		ioutil.WriteFile(SETTING_FILE, []byte(SETTING_INIT), os.ModePerm)
		fmt.Println("Please configurate setting.conf.")
		return
	}
	config := Config{}
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		fmt.Println("The file setting.conf is not a valid json file.", err)
		return
	}

	fmt.Println("Deamon is running.")
	http.HandleFunc("/", defaultHandler(config.DistHost, config.Port))
	for k, v := range config.UrlMapping {
		http.HandleFunc(k, fileHandler(v))
	}
	http.ListenAndServe(":"+config.Port, nil)
	os.Exit(0)
}
