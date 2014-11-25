package main

/*
#cgo LDFLAGS: -L/home/jeff/SSC/VAService/libs/ -lvaserver  
#cgo CFLAGS: -I/home/jeff/SSC/VAService/

#include <stdio.h>
#include <stdlib.h>
#include "server.h"
 
void CB(sVAParas *vaParas, int vaParasCount)
{
 	if(vaParasCount>0)
 		printf("<<<<%d,%d>>>>>\n", vaParasCount, vaParas[0].bottom);
}
 
*/
import "C"

import (
	"fmt"
	"io"
 	"net/http"
 //	"net/http/cgi"
 	"runtime"
 		
 	"time"

 	"database/sql"
 	_ "github.com/go-sql-driver/mysql"
)

var datetime = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

func hello(rw http.ResponseWriter, req *http.Request) {
    io.WriteString(rw, "hello~~~")
 

}

func Q(rw http.ResponseWriter, req *http.Request) {
    io.WriteString(rw, "QQ~~~")
}


type DBStruct struct {
 	db *sql.DB
}

// Create the database handle, confirm driver is present
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// username:password@protocol(address)/dbname?param=value
// DROP TABLE times;
func (dbs *DBStruct) dbOpen() (err error) {
	 fmt.Println("  dbOpen" ) 
	dbs.db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/jeff")
	if err != nil {
       fmt.Println("ERROR dbOpen" )   
    }

    return err
} 

func (dbs *DBStruct) dbInsert() (err error) {
	 
 	err = dbs.db.Ping()
    if err != nil {
        fmt.Println(err)
    }

 	if _, err = dbs.db.Exec( "CREATE TABLE IF NOT EXISTS times (id INT AUTO_INCREMENT PRIMARY KEY, datetime TIMESTAMP )" ); err != nil {
	  	panic(err)
	}
 	
 	t := time.Now()
//datetime.Format(time.RFC3339)
	var INSERT = fmt.Sprintf("INSERT INTO times (datetime) VALUES( '%s')",t.Format("20060102150405") )

 	if _, err = dbs.db.Exec( INSERT ); err != nil {
 		panic(err)
	}

    return err
}

func (dbs *DBStruct) dbCheckVersion() {
	// Connect and check the server version
	var version string
	dbs.db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

  

func main() {
 //runtime.NumCPU()
	
 	runtime.GOMAXPROCS(2)
 
 
	//fmt.printfln("format " )
	var dbs DBStruct

   dbs.dbOpen()
   dbs.dbInsert()
   dbs.dbCheckVersion()

/* 
    C.regCB((*[0]byte)(C.CB))
 	go C.runServer()
*/

	http.HandleFunc("/", hello)
	http.HandleFunc("/Q", Q)
	http.ListenAndServe(":8080", nil)

}