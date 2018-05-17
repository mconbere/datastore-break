# datastore-break

Go App Engine Datastore library code does not work for bytes fields containing
non-utf-8 arrays.

This is caused by two errors in the appengine Datastore library at:

    http://google.golang.org/appengine/datastore

Datastore is treating byte slices as strings. The latest Go Protobuf libraries
check for UTF-8 validity in these trhings. Latest Go Protobuf turns bytes
fields in `.proto`s into `[]byte` in the resulting `.pb.go`, but the
`datastore.proto` definition of a byte slice needs to be updated.

Reproducing this code was easiest for us when using the `remote_api` to
run datastore commands locally, as this appears to work correctly with a
deployed App Engine app.

## Reproduce

`app/` contains a small go appengine app that will run successful and add
invalid UTF-8 to the datastore.

`cmd/` contains a small go command that uses remote_api to fetch these
invalid UTF-8 objects. Running it with correct permissions will panic with:

    panic: proto: invalid UTF-8 string

    goroutine 1 [running]:
    main.main()
        /Users/morgan/mconbere/datastore-break/cmd/main.go:37 +0x2d5
