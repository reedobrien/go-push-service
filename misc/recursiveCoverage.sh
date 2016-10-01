#!/bin/sh 
echo 'mode: atomic' > coverage.txt && touch coverage.tmp &&  go list gitlab.com/fffd/wps/... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp && go tool cover -func=coverage.txt

