package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type bit byte

const (
	b0 bit = 0b0000_0001
	b1 bit = 0b0000_0010
	b2 bit = 0b0000_0100
	b3 bit = 0b0000_1000
	b4 bit = 0b0001_0000
	b5 bit = 0b0010_0000
	b6 bit = 0b0100_0000
	b7 bit = 0b1000_0000
)

type bits struct {
	timestamp int64
	value     int64
}

func FirstPass(data []byte) map[int64][]bits {

	dict := make(map[int64][]bits)
	var leadingZeros int = 0
	var value int64 = 0
	var run int = 1
	var values []bits
	var DEBUG int

	for _, Byte := range data {
		one := func() bool {
			if Byte&byte(b0) > 0 {
				return true
			} else {
				return false
			}
		}
		two := func() bool {
			if Byte&byte(b1) > 0 {
				return true
			} else {
				return false
			}
		}
		three := func() bool {
			if Byte&byte(b2) > 0 {
				return true
			} else {
				return false
			}
		}
		four := func() bool {
			if Byte&byte(b3) > 0 {
				return true
			} else {
				return false
			}
		}
		five := func() bool {
			if Byte&byte(b4) > 0 {
				return true
			} else {
				return false
			}
		}
		six := func() bool {
			if Byte&byte(b5) > 0 {
				return true
			} else {
				return false
			}
		}
		seven := func() bool {
			if Byte&byte(b6) > 0 {
				return true
			} else {
				return false
			}
		}
		eight := func() bool {
			if Byte&byte(b7) > 0 {
				return true
			} else {
				return false
			}
		}

		if eight() {
			value = value + 10000000
		} else {
			leadingZeros++
		}

		if seven() {
			value = value + 1000000
		} else {
			if !eight() {
				leadingZeros++
			}
		}

		if six() {
			value = value + 100000
		} else {
			if !eight() && !seven() {
				leadingZeros++
			}
		}

		if five() {
			value = value + 10000
		} else {
			if !eight() && !seven() && !six() {
				leadingZeros++
			}
		}

		if four() {
			value = value + 1000
		} else {
			if !eight() && !seven() && !six() && !five() {
				leadingZeros++
			}
		}

		if three() {
			value = value + 100
		} else {
			if !eight() && !seven() && !six() && !five() && !four() {
				leadingZeros++
			}
		}

		if two() {
			value = value + 10
		} else {
			if !eight() && !seven() && !six() && !five() && !four() && !three() {
				leadingZeros++
			}
		}
		if one() {
			value = value + 1
		} else {
			if !eight() && !seven() && !six() && !five() && !four() && !three() && !two() {
				leadingZeros++
			}
		}
		run++
		//fmt.Println(value) //DEBUG ONLY
		if DEBUG == 2 {
			//os.Exit(1)
		}
		//os.Exit(1)
		values = append(values, bits{time.Now().UnixNano(), value})

		if run == 9 {
			dict[time.Now().UnixNano()] = values
			//fmt.Println(values)
			value = 0
			run = 0
			DEBUG++
			//os.Exit(1) ///USEFULL! DEBUG
		}
	}
	return dict
}

func SecondPass(dict map[int64][]bits) []byte {
	//var int64_array []int64
	var data []byte
	var done bool = false

	//t := make([]byte, 8)
	v := make([]byte, 8)

	var i int64 = 0
	var j int64 = 0

	for _, n := range dict {
		for i = 0; i < int64(len(n)); i = i + 9 {
			for j = 0; j < 9; j++ {
				if i+j > int64(len(n)-1) {
					done = true
					break
				}
				//m := n[i+j].timestamp
				o := n[i+j].value

				//fmt.Println(o)

				/*t = []byte{
				byte(0xff & m),
				byte(0xff & (m >> 8)),
				byte(0xff & (m >> 16)),
				byte(0xff & (m >> 24)),
				byte(0xff & (m >> 32)),
				byte(0xff & (m >> 40)),
				byte(0xff & (m >> 48)),
				byte(0xff & (m >> 56))}
				*/
				v = []byte{
					byte(0xff & o),
					byte(0xff & (o >> 8)),
					byte(0xff & (o >> 16)),
					byte(0xff & (o >> 24)),
					byte(0xff & (o >> 32)),
					byte(0xff & (o >> 40)),
					byte(0xff & (o >> 48)),
					byte(0xff & (o >> 56)),
				}

				//binary.LittleEndian.PutUint64(b, uint64(n))

				/*fmt.Println(v[0])
				fmt.Println(v[1])
				fmt.Println(v[2])
				fmt.Println(v[3])
				fmt.Println(v[4])
				fmt.Println(v[5])
				fmt.Println(v[6])
				fmt.Println(v[7])
				os.Exit(1)
				*/

				//data = append(data, t...)
				data = append(data, v...)

			}
			//if i == 9 { //DEBUG
			//	os.Exit(1)
			//}
			if done {
				break
			}
		}
	}
	return data
}

