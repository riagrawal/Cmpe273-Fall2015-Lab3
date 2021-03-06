package main 

import (
    "net/http"
    "log"
    "io/ioutil"
    "strconv"
    "fmt"
    "os"
    "strings"
    //"encoding/json"
    //"bytes"
)


type Response struct{
	Key  int 		`json:"key"`
	Value string    `json:"value"`

}

var key_value map[int] string

func main(){


	//client code
	var port string
	var hash int
	if(len(os.Args)==1){
		log.Println("Not enough arguments")
		os.Exit(1)
	}

	if(len(os.Args)==2){

		url := fmt.Sprintf("http://localhost:3000/keys")
		get, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(get.Body)
		get.Body.Close()
		log.Println("Key Value pairs are  : ", string(data))
	}else if os.Args[1]== "GET" {
		request_string := os.Args[2]
		key := strings.Split(request_string,"/")
		key_integer,_ := strconv.Atoi(key[2])
		hash = key_integer % 3
		if(hash == 0){
			port = "3000"
		}else if(hash == 1){
			port = "3001"
		}else {
			port = "3002"
		}
		url := fmt.Sprintf("http://localhost:%s/keys/%s",port,key[2])
		get_key, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		data, err := ioutil.ReadAll(get_key.Body)
		get_key.Body.Close()
		log.Println("Key Value pairs are  : ", string(data))
	}else{
		req_string := os.Args[2]
		key_put := strings.Split(req_string,"/")
		key_int,_ := strconv.Atoi(key_put[2])
		hash = key_int % 3
		if(hash == 0){
			port = "3000"
		}else if(hash == 1){
			port = "3001"
		}else {
			port = "3002"
		}
		put_url:= fmt.Sprintf("http://localhost:%s/keys/%s/%s",port,key_put[2],key_put[3])
		
		client := &http.Client{}
		req, _ := http.NewRequest("PUT", put_url, nil)
		resp, _ := client.Do(req)
		//out, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		log.Println(" Response : ", 200)
	}
	
}