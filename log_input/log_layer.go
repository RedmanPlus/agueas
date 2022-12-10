package log_input

import (
  "log"
  "net"
  "bufio"
)

func dialTCP() (net.Conn, error) {
  if ln, err := net.Listen("tcp", ":4510"); err != nil {
    return nil, err
  } else {
    return ln.Accept()
  }
}

func InputLoop() {
  conn, err := dialTCP()
  if err != nil {
    log.Fatal(err)
  }
  for {
    msg, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    log.Print(msg)
    conn.Write([]byte{})
    if string(msg) == "Quit" {
      break
    }
  }
}
