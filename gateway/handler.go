package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const ms1Addr = ""

func handleConnection(conn net.Conn) {
	var tran = &Transaction{}
	var err error
	decoder := json.NewDecoder(conn)
	for {
		err = decoder.Decode(tran)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Println(*tran)
		output := Output{
			ID:        encodeAddr(conn.RemoteAddr()),
			Timestamp: time.Now(),
			Tran:      *tran,
		}
		if err = sendOutput(&output); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func encodeAddr(addr net.Addr) string {
	return base64.StdEncoding.EncodeToString([]byte(addr.String()))
}

func decodeAddr(code string) (addr []byte, err error) {
	addr, err = base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func sendOutput(output *Output) error {
	//fmt.Println((*output).ID)
	addr, err := decodeAddr(output.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(addr))
	//conn, err := net.Dial("tcp", ms1Addr)
	//if err != nil {
	//	return err
	//}
	//b, err := json.Marshal(output)
	//if err != nil {
	//	return err
	//}
	//if _, err = conn.Write(b); err != nil {
	//	return err
	//}
	return nil
}
