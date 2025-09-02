package entity

import "encoding/xml"

type C05 struct {
	XMLName        xml.Name `xml:"IPSEnvelope"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Header         string   `xml:"header,attr"`
	AttrDocument   string   `xml:"document,attr"`
	Sign           string   `xml:"sign,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	AppHdr         struct {
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
		Sgntr     struct {
			Text string `xml:",chardata"`
			Sign string `xml:"Sign"`
		} `xml:"Sgntr"`
	} `xml:"AppHdr"`
	Document struct {
		Text              string `xml:",chardata"`
		FIToFICstmrCdtTrf struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text     string `xml:",chardata"`
				MsgId    string `xml:"MsgId"`
				CreDtTm  string `xml:"CreDtTm"`
				NbOfTxs  string `xml:"NbOfTxs"`
				SttlmInf struct {
					Text     string `xml:",chardata"`
					SttlmMtd string `xml:"SttlmMtd"`
				} `xml:"SttlmInf"`
				PmtTpInf struct {
					Text      string `xml:",chardata"`
					LclInstrm struct {
						Text  string `xml:",chardata"`
						Prtry string `xml:"Prtry"`
					} `xml:"LclInstrm"`
				} `xml:"PmtTpInf"`
			} `xml:"GrpHdr"`
			CdtTrfTxInf struct {
				Text  string `xml:",chardata"`
				PmtId struct {
					Text       string `xml:",chardata"`
					EndToEndId string `xml:"EndToEndId"`
					TxId       string `xml:"TxId"`
				} `xml:"PmtId"`
				IntrBkSttlmAmt struct {
					Text string `xml:",chardata"`
					Ccy  string `xml:"Ccy,attr"`
				} `xml:"IntrBkSttlmAmt"`
				AccptncDtTm string `xml:"AccptncDtTm"`
				ChrgBr      string `xml:"ChrgBr"`
				Dbtr        struct {
					Text    string `xml:",chardata"`
					Nm      string `xml:"Nm"`
					PstlAdr struct {
						Text    string   `xml:",chardata"`
						AdrLine []string `xml:"AdrLine"`
					} `xml:"PstlAdr"`
					ID struct {
						Text   string `xml:",chardata"`
						PrvtId struct {
							Text string `xml:",chardata"`
							Othr []struct {
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
					Text    string `xml:",chardata"`
					Nm      string `xml:"Nm"`
					PstlAdr struct {
						Text    string   `xml:",chardata"`
						AdrLine []string `xml:"AdrLine"`
					} `xml:"PstlAdr"`
					ID struct {
						Text   string `xml:",chardata"`
						PrvtId struct {
							Text string `xml:",chardata"`
							Othr []struct {
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
				} `xml:"Cdtr"`
				RmtInf struct {
					Text  string `xml:",chardata"`
					Ustrd string `xml:"Ustrd"`
				} `xml:"RmtInf"`
				SplmtryData struct {
					Text  string `xml:",chardata"`
					Envlp struct {
						Text  string `xml:",chardata"`
						IpsDt struct {
							Text     string `xml:",chardata"`
							TrnCV    string `xml:"TrnCV"`
							RefTrnId string `xml:"RefTrnId"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"CdtTrfTxInf"`
		} `xml:"FIToFICstmrCdtTrf"`
	} `xml:"Document"`
}
