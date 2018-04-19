package processor

import (
	"fmt"
	"github.com/godlet-cn/HowGodletWorks/exp02/request"
	"github.com/godlet-cn/HowGodletWorks/exp02/response"
)

type StaticResourceProcessor struct {
}

func (processor StaticResourceProcessor) Process(request request.Request, response response.Response) {
	fmt.Println("Handle static resource request")
	response.SendResponse()
}
