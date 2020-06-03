package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var temp *template.Template

type employee struct {
	ID     int    `json:"ID"`
	Fname  string `json:"FNAME"`
	Mname  string `json:"MNAME"`
	Lname  string `json:"LNAME"`
	Mailid string `json:"MAILID"`
	Exten  string `json:"EXTEN"`
	Mobile int    `json:"MOBILE"`
	Depat  string `json:"DEPAT"`
}
type results struct {
	Result []employee `json:"results"`
}
type root struct {
	Root results `json:"d"`
}
type person struct {
	Name string
	Age  int
}

func init() {
	temp = template.Must(template.ParseFiles("rawtemp.gohtml"))
}
func main() {
	lUrl := "https://hxehost.com:4390/emp_masterSrv/emp_master.xsodata/EMPLOYEE?$format=json"
	var username string = "SYSTEM"
	var passwd string = "Abap@2113"
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(http.MethodGet, lUrl, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var responseObj root

		json.Unmarshal(data, &responseObj)

		//fmt.Println(responseObj)

		var empData []employee

		for i := 0; i < len(responseObj.Root.Result); i++ {
			//fmt.Println(responseObj.Root.Result[i])
			empData = append(empData, responseObj.Root.Result[i])
		}
		err := temp.Execute(os.Stdout, empData)
		if err != nil {
			log.Println(err)
		}
	}

}
