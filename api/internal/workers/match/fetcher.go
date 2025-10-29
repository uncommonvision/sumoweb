package match

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sumoweb/server/internal/models/sumojp"
)

func FetchTorikumiData(division, day int) (*sumojp.Response, error) {
	url := fmt.Sprintf("https://www.sumo.or.jp/EnHonbashoMain/torikumiAjax/%d/%d/", division, day)

	formData := fmt.Sprintf("basho_id=631&kakuzuke_id=%d&day=%d", division, day)

	req, err := http.NewRequest("POST", url, strings.NewReader(formData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch torikumi data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response sumojp.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
