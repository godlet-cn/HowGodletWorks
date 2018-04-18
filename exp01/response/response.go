package response

import(
	"net"
	"fmt"
	"github.com/godlet-io/HowGodletWorks/exp01/request"
)

const BUFFER_SIZE int =1024

type Response struct{
	Conn net.Conn
    Req request.Request
}

func NewResponse(conn net.Conn) Response{
	var resp Response
	resp.Conn=conn
	return resp
}

func (response *Response)SetRequest(request request.Request){
	response.Req=request;
}

func (response *Response) SendResponse(){
	//var bytes [BUFFER_SIZE]byte
	fmt.Println("Data send to client:"+response.Req.Uri) 
    response.Conn.Write([]byte(response.Req.Uri))
}