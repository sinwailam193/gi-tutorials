package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	aaron := &Person{
		Name: "Aaron",
		Age:  25,
		SocialFollowers: &SocialFollowers{
			Twitter: 1400,
			Youtube: 2500,
		},
	}

	data, err := proto.Marshal(aaron)
	if err != nil {
		log.Fatal("Marchalling error: ", err)
	}

	newAaron := &Person{}
	if err = proto.Unmarshal(data, newAaron); err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Println(newAaron.GetAge())
	fmt.Println(newAaron.GetName())
	fmt.Println(newAaron.SocialFollowers.GetTwitter())
	fmt.Println(newAaron.SocialFollowers.GetYoutube())
}
