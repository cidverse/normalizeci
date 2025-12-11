#!/usr/bin/env bash
set -euo pipefail

# parameters
CID_VERSION="${1:-latest}"
CID_SHA256="${2:-}"

# variables
BIN_DIR="/usr/local/bin"
DOWNLOAD_URL="https://github.com/cidverse/cid/releases/download/v${CID_VERSION}/linux_amd64"
[[ "$CID_VERSION" == "latest" ]] && DOWNLOAD_URL="https://github.com/cidverse/cid/releases/latest/download/linux_amd64"

# github actions
if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
    echo "Detected GitHub Actions ..."

    TOOL_DIR="${RUNNER_TOOL_CACHE}/cid-${CID_VERSION}"
    BIN_DIR="${TOOL_DIR}/bin"
fi

# prepare bin dir
mkdir -p "$BIN_DIR"

# download
echo "Downloading CID $CID_VERSION ..."
curl -sSL -o "${BIN_DIR}/cid" "$DOWNLOAD_URL"

# optional hash verification
if [[ -n "${CID_SHA256}" ]]; then
    echo "Verifying SHA256..."
    echo "${CID_SHA256}  ${BIN_DIR}/cid" | sha256sum -c -
fi

# make executable, after verification
chmod +x "${BIN_DIR}/cid"

# github actions
if [[ "${GITHUB_ACTIONS:-}" == "true" ]]; then
    echo "${BIN_DIR}" >> "$GITHUB_PATH"
    echo "CID_VERSION=${CID_VERSION}" >> "$GITHUB_ENV"
fi

# export to path
export PATH="${BIN_DIR}:${PATH}"
echo "CID version: ${CID_VERSION}"
