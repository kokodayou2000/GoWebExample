package main

type P struct {
	id   string
	name string
}

func logger(p1 P) string {
	println("p1.id = " + p1.id + "\np1.name = " + p1.name)
	return "logger"
}
func do(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	p1 := P{
		id:   "1",
		name: "tom",
	}
	_ = do("string", logger(p1))

}
