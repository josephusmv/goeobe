TestSingleThreadClientSession Smoke test for Single Thread
    go test -v -cover -coverprofile cover.out -run SingleThreadClientSession
    go tool cover -html=cover.out -o cover.html
    Current coverage: 84.4%
Cover session single threading white box tests.
Session package should NOT be directly referenced by multithreading environment, 
Use cltmgmt package to do multi-threading test, which has encapsulated all requests into a serialized worker.