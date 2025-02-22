package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"bigo/model"
	"bigo/utils"
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
		cmd, err := client.readCommand()
		if err != nil {
			fmt.Println("[client.Serve]", err)
			continue
		}
		if err = client.SendCommand(cmd); err != nil {
			fmt.Println("[client.Serve]", err)
			continue
		}

		// read response
		respon, err := client.ReadResponse()
		if err != nil {
			if err == io.EOF { return } else { continue }
		}
		fmt.Fprintln(os.Stdout, strings.TrimRight(respon, "\t"))
	}
}

func (client *Client) readCommand() (aCommand model.BigoRequest, err error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(client.PS)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			return aCommand, err
		}
		cmdStr = strings.TrimLeft(cmdStr, " ")
		cmdStr = strings.TrimRight(cmdStr, " \n")
		if cmdStr == "\n" || cmdStr == "" {
			continue
		}

		cmd, err := client.parseCommand(cmdStr)
		if err != nil {
			return aCommand, err
		}
		return cmd, nil
	}
}

func (client *Client) SendCommand(cmd model.BigoRequest) error {
	bytes, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	bytes = append(bytes, '\t') // the end of bytes stream
	_, err = client.Conn.Write(bytes)
	if err != nil {
		log.Println("client.Conn.Write(bytes) err: ", err)
	}

	return nil
}

func (client *Client) ReadResponse() (string, error) {
	reader := bufio.NewReader(client.Conn)
	respon, err := reader.ReadString('\t')
	if err != nil {
		if err == io.EOF {
			fmt.Fprintln(os.Stdout, "Connection closed")
			return "", io.EOF
		}
		log.Println("[client.Serve](An error happened when read response from server)", err)
		return "", err
	}
	return respon, nil
}

func (client *Client) parseCommand(cmdStr string) (model.BigoRequest, error) {
	strs, err := utils.Split(cmdStr, ' ')
	if err != nil {
		return model.BigoRequest{}, err
	}

	request := model.BigoRequest{
		CommandName: strings.ToUpper(strs[0]),
		Args:        strs[1:],
		ClientInfo: model.ClientInfo{
			ClientId: client.Id,
		},
	}

	return request, nil
}

func connectToServer(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Connect2Server error %v, ", err)
	}
	log.Printf("Connect to %s", conn.RemoteAddr())
	return conn
}
