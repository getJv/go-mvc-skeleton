package exception

import (
	"getjv/backend/kernel"
)

func ValidationMessage(err error) ErrorMessage {

    return ErrorMessage{
      Error: true,
      Code: 422,
      Message: "Validation Error",
      Details: err.Error(),
      Origin: kernel.ErrorOrigin(),
      
    }
}