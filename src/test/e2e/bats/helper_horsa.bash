#!/usr/bin/env bash

PROJECT_DIR="$( cd "$( dirname "${BATS_TEST_DIRNAME}/../../../../../" )" >/dev/null 2>&1 && pwd )"

export HORSA=${PROJECT_DIR}/target/horsa/horsa

function horsa_compiled_version_compare() {
    local OPERATOR=${1}
    local VERSION=${2}

    local COMPILED_VERSION; read COMPILED_VERSION < <(go version ${HORSA} | cut -d ' ' -f 2)

    if test "${COMPILED_VERSION}" = "${VERSION}" && test "${OPERATOR}" = "=" -o "${OPERATOR}" = ">=" -o "${OPERATOR}" = ">="
    then
        return 0
    fi


    if test "${OPERATOR}" = "=" -a "${COMPILED_VERSION}" != "${VERSION}"
    then
        return 1
    fi

    local GREATER_VERSION; read GREATER_VERSION < <(echo -e "${COMPILED_VERSION}\ngo${VERSION}" | sort --version-sort | tail -n 1)

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

    echo "unknown operator: ${OPERATOR}"
    exit 1

}