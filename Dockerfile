FROM golang:1.19.2-alpine
WORKDIR /app
COPY . . 
RUN apk add --no-cache git
RUN go install github.com/shurcooL/goexec@latest
RUN go get github.com/shurcooL/go-goon
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
RUN GOOS=js GOARCH=wasm go build -o main.wasm
CMD ["goexec", "http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))"]

