package log_input

import (
  "log"
  "net"
  "github.com/nlpodyssey/gopickle/pickle"
  "github.com/nlpodyssey/gopickle/types"
)

type InputLayer struct {
  connType string
  tcpPort  string
}

func NewInputLayer(connType string, tcpPort string) *InputLayer {
  return &InputLayer{
    connType: connType,
    tcpPort:  tcpPort,
  }
}

func (self *InputLayer) dialTCP() (net.Listener, error) {
  if ln, err := net.Listen(self.connType, self.tcpPort); err != nil {
    return nil, err
  } else {
    return ln, nil
  }
}

func (self InputLayer) InputLoop() {
  ln, err := self.dialTCP()
  defer ln.Close()
  if err != nil {
    log.Fatal(err)
  }
  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  buffer := make([]byte, 1024)
  _, err := conn.Read(buffer)
  if err != nil {
    log.Fatal(err)
  }
  go handleLog(buffer)
  conn.Close()
}

func handleLog(input []byte) {
  if msgInterface, err := pickle.Loads(string(input[4:])); err != nil {
    log.Fatal(err)
  } else {
    decodedMsg := msgInterface.(*types.Dict)
    DictDecoder{}.Decode(*decodedMsg)
  } 
}

