package main

import "challenge_sesi_2_api/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
