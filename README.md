# healthchecks

Simple utils methods that helps to notify [healthchecks.io](https://healthchecks.io/) healthchecks system.

## Usage

```go
// initialize object
hc := healthchecks.Config{
Url: "https://hc-ping.com/<uuid>",
}

// notify start
err := hc.Notify(healthchecks.START)
if err != nil {
panic(err)
}

// notify succesful end
err := hc.Notify(healthchecks.OK)
if err != nil {
panic(err)
}

// notify fail
err := hc.NotifyWithLog(healthchecks.FAIL, "my-message")
if err != nil {
panic(err)
}

```