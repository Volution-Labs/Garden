## garden-cli config

Configure on device settings.

### Synopsis

Configure device setting via network (Coap) or serial connections.
First device found will be used, or use 'garden-cli device' to select device.

```
garden-cli config [flags]
```

### Options

```
  -h, --help              help for config
  -i, --interval string   Set the rate sensor data is sent: [milliseconds]
  -s, --serverIP string   Set the serving ip address: [ipaddress, auto]
```

### SEE ALSO

- [garden-cli](README.md) - Cli tool to config and debug garden devices.
