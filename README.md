## ccli-tz

A cardano-cli wrapper with send funds function and time zone adjusted leadership-schedule. Settings are controlled via config to simplify usage.  

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

### Basic Usage

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

### Send funds

Allows you to send funds to multiple wallets at once defined in a simple yaml file. Example file:

```yaml
sourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
targetAddresses:
  addr1q8a5gtz7qv8cccy5tymwwansn9m5zwm9kjkt55eqyukrm4fk9mk69u550yut4hhf5cyqu5nmh8jpw57lhxvhwqgx5sxqcydlqg:
    lovelaceAmount: 1150770
  addr1qyl66psd5nrwpd85ddne2x5reg006sqyzfallkeeuenydkh0ays4l5jylz7v4cwvgrwnvqcthn4tjk4g6lcuw567js6sphzc2m:
    lovelaceAmount: 1150770
    paymentTokens:
      - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.484f4e4559
        tokenAmount: 1
```

This builds, signs and submits the transaction in one step. It requires the path to the signing key file to sign the transaction - this is supplied from an environment variable like so:
```shell
$ SIGNING_KEY_FILE=/path/to/source/payment.skey ccli-tz sendfunds --payment-file ~/some/path/to/payment.yml
```

You could also export the environment variable separately:

```shell
$ export SIGNING_KEY_FILE=/path/to/source/payment.skey 
$ ccli-tz sendfunds --payment-file ~/some/path/to/payment.yml --testnet-magic 1
```
