package main

func main() {
	api := newAPIServer(":9595", nil)

	store := NewStore(db)

}
