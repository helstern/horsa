#!/usr/bin/env bash

PROJECT_DIR="$( cd "$( dirname "${BATS_TEST_DIRNAME}/../../../../../" )" >/dev/null 2>&1 && pwd )"

export HORSA=${PROJECT_DIR}/target/horsa/horsa

# compares the default go version against a supplied one using a specified operator
# it returnins the supplied version if the comparison returns true or an empty string
# example: horsa_supported_go_version < 1.14
function horsa_supported_go_version() {
    local GO_VERSION; read GO_VERSION < <(go version | cut -d ' ' -f 3)

    local GREATER_GO_VERSION; read GREATER_GO_VERSION < <(echo -e "go1.13\n${GO_VERSION}" | sort --version-sort | tail -n 1)
    if test "${GO_VERSION}" = "${GREATER_GO_VERSION}"
    then
        echo ${GO_VERSION}
    else
        echo
    fi
}

function horsa_compiled_version_compare() {
    local OPERATOR=${1}
    local VERSION=go${2}

    local COMPILED_VERSION; read COMPILED_VERSION < <(go version ${HORSA} | cut -d ' ' -f 2)

    if test "${COMPILED_VERSION}" = "${VERSION}" && test "${OPERATOR}" = "=" -o "${OPERATOR}" = ">=" -o "${OPERATOR}" = ">="
    then
        return 0
    fi


    if test "${OPERATOR}" = "=" -a "${COMPILED_VERSION}" != "${VERSION}"
    then
        return 1
    fi

    local GREATER_VERSION; read GREATER_VERSION < <(echo -e "${COMPILED_VERSION}\n${VERSION}" | sort --version-sort | tail -n 1)

    case "${OPERATOR}" in
        '>')
            if test "${COMPILED_VERSION}" = "${GREATER_VERSION}"
            then
                return 0
            else
                return 1
            fi
        ;;
        '<')
            if test "${VERSION}" = "${GREATER_VERSION}"
            then
                return 0
            else
                return 1
            fi
        ;;
    esac

    exit 1
}