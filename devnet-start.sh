#!/bin/bash -e

function cleanup {
    echo ""
    echo "[devnet] stopping devnet"
    make devnet-down
    echo "[devnet] devnet stoped"
    echo ""
}

trap cleanup EXIT

echo "[devnet] starting devnet"
make devnet-up-deploy
echo "[devnet] devnet started"

make devnet-logs
