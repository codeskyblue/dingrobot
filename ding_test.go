package dingrobot

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	token string
	robot *Robot
)

func init() {
	token = os.Getenv("DINGROBOT_TOKEN")
	if token == "" {
		panic("DINGROBOT_TOKEN env is empty")
	}
	robot = New(token)
}

func TestText(t *testing.T) {
	err := robot.Text("Hello\nwhat")
	assert.NoError(t, err)
}

func TestMarkdown(t *testing.T) {
	err := robot.Markdown("test markdown", ">Nice tools\n\n# Title here\n<https://www.google.com> ![screenshot](@lADOpwk3K80C0M0FoA) ")
	assert.NoError(t, err)
}

func TestLink(t *testing.T) {
	err := robot.Link("test link", "Baidu link", "https://www.baidu.com", "")
	assert.NoError(t, err)
}

func TestAtText(t *testing.T) {
	newRobot := robot.AtAll(true)
	assert.NotEqual(t, newRobot, robot)
	newRobot.Text("This is at all message")

	robot.AtMobiles("11111111112").Text("At someone")
}
