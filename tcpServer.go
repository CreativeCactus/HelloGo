package main

import (
  "net"
  "fmt"
  "bufio"
  "os/exec"
  "log"
  "sort"
  "reflect"
  "strconv"
	"github.com/HouzuoGuo/tiedot/db" //Replace with file io
//	"github.com/HouzuoGuo/tiedot/dberr"
)




func main() {
  // Config
  proto := "tcp"
  port  := ":8081"
  dbPath:= "./godb"
  
  // Storage
	DB, err := db.OpenDB(dbPath)
	cols:=DB.AllCols()
	print("Loaded DB with cols: %v",cols);
	print("DBG:%v",in(cols,"Users"))
	if len(cols)==0 {
	  if err := DB.Create("Users"); err != nil {		panic(err)	}
	  print("INIT:Created Users collection.");
	}
	
	if err != nil {		panic(err)	}
  
  // Blocking Server
  print("Waiting for connections on %v %v",proto,port)
  session, _ := net.Listen(proto,port)
  for {
    conn, _ := session.Accept()
    go fmt.Println(handleComms(conn,DB))
  }
}

func handleComms(conn  net.Conn, DB *db.DB) string{
  fmt.Printf("User connected from %v\n",conn.RemoteAddr());
  conn.Write([]byte("Hi. Do you have ID?\n"))
  ID, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("%v assumed ID:%v\n", conn.RemoteAddr(), ID)
  UID:=µ(strconv.Atoi(ID))[0].(int);
  print("DBG:%v",UID)
  
  //if err := DB.Create("Feeds"); err != nil {		panic(err)	}
  DBusers := DB.Use("Users")
  if value,err := DBusers.Read(UID); err!=nil||value==nil {
    //conn.Close(); 
    //return fmt.Sprintf("Declined new user: %v",value);
  } else {
    fmt.Sprintf("Accepted user: %v",value); 
    
  }
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
    Cmd=string(Cmd);
    res, err  := exec.Command(Cmd).Output()
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

func anySlice(arg interface{}) (out []interface{}) {
  slice, success := func (arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
      val = reflect.ValueOf(arg)
      ok = (val.Kind() == kind)
      return
    }(arg, reflect.Slice)
  if !success {        return nil  }
  c := slice.Len()
  out = make([]interface{}, c)
  for i := 0; i < c; i++ {
      out[i] = slice.Index(i).Interface()
  }
  return out
}














//func ĳ(){}//T* to interface









func µ(a ...interface{}) []interface{} {  return a } // pasta


