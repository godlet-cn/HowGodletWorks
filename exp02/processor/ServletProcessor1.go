package processor

import (
	"fmt"
	"github.com/godlet-cn/HowGodletWorks/exp02/request"
	"github.com/godlet-cn/HowGodletWorks/exp02/response"
)

type ServletProcessor1 struct {
}

func (processor ServletProcessor1) Process(request request.Request, response response.Response) {
	fmt.Println("Handle servlet request")
	response.SendResponse()
}
