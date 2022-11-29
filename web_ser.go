package main
import (
	"io/ioutil"
	"log"
	"net/http"
)

func web_server() {
    http.HandleFunc("/", sindex) // index 为向 url发送请求时，调用的函数
    log.Fatal(http.ListenAndServe("172.17.0.2:8181", nil))
}
func sindex(w http.ResponseWriter, r *http.Request) {
    content, _ := ioutil.ReadFile("./index.html")
    w.Write(content)
}
