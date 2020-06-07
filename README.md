# pb-neopixel

Control neopixel with protobuf between;

    go (tcp client) --protobuf_encoded_msg--> esp8266-neopixel


The message format is `neopixel.proto`.

In `esp8266`, [nanopb](https://jpa.kapsi.fi/nanopb/) is used to decode protobuf.
`esp8666` directory contains Arduino sketch for this.