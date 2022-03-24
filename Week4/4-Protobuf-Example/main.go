package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf_example/homework/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	"protobuf_example/src/simple/simplepb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	/*
		sm := DoSimple()
		fmt.Println(sm)
	*/

	// ReadAndWriteDemo(sm)
	// jsonDemo(sm)

	/*
		em := enumpb.EnumMessage{
			Id:           12,
			DayOfTheWeek: enumpb.DayOfTheWeek_SUNDAY,
		}
		fmt.Println(em.String())
		em.DayOfTheWeek = enumpb.DayOfTheWeek_TUESDAY
		fmt.Println(em.String())
	*/

	/*
		cm := complexpb.ComplexMessage{
			OneDummy: &complexpb.DummyMessage{
				Id:   1,
				Name: "Sanju",
			},
			MultipleDummy: []*complexpb.DummyMessage{
				{
					Id:   2,
					Name: "Sameer",
				},
				{
					Id:   3,
					Name: "Karan",
				},
			},
		}
		fmt.Println(cm.String())
	*/

	homework := &tutorialpb.Person{
		Name:  "Sanju",
		Id:    03,
		Email: "abcd3066@gmail.com",
		Phones: []*tutorialpb.Person_PhoneNumber{
			{
				Number: "987654321",
				Type:   tutorialpb.Person_WORK,
			},
			{
				Number: "123456789",
				Type:   tutorialpb.Person_MOBILE,
			},
		},
	}
	fmt.Println(homework.String())

	WriteTo("homework/solution.bin", homework)
	homework2 := &tutorialpb.Person{}
	ReadFrom("homework/solution.bin", homework2)
	fmt.Println(homework2.String())

}

func WriteTo(fname string, pb *tutorialpb.Person) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	fmt.Println("Successfully Write to the File")
}

func ReadFrom(fname string, pb *tutorialpb.Person) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	fmt.Println("Successfully Read From File and write back to pb")
}

func JsonDemo(sm proto.Message) {
	smAsString := ToJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	FromJSON(smAsString, sm2)
	fmt.Println(sm2)
}

func ToJSON(pb proto.Message) string {
	/*
		marshaler := jsonpb.Marshaler{}
		out, err := marshaler.MarshalToString(pb)
	*/
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Error: ", err)
		return ""
	}
	return string(out)
}

func FromJSON(str string, pb proto.Message) {
	/*
		err := jsonpb.UnmarshalString(str, pb)
		if err != nil {
			log.Fatalln("Error: ", err)
			return
		}
	*/
	protojson.Unmarshal([]byte(str), pb)
}

func ReadAndWriteDemo(sm proto.Message) {
	WriteToFile("data.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	ReadFromFile("data.bin", sm2)
	fmt.Println(sm2)
}

func WriteToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	fmt.Println("Data has been Written")
}

func ReadFromFile(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
}
func DoSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		Name:       "Sanju",
		IsSimple:   true,
		SampleList: []int32{1, 45, 6, 3},
	}
	fmt.Println(sm.GetId())
	fmt.Println(sm.GetName())
	sm.Name = "Sameer"
	fmt.Println(sm.GetName())
	return &sm
}
