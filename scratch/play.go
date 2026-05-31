package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"time"
)

func test1() {
	var i int = 20
	var f float32 = float32(i)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("f = %f\n", f)
	//fmt.Println("i = ", i, "\nf = ", f)
}

func test2() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	println(b)
	println(smallI)
	println(bigI)
	println("--------")
	println(b + 1)
	println(smallI + 1)
	println(bigI + 1)

}

func test3() {
	intSet := make(map[int]bool)
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println((intSet[5]))
	// fmt.Println(intSet[500])
	// if intSet[100] {
	// 	fmt.Println("100 is in the Set")
	// }
	if _, ok := intSet[2]; ok {
		fmt.Println(ok)
	}
}

func test4() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:5]
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
}

func test5() {
	var message = "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Printf("%c \n", runes[3])
}

func test6() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	pedro := Employee{
		"pedro",
		"mango",
		12,
	}

	joao := Employee{
		"joao",
		"manuel",
		31,
	}

	var carlos Employee
	carlos.firstName = "carlos"
	carlos.lastName = "emanuel"
	carlos.id = 51

	fmt.Println(carlos, pedro, joao)
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type testc struct {
	num int64
	stg string
}

func parseFlags() (string, string) {
	addr := flag.String("addr", "portMango", "HTTP network address")

	dbName := flag.String("dbName", "mango1", "DB name")
	dbPort := flag.String("dbPort", "mango2", "DB port")
	dbHost := flag.String("dbHost", "mango3", "DB host address")
	dbUser := flag.String("dbUser", "mango4", "DB User")
	dbPass := flag.String("dbPass", "mango5", "DB password")
	flag.Parse()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5", *dbHost, *dbPort, *dbUser, *dbPass, *dbName)
	*addr = fmt.Sprintf(":%s", *addr)

	return dsn, *addr
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// Have to build to run this function, and then direct the result to a file
// ex: ./play > out.gif
func lissajous(out io.Writer) {
	const (
		cycles = 5
		// number of complete x oscillator revolutions
		res  = 0.001 // angular resolution
		size = 100
		// image canvas covers [-size..+size]
		nframes = 64
		// number of animation frames
		delay = 8
	// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
