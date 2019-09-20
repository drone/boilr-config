A config extension to {{Description}}. _Please note this project requires Drone server version 1.4 or higher._

## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --restart=always \
  --name=config {{DockerRepository}}
```

Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6
```