package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jlaffaye/ftp"
	"golang.org/x/crypto/ssh"
)

var addr = flag.String("addr", "localhost:8010", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func get(conn *ftp.ServerConn, pathToFile string) []byte {
	resp, err := conn.Retr(pathToFile)
	if err != nil {
		panic(err)
	}
	defer func(resp *ftp.Response) {
		err := resp.Close()
		if err != nil {
			panic(err)
		}
	}(resp)
	buf, err := ioutil.ReadAll(resp)
	if err != nil {
		panic(err)
	}
	return buf
}
func task1(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	username := "test"
	password := "SDHBCXdsedfs222"
	hostname := "151.248.113.144"
	port := "443"

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, _ := ssh.Dial("tcp", hostname+":"+port, config)

	for {

		err = conn.WriteMessage(websocket.TextMessage, []byte(norm(client)))

		time.Sleep(2 * time.Second)
	}
}
func RunCommandAndWait(cmd string, client *ssh.Client) []byte {
	var (
		session *ssh.Session
		err     error
		baData  []byte
	)

	if session, err = client.NewSession(); err != nil {
		return nil
	}
	defer session.Close()

	baData, err = session.Output(cmd)
	return baData
}

func norm(client *ssh.Client) string {

	str := string(RunCommandAndWait(" find . -name \"achtung.txt\"", client))

	if str == "" {
		str = "norm"
		return str
	}
	str = string(RunCommandAndWait("cat "+str, client))
	return str

}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task1", task1)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
