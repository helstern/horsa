# horsa
> socket server proxy

## description

horsa is a socket server proxy inspired by and very similar with the `socket` utility. it was originally built to help using the systemd's socket activation. the main difference between `horsa` and `socket` is that `horsa` can also accept a file descriptor to a unix socket rather than the full path to the socket 