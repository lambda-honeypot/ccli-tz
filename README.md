## ccli-tz

A cardano-cli leadership-schedule wrapper with time zone adjust. Settings are controlled via config to simplify usage.  

### Installation

TBD

### Config File Sample

```yaml
VRFSigningKeyFile: /path/to/key/vrf.skey
stakePoolID: 217e45e759ef5d132dd47d4b8535327d897134ee6803f6d1383a0b50
shelleyGenesisFile: /path/to/configs/shelley-genesis.json
timeZone: Europe/London
```

### Usage

Simple (defaults to mainnet):
```shell
$ ./ccli-tz next
```

Testnet:
```shell
$ ./ccli-tz current --testnet-magic 1
```

Override config:
```shell
$ ./ccli-tz current --testnet-magic 1 --config ~/other_path/.ccli-tz.yaml
```
