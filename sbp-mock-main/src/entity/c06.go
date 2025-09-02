package entity

import (
	"encoding/xml"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	responseTemplateC06 string
)

func init() {
	b, err := embedXML.ReadFile("xml/C06.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	responseTemplateC06 = string(b)
}

func NewC06(transactionNumber string, c05 *C05) string {
	ident13 := c05.AppHdr.CreDt
	ident79 := c05.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId
	amount := c05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Text
	dbtrNm := c05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.Nm
	var dbtrID, dbtrINN string
	for _, id := range c05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.ID.PrvtId.Othr {
		switch id.SchmeNm.Prtry {
		case "MTEL":
			dbtrID = id.ID
		case "TIDN":
			dbtrINN = id.ID
		}
	}
	dbtrAcct := c05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.ID.Othr.ID

	var crdtID, crdtINN string
	for _, id := range c05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.ID.PrvtId.Othr {
		switch id.SchmeNm.Prtry {
		case "MTEL":
			crdtID = id.ID
		case "TIDN":
			crdtINN = id.ID
		}
	}

	uuid := strings.Replace(strings.ToUpper((uuid.New()).String()), "-", "", -1)

	response := responseTemplateC06
	response = strings.Replace(response, "transactionNumber", transactionNumber, -1)
	response = strings.Replace(response, "uuid", uuid, -1)
	response = strings.Replace(response, "79_ident", ident79, -1)
	response = strings.Replace(response, "13_ident", ident13, -1)
	response = strings.Replace(response, "amount", amount, -1)
	response = strings.Replace(response, "dbtrNm", dbtrNm, -1)
	response = strings.Replace(response, "dbtrID", dbtrID, -1)
	response = strings.Replace(response, "dbtrINN", dbtrINN, -1)
	response = strings.Replace(response, "dbtrAcct", dbtrAcct, -1)
	response = strings.Replace(response, "crdtID", crdtID, -1)
	response = strings.Replace(response, "crdtINN", crdtINN, -1)

	return response
}

type C06 struct {
	XMLName xml.Name `xml:"IPSEnvelope"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
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
			Sign struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
			} `xml:"Sign"`
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
					Dbtr struct {
						Text string `xml:",chardata"`
						Pty  struct {
							Text    string `xml:",chardata"`
							Nm      string `xml:"Nm"`
							PstlAdr struct {
								Text    string `xml:",chardata"`
								AdrLine string `xml:"AdrLine"`
							} `xml:"PstlAdr"`
							ID struct {
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
							CtctDtls struct {
								Text string `xml:",chardata"`
								Nm   string `xml:"Nm"`
							} `xml:"CtctDtls"`
						} `xml:"Pty"`
					} `xml:"Dbtr"`
					DbtrAcct struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text    string `xml:",chardata"`
								ID      string `xml:"Id"`
								SchmeNm struct {
									Text  string `xml:",chardata"`
									Prtry string `xml:"Prtry"`
								} `xml:"SchmeNm"`
							} `xml:"Othr"`
						} `xml:"Id"`
					} `xml:"DbtrAcct"`
					DbtrAgt struct {
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
							Text    string `xml:",chardata"`
							Nm      string `xml:"Nm"`
							PstlAdr struct {
								Text    string `xml:",chardata"`
								AdrLine string `xml:"AdrLine"`
							} `xml:"PstlAdr"`
							ID struct {
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
							CtctDtls struct {
								Text string `xml:",chardata"`
								Nm   string `xml:"Nm"`
							} `xml:"CtctDtls"`
						} `xml:"Pty"`
					} `xml:"Cdtr"`
				} `xml:"OrgnlTxRef"`
				SplmtryData struct {
					Text  string `xml:",chardata"`
					Envlp struct {
						Text  string `xml:",chardata"`
						IpsDt struct {
							Text     string `xml:",chardata"`
							Document string `xml:"document,attr"`
							OperDate string `xml:"OperDate"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}
