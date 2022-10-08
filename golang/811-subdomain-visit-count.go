package leetcode

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) (res []string) {

	mm := make(map[string]int)
	for _, strs := range cpdomains {
		demo := strings.SplitN(strs, " ", 2)

		countStr, domain := demo[0], demo[1]

		count, _ := strconv.Atoi(countStr)

		domainList := strings.Split(domain, ".")

		for i := len(domainList) - 1; i >= 0; i-- {
			prev := strings.Join(domainList[i:], ".")
			mm[prev] += count
		}
	}

	for k, v := range mm {
		countStr := strconv.Itoa(v)
		res = append(res, countStr+" "+k)
	}
	return res
}
