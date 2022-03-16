package weather

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetWeather(t *testing.T) {
	testData := []struct {
		name          string
		server        *httptest.Server
		response      *Weather
		expectedError error
	}{
		{
			name: "Basic Request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"city": "Chandigarh", "forecast": "Sunny"}`))
			})),
			response: &Weather{
				City:     "Chandigarh",
				Forecast: "Sunny",
			},
			expectedError: nil,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			resp, err := GetWeather(test.server.URL)
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("FAILED: expected %v, got %v\n", test.expectedError, err)
			}
		})
	}
}
