package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var kLength = 256
var kTargetKey = make([]byte, kLength)
var kFound = make(chan bool)
var kResult []byte

func Init() {
    // Define printable letters for easy read.
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().Unix())
    //  rand.Read(kTargetKey)
	for i := 0; i < kLength; i++ {
		kTargetKey[i] = letters[rand.Intn(len(letters))]
	}
}

func HalfKey(lowKey, highKey []byte) []byte {
	var carry int32 = 0
	var halfKey = make([]byte, kLength)
	var sumKey = make([]byte, kLength)

	for i := 255; i >= 0; i-- {
		tmp := int32(lowKey[i]) + int32(highKey[i]) + carry
		if tmp >= 256 {
			carry = 1
			tmp -= 256
		} else {
			carry = 0
		}
		sumKey[i] = byte(tmp)
	}

	for i := 0; i < kLength; i++ {
		tmp := carry*256 + int32(sumKey[i])
		if tmp&1 == 1 {
			carry = 1
		} else {
			carry = 0
		}
		halfKey[i] = byte(tmp >> 1)
	}

	return halfKey
}

func Search(key []byte) []byte {
	time.Sleep(time.Millisecond * 10)
	if bytes.Compare(key, kTargetKey) > 0 {
		return nil
	}
	return key
}

// Simple binary search.
func BinarySearch(low, high []byte) []byte {
	half := HalfKey(low, high)
	if bytes.Compare(half, kTargetKey) == 0 {
		return half
	}

	if Search(half) == nil {
		high = half
	} else {
		low = half
	}
	return BinarySearch(low, high)
}

// Concurrent binary search.
func ConcurrentBinarySearch(low, high []byte) {
	if bytes.Compare(low, kTargetKey) > 0 ||
		bytes.Compare(high, kTargetKey) < 0 {
		return
	}

	half := HalfKey(low, high)
	if bytes.Compare(half, kTargetKey) == 0 {
		kResult = half
		kFound <- true
		return
	}

	go ConcurrentBinarySearch(low, half)
	go ConcurrentBinarySearch(half, high)
}

func main() {
	Init()

	low := make([]byte, kLength)
	high := make([]byte, kLength)
	for i := 0; i < kLength; i++ {
		low[i] = 0
		high[i] = 255
	}

	t := time.Now()
	ConcurrentBinarySearch(low, high)
	<-kFound
	elapsed := time.Since(t)

	fmt.Println("TargetKey: ", string(kTargetKey))
	fmt.Println("Result:    ", string(kResult))
	fmt.Println("App elapsed: ", elapsed)
}
