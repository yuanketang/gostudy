package bird

import "fmt"

type (
	Bird1 struct {
		Name string
	}
)

// 指针接收者
func (b *Bird1) Fly2() {
	fmt.Printf("%s 飞\n", b.Name)
}
