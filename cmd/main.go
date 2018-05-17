package main

import (
	"fmt"
	"context"

	"golang.org/x/oauth2/google"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/remote_api"
)

type ByteBucket struct {
	Bytes []byte // Raw bytes
}

func main() {
	ctx := context.Background()

	hc, err := google.DefaultClient(ctx,
		"https://www.googleapis.com/auth/appengine.apis",
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/cloud-platform",
	)
	if err != nil {
		panic(err)
	}

	remoteCtx, err := remote_api.NewRemoteContext("datastore-break.appspot.com", hc)
	if err != nil {
		panic(err)
	}

	q := datastore.NewQuery("ByteBucket")
	bbs := []ByteBucket{}
	keys, err := q.GetAll(remoteCtx, &bbs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success: %v\n", keys)
}
