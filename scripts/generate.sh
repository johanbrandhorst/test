#!/usr/bin/env bash

set -euo pipefail

cp assert.go must/assert.go
cp assert_test.go must/assert_test.go
cp settings.go must/settings.go
cp settings_test.go must/settings_test.go
cp scripts.go must/scripts.go
cp scripts_test.go must/scripts_test.go
cp test.go must/must.go
cp test_test.go must/must_test.go
sed -i "s|package test|package must|g" must/assert.go
sed -i "s|package test|package must|g" must/assert_test.go
sed -i "s|package test|package must|g" must/settings.go
sed -i "s|package test|package must|g" must/settings_test.go
sed -i "s|package test|package must|g" must/scripts.go
sed -i "s|package test|package must|g" must/scripts_test.go
sed -i "s|package test|package must|g" must/must.go
sed -i "s|package test|package must|g" must/must_test.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/assert.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/assert_test.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/settings.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/settings_test.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/scripts.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/scripts_test.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/must.go
sed -i -e "1s|^|// Code generated via scripts/generate.sh. DO NOT EDIT.\n\n|g" must/must_test.go
