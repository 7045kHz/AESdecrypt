package main

import (
	// "encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/mervick/aes-everywhere/go/aes256"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("decrypt -key \"passwd\" -enc \"encrypted string\"")
		os.Exit(1)
	}
	hash := flag.String("key", "", "Decryption Key")
	msg := flag.String("enc", "", "Quoted Message to encrypt")
	flag.Parse()

	// msg = "server=DC1.RSYSLAB.COM;database=OSDISCOVERY;user id=RSYSLAB\\U00001;password=Na0mi@1782;port=1433"
	//msg:="Na0mi@1782"

	// hash := "5fa4bf1b469b2745a4902e569ae14292"
	if len(*hash) > 0 && len(*msg) > 0 {
		dencrypted := aes256.Decrypt(*msg, *hash)
		fmt.Printf("%s\n", dencrypted)

	}

}
