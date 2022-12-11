package log_input

import (
  "log"
  "github.com/nlpodyssey/gopickle/types"
  "augeas/utils"
)

type DictDecoder struct {}

func (self DictDecoder) Decode (input types.Dict) {
  logMap := make(map[string]string)
  logKeys := utils.MakeLogKeys()
  for i := 0; i <= input.Len(); i++ {
    val, _ := input.Get(logKeys[i])
    logMap[logKeys[i]] = val.(string)
  }
  log.Print(logMap)
}
