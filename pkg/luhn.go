package pkg

import (
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
)

type Request struct {
	Number string `json:"number"`
}

type Response struct {
	Valid bool `json:"valid"`
}

func reverse(s string) (r string) {
	for _, i := range s {
		r = string(i) + r
	}

	return r
}

func luhn(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) < 16 || len(s) > 16 {
		return false
	}

	s = reverse(s)
	sum := 0
	for i, c := range s {
		n := int(c - '0')
		if i%2 == 1 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	if sum%10 != 0 {
		return false
	}
	
	return true
}

func LunhHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	fmt.Println("r.body", r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error decoding request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := luhn(req.Number)
	response := Response{Valid: result}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	return
}
