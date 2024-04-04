package cmd

func main() {
	server := NewAPIServer(":8080")
	server.Run()
}