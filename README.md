# horsa
> socket server proxy

## description

horsa is a socket server proxy inspired by and very similar with the `socket` utility. it was originally built to help using the systemd's socket activation. the main difference between `horsa` and `socket` is that `horsa` can also accept a file descriptor to a unix socket rather than the full path to the socket

- [Installation](#installation)

## Installation

using the following:

- `wget --version`: GNU Wget 1.20.3 built on linux-gnu.
- `tar --version`: tar (GNU tar) 1.30
 
run the following commands: 

```
    export HORSA_RELEASE=$(wget -qO - https://api.github.com/repos/helstern/horsa/releases/latest | grep tag_name | grep --only-matching -E 'v[^"]+')
    wget -qO - https://github.com/helstern/horsa/releases/download/${HORSA_RELEASE}/horsa-${HORSA_RELEASE}-linux-amd64.tar.gz | sudo tar xvz -C / 
``` 

which will install `horsa` at: `/usr/local/bin/horsa`

verify the installation by running: `horsa help`