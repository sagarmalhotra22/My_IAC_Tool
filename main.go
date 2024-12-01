package main

//go get google.golang.org/protobuf/proto
// go get cloud.google.com/go/compute/apiv1
//go get gopkg.in/yaml.v3
//go get github.com/spf13/cobra@latest
import (
	"log"

	"github.com/sagarmalhotra22/My_IAC_Tool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
