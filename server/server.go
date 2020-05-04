package server

import (
	"bigo/controller"
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"

	"bigo/model"
	"bigo/utils"
)

var (
	host           = ""
	port           = ""
	addr           = ""
	configFilePath = "/home/mpb/work/bigProject/bigo/apps/server/config/config.ini"
)

func init() {
	ini := utils.IniParser{}
	ini.Load(configFilePath)
	host = ini.GetString("base", "host")
	port = ini.GetString("base", "port")
	addr = net.JoinHostPort(host, port)
}

func Start() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Start server listener failed, err: %v", err)
		os.Exit(1)
	}
	defer listener.Close()
	log.Printf("Start listening at %s...", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener accepts connection err")
		}
		log.Printf("Accepted a connection from %s", conn.RemoteAddr())
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		cmdBytes, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("An error happened when reader.ReadBytes", err)
			return
		}

		// request bytes stream parse
		bigoRequest := model.BigoRequest{}
		if err = json.Unmarshal(cmdBytes, &bigoRequest); err != nil {
			log.Println("An error happened when json.Unmarshal()", err)
			return
		}
		log.Printf("[client id] %s, [command] %s, [args] %s\n", bigoRequest.ClientInfo.ClientId, bigoRequest.CommandName, bigoRequest.Args)

		keyType, err := typeByBigoRequest(bigoRequest)
		if err != nil {
			log.Println(err)
		}
		// address to the specific namespace by command name
		res := controller.Bigo[keyType].Call(bigoRequest.CommandName, bigoRequest.Args)

		// Write response
		n, err := conn.Write(res)
		if err != nil {
			log.Println("An error happened when write data to conn")
		} else {
			log.Println(n, "bytes written")
		}
	}
}


//func NewBigoObjectByRequest(req model.BigoRequest) *model.BigoObject {
//	bigoObject := &model.BigoObject{}
//
//	// bigoObject.Type
//	_type, err := typeByBigoRequest(req)
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//	bigoObject.Type = _type
//
//
//	return bigoObject
//}
