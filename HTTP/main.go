package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	//bs := make([]byte, 99999)
	//bs byte sliceba fogjuk beleolvasni a Readnél
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//shorter version:
	//Copy(Writer, Reader) -> kiírja az elsőre, amit beolvasott a másodikról
	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

//Implement Write interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(" bytes written: ", len(bs))
	return len(bs), nil
}