func ThirdPass(folded_data []byte) []byte {
	return RebuildFile(folded_data)
}

func FourthPass(serialized []byte) []byte {
	return BuildBlocks(serialized)
}

func FifthPass(data []byte) []byte {
	return data
}

func RebuildBits(Digits ...int) []byte {

	var Rebuilt []byte

	NbDigits := len(Digits)

	switch NbDigits {
	case 1:
		var one int = Digits[0]
		//fmt.Println(one)
		if one == 0 {
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
		} else {
			//for i := 1; i <= 9; i++ {
			switch one {
			case 1:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
			case 2:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 3:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 4:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 5:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 6:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 7:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 8:
				Rebuilt = append(Rebuilt, byte(0))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			case 9:
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
				Rebuilt = append(Rebuilt, byte(1))
			}
			//}
		}
	case 2:
		var one, two int = Digits[0], Digits[1]
		//fmt.Println(one, two)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
	case 3:
		var one, two, three int = Digits[0], Digits[1], Digits[2]
		//fmt.Println(one, two, three)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
	case 4:
		var one, two, three, four int = Digits[0], Digits[1], Digits[2], Digits[3]
		//fmt.Println(one, two, three, four)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
		switch four {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 9:
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		}
	case 5:
		var one, two, three, four, five int = Digits[0], Digits[1], Digits[2], Digits[3], Digits[4]
		//fmt.Println(one, two, three, four, five)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
		switch four {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 9:
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		}
		switch five {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 16:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 9:
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		}
	case 6:
		var one, two, three, four, five, six int = Digits[0], Digits[1], Digits[2], Digits[3], Digits[4], Digits[5]
		//fmt.Println(one, two, three, four, five, six)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
		switch four {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 9:
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		}
		switch five {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 16:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 9:
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		}
		switch six {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 32:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 9:
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		}
	case 7:
		var one, two, three, four, five, six, seven int = Digits[0], Digits[1], Digits[2], Digits[3], Digits[4], Digits[5], Digits[6]
		//fmt.Println(one, two, three, four, five, six, seven)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
		switch four {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 9:
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		}
		switch five {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 16:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 9:
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		}
		switch six {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 32:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 9:
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		}
		switch seven {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 64:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 9:
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		}
	case 8:
		var one, two, three, four, five, six, seven, eight int = Digits[0], Digits[1], Digits[2], Digits[3], Digits[4], Digits[5], Digits[6], Digits[7]
		//fmt.Println(one, two, three, four, five, six, seven, eight)
		switch one {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		case 9:
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
			Rebuilt = append(Rebuilt, byte(1))
		}
		switch two {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		case 9:
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
			Rebuilt = append(Rebuilt, byte(2))
		}
		switch three {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		case 9:
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
			Rebuilt = append(Rebuilt, byte(4))
		}
		switch four {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		case 9:
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
			Rebuilt = append(Rebuilt, byte(8))
		}
		switch five {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		case 9:
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
			Rebuilt = append(Rebuilt, byte(16))
		}
		switch six {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		case 9:
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
			Rebuilt = append(Rebuilt, byte(32))
		}
		switch seven {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		case 9:
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
			Rebuilt = append(Rebuilt, byte(64))
		}
		switch eight {
		case 1:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
		case 2:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 3:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 4:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 5:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 6:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 7:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 8:
			Rebuilt = append(Rebuilt, byte(0))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		case 9:
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
			Rebuilt = append(Rebuilt, byte(128))
		}
	}
	return Rebuilt
}

