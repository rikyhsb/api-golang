// main.go

package main

func main() {
	a := App{}
	// konfigurasi database disini
	a.Initialize("DB_USERNAME", "DB_PASSWORD", "DB_NAME")
	a.Run(":8080")
}
