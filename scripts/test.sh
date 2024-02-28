#!/bin/bash
# Colors
GREEN_COLOR=$'\e[92m'
RED_COLOR=$'\e[91m'
BLUE_COLOR=$'\e[94m'
END_COLOR=$'\e[0m'
YELLOW_COLOR=$'\e[93m'
# Variables
COVERAGE_REPORT=cover.out
COVERAGE_DESTINATION=./
ENV=test

# Run tests
printf "${BLUE_COLOR}Running unit tests...${END_COLOR}\n"
ginkgo run -r ./test/unit

if [ $? -eq 0 ]; then
  printf "${GREEN_COLOR}Unit tests passed!${END_COLOR}\n"
else
  printf "${RED_COLOR}Unit tests failed!${END_COLOR}\n"
  exit 1
fi

# Generate coverage report
printf '%s\n\n' "${BLUE_COLOR}>>> Running application tests coverage with ${YELLOW_COLOR}[${ENV}] ${BLUE_COLOR}environment...${RESET_COLOR}" &&
ENV=${ENV} go test -v -coverprofile ${COVERAGE_REPORT} -coverpkg ./... ./... 1>/dev/null &&
echo "${YELLOW_COLOR}>>> Cleaning cover ignored files ${RESET_COLOR}" && ./scripts/exclude_from_code_coverage.sh &&
go tool cover -html=${COVERAGE_REPORT} &&
printf '%s\n\n'

if [ $? -eq 0 ] 
then
  echo "${GREEN_COLOR}>>> Launching web coverage report ${RESET_COLOR}"
  gocov convert "$COVERAGE_REPORT" > "$COVERAGE_DESTINATION/coverage.json"
  gocov report "$COVERAGE_DESTINATION/coverage.json" 1>"$COVERAGE_DESTINATION/report.txt" 
  TOTAL_COVERAGE=$(grep "Total Coverage" "$COVERAGE_DESTINATION/report.txt")
  rm ${COVERAGE_REPORT} "$COVERAGE_DESTINATION/coverage.json" "$COVERAGE_DESTINATION/report.txt"
  echo "${GREEN_COLOR}>>> Tests ran successfully with ${TOTAL_COVERAGE} ${RESET_COLOR}"
else
  echo "${RED_COLOR}>>> Error occured running tests ${RESET_COLOR}"
fi