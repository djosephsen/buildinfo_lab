#!/bin/sh

set -x
go run -ldflags="-X 'main.version=1.2.3'" .
