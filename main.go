package main

import (
	hang "main/scripts"
	web "main/web"
)

func main() {
	hang.Initweb()
	web.CreateWebsite()
}
