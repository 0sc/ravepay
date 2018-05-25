package ravepay

import (
	"fmt"
	"testing"
)

func TestSwitchToLiveMode(t *testing.T) {
	currentMode = ""
	baseURL = ""

	SwitchToLiveMode()

	if currentMode != "live" {
		t.Errorf("want %s got %s", "live", currentMode)
	}

	if baseURL != liveModeBaseURL {
		t.Errorf("want %s got %s", liveModeBaseURL, baseURL)
	}
}

func TestSwitchToTestMode(t *testing.T) {
	currentMode = ""
	baseURL = ""

	SwitchToTestMode()

	if currentMode != "test" {
		t.Errorf("want %s got %s", "test", currentMode)
	}

	if baseURL != testModeBaseURL {
		t.Errorf("want %s got %s", testModeBaseURL, baseURL)
	}
}

func Test_buildURL(t *testing.T) {
	type args struct {
		path string
		mode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "builds url with the live mode base url",
			args: args{
				path: "to/rave/endpoint",
				mode: "live",
			},
			want: fmt.Sprintf("%s%s", liveModeBaseURL, "to/rave/endpoint"),
		},
		{
			name: "builds url with the test mode base url",
			args: args{
				path: "to/rave/endpoint",
				mode: "test",
			},
			want: fmt.Sprintf("%s%s", testModeBaseURL, "to/rave/endpoint"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.args.mode {
			case "live":
				SwitchToLiveMode()
			case "test":
				SwitchToTestMode()
			}

			if got := buildURL(tt.args.path); got != tt.want {
				t.Errorf("buildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentMode(t *testing.T) {
	type args struct {
		mode string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns the current mode when in live mode",
			args: args{"live"},
			want: "live",
		},
		{
			name: "returns the current mode when in test mode",
			args: args{"test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.args.mode {
			case "live":
				SwitchToLiveMode()
			case "test":
				SwitchToTestMode()
			}

			if got := CurrentMode(); got != tt.want {
				t.Errorf("CurrentMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
