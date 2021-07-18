package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
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
		log.Println(*tran)
		//output := Output{
		//	ID:        encodeAddr(conn.RemoteAddr()),
		//	Timestamp: time.Now(),
		//	Tran:      *tran,
		//}
		//if err = sendOutput(&output); err != nil {
		//	log.Fatal(err.Error())
		//}
	}
}

func decodeTran(payload []byte) (tran *Transaction, err error) {
	err = json.Unmarshal(payload, tran)
	if err != nil {
		return nil, err
	}
	fmt.Println("here", tran)
	return tran, nil
}

func encodeAddr(addr net.Addr) (code []byte) {
	base64.StdEncoding.Encode(code, []byte(addr.String()))
	return code
}

func sendOutput(output *Output) error {
	log.Println(*output)
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
