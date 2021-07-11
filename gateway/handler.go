package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"time"
)

const ms1Addr = ""

func handleConnection(conn net.Conn) {
	tran, err := decodeTran(conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	output := Output{
		ID:        encodeAddr(conn.RemoteAddr()),
		Timestamp: time.Now(),
		Tran:      *tran,
	}
	if err = sendOutput(&output); err != nil {
		log.Fatal(err.Error())
	}
}

func decodeTran(conn net.Conn) (*Transaction, error) {
	var tran Transaction
	err := json.NewDecoder(conn).Decode(&tran)
	if err != nil {
		return nil, err
	}
	return &tran, nil
}

func encodeAddr(addr net.Addr) (code []byte) {
	base64.StdEncoding.Encode(code, []byte(addr.String()))
	return
}

func sendOutput(output *Output) error {
	conn, err := net.Dial("tcp", ms1Addr)
	if err != nil {
		return err
	}
	b, err := json.Marshal(output)
	if err != nil {
		return err
	}
	if _, err = conn.Write(b); err != nil {
		return err
	}
	return nil
}
