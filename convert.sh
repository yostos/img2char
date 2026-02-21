#!/bin/bash
# Sample script: Resize and binarize an image using ImageMagick, then run img2char.
# Requires: ImageMagick (magick command) and img2char binary in the same directory.
set -euo pipefail

usage() {
  echo "Usage: $0 [-w 640|320] [-t threshold] [-n] <input-image>" >&2
  echo "  -w  output width (640 or 320, default: 640)" >&2
  echo "  -t  binarize threshold in % (0-100, default: 50)" >&2
  echo "  -n  negate (use when dark areas are the subject)" >&2
  exit 1
}

WIDTH=640
THRESHOLD=50
NEGATE=""

while getopts "w:t:nh" opt; do
  case $opt in
    w) WIDTH=$OPTARG ;;
    t) THRESHOLD=$OPTARG ;;
    n) NEGATE="-negate" ;;
    *) usage ;;
  esac
done
shift $((OPTIND - 1))

[ $# -ne 1 ] && usage

INPUT="$1"
[ ! -f "$INPUT" ] && echo "error: file not found: $INPUT" >&2 && exit 1

if [ "$WIDTH" != "640" ] && [ "$WIDTH" != "320" ]; then
  echo "error: width must be 640 or 320" >&2
  exit 1
fi

TMPFILE=$(mktemp /tmp/img2char_XXXXXX.png)
trap 'rm -f "$TMPFILE"' EXIT

# Resize to WIDTHx200 and binarize
magick "$INPUT" \
  -resize "${WIDTH}x200!" \
  -colorspace Gray \
  $NEGATE \
  -threshold "${THRESHOLD}%" \
  -depth 1 \
  "$TMPFILE"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
"${SCRIPT_DIR}/img2char" "$TMPFILE"
