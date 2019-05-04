#include <Wire.h>
#include <Adafruit_Sensor.h>
#include <Adafruit_TSL2561_U.h>
#include <WiFi.h>
#include <WiFiClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>

const char *ssid = "********";
const char *password = "*******";
WebServer server(80);

const int valve = 25;
bool valveState = false;

void setup()
{
  pinMode(valve, OUTPUT);
  digitalWrite(valve, LOW);
  Serial.begin(9600);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  if (MDNS.begin("garden"))
  {
    Serial.println("MDNS responder started");
  }

  server.on("/", handleRoot);
  server.on("/water/{}", handleWater);
  server.on("/sensor", handleSensorData);
  server.onNotFound(handleNotFound);

  server.begin();

  setupLight();
}

void loop()
{
  server.handleClient();
  //printLight();

  if (valveState)
  {
    digitalWrite(valve, HIGH);
  }
  else if (!valveState)
  {
    digitalWrite(valve, LOW);
  }
}