func RebuildFile(data []byte) []byte {
	var Rebuilt_Data []byte
	b := make([]uint8, 8)
	var count = 1
	var value int64
	var done bool

	for _, folded_byte := range data {
		if count <= 8 {
			//fmt.Println(folded_byte) //yes there is data there!
			//os.Exit(1)
			b = append(b, folded_byte)
		} else {
			for i := 7; i >= 0; i-- {
				switch b[i] {
				case 255:
					value += 11111111
				case 254:
					value += 11111110
				case 253:
					value += 11111101
				case 252:
					value += 11111100
				case 251:
					value += 11111011
				case 250:
					value += 11111010
				case 249:
					value += 11111001
				case 248:
					value += 11111000
				case 247:
					value += 11110111
				case 246:
					value += 11110110
				case 245:
					value += 11110101
				case 244:
					value += 11110100
				case 243:
					value += 11110011
				case 242:
					value += 11110010
				case 241:
					value += 11110001
				case 240:
					value += 11110000
				case 239:
					value += 11101111
				case 238:
					value += 11101110
				case 237:
					value += 11101101
				case 236:
					value += 11101100
				case 235:
					value += 11101011
				case 234:
					value += 11101010
				case 233:
					value += 11101001
				case 232:
					value += 11101000
				case 231:
					value += 11100111
				case 230:
					value += 11100110
				case 229:
					value += 11100101
				case 228:
					value += 11100100
				case 227:
					value += 11100011
				case 226:
					value += 11100010
				case 225:
					value += 11100001
				case 224:
					value += 11100000
				case 223:
					value += 11011111
				case 222:
					value += 11011110
				case 221:
					value += 11011101
				case 220:
					value += 11011100
				case 219:
					value += 11011011
				case 218:
					value += 11011010
				case 217:
					value += 11011001
				case 216:
					value += 11011000
				case 215:
					value += 11010111
				case 214:
					value += 11010110
				case 213:
					value += 11010101
				case 212:
					value += 11010100
				case 211:
					value += 11010011
				case 210:
					value += 11010010
				case 209:
					value += 11010001
				case 208:
					value += 11010000
				case 207:
					value += 11001111
				case 206:
					value += 11001110
				case 205:
					value += 11001101
				case 204:
					value += 11001100
				case 203:
					value += 11001011
				case 202:
					value += 11001010
				case 201:
					value += 11001001
				case 200:
					value += 11001000
				case 199:
					value += 11000111
				case 198:
					value += 11000110
				case 197:
					value += 11000101
				case 196:
					value += 11000100
				case 195:
					value += 11000011
				case 194:
					value += 11000010
				case 193:
					value += 11000001
				case 192:
					value += 11000000
				case 191:
					value += 10111111
				case 190:
					value += 10111110
				case 189:
					value += 10111101
				case 188:
					value += 10111100
				case 187:
					value += 10111011
				case 186:
					value += 10111010
				case 185:
					value += 10111001
				case 184:
					value += 10111000
				case 183:
					value += 10110111
				case 182:
					value += 10110110
				case 181:
					value += 10110101
				case 180:
					value += 10110100
				case 179:
					value += 10110011
				case 178:
					value += 10110010
				case 177:
					value += 10110001
				case 176:
					value += 10110000
				case 175:
					value += 10101111
				case 174:
					value += 10101110
				case 173:
					value += 10101101
				case 172:
					value += 10101100
				case 171:
					value += 10101011
				case 170:
					value += 10101010
				case 169:
					value += 10101001
				case 168:
					value += 10101000
				case 167:
					value += 10100111
				case 166:
					value += 10100110
				case 165:
					value += 10100101
				case 164:
					value += 10100100
				case 163:
					value += 10100011
				case 162:
					value += 10100010
				case 161:
					value += 10100001
				case 160:
					value += 10100000
				case 159:
					value += 10011111
				case 158:
					value += 10011110
				case 157:
					value += 10011101
				case 156:
					value += 10011100
				case 155:
					value += 10011011
				case 154:
					value += 10011010
				case 153:
					value += 10011001
				case 152:
					value += 10011000
				case 151:
					value += 10010111
				case 150:
					value += 10010110
				case 149:
					value += 10010101
				case 148:
					value += 10010100
				case 147:
					value += 10010011
				case 146:
					value += 10010010
				case 145:
					value += 10010001
				case 144:
					value += 10010000
				case 143:
					value += 10001111
				case 142:
					value += 10001110
				case 141:
					value += 10001101
				case 140:
					value += 10001100
				case 139:
					value += 10001011
				case 138:
					value += 10001010
				case 137:
					value += 10001001
				case 136:
					value += 10001000
				case 135:
					value += 10000111
				case 134:
					value += 10000110
				case 133:
					value += 10000101
				case 132:
					value += 10000100
				case 131:
					value += 10000011
				case 130:
					value += 10000010
				case 129:
					value += 10000001
				case 128:
					value += 10000000
				case 127:
					value += 01111111
				case 126:
					value += 01111110
				case 125:
					value += 01111101
				case 124:
					value += 01111100
				case 123:
					value += 01111011
				case 122:
					value += 01111010
				case 121:
					value += 01111001
				case 120:
					value += 01111000
				case 119:
					value += 01110111
				case 118:
					value += 01110110
				case 117:
					value += 01110101
				case 116:
					value += 01110100
				case 115:
					value += 01110011
				case 114:
					value += 01110010
				case 113:
					value += 01110001
				case 112:
					value += 01110000
				case 111:
					value += 01101111
				case 110:
					value += 01101110
				case 109:
					value += 01101101
				case 108:
					value += 01101100
				case 107:
					value += 01101011
				case 106:
					value += 01101010
				case 105:
					value += 01101001
				case 104:
					value += 01101000
				case 103:
					value += 01100111
				case 102:
					value += 01100110
				case 101:
					value += 01100101
				case 100:
					value += 01100100
				case 99:
					value += 01100011
				case 98:
					value += 01100010
				case 97:
					value += 01100001
				case 96:
					value += 01100000
				case 95:
					value += 01011111
				case 94:
					value += 01011110
				case 93:
					value += 01011101
				case 92:
					value += 01011100
				case 91:
					value += 01011011
				case 90:
					value += 01011010
				case 89:
					value += 01011001
				case 88:
					value += 01011000
				case 87:
					value += 01010111
				case 86:
					value += 01010110
				case 85:
					value += 01010101
				case 84:
					value += 01010100
				case 83:
					value += 01010011
				case 82:
					value += 01010010
				case 81:
					value += 01010001
				case 80:
					value += 01010000
				case 79:
					value += 01001111
				case 78:
					value += 01001110
				case 77:
					value += 01001101
				case 76:
					value += 01001100
				case 75:
					value += 01001011
				case 74:
					value += 01001010
				case 73:
					value += 01001001
				case 72:
					value += 01001000
				case 71:
					value += 01000111
				case 70:
					value += 01000110
				case 69:
					value += 01000101
				case 68:
					value += 01000100
				case 67:
					value += 01000011
				case 66:
					value += 01000010
				case 65:
					value += 01000001
				case 64:
					value += 01000000
				case 63:
					value += 00111111
				case 62:
					value += 00111110
				case 61:
					value += 00111101
				case 60:
					value += 00111100
				case 59:
					value += 00111011
				case 58:
					value += 00111010
				case 57:
					value += 00111001
				case 56:
					value += 00111000
				case 55:
					value += 00110111
				case 54:
					value += 00110110
				case 53:
					value += 00110101
				case 52:
					value += 00110100
				case 51:
					value += 00110011
				case 50:
					value += 00110010
				case 49:
					value += 00110001
				case 48:
					value += 00110000
				case 47:
					value += 00101111
				case 46:
					value += 00101110
				case 45:
					value += 00101101
				case 44:
					value += 00101100
				case 43:
					value += 00101011
				case 42:
					value += 00101010
				case 41:
					value += 00101001
				case 40:
					value += 00101000
				case 39:
					value += 00100111
				case 38:
					value += 00100110
				case 37:
					value += 00100101
				case 36:
					value += 00100100
				case 35:
					value += 00100011
				case 34:
					value += 00100010
				case 33:
					value += 00100001
				case 32:
					value += 00100000
				case 31:
					value += 00011111
				case 30:
					value += 00011110
				case 29:
					value += 00011101
				case 28:
					value += 00011100
				case 27:
					value += 00011011
				case 26:
					value += 00011010
				case 25:
					value += 00011001
				case 24:
					value += 00011000
				case 23:
					value += 00010111
				case 22:
					value += 00010110
				case 21:
					value += 00010101
				case 20:
					value += 00010100
				case 19:
					value += 00010011
				case 18:
					value += 00010010
				case 17:
					value += 00010001
				case 16:
					value += 00010000
				case 15:
					value += 00001111
				case 14:
					value += 00001110
				case 13:
					value += 00001101
				case 12:
					value += 00001100
				case 11:
					value += 00001011
				case 10:
					value += 00001010
				case 9:
					value += 00001001
				case 8:
					value += 00001000
				case 7:
					value += 00000111
				case 6:
					value += 00000110
				case 5:
					value += 00000101
				case 4:
					value += 00000100
				case 3:
					value += 00000011
				case 2:
					value += 00000010
				case 1:
					value += 00000001
				case 0:
					value += 00000000
				}
			}
			done = true
		}

		count++

		if done {
			var digit_1, digit_2, digit_3, digit_4, digit_5, digit_6, digit_7, digit_8 string = "", "", "", "", "", "", "", ""
			//fmt.Println(value)
			str_value := fmt.Sprintf("%d", value)
			//fmt.Println(str_value)
			//fmt.Println(len(str_value))
			len_str_value := len(str_value)
			switch len_str_value {
			case 8:
				digit_1, digit_2, digit_3, digit_4, digit_5, digit_6, digit_7, digit_8 = string(str_value[0]), string(str_value[1]), string(str_value[2]), string(str_value[3]), string(str_value[4]), string(str_value[5]), string(str_value[6]), string(str_value[7])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				four, _ := strconv.Atoi(digit_4)
				five, _ := strconv.Atoi(digit_5)
				six, _ := strconv.Atoi(digit_6)
				seven, _ := strconv.Atoi(digit_7)
				eight, _ := strconv.Atoi(digit_8)
				//fmt.Printf("%d%d%d%d%d%d%d%d\n", one, two, three, four, five, six, seven, eight)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three, four, five, six, seven, eight)...)
			case 7:
				digit_1, digit_2, digit_3, digit_4, digit_5, digit_6, digit_7 = string(str_value[0]), string(str_value[1]), string(str_value[2]), string(str_value[3]), string(str_value[4]), string(str_value[5]), string(str_value[6])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				four, _ := strconv.Atoi(digit_4)
				five, _ := strconv.Atoi(digit_5)
				six, _ := strconv.Atoi(digit_6)
				seven, _ := strconv.Atoi(digit_7)
				//fmt.Printf("%d%d%d%d%d%d%d\n", one, two, three, four, five, six, seven)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three, four, five, six, seven)...)
			case 6:
				digit_1, digit_2, digit_3, digit_4, digit_5, digit_6 = string(str_value[0]), string(str_value[1]), string(str_value[2]), string(str_value[3]), string(str_value[4]), string(str_value[5])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				four, _ := strconv.Atoi(digit_4)
				five, _ := strconv.Atoi(digit_5)
				six, _ := strconv.Atoi(digit_6)
				//fmt.Printf("%d%d%d%d%d%d\n", one, two, three, four, five, six)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three, four, five, six)...)
			case 5:
				digit_1, digit_2, digit_3, digit_4, digit_5 = string(str_value[0]), string(str_value[1]), string(str_value[2]), string(str_value[3]), string(str_value[4])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				four, _ := strconv.Atoi(digit_4)
				five, _ := strconv.Atoi(digit_5)
				//fmt.Printf("%d%d%d%d%d\n", one, two, three, four, five)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three, four, five)...)
			case 4:
				digit_1, digit_2, digit_3, digit_4 = string(str_value[0]), string(str_value[1]), string(str_value[2]), string(str_value[3])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				four, _ := strconv.Atoi(digit_4)
				//fmt.Printf("%d%d%d%d\n", one, two, three, four)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three, four)...)
			case 3:
				digit_1, digit_2, digit_3 = string(str_value[0]), string(str_value[1]), string(str_value[2])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				three, _ := strconv.Atoi(digit_3)
				//fmt.Printf("%d%d%d\n", one, two, three)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two, three)...)
			case 2:
				digit_1, digit_2 = string(str_value[0]), string(str_value[1])
				one, _ := strconv.Atoi(digit_1)
				two, _ := strconv.Atoi(digit_2)
				//fmt.Printf("%d%d\n", one, two)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one, two)...)
			case 1:
				digit_1 = string(str_value[0])
				one, _ := strconv.Atoi(digit_1)
				//fmt.Printf("%d\n", one)
				Rebuilt_Data = append(Rebuilt_Data, RebuildBits(one)...)
			}

			//os.Exit(1)

			value = 0
			count = 1
			b = nil
			done = false
		}
	}
	return Rebuilt_Data
}

