# BenchHub cleanup

Xephon-B was merged into Xephon-K, and now we [split it out for BenchHub](https://github.com/xephonhq/xephon-b/issues/31)

## History

- it was using viper for config, we no longer use it, just unmarshal YAML to config struct using gommon/config
- it generates time and value using different generator
- it had clients, but most are just pinging, and have a payload interface that return bytes
- there are docker scripts for setting up different databases
  - `wait-for-it.sh` is used for docker-compose