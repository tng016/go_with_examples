#!/usr/bin/env bash

protoc -I=${PWD} --go_out=${PWD} ${PWD}/integers.proto