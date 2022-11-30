= Sample Golang WASM
:dockinfo: shared
:!toc:

Small sample application of Golang with WASM

Requirement:

Must have your GOPATH set even if using modules as wasm compilation will utilize it. e.g. $HOME/go
Must have your GOROOT set.  e.g. /usr/lib/go-<version number>

Compiling:

docker run -d -p 8081:8080 wasm

Testing:

Curl the result: curl http://localhost:8081

Or use a browser:

http://localhost:8081



