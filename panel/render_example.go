package panel

import (
	"fmt"
	"time"
)

func ExampleRender() {
	const s = `
	First Line
	Second Line
	Third Line
	Hello
	This is go-music`

	fmt.Println(RenderText(s, Spring))
	fmt.Println(RenderText(s, Autumn))
	fmt.Println(RenderText(s, Winter))
	fmt.Println(RenderText(s, Rose))
	fmt.Println(RenderText(s, Valentine))
}

func ExampleShowProgressBar() {
	for {
		time.Sleep(10 * time.Millisecond)
		ShowProgressBar()
	}
}
