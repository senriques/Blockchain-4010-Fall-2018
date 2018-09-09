package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// -------------------------------------------------------------------------------------------------
// Return true if file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// -----------------------------------------------------------------------------------------------------------

var invalidMode = errors.New("Invalid Mode")

func Fopen(fn string, mode string) (file *os.File, err error) {
	file = nil
	if mode == "r" {
		file, err = os.Open(fn) // For read access.
	} else if mode == "w" {
		file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	} else if mode == "a" {
		file, err = os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		}
	} else {
		err = invalidMode
	}
	return
}

// ReadConfig will read a configuration file into the global congiguration structure.
func ReadConfig(filename string) (gCfg ConfigType, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read config file %s,  %s\n", filename, err)
		os.Exit(1)
	}
	err = json.Unmarshal(buf, &gCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config %s\n", err)
		os.Exit(1)
	}
	return
}

// -------------------------------------------------------------------------------------------------
func SVar(v interface{}) string {
	s, err := json.Marshal(v)
	// s, err := json.MarshalIndent ( v, "", "\t" )
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

// -------------------------------------------------------------------------------------------------
func SVarI(v interface{}) string {
	// s, err := json.Marshal ( v )
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

// -------------------------------------------------------------------------------------------------
func respHandlerStatus(res http.ResponseWriter, req *http.Request) {
	q := req.RequestURI

	var rv string
	res.Header().Set("Content-Type", "application/json")
	rv = fmt.Sprintf(`{"status":"success","name":"go-server version 1.0.0","URI":%q,"req":%s, "response_header":%s}`,
		q, SVarI(req), SVarI(res.Header()))

	io.WriteString(res, rv)
}
