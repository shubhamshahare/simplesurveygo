package main

import (
        mgo "gopkg.in/mgo.v2"
        "log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
        "sync"
      )


var MgoSession *mgo.Session

type KeysResponse struct {
      Collection []Movie
}
type Movie struct {
	Title   string `json:"name"`
	Year  int `json:"year"`
	Director    string    `json:"director"`
	Cast    string    `json:"cast"`
	Genre    string    `json:"genre"`
	Notes    string    `json:"notes"`
}

func main() {

        jsonFile, err := os.Open("movie1.json")
	if err != nil {
		fmt.Println(err)
	}

        fmt.Println(jsonFile)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
        fmt.Println(byteValue)

         keys := make([]Movie,0)

	json.Unmarshal(byteValue, &keys)

        c := make(chan Movie)

        db, err := mgo.Dial("localhost")
        if err != nil {
        log.Fatal("cannot dial mongo", err)
        }
        defer db.Close() 

    db.SetMode(mgo.Monotonic, true)
    var waitGroup sync.WaitGroup
    for instance := 0; instance < 4; instance++ {
        go RunQuery(c,&waitGroup, db)
    }
   for _,val:= range keys {
     c <- val
   }
    waitGroup.Wait()
    log.Println("Completed")


}


func RunQuery(movie chan Movie, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
  for {
  select{
         case m := <- movie:
          defer waitGroup.Done()
          sessionCopy := mongoSession.Copy()
          defer sessionCopy.Close()
          collection := sessionCopy.DB("movie").C("movieinfo")
          err := collection.Insert(m)
          if err != nil {
              log.Printf("RunQuery : ERROR : %s\n", err)
              return
           }
        }
      }
}
