package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"captcha/process"
)

func EncodeFile(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func CaptchaStream(w http.ResponseWriter, req *http.Request) {
	str, err := process.CaptchaContainer.Next()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	strArr := strings.Split(str, "|")
	file, err := ioutil.ReadFile("../tmp/" + strArr[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write([]byte(EncodeFile(file) + "|" + strArr[1]))
}

func main() {
	process.Start()

	http.HandleFunc("/", CaptchaStream)
	http.ListenAndServe(":8001", nil)
}
