#!/usr/bin/env bash

export HORSA_ECHO_SERVER=${PROJECT_DIR}/src/test/e2e/bats/mock_server.sh

socket_client() {
    local INPUT=$(cat -)
    PORT=${PORT:-8080}

    echo -e "${INPUT}\n" | { pid=$(exec sh -c 'echo "$PPID"'); echo ${pid}; socket localhost ${PORT}; } | {
        read PID
        while read LINE
        do
            echo ${LINE}
        done
    }
}