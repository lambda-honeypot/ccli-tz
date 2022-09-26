## ccli-tz

A cardano-cli leadership-schedule wrapper with time zone adjust. Settings are controlled via config to simplify usage.  

### Installation

You can visit the [github releases](https://github.com/lambda-honeypot/ccli-tz/releases) page for the project and install manually or use the below snippet:

```shell
mkdir -p ~/ccli-tz && cd ~/ccli-tz && \
VERSION="$(curl -s https://api.github.com/repos/lambda-honeypot/ccli-tz/releases/latest | jq -r .tag_name)" && \
OS="$(uname)" && \
ARCH="$(uname -m | sed -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" && \
curl -fsSLO "https://github.com/lambda-honeypot/ccli-tz/releases/download/${VERSION}/ccli-tz_${VERSION:1}_${OS}_${ARCH}.tar.gz" && \
tar zxvf ccli-tz_${VERSION:1}_${OS}_${ARCH}.tar.gz
```

This will install the latest version of `ccli-tz` to `~/ccli-tz`. You may wish to add this to your path by adding the below line to your `.bashrc` or similar:

```shell
export PATH=~/ccli-tz:$PATH
```

You will need a config file for `ccli-tz` to work. To create a sample config you can run: 

```shell
ccli-tz init
```

This will create a sample file at `~/.ccli-tz.yaml` (see below for an example).

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
