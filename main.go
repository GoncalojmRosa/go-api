package main

func main() {
	svr := NewAPIServer(":8080")
	svr.Start()
}