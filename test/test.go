package main

import (
	"fmt"
	"github.com/jiguorui/bitcask"
)

func test_bucket() {
	bucket, err := bitcask.NewBucket("001.ar", 1)
	if err != nil {
		fmt.Println(err)
	}

	defer bucket.Close()

	n, err := bucket.Write([]byte("abcdef"))
	fmt.Printf("%d,%v", n, err)
	offset, err := bucket.GetWriteOffset()
	fmt.Printf("%d,%v", offset, err)

	b, err := bucket.Read(12, 2)
	fmt.Println(string(b))

}

func test_keydir() {
	kd := bitcask.NewKeyDir()
	kd.Set("abc", uint32(16), uint32(0), int32(0), int32(0))

	e, _:= kd.Get("abc")
	fmt.Printf("%s,%d\n", e.Key, e.Total_size)

}

func test_bitcask() {
	bc, err := bitcask.Open(".")
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer bc.Close()

	//bc.Set("abc", []byte("defghijklmnopqrstuvwxyz"))
	b, err := bc.Get("abc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", b)
}

func main() {
	//test_bucket()
	//test_keydir()
	test_bitcask()
	// fmt.Printf("Hello, bitcast\n")
	// bc, err := bitcask.Open(".")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
	// defer bc.Close()

	// bc.Put("abc", []byte("defghijklmnopqrstuvwxyz"))

	// fmt.Printf("%d\n", bitcask)
}
