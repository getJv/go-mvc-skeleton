package exception

import (
	"getjv/backend/kernel"
)

func NotFoundMessage() ErrorMessage {

    return ErrorMessage{
      Error: true,
      Code: 404,
      Message: "Entity not found",
      Details: "",
      Origin: kernel.ErrorOrigin(),
      
    }
}