package main

import (
	"fmt"
	"net/http"

	_ "google.golang.org/appengine/remote_api"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type ByteBucket struct {
	Bytes []byte // Raw bytes
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	incKey := datastore.NewIncompleteKey(ctx, "ByteBucket", nil)
	key, err := datastore.Put(ctx, incKey, &ByteBucket{
		Bytes: []byte{0xae, 0x6c, 0x6e, 0xa8, 0x39, 0xbd, 0x6e, 0x12, 0x8f, 0x21, 0x76, 0x70, 0xdc, 0x28, 0x6a, 0x42},
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to put simple ByteBucket in datastore: %v", err)
		return
	}

	bb := ByteBucket{}
	if err := datastore.Get(ctx, key, &bb); err != nil {
		fmt.Fprintf(w, "Failed to get simple ByteBucket from datastore: %v", err)
		return
	}

	fmt.Fprintf(w, "Success")
}
