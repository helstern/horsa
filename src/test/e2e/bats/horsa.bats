#!/usr/bin/env bash

load helper_horsa
load helper_socket

@test "invoking foo with a nonexistent file prints an error" {
    local PORT=8000
    systemd-socket-activate --listen ${PORT} ${HORSA} systemd ${HORSA_ECHO_SERVER} &
    local SERVER_PID=$!

    local OPENED_SOCKETS; read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    echo "2 == ${OPENED_SOCKETS}"
    test 0 == ${OPENED_SOCKETS}

    local RESPONSE; read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 'hello' -eq "${RESPONSE}"
    test 2 == ${OPENED_SOCKETS}

    read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 'hello' -eq "${RESPONSE}"
    test 2 == ${OPENED_SOCKETS}

#    kill ${SERVER_PID}
}

teardown() {
    pkill -f horsa
}