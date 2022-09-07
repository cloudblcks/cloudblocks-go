package cloudblocks_test

import (
	"bufio"
	"fmt"
	"os"

	cloudblocks_go "github.com/cloudblcks/cloudblocks-go"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Cloudblocks username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Cloudblocks password: ")
	password, _ := reader.ReadString('\n')
	client, err := cloudblocks_go.NewClient(username, password)
	if err != nil {
		panic(err)
	}
	fmt.Print("Origin resource ID: ")
	originId, _ := reader.ReadString('\n')
	fmt.Print("Target resource ID: ")
	targetId, _ := reader.ReadString('\n')
	creds, err := client.RequestCredentials(originId, targetId)
	if err != nil {
		panic(err)
	}
	fmt.Println(creds)
}
