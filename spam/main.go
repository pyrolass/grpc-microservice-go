package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	var data = map[string]any{
		"driver_id": 1,
		"lat":       1,
		"lon":       1,
	}

	for {
		b, err := json.Marshal(data)

		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/drivers", bytes.NewReader(b))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Status)
		// It's good practice to close the response body when you're done with it to avoid leaking resources.
		resp.Body.Close()
	}

}
