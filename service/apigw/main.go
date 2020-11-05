package main

import "github.com/yguilai/go-cloud-storage/service/apigw/route"

func main() {
	r := route.Router()
	r.Run(":8080")
}
