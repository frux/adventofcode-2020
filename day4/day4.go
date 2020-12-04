package day4

import (
	"adventofcode-2020/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const INPUT = "./day4/input.txt"

type tPassportData = map[string]string
type tIsValidFn = func(passport tPassportData) bool

func Part1() string {
	entries := input.ReadStrings(INPUT)

	return fmt.Sprint(getValidEntriesCount(entries, hasRequiredFields))
}

func Part2() string {
	entries := input.ReadStrings(INPUT)

	return fmt.Sprint(getValidEntriesCount(entries, isValidPassport))
}

func getValidEntriesCount(entries []string, isValidFn tIsValidFn) int {
	var count int
	var passportEntries []string

	for i, entry := range entries {
		if entry != "" {
			passportEntries = append(passportEntries, entry)
		}

		if entry == "" || i == len(entries)-1 {
			passport := parsePassportData(strings.Join(passportEntries, " "))
			passportEntries = []string{}

			if isValidFn(passport) {
				count++
			}
		}
	}

	return count
}

func parsePassportData(data string) tPassportData {
	result := make(tPassportData)
	pairs := strings.Split(data, " ")

	for _, pair := range pairs {
		keyVal := strings.Split(pair, ":")

		result[keyVal[0]] = keyVal[1]
	}

	return result
}

func hasRequiredFields(passport tPassportData) bool {
	return hasField(passport, "byr") &&
		hasField(passport, "iyr") &&
		hasField(passport, "eyr") &&
		hasField(passport, "hgt") &&
		hasField(passport, "hcl") &&
		hasField(passport, "ecl") &&
		hasField(passport, "pid")
}

func isValidPassport(passport tPassportData) bool {
	return isValidByr(passport["byr"]) &&
		isValidIyr(passport["iyr"]) &&
		isValidEyr(passport["eyr"]) &&
		isValidHgt(passport["hgt"]) &&
		isValidHcl(passport["hcl"]) &&
		isValidEcl(passport["ecl"]) &&
		isValidPid(passport["pid"])
}

func isValidByr(val string) bool {
	return checkNumber(val, 1920, 2002)
}

func isValidIyr(val string) bool {
	return checkNumber(val, 2010, 2020)
}

func isValidEyr(val string) bool {
	return checkNumber(val, 2020, 2030)
}

func isValidHgt(val string) bool {
	if len(val) < 3 {
		return false
	}

	if strings.HasSuffix(val, "cm") {
		return checkNumber(val[:len(val)-2], 150, 193)
	}

	if strings.HasSuffix(val, "in") {
		return checkNumber(val[:len(val)-2], 59, 76)
	}

	return false
}

func isValidHcl(val string) bool {
	if len(val) != 7 {
		return false
	}

	_, err := strconv.ParseUint(val[1:], 16, 32)

	return err == nil
}

func isValidEcl(val string) bool {
	return val == "amb" ||
		val == "blu" ||
		val == "brn" ||
		val == "gry" ||
		val == "grn" ||
		val == "hzl" ||
		val == "oth"
}

func isValidPid(val string) bool {
	matched, err := regexp.MatchString(`^\d{9}$`, val)

	return matched && err == nil
}

func checkNumber(value string, min int, max int) bool {
	num, err := strconv.Atoi(value)

	return err == nil && num >= min && num <= max
}

func hasField(passport tPassportData, fieldName string) bool {
	_, ok := passport[fieldName]

	return ok
}
