package exception

import (
	"getjv/backend/kernel"
)

func DatabaseMessage(err error) ErrorMessage {

    return ErrorMessage{
      Error: true,
      Code: 503,
      Message: "Database Error",
      Details: err.Error(),
      Origin: kernel.ErrorOrigin(),
      
    }
}