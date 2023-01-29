package tools_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/mdigger/tools"
)

func ExampleIsEmpty() {
	data := struct{ test string }{test: "test"}
	if !tools.IsEmpty(data) {
		fmt.Println(data.test)
	}

	// Output: test
}

func ExampleTry() {
	if !tools.Try(func() error {
		panic("test")
	}) {
		fmt.Println("panic")
	}

	if !tools.Try(func() error {
		return errors.New("error")
	}) {
		fmt.Println("error")
	}

	if tools.Try(func() error {
		return nil
	}) {
		fmt.Println("ok")
	}

	// Output:
	// panic
	// error
	// ok
}

func ExampleAsync() {
	result := tools.Async(func() string {
		time.Sleep(time.Second)
		return "ok"
	})

	fmt.Println(<-result)

	// Output: ok
}

func ExampleToAnySlice() {
	result := tools.ToAnySlice([]string{"one", "two", "three"})
	fmt.Printf("%#+v\n", result)

	// Output: []interface {}{"one", "two", "three"}
}

func ExampleMax() {
	max := tools.Max(99, 13, 9, 0, 100, 11, 11)
	fmt.Println(max)

	// Output: 100
}

func ExampleMin() {
	min := tools.Min(99, 13, 9, 0, 100, -10, 11)
	fmt.Println(min)

	// Output: -10
}

func ExampleNth() {
	three := tools.Nth(2, "one", "too", "three", "four")
	fmt.Println(three)

	// Output: three
}

func ExampleSample() {
	result := tools.Sample("one", "too", "three", "four")
	fmt.Println(result)

	// Output: too
}

func ExampleTernary() {
	result := tools.Ternary(1 != 2, "true", "false")
	fmt.Println(result)

	// Output: true
}

func ExampleCoalesce() {
	result := tools.Coalesce("", "str")
	fmt.Println(result)

	// Output: str
}

func ExampleUnique() {
	result := tools.Unique("s1", "s2", "s1", "s3", "s2", "s1")
	fmt.Println(result)

	// Output: [s1 s2 s3]
}

func ExampleCompact() {
	result := tools.Compact("s1", "", "s1", "", "", "s1")
	fmt.Println(result)

	// Output: [s1 s1 s1]
}

func ExampleLen() {
	s := "строка из символов и 🤫"
	l := tools.Len(s)
	fmt.Println(l, len(s))

	// Output: 22 42
}

func ExampleTruncate() {
	s := "строка из символов и 🤫"
	result := tools.Truncate(s, 12)
	fmt.Println(result)

	// Output: строка из…
}

func ExampleMap() {
	var m tools.Map[string, int]

	go func() {
		for i := 5; i < 15; i++ {
			key := fmt.Sprintf("i:%02d", i)
			m.Store(key, i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("i:%02d", i)
			m.Store(key, i*2)
		}
	}()

	time.Sleep(time.Second)

	m.Range(func(key string, value int) bool {
		fmt.Println(key, value)
		return true
	})
}

func ExampleFanout() {
	var l tools.Fanout[string, int]

	ch1, ch2 := l.Listen("01"), l.Listen("02")

	go func() {
		for i1 := range ch1 {
			fmt.Print(i1, " ")
		}
	}()

	go func() {
		for i2 := range ch2 {
			fmt.Print(i2, " ")
		}
	}()

	l.Notify(1, 2, 3, 4, 5)

	time.Sleep(time.Second)

	// Output: 1 2 3 4 5 1 2 3 4 5
}
