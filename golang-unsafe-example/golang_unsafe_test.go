package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceLengthAndCapacity(_ *testing.T) {
	// type slice struct {
	//     array unsafe.Pointer
	//     len   int
	//     cap   int
	// }
	//

	s := make([]int32, 4, 10)

	array := *(*unsafe.Pointer)(unsafe.Pointer(&s))
	length := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(uintptr(0))))
	capacity := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 2*unsafe.Sizeof(uintptr(0))))

	fmt.Println("Array pointer:", array)        // 输出: Array pointer: 0x...
	fmt.Println("Length of slice:", length)     // 输出: Length of slice: 4
	fmt.Println("Capacity of slice:", capacity) // 输出: Capacity of slice: 10
}

func TestStructFields(_ *testing.T) {
	type person struct {
		name string
		age  int
	}

	p := person{name: "Alice", age: 30}
	fmt.Println(p) // 输出: {Alice 30}

	namePtr := (*string)(unsafe.Pointer(&p))
	agePtr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.age)))

	*namePtr = "Bob"
	*agePtr = 35
	fmt.Println(p) // 输出: {Bob 35}
}

func TestMapLength(_ *testing.T) {
	// type hmap struct {
	// 	count      int
	// 	flags      uint8
	// 	B          uint8
	// 	noverflow  uint16
	// 	hash0      uint32
	// 	buckets    unsafe.Pointer
	// 	oldbuckets unsafe.Pointer
	// 	nevacuate  uintptr
	// 	extra      *mapextra
	// }

	// NOTE: make map 返回的是一个指向 hmap 的指针
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2

	countPtr := *(**int)(unsafe.Pointer(&m))
	fmt.Println("Length of map:", *countPtr) // 输出: Length of map: 2
}

// Deprecated:
func stringToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}

// Deprecated:
func bytesToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}

func stringToBytes1(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func bytesToString1(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func TestStringToBytesAndBack(t *testing.T) {
	original := "Hello, World!"
	fmt.Println("Original string:", original)

	// Convert string to []byte
	// bytes := stringToBytes(original)
	bytes := stringToBytes1(original)
	fmt.Println("Converted to bytes:", bytes)

	// Convert []byte back to string
	// converted := bytesToString(bytes)
	converted := bytesToString1(bytes)
	fmt.Println("Converted back to string:", converted)

	if original != converted {
		t.Errorf("Expected %s, got %s", original, converted)
	}
}
