package client

import (
	"bigo/model"
	"bigo/utils"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Client struct {
	Id   string
	PS   string
	Conn net.Conn
}

func NewClient(address string) Client {
	client := Client{
		Id:   utils.Uuid(),
		Conn: connectToServer(address),
	}
	client.PS = fmt.Sprintf("%s> ", client.Conn.RemoteAddr())
	return client
}

func (client *Client) Serve() {
	defer client.Conn.Close()

	for {
		cmd, _ := client.readCommand()
		client.sendCommand(cmd)

		// read response
		reader := bufio.NewReader(client.Conn)
		respon, err := reader.ReadString('\n')
		if err != nil {
			log.Println("An error happened when read response from server")
		}
		fmt.Fprint(os.Stdout, respon)
	}
}

func (client *Client) readCommand() (aCommand model.BigoRequest, err error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(client.PS)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("ReadCommand: An error happened", err)
			return aCommand, err
		}

		cmd, err := client.parseCommand(cmdStr)
		if err != nil {
			log.Fatal("an error happened when parses command")
			return aCommand, err
		}
		return cmd, nil
	}
}

func (client *Client) sendCommand(cmd model.BigoRequest) error {
	bytes, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal("json.Marshal(cmd) err:", err)
		return err
	}

	bytes = append(bytes, '\n') // the end of bytes stream
	_, err = client.Conn.Write(bytes)
	if err != nil {
		log.Fatal("client.Conn.Write(bytes) err: ", err)
	}

	return nil
}

func (client *Client) parseCommand(cmdStr string) (model.BigoRequest, error) {
	cmdStr = strings.TrimRight(cmdStr, "\n")
	strs := strings.Split(cmdStr, " ")
	if len(strs) < 2 {
		log.Fatal("ParseCommand error, command format error(length is less than 2)")
		return model.BigoRequest{}, errors.New("Command format error")
	}

	request := model.BigoRequest{
		CommandName: strings.ToUpper(strs[0]),
		Args:        []byte(cmdStr[len(strs[0])+1:]),
		ClientInfo: model.ClientInfo{
			ClientId: client.Id,
		},
	}

	if !isCommandFormatValid(request) {
		return model.BigoRequest{}, errors.New("Command not support!, Please check and retry")
	}

	return request, nil
}

func connectToServer(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Connect2Server error %v, ", err)
		os.Exit(1)
	}
	log.Printf("Connect to %s", conn.RemoteAddr())
	return conn
}

func isCommandFormatValid(commandRequest model.BigoRequest) bool {
	strings.ToUpper(commandRequest.CommandName)

	return true
}
