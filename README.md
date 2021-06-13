# boilr-config

This is a [boilr template](https://github.com/tmrts/boilr) for creating a configuration extension. Use a configuration extension to customize how Drone fetches the configuration file. Get started by cloning the project and installing the template:

```console
$ cd boilr-config
$ boilr template save . drone-config -f
```

create a project in directory my-config:

```console
$ boilr template use drone-config my-config
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar"]:
```
