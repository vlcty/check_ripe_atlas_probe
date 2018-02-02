# check_ripe_atlas_probe

A perl plugin to monitor the connection state of your RIPE Atlas probe. This plugin provides no performance data. It contacts the public RIPE Atlas API using the anonymous account.

**Please don't put too much stress ton the API. A call every 15 minute for every node is enough!**

## Usage from the CLI

Just run the script and set the `--probe` parameter with your probe ID. Example for my probe with the ID 34088:
```
:~$ ./check_ripe_atlas_probe --probe 34088
OK - Probe 34088 is connected since 2018-01-30T18:44:47Z
```

## Possible outputs
Output if your probe is connected (OK state):
```
OK - Probe 34088 is connected since 2018-01-30T18:44:47Z
```


Output if your probe is not connected (CRITICAL state):
```
CRITICAL - Probe 34088 is not connected
```


Output if your probe was not found (UNKNOWN state):
```
CRITICAL - Was not able to contact API: 404 Not Found
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
