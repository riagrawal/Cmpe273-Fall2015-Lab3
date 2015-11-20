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
	/*server := rpc.NewServer()
	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":3000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}*/
	key_value = make(map[int] string)
	mux := httprouter.New()
    mux.PUT("/keys/:id/:value",put)
    mux.GET("/keys/:id",get)
    mux.GET("/keys",getall)
    server := http.Server{
            Addr:        "0.0.0.0:3001",
            Handler: mux,
    }
    server.ListenAndServe()
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