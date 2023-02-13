package exception

import (
	"getjv/backend/kernel"
)

func ServerMessage(err error) ErrorMessage {

    return ErrorMessage{
      Error: true,
      Code: 500,
      Message: "Server Error",
      Details: err.Error(),
      Origin: kernel.ErrorOrigin(),
    }
}