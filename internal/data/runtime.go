package data

import (
	"fmt"
	"strconv"
)


type Runtime int32

func(r Runtime) MarshalJson() ([]byte, error){
  jsonValue := fmt.Sprintf("%d min", r)

  Quotedjsonvalue := strconv.Quote(jsonValue)

  return []byte(Quotedjsonvalue), nil
}