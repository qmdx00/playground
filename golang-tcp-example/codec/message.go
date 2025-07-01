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

	fmt.Printf("[ENCODE: length=%d] bytes=%v\n", length, buffer.Bytes())
	return buffer.Bytes(), nil
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
	data := buffered[headerLength:]

	fmt.Printf("[DECODE: length=%d] data=%v\n", length, string(data))

	return length, string(data), nil
}
