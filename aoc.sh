#!/usr/bin/env bash

set -o errexit
set -o pipefail

usage() {
    local name
    name="$(basename "${0}")"
    echo "Fetch inputs and send answers for Advent of Code.

Usage:
  ${name} (input|inputs)  [-d=<day>] [-y=<year>] [-c=<cookie>] [<path>]
  ${name} answer          [-d=<day>] [-y=<year>] [-c=<cookie>] <level> <answer>

Options:
  -d    Day    [default: \$AOC_DAY or current day]
  -y    Year   [default: \$AOC_YEAR or current year]
  -c    Cookie [default: \$AOC_COOKIE or cookie.txt]"

    exit 1
}

day="${AOC_DAY:-$(date +"%Oe" | sed 's/[[:blank:]]//')}"
year="${AOC_YEAR:-$(date +"%Y")}"
cookie="${AOC_COOKIE:-$(if test -r "cookie.txt"; then cat "cookie.txt"; fi)}"

while getopts d:y:c: opt; do
    case ${opt} in
    d) day="${OPTARG}" ;;
    y) year="${OPTARG}" ;;
    c) cookie="${OPTARG}" ;;
    *) break ;;
    esac
done
shift $((OPTIND - 1))

if test -z "${1}"; then usage; fi
mode="${1}"
shift

curl_flags=(
    "--fail"
    "--silent"
    "--show-error"
    "--location"
    "--header" "cookie: session=${cookie}"
    "--user-agent" "github.com/nothub/AoC2023"
)

case ${mode} in
input | inputs)
    path="${1:-"day${day}/input.txt"}"
    echo >&2 "Downloading input file for ${year}/${day} to: ${path}"
    if test -f "${path}"; then
        echo >&2 "File already present, skipping download..."
        exit 0
    fi
    curl "${curl_flags[@]}" \
        --create-dirs \
        --output "${path}" \
        "https://adventofcode.com/${year}/day/${day}/input"
    ;;
answer)
    if test ${#} -lt 2; then usage; fi
    level="${1}"
    answer="${2}"
    echo >&2 "Posting answer for ${year}/${day}/${level}: ${answer}"
    curl "${curl_flags[@]}" \
        --data "level=${level}&answer=${answer}" \
        "https://adventofcode.com/${year}/day/${day}/answer" |
        grep "<article>" |
        sed -E "s/<[^<>]*>|\[.*\]//g"
    ;;
*)
    usage
    ;;
esac
