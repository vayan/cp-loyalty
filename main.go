package main

func main() {
	a := App{}
	a.Initialize("development.db")
	a.Run(":8080")
}
