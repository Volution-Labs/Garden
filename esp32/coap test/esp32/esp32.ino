#include <WiFi.h>
#include <WiFiUdp.h>
#include <coap-simple.h>

const char *ssid = "********";
const char *password = "********";

// CoAP client response callback
void callback_response(CoapPacket &packet, IPAddress ip, int port);

// CoAP server endpoint url callback
void callback_light(CoapPacket &packet, IPAddress ip, int port);

// UDP and CoAP class
WiFiUDP udp;
Coap coap(udp);

// LED STATE
bool LEDSTATE;

// CoAP server endpoint URL
void callback_light(CoapPacket &packet, IPAddress ip, int port)
{
  Serial.println("[Light] ON/OFF");
  Serial.println(ip);
  // send response
  char p[packet.payloadlen + 1];
  memcpy(p, packet.payload, packet.payloadlen);
  p[packet.payloadlen] = NULL;

  String message(p);
  Serial.println(message);

  if (message.equals("0"))
    LEDSTATE = false;
  else if (message.equals("1"))
    LEDSTATE = true;

  Serial.println(LEDSTATE);

  if (LEDSTATE)
  {
    //digitalWrite(5, HIGH) ;
    Serial.println("Fire");
    coap.sendResponse(ip, port, packet.messageid, "Help Me!!!!");
  }
  else
  {
    //digitalWrite(5, LOW) ;
    Serial.println("Fire");
    coap.sendResponse(ip, port, packet.messageid, "Help Me!!!!!");
  }
}

// CoAP client response callback
void callback_response(CoapPacket &packet, IPAddress ip, int port)
{
  Serial.println("Sending Response");
  Serial.println("[Coap Response got]");

  char p[packet.payloadlen + 1];
  memcpy(p, packet.payload, packet.payloadlen);
  p[packet.payloadlen] = NULL;

  Serial.println(p);
}

void setup()
{
  Serial.begin(9600);

  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());

  // LED State
  pinMode(5, OUTPUT);
  digitalWrite(5, LOW);
  LEDSTATE = true;

  // add server url endpoints.
  // can add multiple endpoint urls.
  // exp) coap.server(callback_switch, "switch");
  //      coap.server(callback_env, "env/temp");
  //      coap.server(callback_env, "env/humidity");
  Serial.println("Setup Callback Light");
  coap.server(callback_light, "light");

  // client response callback.
  // this endpoint is single callback.
  Serial.println("Setup Response Callback");
  coap.response(callback_response);

  // start coap server/client
  coap.start();
}

void loop()
{
  delay(1000);
  coap.loop();
}
/*
if you change LED, req/res test with coap-client(libcoap), run following.
coap-client -m get coap://(arduino ip addr)/light
coap-client -e "1" -m put coap://(arduino ip addr)/light
coap-client -e "0" -m put coap://(arduino ip addr)/light
*/
