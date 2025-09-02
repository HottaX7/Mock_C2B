package entity

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	responseTemplateC23 string
)

func init() {
	b, err := embedXML.ReadFile("xml/C23.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	responseTemplateC23 = string(b)
}

func NewC23(transactionNumber string, c04 *C04) string {
	amount := c04.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Text
	dbtrID := c04.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.Dbtr.Pty.ID.PrvtId.Othr.ID
	crdtID := c04.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.Cdtr.Pty.ID.PrvtId.Othr.ID
	dtTm := time.Now().Format("2006-01-02T15:04:05.123Z")

	uuid := strings.Replace(strings.ToUpper((uuid.New()).String()), "-", "", -1)

	response := responseTemplateC23
	response = strings.Replace(response, "transactionNumber", transactionNumber, -1)
	response = strings.Replace(response, "uuid", uuid, -1)
	response = strings.Replace(response, "amount", amount, -1)
	response = strings.Replace(response, "dbtrID", dbtrID, -1)
	response = strings.Replace(response, "crdtID", crdtID, -1)
	response = strings.Replace(response, "dtTm", dtTm, -1)

	return response
}

type C23 struct {
	XMLName xml.Name `xml:"IPSEnvelope"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Sign    string   `xml:"sign,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	AppHdr  struct {
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
	} `xml:"AppHdr"`
	Document struct {
		Text            string `xml:",chardata"`
		Xmlns           string `xml:"xmlns,attr"`
		FIToFIPmtStsRpt struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
			} `xml:"GrpHdr"`
			TxInfAndSts struct {
				Text            string `xml:",chardata"`
				OrgnlEndToEndId string `xml:"OrgnlEndToEndId"`
				OrgnlTxId       string `xml:"OrgnlTxId"`
				TxSts           string `xml:"TxSts"`
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
					RSN struct {
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
					IntrBkSttlmDt string `xml:"IntrBkSttlmDt"`
					Dbtr          struct {
						Text string `xml:",chardata"`
						Pty  struct {
							Text string `xml:",chardata"`
							ID   struct {
								Text   string `xml:",chardata"`
								PrvtId struct {
									Text string `xml:",chardata"`
									Othr struct {
										Text    string `xml:",chardata"`
										ID      string `xml:"Id"`
										SchmeNm struct {
											Text  string `xml:",chardata"`
											Prtry string `xml:"Prtry"`
										} `xml:"SchmeNm"`
									} `xml:"Othr"`
								} `xml:"PrvtId"`
							} `xml:"Id"`
						} `xml:"Pty"`
					} `xml:"Dbtr"`
					DbtrAgt struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"DbtrAgt"`
					CdtrAgt struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text        string `xml:",chardata"`
							ClrSysMmbId struct {
								Text  string `xml:",chardata"`
								MmbId string `xml:"MmbId"`
							} `xml:"ClrSysMmbId"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"CdtrAgt"`
					Cdtr struct {
						Text string `xml:",chardata"`
						Pty  struct {
							Text string `xml:",chardata"`
							ID   struct {
								Text   string `xml:",chardata"`
								PrvtId struct {
									Text string `xml:",chardata"`
									Othr struct {
										Text    string `xml:",chardata"`
										ID      string `xml:"Id"`
										SchmeNm struct {
											Text  string `xml:",chardata"`
											Prtry string `xml:"Prtry"`
										} `xml:"SchmeNm"`
									} `xml:"Othr"`
								} `xml:"PrvtId"`
							} `xml:"Id"`
						} `xml:"Pty"`
					} `xml:"Cdtr"`
				} `xml:"OrgnlTxRef"`
				SplmtryData struct {
					Text  string `xml:",chardata"`
					Envlp struct {
						Text  string `xml:",chardata"`
						IpsDt struct {
							Text         string `xml:",chardata"`
							Document     string `xml:"document,attr"`
							TrnCnlRsn    string `xml:"TrnCnlRsn"`
							OrgnlTxId    string `xml:"OrgnlTxId"`
							OperDate     string `xml:"OperDate"`
							OperProcDate string `xml:"OperProcDate"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}