func BuildBlocks(serialized []byte) []byte {
	series := make([]byte, 9)
	var value int
	var bit8 byte
	var bits8 []byte

	for i := range serialized {
		if i%8 == 0 {
			series = append(series, serialized[i])
		} else {
			for j := range series {
				value += int(series[j])
			}
			//fmt.Println(value)

			switch value {
			case 1:
				bit8 += byte(b0)
			case 2:
				bit8 += byte(b1)
			case 4:
				bit8 += byte(b2)
			case 8:
				bit8 += byte(b3)
			case 16:
				bit8 += byte(b4)
			case 32:
				bit8 += byte(b5)
			case 64:
				bit8 += byte(b6)
			case 128:
				bit8 += byte(b7)
			default:
				bit8 += byte(0b0000_0000)
			}

			bits8 = append(bits8, bit8)
			series = nil //check that this works.
			value = 0
		}
	}
	return bits8
}

func readBytesFromFile(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// writeBytesToFile self-explanatory
func writeBytesToFile(filename string, data []byte) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfile.Close()
	outfile.Write(data)

	return nil
}

func hide(data []byte, block []byte) []byte {
	var KeyMap []int64
	var i int64
	for _, with := range data {
		for _, entropy := range block {
			if with == entropy {
				KeyMap = append(KeyMap, i)
				break
			}
			i++
		}
		i = 0
	}

	jsonData, err := json.Marshal(&KeyMap)
	if err != nil {
		panic(err)
	}

	return jsonData
}

