# check_ripe_atlas_probe

A plugin to monitor the connection state of your RIPE Atlas probe. This plugin provides no performance data. It uses the public RIPE Atlas API.

**Please don't put too much stress on the API. A call every 30 minute is enough!**

## Installation

You need go (golang) to compile or run. After installing go you can run:

```
go run check_ripe_atlas_probe.go --probe 34088
```

Or compile it:

```
go build check_ripe_atlas_probe.go
```

You should then have a binary called `check_ripe_atlas_probe`. If you want to cross-compile it append a `GOOS=linux GOARCH=amd64 ` before the compile command to compile it for a 64bit linux. Usefull/Needed if your target arch or OS is not the one you are developing on.

## Possible outputs

Output if your probe is connected:
```
OK - Probe 34088 ("my fancy probe") is connected since 442.3 hours
```


Output if your probe is not connected:

```
OK - Probe 34088 ("my fancy probe") is disconnected since 2 hours
```

Output if your probe was not found:

```
UNKNOWN - Status was 404 Not Found. Expected 200 OK
```

## Integration into Icinga 2

Deploy the script into your PluginDir. Define the check command according to your Icinga 2 installation:

```
object CheckCommand "ripe_atlas_probe" {
    import "plugin-check-command"

    command = [ PluginDir + "/check_ripe_atlas_probe" ]

    arguments = {
        "--probe" = {
            value = "$ripe_atlas_probeid$"
            description = "The RIPE Atlas probe ID"
            required = true
        }
    }
}
```

Create a Service and restart Icinga 2.
