package maps

import "fmt"

func Maps() {
	websites := map[string]string{
		"Google":              "https://googl.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
}
