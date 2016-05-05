package main

import (
  "net"
  "fmt"
  "bufio"
  "os/exec"
  "log"
  "sort"
  "strconv"
  "strings"
  
  /*
  "io/ioutil"
  "os"
  // */
)




func main() {
  // Config
  proto := "tcp"
  port  := ":20000"
  
  // Storage
  //......
  
  // Blocking Server
  print("Waiting for connections on %v %v",proto,port)
  session, _ := net.Listen(proto,port)
  for {
    conn, _ := session.Accept()
    go fmt.Println(handleComms(conn))
  }
}

func handleComms(conn  net.Conn) string{
  fmt.Printf("User connected from %v\n",conn.RemoteAddr());
  /*
    Look for browser user agent and send html with term callback
  */
  conn.Write([]byte("Hi. Do you have ID?\n"))
  ID, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("%v assumed ID:%v\n", conn.RemoteAddr(), ID)
  UID:=µ(strconv.Atoi(ID))[0].(int);
  print("DBG:%v",UID)
  
  //if err := DB.Create("Feeds"); err != nil {		panic(err)	}
  //conn.Close(); 
  //return fmt.Sprintf("Declined new user: %v",value);
  fmt.Sprintf("Accepted user: %v",UID); 
  
  //fmt.Printf("%v\n", value)
  //db.Erase(key)
  
  
  
  
  for {
    // listen for message ending in \n
    message, err := bufio.NewReader(conn).ReadString('\n') //Simplify?
    if err!= nil { panic(err); }
    print("Message Received:%v", string(message))
    newmessage := Command(message)
    conn.Write([]byte(newmessage))
  }
}

func Command(Cmd string)string{
  c   := string( Cmd[0  ] )//ASCII only
  Cmd  = string( Cmd[1: ] ) 
  print("C:%v\tCmd:%v",c,Cmd)
  if c=="_" {
    Fields := strings.Fields(string(Cmd))
    if len(Fields)==0 { return "" }
    Args := " "
    if len(Fields)>1 {    Args = strings.Join(Fields[1:]," ")  }
    
    //Why does _ls cause exit status 2?
    
    res, err  := exec.Command(Fields[0],Args).Output()
	  if err != nil {		log.Fatal(err)	}
	  return string(res);
  }
	return "*eyes glaze over\n";
}





// Helpers

func print(line string, varargs ...interface{}) (s string){
  s = fmt.Sprintf(line, varargs...)
  fmt.Println(s)
  return
}

func instr(haystack []string, needle string) int{
  return sort.Search(len(haystack),func(i int) bool {
    print("DBG:hay[%v]=%v, need %v [%v]",i,haystack[i],needle,haystack[i]==needle)
    return haystack[i]==needle;
  })
}




func µ(a ...interface{}) []interface{} {  return a } // pasta


