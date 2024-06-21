# Opertizen - Operate Samsung model TV with Tizen OS through SmartThings API 

Opertizen is a lazy tool that lazily automates the TV viewing process.
It does this by directly interfacing with the SmartThings API through a series of commands defined in a runflow file.
A runflow is a simple workflow that executes a series of steps, calling the SmartThings API for each command.

# Runflow file

```yaml
runflow:
  name: Name of runflow
  description: Description of runflow
  steps:
    - name: Name of step
      command: <capability>;<command(arguments)>
    - name: Example step. Send remote control DOWN command
      command: samsungvd.remoteControl;send(UP, PRESS_AND_RELEASED)
    - name: Example step. Turn off TV
      command: switch;off
```

# Command line usage

Running from source:

The TV configuration file `opertizen/configs/tvconf.yaml` must contain the device ID and SmartThings access token to call the SmartThings API

You can run the tool using the command line script `opertizen/scripts/build_and_run_with_args.sh` and specifying a custom runflow file:
```
$ opertizen/scripts/build_and_run_with_args.sh -runflow=<runflow_file>
```
`runflow` is the path to the custom runflow file

# Supported SmartThings Capabilities and Commands

| Capability              | Command                          | 
|-------------------------|----------------------------------|
| switch                  | on/off                           |
| audioVolume             | volumnDown/volumeUp/setVolume    |
| audioMute               | mute/unmute                      |
| samsungvd.remoteControl | send                             |

# References

- Command-line interface for the SmartThings API: https://github.com/SmartThingsCommunity/smartthings-cli
