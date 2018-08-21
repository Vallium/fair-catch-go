package main

import (
	"log"

	fc "git.vpgrp.io/channels/search/fair-catch-go"
)

func main() {
	log.Println("Test started")
	fc.TCR{
		Try: func() {
			log.Println("Try")
			fc.Throw("Throw exception")
		},
		Catch: func(e fc.Exception) {
			log.Printf("Caught %v\n", e)
		},
		Recover: func() {
			log.Println("Recovered")
		},
	}.Do()
	log.Println("Here we are")
}
