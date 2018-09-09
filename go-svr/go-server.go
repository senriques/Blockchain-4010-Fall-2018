package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Cfg = flag.String("cfg", "cfg.json", "config file for this call")
var Dir = flag.String("dir", "./static", "Directory to server static files form")
var Port = flag.String("port", ":15021", "Default port, 15021")
var Debug = flag.String("debug", "", "Debug flags - list of CSV flags")
var ErrorStatus = flag.Bool("errorStatus", false, "set 404/500 errors, else JSON 200s")

type ConfigType struct {
	Dir         string
	Port        string
	Debug       []string
	ErrorStatus string
}

func ParseCmdLineArgs() {

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) > 0 {
		fmt.Printf("Unexpected arguments")
		os.Exit(1)
	}

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	gCfg, err := ReadConfig(*Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	if *Dir == "./static" && gCfg.Dir != "" {
		Dir = &gCfg.Dir
	}
	if *Port == ":15021" && gCfg.Port != "" {
		Port = &gCfg.Port
	}

	// xyzzy - ./cfg.json Debug
	// xyzzy - ./cfg.json ErrorStatus

	if *Dir == "." {
		tDir, _ := os.Getwd()
		Dir = &tDir
	}

}

/*

6. Wallet Homework (6)
	1. Make the system a server
		RESTful
			/api/SendTransaction?from= to= amount= message= messageHash= signature= comment=
			/api/ListAccounts?
			/api/ShowBalance?acct=
			/api/ChainStatus?
			/files...

7. Smart Contract (7)
		RESTful
			POST: /api/CallContract?from= to= amount= message= messageHash= signature= comment= contractToCall= p01= ...
			POST: /api/LoadContract?from= message= messageHash= signature= code=

*/

// -------------------------------------------------------------------------------------------------
func main() {
	ParseCmdLineArgs()

	http.HandleFunc("/api/status", respHandlerStatus)
	http.HandleFunc("/api/SendTransaction", respHandlerStatus) // xyzzy
	http.HandleFunc("/api/ListAccounts", respHandlerStatus)    // xyzzy
	http.HandleFunc("/api/ShowBalance", respHandlerStatus)     // xyzzy
	http.HandleFunc("/api/ChainStatus", respHandlerStatus)     // xyzzy
	http.HandleFunc("/api/CallContract", respHandlerStatus)    // xyzzy
	http.HandleFunc("/api/LoadContract", respHandlerStatus)    // xyzzy
	http.Handle("/", http.FileServer(http.Dir(*Dir)))
	log.Fatal(http.ListenAndServe(*Port, nil))
}
