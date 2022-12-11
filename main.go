package main

import (
  "augeas/log_input"
)

func main() {
  inputLayer := log_input.NewInputLayer("tcp", ":4510")
  inputLayer.InputLoop()
}
