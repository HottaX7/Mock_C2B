package entity

import "encoding/xml"

type A01 struct {
	XMLName        xml.Name `xml:"IPSEnvelope"`
	Text           string   `xml:",chardata"`
	Sign           string   `xml:"sign,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	AppHdr         struct {
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
	} `xml:"AppHdr"`
	Document struct {
		Text        string `xml:",chardata"`
		Xmlns       string `xml:"xmlns,attr"`
		IdVrfctnReq struct {
			Text    string `xml:",chardata"`
			Assgnmt struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
				Cretr   struct {
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
				} `xml:"Cretr"`
				Assgnr struct {
					Text string `xml:",chardata"`
					Agt  struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"Agt"`
				} `xml:"Assgnr"`
				Assgne struct {
					Text string `xml:",chardata"`
					Agt  struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"Agt"`
				} `xml:"Assgne"`
			} `xml:"Assgnmt"`
			Vrfctn struct {
				Text         string `xml:",chardata"`
				ID           string `xml:"Id"`
				PtyAndAcctId struct {
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
				} `xml:"PtyAndAcctId"`
			} `xml:"Vrfctn"`
			SplmtryData struct {
				Text  string `xml:",chardata"`
				Envlp struct {
					Text  string `xml:",chardata"`
					IpsDt struct {
						Text     string `xml:",chardata"`
						Document string `xml:"document,attr"`
						TrnTp    string `xml:"TrnTp"`
					} `xml:"IpsDt"`
				} `xml:"Envlp"`
			} `xml:"SplmtryData"`
		} `xml:"IdVrfctnReq"`
	} `xml:"Document"`
}
