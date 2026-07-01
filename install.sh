#!/bin/sh
# Ponytail-for-Python installer.
# Downloads the latest pnt binary from GitHub Releases, then runs `pnt install`
# in the current directory. Re-run any time to upgrade.
#
#   curl -sSL https://raw.githubusercontent.com/Danush-Aries/ponytail-for-python/main/install.sh | sh
#
set -eu

REPO="Danush-Aries/ponytail-for-python"
INSTALL_DIR="${PNT_INSTALL_DIR:-$HOME/.local/bin}"

log() { printf '\033[1;34m==>\033[0m %s\n' "$*"; }
err() { printf '\033[1;31merror:\033[0m %s\n' "$*" >&2; exit 1; }

os_name() {
  case "$(uname -s)" in
    Darwin) echo darwin ;;
    Linux)  echo linux ;;
    *) err "unsupported os: $(uname -s). Try building from source: https://github.com/$REPO" ;;
  esac
}

arch_name() {
  case "$(uname -m)" in
    x86_64|amd64) echo x86_64 ;;
    arm64|aarch64) echo arm64 ;;
    *) err "unsupported arch: $(uname -m)" ;;
  esac
}

latest_version() {
  # Follow the Releases /latest redirect to get the tag.
  curl -fsSL -o /dev/null -w '%{url_effective}' \
    "https://github.com/$REPO/releases/latest" \
  | sed -E 's|.*/tag/(v[0-9A-Za-z._-]+)/?$|\1|'
}

main() {
  OS="$(os_name)"
  ARCH="$(arch_name)"
  VER="${PNT_VERSION:-$(latest_version)}"
  [ -n "$VER" ] || err "could not resolve latest version"

  TARBALL="pnt_${VER#v}_${OS}_${ARCH}.tar.gz"
  URL="https://github.com/$REPO/releases/download/$VER/$TARBALL"

  TMP="$(mktemp -d)"
  trap 'rm -rf "$TMP"' EXIT

  log "downloading $TARBALL"
  curl -fsSL "$URL" -o "$TMP/$TARBALL" || err "download failed: $URL"

  log "extracting"
  tar -xzf "$TMP/$TARBALL" -C "$TMP"

  mkdir -p "$INSTALL_DIR"
  install -m 0755 "$TMP/pnt" "$INSTALL_DIR/pnt"
  log "installed pnt $VER to $INSTALL_DIR/pnt"

  case ":$PATH:" in
    *":$INSTALL_DIR:"*) : ;;
    *)
      printf '\n\033[1;33mnote:\033[0m %s is not on your PATH.\n' "$INSTALL_DIR"
      printf 'add this to your shell rc:\n  export PATH="%s:$PATH"\n\n' "$INSTALL_DIR"
      ;;
  esac

  log "running: pnt install"
  "$INSTALL_DIR/pnt" install
}

main "$@"
