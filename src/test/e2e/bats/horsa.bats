#!/usr/bin/env bats

load helper_horsa
load helper_socket

@test "test os resources are released after each connection is closed" {

    # it appears that when compiled under go 1.14 after the initial connection, lsof reports two pipes
    local EXPECTED_OPENPIPES=2
    if horsa_compiled_version_compare '<' 1.14
    then
        EXPECTED_OPENPIPES=0
    fi

    local SYSTEMD; read SYSTEMD < <(which systemd-socket-activate || echo /lib/systemd/systemd-activate)

    local PORT=8000
    ${SYSTEMD} --listen ${PORT} ${HORSA} systemd ${HORSA_ECHO_SERVER} &
    local SERVER_PID=$!

    local OPEN_PIPES; read OPEN_PIPES < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 0 = ${OPEN_PIPES}

    local RESPONSE; read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPEN_PIPES < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test ${EXPECTED_OPENPIPES} = ${OPEN_PIPES}
    test 'hello' = "${RESPONSE}"

    read RESPONSE < <(echo "hello" | PORT=${PORT} socket_client)
    read OPEN_PIPES < <(lsof -c horsa +E | grep -E '^horsa.+FIFO' | wc -l)
    test 'hello' = "${RESPONSE}"
    test ${EXPECTED_OPENPIPES} = ${OPEN_PIPES}

    kill ${SERVER_PID}
}

teardown() {
    pkill -f ${HORSA} || true
}