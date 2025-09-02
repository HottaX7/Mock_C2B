package entity

import (
	"encoding/xml"
	"github.com/rs/zerolog/log"
)

// SendPaymentRequest — структура для парсинга входящего запроса
type SendPaymentRequest struct {
	XMLName            xml.Name `xml:"Request"`
	Id                 string   `xml:"Id,attr"`
	Service            string   `xml:"Service,attr"`
	Time               string   `xml:"Time,attr"`
	Value              string   `xml:"Value,attr"`
	Commission         string   `xml:"Commission,attr"`
	Currency           string   `xml:"Currency,attr"`
	PaymentTool        string   `xml:"PaymentTool,attr"`
	PaymentToolSubType string   `xml:"PaymentToolSubType,attr"`
	Session            string   `xml:"Session,attr"`

	PaymentParameters []struct {
		Name  string `xml:"Name,attr"`
		Value string `xml:"Value,attr"`
	} `xml:"PaymentParameters>Parameter"`
}

// SendPaymentResponse — структура для формирования ответа
type SendPaymentResponse struct {
	XMLName             xml.Name `xml:"Response"`
	PaymentId           string   `xml:"PaymentId,attr"`
	PaymentTime         string   `xml:"PaymentTime,attr"`
	State               string   `xml:"State,attr"`
	StateComment        string   `xml:"StateComment,attr"`
	ProcessingErrorCode string   `xml:"ProcessingErrorCode,attr"`
	Currency            string   `xml:"Currency,attr"`
}

// ParseSendPaymentRequest парсит XML в структуру SendPaymentRequest с логированием
func ParseSendPaymentRequest(data []byte) (*SendPaymentRequest, error) {
	var req SendPaymentRequest
	err := xml.Unmarshal(data, &req)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка парсинга SendPaymentRequest из XML")
		return nil, err
	}
	log.Info().
		Str("Id", req.Id).
		Str("Service", req.Service).
		Str("Value", req.Value).
		Msg("SendPaymentRequest успешно разобран из XML")
	return &req, nil
}

// NewSendPaymentResponse формирует SendPaymentResponse с логированием
func NewSendPaymentResponse(paymentId, paymentTime, state, stateComment, processingErrorCode, currency string) *SendPaymentResponse {
	log.Info().
		Str("PaymentId", paymentId).
		Str("State", state).
		Str("Currency", currency).
		Msg("SendPaymentResponse формируется")
	return &SendPaymentResponse{
		PaymentId:           paymentId,
		PaymentTime:         paymentTime,
		State:               state,
		StateComment:        stateComment,
		ProcessingErrorCode: processingErrorCode,
		Currency:            currency,
	}
}
