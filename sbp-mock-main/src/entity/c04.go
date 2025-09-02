package entity

import "encoding/xml"

type C04 struct {
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
		Rltd      struct {
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
						} `xml:"Pty"`
					} `xml:"Cdtr"`
					CdtrAcct struct {
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
					} `xml:"CdtrAcct"`
				} `xml:"OrgnlTxRef"`
				SplmtryData struct {
					Text  string `xml:",chardata"`
					Envlp struct {
						Text  string `xml:",chardata"`
						IpsDt struct {
							Text     string `xml:",chardata"`
							Document string `xml:"document,attr"`
							TrnCV    string `xml:"TrnCV"`
							RefTrnId string `xml:"RefTrnId"`
							RcvrDt   struct {
								Text      string `xml:",chardata"`
								FrScrRcvr string `xml:"FrScrRcvr"`
							} `xml:"RcvrDt"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}