func uncover(jsonData []byte, block []byte) []byte {
	var KeyMap []int64
	var Recovered []byte
	err := json.Unmarshal(jsonData, &KeyMap)
	if err != nil {
		panic(err)
	}

	for _, Key := range KeyMap {
		Recovered = append(Recovered, block[Key])
	}
	return Recovered
}

func main() {
	var e string
	var d string
	var b string
	var k string
	var seed string
	var block string
	var key string
	var outFile string

	flag.StringVar(&e, "e", "", "encode file")
	flag.StringVar(&d, "d", "", "decode file")
	flag.StringVar(&b, "b", "", "decode file")
	flag.StringVar(&k, "k", "", "decode file")

	flag.Parse()

	if len(os.Args) < 3 {
		fmt.Println("Usage: Cypher -e <original file> [OR] -b <.block file> -k <.key file> [Respectively Encode and Decode]")
		os.Exit(1)
	}
	if e != "" {
		seed = e
		block = seed + ".block"
		key = seed + ".key"
		outFile = seed + ".recovered"

		data, err := readBytesFromFile(seed)
		if err != nil {
			panic(err)
		}

		bits := FirstPass(data)
		folded_data := SecondPass(bits)
		serialized := ThirdPass(folded_data)
		BuiltEntropy := FourthPass(serialized)

		writeBytesToFile(block, BuiltEntropy)

		//------------------ENCODE----------------\\
		jsonData := hide(data, BuiltEntropy)

		writeBytesToFile(key, jsonData)
	}

	if b != "" && k != "" {
		block = b
		key = k
		ext := strings.Split(b, ".")[1]
		outFile = strings.Split(b, ".")[0] + "." + ext

		BuiltEntropy, err := readBytesFromFile(block)
		if err != nil {
			panic(err)
		}

		jsonData, err := readBytesFromFile(key)
		if err != nil {
			panic(err)
		}

		RecoveredData := uncover(jsonData, BuiltEntropy)

		writeBytesToFile(outFile, RecoveredData)
	}
}
