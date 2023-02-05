package main

func main() {
	server := NewServer(":3550")
	server.RunServer()
}
