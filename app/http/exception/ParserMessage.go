package exception

import (
	"getjv/backend/kernel"
)

func ParserMessage( err error ) ErrorMessage {
  
   
  return ErrorMessage{
      Error: true,
      Code: 503,
      Message: "Body parser Error",
      Details: err.Error(),
      Origin: kernel.ErrorOrigin(),
    }

}