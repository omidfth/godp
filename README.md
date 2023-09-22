# Godp

**Godp** is a UDP library for Golang that tries to make sending and receiving data easy. Using this library, you can quickly design a app under the UDP protocol. This project is used in online games that require high speed data transmission.

## Getting started
### Prerequisites
- [Go]("https://go.dev/"): any one of the three latest **major** releases.

## Getting Godp
With [Go module]("https://github.com/golang/go/wiki/Modules") support, simply add the following import

```
import "github.com/omidfth/godp"
```
to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `godp` package:
```sh
$ go get -u github.com/omidfth/godp
```

## How to Use?
### Quick Start
First you need to import Godp package for using Godp, one simplest example likes the follow:

**Server:**
```go
import (
    "encoding/json"
    "github.com/omidfth/godp"
)

func main() {
    r := godp.NewRouter()
    r.NewRoute(godp.PING, OnPing)
    r.ListenAndServe(1055, godp.MaxBufferSize)
}

func OnPing(c *godp.Context) {
    c.GetPingService().Ping(c.Packet.SocketID)
    packet := godp.MakePacket(c.Packet.SocketID, c.Packet.RoomID, godp.PING, nil)
    j, _ := json.Marshal(packet)
    c.Emit(c.Address, j)
}
```

**Test:**
```go
import (
    "log"
    "net"
    "os"
)

const (
    HOST = "localhost"
    PORT = "1055"
    TYPE = "udp"
)

func main() {
    udpServer, err := net.ResolveUDPAddr(TYPE, HOST+":"+PORT)
    log.Println("Attempt Connect to:", HOST+":"+PORT, TYPE)

    if err != nil {
        println("ResolveUDPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialUDP("udp", nil, udpServer)
        if err != nil {
            println("Listen failed:", err.Error())
        os.Exit(1)
    }

    //close the connection
    defer conn.Close()
        _, err = conn.Write([]byte("{\"s\":null,\"e\":0,\"d\":{\"u\":\"\",\"s\":\"\"},\"r\":null}"))
        if err != nil {
            println("Write data failed:", err.Error())
        os.Exit(1)
    }

    // buffer to get data
    received := make([]byte, 1024)
        _, err = conn.Read(received)
        if err != nil {
            println("Read data failed:", err.Error())
        os.Exit(1)
    }

    println(string(received))
}
```

