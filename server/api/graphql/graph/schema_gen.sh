#!/bin/bash
cd "$(dirname "$0")" || exit 1
go run github.com/99designs/gqlgen generate