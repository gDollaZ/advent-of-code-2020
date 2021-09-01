package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TODO: this is showing 1 too many in part 2, need to fix
func Run(passportList string) {
	valid := 0
	valid2 := 0
	passports := getPassports(passportList)
	fmt.Println(len(passports))
	for _, p := range passports {
		isValid := checkValidity(p, 1)
		isValid2 := checkValidity(p, 2)
		if isValid {
			valid += 1
		}
		if isValid2 {
			valid2 += 1
		}
	}
	fmt.Println("Day 4 Part I: ", valid)
	fmt.Println("Day 4 Part II: ", valid2)
}

func checkValidity(passport map[string]string, part int) bool {
	var fields []string
	for k, v := range passport {
		switch k {
		case "cid":
			continue
		default:
			if part == 1 {
				if !contains(fields, k) {
					fields = append(fields, k)
				}
			} else {

				if !contains(fields, k) && isValidValue(k, v) {
					fields = append(fields, k)
				}
			}
		}
	}
	if len(fields) == 7 {
		return true
	}
	return false
}

func contains(fields []string, key string) bool {
	for _, v := range fields {
		if v == key {
			return true
		}
	}
	return false
}

func isValidValue(field string, value string) bool {
	switch field {
	case "byr":
		birth_year, err := strconv.Atoi(value)
		if err == nil && birth_year >= 1920 && birth_year <= 2002 {
			return true
		}
	case "iyr":
		issue_year, err := strconv.Atoi(value)
		if err == nil && issue_year >= 2010 && issue_year <= 2020 {
			return true
		}
	case "eyr":
		expiration_year, err := strconv.Atoi(value)
		if err == nil && expiration_year >= 2020 && expiration_year <= 2030 {
			return true
		}
	case "hgt":
		val_split := strings.Split(value, "")
		length := len(value)
		switch length {
		case 5:
			unit := val_split[3] + val_split[4]
			hgt_str := val_split[0] + val_split[1] + val_split[2]
			hgt, err := strconv.Atoi(hgt_str)
			if unit == "cm" && err == nil && hgt >= 150 && hgt <= 193 {
				return true
			}
		case 4:
			unit := val_split[2] + val_split[3]
			hgt_str := val_split[0] + val_split[1]
			hgt, err := strconv.Atoi(hgt_str)
			if unit == "in" && err == nil && hgt >= 59 && hgt <= 76 {
				return true
			}
		}
	case "hcl":
		hcl_split := strings.Split(value, "")
		if hcl_split[0] != "#" {
			return false
		}
		for i, v := range hcl_split {
			if i == 0 {
				continue
			}
			matched, err := regexp.MatchString(`[0-9]|[a-f]`, v)
			if err != nil || matched == false {
				return false
			}
		}
		return true
	case "ecl":
		colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, v := range colors {
			if value == v {
				return true
			}
		}
	case "pid":
		pid := strings.Split(value, "")
		if len(pid) != 9 {
			return false
		}
		for i, v := range pid {
			_, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if i == 8 {
				return true
			}
		}
	case "cid":
		return true
	}
	return false
}

func getPassports(passportList string) []map[string]string {
	var passports []map[string]string
	file, err := os.Open(passportList)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passport := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		chunks := strings.Split(line, " ")
		for _, chunk := range chunks {
			s_chunk := strings.Split(chunk, ":")
			k := s_chunk[0]
			v := s_chunk[1]
			passport[k] = v
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return passports
}
