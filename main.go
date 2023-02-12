package main

import routers "localhost.com/clean-gin/config"

func main() {
	r := routers.InitRouter()
	r.Run(":9501")
}
