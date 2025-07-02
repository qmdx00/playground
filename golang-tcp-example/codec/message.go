package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const headerLength uint32 = 4 // uint32 4 字节用于记录消息长度

var byteOrder binary.ByteOrder = binary.BigEndian

var debug bool = true

type MessagePacket []byte

func Encode(message string) (MessagePacket, error) {
	buffer := bytes.NewBuffer([]byte{})

	length := uint32(len(message))
	if err := binary.Write(buffer, byteOrder, length); err != nil {
		return nil, err
	}

	data := []byte(message)
	if err := binary.Write(buffer, byteOrder, data); err != nil {
		return nil, err
	}

	bts := buffer.Bytes()

	if debug {
		fmt.Println("======================= ENCODE BLOCK START ========================")
		fmt.Printf("|| length=%d\n", length)
		fmt.Printf("|| bytes=%v\n", bts)
		fmt.Printf("|| bits=%08b\n", bts)
		fmt.Println("======================= ENCODE BLOCK END ==========================")
	}

	return bts, nil
}

func Decode(reader io.Reader) (uint32, string, error) {
	buffer := bufio.NewReader(reader)

	var length uint32
	lenBytes, _ := buffer.Peek(int(headerLength))
	if err := binary.Read(bytes.NewBuffer(lenBytes), byteOrder, &length); err != nil {
		return 0, "", err
	}

	if uint32(buffer.Buffered()) < length+headerLength {
		return 0, "", errors.New("data incomplete")
	}

	// data without length
	buffered := make([]byte, length+headerLength)
	if _, err := buffer.Read(buffered); err != nil {
		return 0, "", err
	}
	data := string(buffered[headerLength:])

	if debug {
		fmt.Println("======================= DECODE BLOCK START ========================")
		fmt.Printf("|| length=%d\n", length)
		fmt.Printf("|| bytes=%v\n", buffered)
		fmt.Printf("|| bits=%08b\n", buffered)
		fmt.Printf("|| data=%s\n", data)
		fmt.Println("======================= DECODE BLOCK END ==========================")
	}

	return length, string(data), nil
}
