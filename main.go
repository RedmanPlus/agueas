package main

import (
  "augeas/log_input"
  "augeas/log_analysis"
)

func main() {
  inputLayer := log_input.NewInputLayer("tcp", ":4510")
  go inputLayer.InputLoop()
  analysisLayer := log_analysis.NewAnalyser(inputLayer.Results)
  analysisLayer.GetInputs()
}
