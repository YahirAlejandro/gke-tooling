# Before building the Go binary, remember to detach its OS bindings with:
# $ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM scratch
ADD main /
CMD ["/main"]
