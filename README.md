# pb-neopixel

![photo](photo/pb-neopixel.jpg)

Control neopixel with protobuf between;

    go (tcp client) --protobuf_encoded_msg--> esp8266-neopixel (tcp server)


The message format is `neopixel.proto`.

In `esp8266`, [nanopb](https://jpa.kapsi.fi/nanopb/) is used to decode protobuf.
`esp8666` directory contains Arduino sketch for this.

## reference

Generate pb handlinging codes;

    $ protoc --go_out=. neopixel.proto
    $ protoc --nanopb_out=./esp8266/ neopixel.proto