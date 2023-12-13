# healthchecks

Simple utils methods that helps to notify [healthchecks.io](https://healthchecks.io/) healthchecks system.

## Usage

```golang
// initialize object
hc := healthchecks.Config{
    Url: "https://hc-ping.com/<uuid>",
}

// notify start
if err := hc.Notify(healthchecks.START); err != nil {
    log.Println(err)
}

// notify succesful end
if err := hc.Notify(healthchecks.OK); err != nil {
    log.Println(err)
}

// notify fail
if err := hc.NotifyWithLog(healthchecks.FAIL, "my-message"); err != nil {
    log.Println(err)
}

```
