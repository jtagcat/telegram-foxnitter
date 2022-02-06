#!/usr/bin/env bash
set -eou pipefail

file="go.mod"
if [ -n "$1" ]; then
	file="$1"
fi

if [[ ! -r "$file" ]]; then
	echo "error: file '${file}' not readable"
	exit 1
fi

matches=0
while read -r line; do
	if [[ "$line" =~ ^go[[:space:]].* ]]; then
		# begins with "go "
		v="$(xargs <<< "${line:3}")"

		if [[ "$matches" != 0 ]]; then
			>&2 echo "error: multiple versions found, 2nd is '${v}'"
			exit 1
		fi

		echo "${v}"
		((matches++))
	fi
done <go.mod

if [[ "$matches" == 0 ]]; then
	>&2 echo "error: no version found"
	exit 1
fi
