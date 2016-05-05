package main 
import (
  "testing" 
  "net"
)
    
func TestNothing(t *testing.T) {
    if testing.Short() {t.Skip("skipping test in short mode.")}//-test.short
}
func TestEcho(t *testing.T) {
    sendRecv("_echo HELLO WORLD")
}

func sendRecv(msg string,dest string) string {
    if dest==nil { dest = "127.0.0.1:20000" }
    conn, _ := net.Dial("tcp", dest)
    fmt.Fprintf(conn, msg+"\n")
    fmt.Println("Sending:"+msg+"\nTo:"+dest)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Response: "+message)
}
