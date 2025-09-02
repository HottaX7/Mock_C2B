package entity

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// B22 — минимальная структура, чтобы сформировать ответ B22
type B22 struct {
	XMLName   xml.Name `xml:"IPSEnvelope"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Header    string   `xml:"header,attr"`
	Document  string   `xml:"document,attr"`
	Sign      string   `xml:"sign,attr"`
	Xsi       string   `xml:"xsi,attr"`
	SchemaLoc string   `xml:"schemaLocation,attr"`

	AppHdr struct {
		Text  string `xml:",chardata"`
		Xmlns string `xml:"xmlns,attr"`
		Fr    struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMsgIdr string `xml:"BizMsgIdr"`
		MsgDefIdr string `xml:"MsgDefIdr"`
		BizSvc    string `xml:"BizSvc"`
		CreDt     string `xml:"CreDt"`
		Sgntr     struct {
			Text string `xml:",chardata"`
			Sign string `xml:"Sign"`
		} `xml:"Sgntr"`
		Rltd struct {
			Text string `xml:",chardata"`
			Fr   struct {
				Text string `xml:",chardata"`
				FIId struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"FIId"`
			} `xml:"Fr"`
			To struct {
				Text string `xml:",chardata"`
				FIId struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"FIId"`
			} `xml:"To"`
			BizMsgIdr string `xml:"BizMsgIdr"`
			MsgDefIdr string `xml:"MsgDefIdr"`
			BizSvc    string `xml:"BizSvc"`
			CreDt     string `xml:"CreDt"`
		} `xml:"Rltd"`
	} `xml:"AppHdr"`

	DocumentContent struct {
		Text            string `xml:",chardata"`
		Xmlns           string `xml:"xmlns,attr"`
		FIToFIPmtStsRpt struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
			} `xml:"GrpHdr"`
			OrgnlGrpInfAndSts struct {
				Text         string `xml:",chardata"`
				OrgnlMsgId   string `xml:"OrgnlMsgId"`
				OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
				OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
			} `xml:"OrgnlGrpInfAndSts"`
			TxInfAndSts struct {
				Text            string `xml:",chardata"`
				OrgnlEndToEndId string `xml:"OrgnlEndToEndId"`
				OrgnlTxId       string `xml:"OrgnlTxId"`
				StsRsnInf       struct {
					Text  string `xml:",chardata"`
					Orgtr struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text  string `xml:",chardata"`
							OrgId struct {
								Text string `xml:",chardata"`
								Othr struct {
									Text string `xml:",chardata"`
									ID   string `xml:"Id"`
								} `xml:"Othr"`
							} `xml:"OrgId"`
						} `xml:"Id"`
					} `xml:"Orgtr"`
					Rsn struct {
						Text  string `xml:",chardata"`
						Prtry string `xml:"Prtry"`
					} `xml:"Rsn"`
					AddtlInf string `xml:"AddtlInf"`
				} `xml:"StsRsnInf"`
				AccptncDtTm string `xml:"AccptncDtTm"`
				OrgnlTxRef  struct {
					Text           string `xml:",chardata"`
					IntrBkSttlmAmt struct {
						Text string `xml:",chardata"`
						Ccy  string `xml:"Ccy,attr"`
					} `xml:"IntrBkSttlmAmt"`
				} `xml:"OrgnlTxRef"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}

const responseTemplateB22 = `<?xml version="1.0" encoding="UTF-8"?>
<IPSEnvelope xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09">
  <AppHdr>
    <Fr>
      <FIId>
        <FinInstnId>
          <Othr>
            <ID>fromId</ID>
          </Othr>
        </FinInstnId>
      </FIId>
    </Fr>
    <To>
      <FIId>
        <FinInstnId>
          <Othr>
            <ID>toId</ID>
          </Othr>
        </FinInstnId>
      </FIId>
    </To>
    <BizMsgIdr>bizMsgIdr</BizMsgIdr>
    <MsgDefIdr>B22</MsgDefIdr>
    <CreDt>time1</CreDt>
    <Sgntr>
      <Sign>dummy-signature</Sign>
    </Sgntr>
  </AppHdr>
  <Document>
    <FIToFIPmtStsRpt>
      <GrpHdr>
        <MsgId>uuid</MsgId>
        <CreDtTm>time2</CreDtTm>
      </GrpHdr>
      <TxInfAndSts>
        <OrgnlEndToEndId>C2CPush</OrgnlEndToEndId>
        <OrgnlTxId>transactionNumber</OrgnlTxId>
        <TxSts>ACWP</TxSts>
        <StsRsnInf>
          <Orgtr>
            <ID>
              <OrgId>
                <Othr>
                  <ID>fromId</ID>
                </Othr>
              </OrgId>
            </ID>
          </Orgtr>
          <Rsn>
            <Prtry>I00000</Prtry>
          </Rsn>
          <AddtlInf>OPKC_OK</AddtlInf>
        </StsRsnInf>
      </TxInfAndSts>
    </FIToFIPmtStsRpt>
  </Document>
</IPSEnvelope>`

// NewB22Response создаёт XML-ответ B22 на основе данных из B21 (только нужные поля)
func NewB22Response(transactionNumber, fromId, toId, bizMsgIdr string) string {
	timeNow := time.Now().Format("2006-01-02T15:04:05Z")
	uuidStr := strings.Replace(strings.ToUpper(uuid.New().String()), "-", "", -1)

	log.Info().
		Str("transactionNumber", transactionNumber).
		Str("fromId", fromId).
		Str("toId", toId).
		Str("bizMsgIdr", bizMsgIdr).
		Str("uuid", uuidStr).
		Msg("B22: формируется ответ")

	response := responseTemplateB22
	response = strings.ReplaceAll(response, "fromId", fromId)
	response = strings.ReplaceAll(response, "toId", toId)
	response = strings.ReplaceAll(response, "bizMsgIdr", bizMsgIdr)
	response = strings.ReplaceAll(response, "time1", timeNow)
	response = strings.ReplaceAll(response, "time2", timeNow)
	response = strings.ReplaceAll(response, "uuid", uuidStr)
	response = strings.ReplaceAll(response, "transactionNumber", transactionNumber)

	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B22: ответ успешно сформирован")

	return response
}
