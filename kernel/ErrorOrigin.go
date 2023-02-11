package kernel

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var ErrorOrigin = func() string{

  isDebugOn,_ := strconv.ParseBool(os.Getenv("DEBUG"))
  if !isDebugOn{
    return "[#PROTECTED#]"
  }

  _, fullFilename , line , _ := runtime.Caller(2)

  filename :=  strings.Split(fullFilename,"/")
  return fmt.Sprintf("%s line: %d", filename[len(filename) - 1],line)
}
