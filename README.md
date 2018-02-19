# telegraf-ethermine
A telegraf plugin for tracking ethermine.org miner stats.

# Telegraf
The plugin reports an `ethermine` measurement.

## Tags
- address: The public ether wallet address

## Metrics
- reported_hashrate
- current_hashrate
- average_hashrate
- valid_shares
- invalid_shares
- stale_shares
- active_workers

# Installation
1. Drop `telegraf-ethermine.exe` in a directory.  For example, you can place it in the same directory that contains `telegraf.exe`.
2. Add an `[[inputs.exec]]` section to your `telegraf.conf` file.  Make sure to add your public ether wallet address, without the leading 0x.

```ini
[[inputs.exec]]
  commands = [
    "c:/progra~1/telegraf/telegraf-ethermine.exe -address YOUREPUBLICTHERWALLETADDRESS"
  ]
  timeout = "5s"
  data_format = "influx"
```

3. Restart the Telegraf service.

# Development
`go build`
