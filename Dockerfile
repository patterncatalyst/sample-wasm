FROM golang:1.14.3-alpine 
WORKDIR /app
COPY . . 
RUN apk add --no-cache git
RUN go get -u github.com/shurcooL/goexec
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
RUN GOOS=js GOARCH=wasm go build -o main.wasm
CMD ["goexec", "http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))"]

