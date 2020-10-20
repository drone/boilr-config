A config extension to {{Description}}. _Please note this project requires Drone server version 1.4 or higher._

## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the extension:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --restart=always \
  --name=config {{DockerRepository}}
```

Update your Drone server configuration to include the extension address and the shared secret.

```text
DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Testing

Use the command line tools to test your extension. _This extension uses http-signatures to authorize client access and will reject unverified requests. You will be unable to test this extension using a simple curl command._

```text
export DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
export DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6

drone plugins config get <repo>
```
