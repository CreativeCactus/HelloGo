package main 
import (
  "testing" 
  "net"
  "fmt"
  "bufio"
)
    
func TestNothing(t *testing.T) {
    if testing.Short() {t.Skip("skipping test in short mode.")}//-test.short
}
func TestEcho(t *testing.T) {
    sendRecv(t,"_echo HELLO WORLD","")
}

// Helpers
func sendRecv(t *testing.T, msg string,dest string) string {
    if dest=="" { dest = "127.0.0.1:20000" }
    conn, _ := net.Dial("tcp", dest)
    t.Logf("Sending:%v\nTo:%v",msg,dest)
    conn.Write([]byte(msg+"\n"))
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Response: "+message)
    return message
}
