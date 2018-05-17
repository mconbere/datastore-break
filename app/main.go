package main

import (
	"fmt"
	"net/http"

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
		Bytes: []byte{0xe2, 0x82, 0x28},
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to put simple ByteBucket in datastore: %v", err)
		return
	}

	bb := ByteBucket{}
	if err := datastore.Get(ctx, key, &bb); err != nil {
		fmt.Fprintf(w, "Failed to get simple ByteBucket from datastore: %v", err)
	}

	fmt.Fprintf(w, "Success")
}
