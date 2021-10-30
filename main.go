package main

func main() {

	// initialize configuration
	InitConfig()

	db := dbConnect()

	defer db.Close()

}
