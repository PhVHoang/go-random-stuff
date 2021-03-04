package rpcServer

import (
	"io"
	"net/http"
	"net/rpc"
	"typeSharing"
)

func main() {
  mit := typeSharing.NewCollege()
  rpc.Register(mit)
  rpc.HandleHTTP()

  http.HandleFunc("/", func(res http.ResponseWriter, req * http.Request){
    io.WriteString(res, "RPC SERVER LIVE")
  })
  http.ListenAndServe(":9000", nil)
}
