//package main
package daemons 

import (
	"log"
_	"fmt"
)


func Controller_process() {
  f := SetupLogging() 
	defer f.Close()
	log.Println("Starting Controller")
	data.processType=3

  
  InitializeParameters()

	HTTP_Server()
	



}
