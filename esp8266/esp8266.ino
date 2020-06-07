
/*
   Copyright (c) 2018, circuits4you.com
   All rights reserved.
   Create a TCP Server on ESP8266 NodeMCU.
   TCP Socket Server Send Receive Demo
*/

#include <ESP8266WiFi.h>

#ifndef STASSID
#define STASSID "your_ssid"
#define STAPSK "your_psk"
#endif

const char *ssid = STASSID;
const char *password = STAPSK;

const int port = 8888; //Port number
WiFiServer server(port);

//=======================================================================

#include "pb_encode.h"
#include "pb_decode.h"
#include "neopixel.pb.h"

#define BUFFSIZE 2048
uint8_t buffer[BUFFSIZE];

#define NUMLEDS 144
LED leds[NUMLEDS];
int ledCnt;

//=======================================================================

#include <Adafruit_NeoPixel.h>
#ifdef __AVR__
#include <avr/power.h> // Required for 16 MHz Adafruit Trinket
#endif

#define LED_PIN 13

// Declare our NeoPixel strip object:
Adafruit_NeoPixel strip(NUMLEDS, LED_PIN, NEO_GRB + NEO_KHZ800);

//=======================================================================
//                    Power on setup
//=======================================================================
void setup()
{
  strip.begin();           // INITIALIZE NeoPixel strip object (REQUIRED)
  strip.show();            // Turn OFF all pixels ASAP
  strip.setBrightness(150); // Set BRIGHTNESS to about 1/5 (max = 255)

  pinMode(LED_BUILTIN, OUTPUT);

  Serial.begin(115200);
  // pinMode(SendKey, INPUT_PULLUP); //Btn to send data
  Serial.println();

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password); //Connect to wifi

  // Wait for connection
  Serial.println("Connecting to Wifi");
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
    delay(500);
  }
  digitalWrite(LED_BUILTIN, HIGH);

  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);

  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
  server.begin();
  Serial.print("Open Telnet and connect to IP:");
  Serial.print(WiFi.localIP());
  Serial.print(" on port ");
  Serial.println(port);
}
//=======================================================================
//                    Loop
//=======================================================================

void loop()
{
  WiFiClient client = server.available();

  if (client)
  {
    if (client.connected())
    {
      Serial.println("Client Connected");
    }

    while (client.connected())
    {
      delay(100); // TODO: right?
      while (client.available() > 0)
      {
        // read data from the connected client
        int n;
        n = client.read(buffer, BUFFSIZE);
        // Serial.println(n);

        ledCnt = 0;
        // parse pb
        NeoPixel neopixel = NeoPixel_init_default;
        neopixel.strip.funcs.decode = read_led;
        pb_istream_t stream = pb_istream_from_buffer(buffer, n);
        if (!pb_decode(&stream, NeoPixel_fields, &neopixel))
        {
          Serial.printf("Decoding failed: %s\n", PB_GET_ERROR(&stream));
          break;
        }
        // Serial.printf("clear %d!\n", (int)neopixel.clear);
        update_neopixel(neopixel.clear);
      }
    }
    client.stop();
    Serial.println("Client disconnected");
  }
}

//=======================================================================

bool read_led(pb_istream_t *stream, const pb_field_iter_t *field, void **arg)
{
  while (stream->bytes_left)
  {
    if (!pb_decode(stream, LED_fields, &leds[ledCnt]))
      return false;

    ledCnt++;
  }
  return true;
}

void update_neopixel(boolean clear) {
  // Serial.printf("update_neopixel; clear: %d, cnt: %d\n", (int)clear, ledCnt);
  if (clear) {
    strip.clear();
  }

  for (int i = 0; i < ledCnt; i++) {
    int32_t idx  = leds[i].index;
    int32_t c = leds[i].color;
    // Serial.printf("led idx: %d, wrgb: 0x%08x\n", idx, c);
    strip.setPixelColor(idx, c);
  }

  strip.show();
}
