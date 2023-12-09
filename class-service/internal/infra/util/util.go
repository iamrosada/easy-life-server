package util

import (
	"time"
)

func GenerateNewID() uint {
	return uint(time.Now().UnixNano())
}

// func FetchStudentFromMicroserviceUser(url string) (int, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return 0, fmt.Errorf("Failed to fetch promotion, status: %d", resp.StatusCode)
// 	}

// 	var response struct {
// 		Students_ids []any `json:"students_ids"`
// 	}

// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		return 0, err
// 	}

// 	fmt.Println("Return of function fetchPromoTypeFromMicroservice:", response.Students_ids)

// 	return response, nil
// }
