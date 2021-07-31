package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

const ms1Addr = "localhost:8081"
var ms1_conn net.Conn = nil

func handleConnection(conn net.Conn) {
	counter := 0
	var tran = &TransactionData{}
	var err error
	decoder := json.NewDecoder(conn)
	for {
		err = decoder.Decode(tran)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF{
				break
			}
			log.Fatal(err.Error())
		}
		counter++
		fmt.Println(counter)
		timestamp := time.Now()
		output := Transaction {
			ID:        encodeAddr(conn.RemoteAddr(), timestamp),
			Timestamp: timestamp,
			Tran:      *tran,
		}
		if err = sendOutput(&output); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func encodeAddr(addr net.Addr, timestamp time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(addr.String() + "@" + strconv.FormatInt(timestamp.Unix(), 10)))
}

func decodeAddr(code string) (addr []byte, err error) {
	addr, err = base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func sendOutput(output *Transaction) error {
	counter := 0
	addr, err := decodeAddr(output.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(addr))
	if ms1_conn == nil {
		ms1_conn, err = net.Dial("tcp", ms1Addr)
		if err != nil {
			return err
		}
	}
	b, err := json.Marshal(output)
	if err != nil {
		return err
	}
	if _, err = ms1_conn.Write(b); err != nil {
		return err
	}
	counter++
	fmt.Println(output)
	return nil
}
