#!/bin/bash

TEST_CMD="go test ./..."

$TEST_CMD
reflex -d none -r '\.go$' -- $TEST_CMD
