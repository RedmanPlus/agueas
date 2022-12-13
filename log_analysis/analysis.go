package log_analysis

import (
  "log"
)

type Analyser struct {
  inputs []chan map[string]string
}

func NewAnalyser(inputs []chan map[string]string) Analyser {
  return Analyser{
    inputs: inputs,
  }
}

func (self Analyser) GetInputs() {
  for {
    for i := 0; i < len(self.inputs); i++ {
      val := <- self.inputs[i]
      log.Print(val)
    }
  }
}
