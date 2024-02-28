#!/bin/bash
mockery --all --recursive --dir=./internal/domain

MOCKS_DESTINATION=./test/utils/mocks/

mv ./mocks/* ${MOCKS_DESTINATION}
rm -rf ./mocks
