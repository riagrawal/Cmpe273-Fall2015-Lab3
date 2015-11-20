package main 

func main(){
	//server 2 code
	/*server := rpc.NewServer()
	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":3001")
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
	mux := httprouter.New()
    mux.PUT("/hello",put)
    server := http.Server{
            Addr:        "0.0.0.0:3001",
            Handler: mux,
    }
    server.ListenAndServe()
}