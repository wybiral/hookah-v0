# hookah
WebSocket pipeline tool.

### Install dependencies
```
go get github.com/gorilla/websocket
go get github.com/wybiral/hookah
```
### Build
```
go build github.com/wybiral/hookah/cmd/hookah
```
### Usage
To see usage instructions run:
```
./hookah
```
To start a relay node:
```
./hookah node host:port
```
Pipe a local program into a remote node:
```
local_program | ./hookah send remote_host:port
```
Read from a remote node:
```
./hookah recv remote_host:port
```
