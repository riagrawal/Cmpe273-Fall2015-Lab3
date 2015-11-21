package main 

import (
	"github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
)

var key_value map[int] string

type Response struct{
	Key  int 		`json:"key"`
	Value string    `json:"value"`

}

func main(){
	//server 1 code
	key_value = make(map[int] string)

	go func(){
	mux1 := httprouter.New()	
    mux1.PUT("/keys/:id/:value",put)
    mux1.GET("/keys/:id",get)
    mux1.GET("/keys",getall)
    server := http.Server{
            Addr:        "0.0.0.0:3000",
            Handler: mux1,
    }
    server.ListenAndServe()
	}()
    

    go func(){
    mux2 := httprouter.New()
    mux2.PUT("/keys/:id/:value",put)
    mux2.GET("/keys/:id",get)
    mux2.GET("/keys",getall)
    server2 := http.Server{
            Addr:        "0.0.0.0:3001",
            Handler: mux2,
    }
    server2.ListenAndServe()
    }()

  	mux3 := httprouter.New()
    mux3.PUT("/keys/:id/:value",put)
    mux3.GET("/keys/:id",get)
    mux3.GET("/keys",getall)
    server3 := http.Server{
            Addr:        "0.0.0.0:3002",
            Handler: mux3,
    }
    server3.ListenAndServe()


}


func put(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	log.Println("Inside Put!!")
	key := p.ByName("id")
	value := p.ByName("value")
	key_int, _ := strconv.Atoi(key)
	key_value[key_int] = value
}

func get(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	log.Println("Inside GET!!")
	key := p.ByName("id")
	key_int, _ := strconv.Atoi(key)
	var response Response
    response.Key = key_int
    response.Value = key_value[key_int]
  	payload, err := json.Marshal(response)  
  	if err != nil {
    	 http.Error(rw,"Bad Request" , http.StatusInternalServerError)
     	return
  	}
  	rw.Header().Set("Content-Type", "application/json")
  	rw.Write(payload)
}


func getall(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	log.Println("Inside GETALL!!")
	var response []Response
	var key_pair Response
	for key, value := range key_value {
    	key_pair.Key = key
    	key_pair.Value = value
		response = append(response, key_pair)
	}
    
  	payload, err := json.Marshal(response)  
  	if err != nil {
    	 http.Error(rw,"Bad Request" , http.StatusInternalServerError)
     	return
  	}
  	rw.Header().Set("Content-Type", "application/json")
  	rw.Write(payload)
}