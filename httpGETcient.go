package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func mainHTTP() {
	fmt.Print("sent:\nHTTP GET:\n")
	res, err := http.Get("https://www.tagesschau.de/"	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("received:")
	io.Copy(os.Stdout, res.Body)
}

