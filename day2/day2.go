package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run(corporatePolicies string) {
	policies := getPasswordPolicies(corporatePolicies)
	count := 0
	count2 := 0
	for _, policy := range policies {
		l, h := getCounts(policy)
		val := getValue(policy)
		password := strings.Split(policy, " ")[2]
		valid := checkPassword(val, l, h, password)
		valid2 := checkPassword2(val, l, h, password)
		if valid {
			count += 1
		}
		if valid2 {
			count2 += 1
		}
	}
	fmt.Println(count)
	fmt.Println(count2)
}

func checkPassword2(val string, l int, h int, pw string) bool {
	good := false
	for i, v := range pw {
		if i+1 == l || i+1 == h {
			if string(v) == val {
				if good == false {
					good = true
				} else {
					good = false
				}
			}
		}
	}
	return good
}

func checkPassword(val string, l int, h int, pw string) bool {
	count := 0
	for _, v := range pw {
		if string(v) == val {
			count += 1
		}
	}

	if count >= l && count <= h {
		return true
	}
	return false
}

func getValue(policy string) string {
	second := strings.Split(policy, " ")[1]
	val := strings.Split(second, "")[0]
	return val
}

func getCounts(policy string) (int, int) {
	nums := strings.Split(policy, " ")[0]
	low, _ := strconv.Atoi(strings.Split(nums, "-")[0])
	high, _ := strconv.Atoi(strings.Split(nums, "-")[1])
	return low, high
}

func getPasswordPolicies(iFile string) []string {
	var policies []string
	file, err := os.Open(iFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		policies = append(policies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return policies
}
