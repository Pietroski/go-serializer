#!/usr/bin/env bash

cat <<EOF >CHANGELOG.md
# CHANGELOG

# [UNRELEASED]
EOF

git --no-pager log --pretty=format:"%C(auto)%d - %s" | sed 's|(tag: \(.*\))|\n# [\1]\n|' | sed 's|(\(.*\))|- [\1]|' | sed '/chore: changelog/d' | sed '/chore: version bump/d' >>CHANGELOG.md

cat <<EOF >>CHANGELOG.md

EOF
