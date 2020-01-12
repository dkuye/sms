package sms

import (
	"gopkg.in/resty.v1"
	"os"
	"strings"
)

func DigitalPulse(receivers, message string) (DigitalPulseResponse, error) {

	var response DigitalPulseResponse

	_, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey": os.Getenv("DIGITAL_PULSE_KEY"),
		}).
		SetBody(map[string]interface{}{
			"sender": os.Getenv("DIGITAL_PULSE_SENDER"),
			"message": strings.Replace(message, "<br>", "\n", -1),
			"receiver": receivers,
		}).
		SetResult(&response).
		Post("https://bss.digitalpulseapi.com/bss/api/v1.0/sms/send")

	return response, err

}


type DigitalPulseResponse struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Data int `json:"data"`
}
