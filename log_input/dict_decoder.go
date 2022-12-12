package log_input

import (
  "log"
  "strconv"
  "fmt"
  "github.com/nlpodyssey/gopickle/types"
  "augeas/utils"
)

type DictDecoder struct {}

func (self DictDecoder) Decode (input types.Dict) {
  logMap := make(map[string]string)
  logKeys := utils.MakeLogKeys()
  for i := 0; i <= input.Len(); i++ {
    val, _ := input.Get(logKeys[i])
    log.Print(val)
    if val != nil {
      switch val.(type) {
      case int:
        logMap[logKeys[i]] = string(val.(int))
      case float32:
        comma, _ := val.(float32)
        vl := fmt.Sprintf("%f", comma)
        logMap[logKeys[i]] = vl
      case bool:
        comma, _ := val.(bool)
        logMap[logKeys[i]] = strconv.FormatBool(comma)
      case string:
        logMap[logKeys[i]] = val.(string)
      default:
        panic("undefined type")
      }
    } else {
      logMap[logKeys[i]] = ""
    }
  }
  log.Print(logMap)
}
