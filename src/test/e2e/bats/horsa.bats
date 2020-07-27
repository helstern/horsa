#!/usr/bin/env bats

load helper_horsa
load helper_socket

@test "invoking foo with a nonexistent file prints an error" {
    local SYSTEMD; read SYSTEMD < <(which systemd-socket-activate || echo /lib/systemd/systemd-activate)

    local PORT=8000
    ${SYSTEMD} --listen ${PORT} ${HORSA} systemd ${HORSA_ECHO_SERVER} &
    local SERVER_PID=$!

    local OPENED_SOCKETS; read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 0 = ${OPENED_SOCKETS}

    local RESPONSE; read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 2 = ${OPENED_SOCKETS}
    test 'hello' = "${RESPONSE}"


    read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPENED_SOCKETS < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 'hello' = "${RESPONSE}"
    test 2 = ${OPENED_SOCKETS}

    kill ${SERVER_PID}
}

teardown() {
    pkill -f ${HORSA} || true
}