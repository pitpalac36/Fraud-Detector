package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var processorAddr = os.Getenv("PROCESSOR_ADDR")
var procConn net.Conn = nil

func handleConnection(loadConn net.Conn) {
	counter := 0
	var tran = &TransactionData{}
	var err error
	decoder := json.NewDecoder(loadConn)
	for {
		err = decoder.Decode(tran)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		counter++
		//fmt.Println(counter)
		timestamp := time.Now()
		output := Transaction {
			ID:        encodeAddr(loadConn.RemoteAddr(), timestamp),
			Tran:      *tran,
		}
		if err = sendOutput(&output); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func encodeAddr(addr net.Addr, timestamp time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(addr.String() + "@" + strconv.FormatInt(timestamp.UnixNano(), 10)))
}

func sendOutput(output *Transaction) error {
	var err error
	if processorAddr == "" {
		processorAddr = os.Getenv("PROCESSOR_ADDR")
	}
	counter := 0
	if procConn == nil {
		procConn, err = net.Dial("tcp", processorAddr)
		if err != nil {
			return err
		}
	}
	b, err := json.Marshal(output)
	if err != nil {
		return err
	}
	if _, err = procConn.Write(b); err != nil {
		return err
	}
	counter++
	//fmt.Println(output)
	return nil
}
