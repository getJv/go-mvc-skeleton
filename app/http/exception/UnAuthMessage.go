package exception

import (
	"getjv/backend/kernel"
)

func UnAuthMessage() ErrorMessage {

    return ErrorMessage{
      Error: true,
      Code: 401,
      Message: "UnAuthorizated",
      Details: "",
      Origin: kernel.ErrorOrigin(),
      
    }
}