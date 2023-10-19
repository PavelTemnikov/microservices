package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/proto"
	desc "week1/grpc/pkg/note_v1"
)

func main() {
	session := &desc.NoteInfo{
		Title:    gofakeit.BeerName(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	dataJson, _ := json.Marshal(session)
	fmt.Printf("\n\ndataJson len %d byte \n%v\n\n", len(dataJson), dataJson)

	dataProto, _ := proto.Marshal(session)
	fmt.Printf("dataProto len %d byte \n%v\n", len(dataProto), dataProto)
}
