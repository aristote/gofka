
/*
 * SIMPLE GOLANG MICROSERVICE TESTER
 * (c) jme 2018
 */

 package main

 import (
	//"crypto/hmac"
	//"crypto/sha256"
	//"crypto/tls"
	"crypto/rand"
	//"crypto/x509"
	//"encoding/base64"
	"fmt"
	"io"
	"time"
	"sync"
 )


 const version_microService = 1.0
 const version_API = 1.0
 const version_dataDOM = 1.0
 const num_proc = 8
 //const hmac_secret string = "89f134ac807ca8e78406bad76d241d3d7fd22"


type AccountBalance struct {
	QueryID  string `json:"queryID"`
	Versions struct {
		Microservice int `json:"microservice"`
		API          int `json:"api"`
		DataDOM      int `json:"dataDOM"`
	} `json:"versions"`
	Customer struct {
		CustomerID string `json:"CustomerID"`
	} `json:"customer"`
	Query struct {
		Command   string `json:"command"`
		AccountNo string `json:"accountNo"`
	} `json:"query"`
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}


func newQuery(buffer chan string)() {
	uuid, err := newUUID()
	if err != nil {
		panic("Error while getting the UUID")
	}
	buffer <- uuid
}

func main() {
	buffer := make(chan string, 10)
	start := time.Now()
	var wg sync.WaitGroup
	
	
	for i := 0; i < num_proc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newQuery(buffer)
			}()
		
	}


	wg.Wait() // WaitGroup so we don't deadlock (MOTHER F*CKER!)
	close(buffer)
	
	for message := range buffer {
		println(message)

	}
elapsed := time.Since(start)
total := fmt.Sprintf("Task took %s", elapsed)
println(total)
} // main
