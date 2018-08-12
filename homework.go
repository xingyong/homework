package main

import (
    "bytes"
    "fmt"
    "math/rand"
    "runtime"
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

// Calculate median.
func HalfKey(lowKey, highKey []byte) []byte {
    var carry int32 = 0
    var halfKey = make([]byte, kLength)
    var sumKey = make([]byte, kLength)

    for i := kLength - 1; i >= 0; i-- {
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

// Add key by one.
func AddByOne(key []byte) []byte {
    var carry int32 = 1
    var ret = make([]byte, kLength)
    for i := kLength - 1; i >= 0; i-- {
        tmp := int32(key[i]) + carry
        if tmp >= 256 {
            carry = 1
            ret[i] = 0
        } else {
            carry = 0
            ret[i] = byte(tmp)
        }
    }
    return ret
}

func Search(key []byte) []byte {
    time.Sleep(time.Millisecond * 10)
    if bytes.Compare(key, kTargetKey) > 0 {
        return nil
    }
    return key
}

// Simple binary search.
func BinarySearch(low, high []byte) {
    half := HalfKey(low, high)

    if bytes.Compare(half, low) == 0 {
        kResult = low
        kFound <- true
        fmt.Println("Found: true")
        return
    }

    if Search(half) == nil {
        high = half
    } else {
        low = half
    }
    fmt.Println("Coroutine ", runtime.NumGoroutine())
    BinarySearch(low, high)
}

// Concurrent binary search.
func ConcurrentBinarySearch(low, high []byte) {
    if Search(low) == nil {
        return
    }
    if Search(AddByOne(high)) != nil {
        return
    }

    // fmt.Println("Coroutine ", runtime.NumGoroutine())
    half0 := HalfKey(low, high)
    if bytes.Compare(half0, low) == 0 {
        kResult = low
        kFound <- true
        fmt.Println("Found: true")
        return
    }

    // 4-way
    half10 := HalfKey(low, half0)
    half11 := HalfKey(half0, high)

    // 8-way
    half20 := HalfKey(low, half10)
    half21 := HalfKey(half10, half0)
    half22 := HalfKey(half0, half11)
    half23 := HalfKey(half11, high)

    // 16-way
    half30 := HalfKey(low, half20)
    half31 := HalfKey(half20, half10)
    half32 := HalfKey(half10, half21)
    half33 := HalfKey(half21, half0)
    half34 := HalfKey(half0, half22)
    half35 := HalfKey(half22, half11)
    half36 := HalfKey(half11, half23)
    half37 := HalfKey(half23, high)

    // 32-way
    half40 := HalfKey(low, half30)
    half41 := HalfKey(half30, half20)
    half42 := HalfKey(half20, half31)
    half43 := HalfKey(half31, half10)
    half44 := HalfKey(half10, half32)
    half45 := HalfKey(half32, half21)
    half46 := HalfKey(half21, half33)
    half47 := HalfKey(half33, half0)
    half48 := HalfKey(half0, half34)
    half49 := HalfKey(half34, half22)
    half410 := HalfKey(half22, half35)
    half411 := HalfKey(half35, half11)
    half412 := HalfKey(half11, half36)
    half413 := HalfKey(half36, half23)
    half414 := HalfKey(half23, half37)
    half415 := HalfKey(half37, high)
	
    // 64-way
    half50 := HalfKey(low, half40)
    half51 := HalfKey(half40, half30)
    half52 := HalfKey(half30, half41)
    half53 := HalfKey(half41, half20)
    half54 := HalfKey(half20, half42)
    half55 := HalfKey(half42, half31)
    half56 := HalfKey(half31, half43)
    half57 := HalfKey(half43, half10)
    half58 := HalfKey(half10, half44)
    half59 := HalfKey(half44, half32)
    half510 := HalfKey(half32, half45)
    half511 := HalfKey(half45, half21)
    half512 := HalfKey(half21, half46)
    half513 := HalfKey(half46, half33)
    half514 := HalfKey(half33, half47)
    half515 := HalfKey(half47, half0)
    half516 := HalfKey(half0, half48)
    half517 := HalfKey(half48, half34)
    half518 := HalfKey(half34, half49)
    half519 := HalfKey(half49, half22)
    half520 := HalfKey(half22, half410)
    half521 := HalfKey(half410, half35)
    half522 := HalfKey(half35, half411)
    half523 := HalfKey(half411, half11)
    half524 := HalfKey(half11, half412)
    half525 := HalfKey(half412, half36)
    half526 := HalfKey(half36, half413)
    half527 := HalfKey(half413, half23)
    half528 := HalfKey(half23, half414)
    half529 := HalfKey(half414, half37)
    half530 := HalfKey(half37, half415)
    half531 := HalfKey(half415, high)
	
    // 128-way
    half60 := HalfKey(low, half50)
    half61 := HalfKey(half50, half40)
    half62 := HalfKey(half40, half51)
    half63 := HalfKey(half51, half30)
    half64 := HalfKey(half30, half52)
    half65 := HalfKey(half52, half41)
    half66 := HalfKey(half41, half53)
    half67 := HalfKey(half53, half20)
    half68 := HalfKey(half20, half54)
    half69 := HalfKey(half54, half42)
    half610 := HalfKey(half42, half55)
    half611 := HalfKey(half55, half31)
    half612 := HalfKey(half31, half56)
    half613 := HalfKey(half56, half43)
    half614 := HalfKey(half43, half57)
    half615 := HalfKey(half57, half10)
    half616 := HalfKey(half10, half58)
    half617 := HalfKey(half58, half44)
    half618 := HalfKey(half44, half59)
    half619 := HalfKey(half59, half32)
    half620 := HalfKey(half32, half510)
    half621 := HalfKey(half510, half45)
    half622 := HalfKey(half45, half511)
    half623 := HalfKey(half511, half21)
    half624 := HalfKey(half21, half512)
    half625 := HalfKey(half512, half46)
    half626 := HalfKey(half46, half513)
    half627 := HalfKey(half513, half33)
    half628 := HalfKey(half33, half514)
    half629 := HalfKey(half514, half47)
    half630 := HalfKey(half47, half515)
    half631 := HalfKey(half515, half0)
    half632 := HalfKey(half0, half516)
    half633 := HalfKey(half516, half48)
    half634 := HalfKey(half48, half517)
    half635 := HalfKey(half517, half34)
    half636 := HalfKey(half34, half518)
    half637 := HalfKey(half518, half49)
    half638 := HalfKey(half49, half519)
    half639 := HalfKey(half519, half22)
    half640 := HalfKey(half22, half520)
    half641 := HalfKey(half520, half410)
    half642 := HalfKey(half410, half521)
    half643 := HalfKey(half521, half35)
    half644 := HalfKey(half35, half522)
    half645 := HalfKey(half522, half411)
    half646 := HalfKey(half411, half523)
    half647 := HalfKey(half523, half11)
    half648 := HalfKey(half11, half524)
    half649 := HalfKey(half524, half412)
    half650 := HalfKey(half412, half525)
    half651 := HalfKey(half525, half36)
    half652 := HalfKey(half36, half526)
    half653 := HalfKey(half526, half413)
    half654 := HalfKey(half413, half527)
    half655 := HalfKey(half527, half23)
    half656 := HalfKey(half23, half528)
    half657 := HalfKey(half528, half414)
    half658 := HalfKey(half414, half529)
    half659 := HalfKey(half529, half37)
    half660 := HalfKey(half37, half530)
    half661 := HalfKey(half530, half415)
    half662 := HalfKey(half415, half531)
    half663 := HalfKey(half531, high)
	
    go ConcurrentBinarySearch(low, half60)
    go ConcurrentBinarySearch(half60, half50)
    go ConcurrentBinarySearch(half50, half61)
    go ConcurrentBinarySearch(half61, half40)
    go ConcurrentBinarySearch(half40, half62)
    go ConcurrentBinarySearch(half62, half51)
    go ConcurrentBinarySearch(half51, half63)
    go ConcurrentBinarySearch(half63, half30)
    go ConcurrentBinarySearch(half30, half64)
    go ConcurrentBinarySearch(half64, half52)
    go ConcurrentBinarySearch(half52, half65)
    go ConcurrentBinarySearch(half65, half41)
    go ConcurrentBinarySearch(half41, half66)
    go ConcurrentBinarySearch(half66, half53)
    go ConcurrentBinarySearch(half53, half67)
    go ConcurrentBinarySearch(half67, half20)
    go ConcurrentBinarySearch(half20, half68)
    go ConcurrentBinarySearch(half68, half54)
    go ConcurrentBinarySearch(half54, half69)
    go ConcurrentBinarySearch(half69, half42)
    go ConcurrentBinarySearch(half42, half610)
    go ConcurrentBinarySearch(half610, half55)
    go ConcurrentBinarySearch(half55, half611)
    go ConcurrentBinarySearch(half611, half31)
    go ConcurrentBinarySearch(half31, half612)
    go ConcurrentBinarySearch(half612, half56)
    go ConcurrentBinarySearch(half56, half613)
    go ConcurrentBinarySearch(half613, half43)
    go ConcurrentBinarySearch(half43, half614)
    go ConcurrentBinarySearch(half614, half57)
    go ConcurrentBinarySearch(half57, half615)
    go ConcurrentBinarySearch(half615, half10)
    go ConcurrentBinarySearch(half10, half616)
    go ConcurrentBinarySearch(half616, half58)
    go ConcurrentBinarySearch(half58, half617)
    go ConcurrentBinarySearch(half617, half44)
    go ConcurrentBinarySearch(half44, half618)
    go ConcurrentBinarySearch(half618, half59)
    go ConcurrentBinarySearch(half59, half619)
    go ConcurrentBinarySearch(half619, half32)
    go ConcurrentBinarySearch(half32, half620)
    go ConcurrentBinarySearch(half620, half510)
    go ConcurrentBinarySearch(half510, half621)
    go ConcurrentBinarySearch(half621, half45)
    go ConcurrentBinarySearch(half45, half622)
    go ConcurrentBinarySearch(half622, half511)
    go ConcurrentBinarySearch(half511, half623)
    go ConcurrentBinarySearch(half623, half21)
    go ConcurrentBinarySearch(half21, half624)
    go ConcurrentBinarySearch(half624, half512)
    go ConcurrentBinarySearch(half512, half625)
    go ConcurrentBinarySearch(half625, half46)
    go ConcurrentBinarySearch(half46, half626)
    go ConcurrentBinarySearch(half626, half513)
    go ConcurrentBinarySearch(half513, half627)
    go ConcurrentBinarySearch(half627, half33)
    go ConcurrentBinarySearch(half33, half628)
    go ConcurrentBinarySearch(half628, half514)
    go ConcurrentBinarySearch(half514, half629)
    go ConcurrentBinarySearch(half629, half47)
    go ConcurrentBinarySearch(half47, half630)
    go ConcurrentBinarySearch(half630, half515)
    go ConcurrentBinarySearch(half515, half631)
    go ConcurrentBinarySearch(half631, half0)
    go ConcurrentBinarySearch(half0, half632)
    go ConcurrentBinarySearch(half632, half516)
    go ConcurrentBinarySearch(half516, half633)
    go ConcurrentBinarySearch(half633, half48)
    go ConcurrentBinarySearch(half48, half634)
    go ConcurrentBinarySearch(half634, half517)
    go ConcurrentBinarySearch(half517, half635)
    go ConcurrentBinarySearch(half635, half34)
    go ConcurrentBinarySearch(half34, half636)
     go ConcurrentBinarySearch(half636, half518)
    go ConcurrentBinarySearch(half518, half637)
    go ConcurrentBinarySearch(half637, half49)
    go ConcurrentBinarySearch(half49, half638)
    go ConcurrentBinarySearch(half638, half519)
    go ConcurrentBinarySearch(half519, half639)
    go ConcurrentBinarySearch(half639, half22)
    go ConcurrentBinarySearch(half22, half640)
    go ConcurrentBinarySearch(half640, half520)
    go ConcurrentBinarySearch(half520, half641)
    go ConcurrentBinarySearch(half641, half410)
    go ConcurrentBinarySearch(half410, half642)
    go ConcurrentBinarySearch(half642, half521)
    go ConcurrentBinarySearch(half521, half643)
    go ConcurrentBinarySearch(half643, half35)
    go ConcurrentBinarySearch(half35, half644)
    go ConcurrentBinarySearch(half644, half522)
    go ConcurrentBinarySearch(half522, half645)
    go ConcurrentBinarySearch(half645, half411)
    go ConcurrentBinarySearch(half411, half646)
    go ConcurrentBinarySearch(half646, half523)
    go ConcurrentBinarySearch(half523, half647)
    go ConcurrentBinarySearch(half647, half11)
    go ConcurrentBinarySearch(half11, half648)
    go ConcurrentBinarySearch(half648, half524)
    go ConcurrentBinarySearch(half524, half649)
    go ConcurrentBinarySearch(half649, half412)
    go ConcurrentBinarySearch(half412, half650)
    go ConcurrentBinarySearch(half650, half525)
    go ConcurrentBinarySearch(half525, half651)
    go ConcurrentBinarySearch(half651, half36)
    go ConcurrentBinarySearch(half36, half652)
    go ConcurrentBinarySearch(half652, half526)
    go ConcurrentBinarySearch(half526, half653)
    go ConcurrentBinarySearch(half653, half413)
    go ConcurrentBinarySearch(half413, half654)
    go ConcurrentBinarySearch(half654, half527)
    go ConcurrentBinarySearch(half527, half655)
    go ConcurrentBinarySearch(half655, half23)
    go ConcurrentBinarySearch(half23, half656)
    go ConcurrentBinarySearch(half656, half528)
    go ConcurrentBinarySearch(half528, half657)
    go ConcurrentBinarySearch(half657, half414)
    go ConcurrentBinarySearch(half414, half658)
    go ConcurrentBinarySearch(half658, half529)
    go ConcurrentBinarySearch(half529, half659)
    go ConcurrentBinarySearch(half659, half37)
    go ConcurrentBinarySearch(half37, half660)
    go ConcurrentBinarySearch(half660, half530)
    go ConcurrentBinarySearch(half530, half661)
    go ConcurrentBinarySearch(half661, half415)
    go ConcurrentBinarySearch(half415, half662)
    go ConcurrentBinarySearch(half662, half531)
    go ConcurrentBinarySearch(half531, half663)
    go ConcurrentBinarySearch(half663, high)
}

func main() {
    Init()

    low := make([]byte, kLength)
    high := make([]byte, kLength)
    for i := 0; i < kLength; i++ {
        low[i] = 0
        high[i] = 255
    }

    if Search(high) != nil {
        kResult = high
    } else {
        high[kLength-1] = 254

        t := time.Now()
        go ConcurrentBinarySearch(low, high)
        // go BinarySearch(low, high)
        <-kFound
        elapsed := time.Since(t)
        fmt.Println("App elapsed: ", elapsed)
    }

    fmt.Println("TargetKey: ", string(kTargetKey))
    fmt.Println("Result:    ", string(kResult))
}
