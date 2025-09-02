package entity

import (
	"encoding/xml"
	"strings"
	"time"

	"math/rand"
	"strconv"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// убираем объявление embedXML

var (
	responseTemplateB06 string
)

func init() {
	b, err := embedXML.ReadFile("xml/B06.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	responseTemplateB06 = string(b)
}

func NewB06(transactionID string, b05 *B05) string {
	fromId := b05.AppHdr.Fr.FIId.FinInstnId.Othr.ID
	toId := b05.AppHdr.To.FIId.FinInstnId.Othr.ID
	bizMsgIdr := b05.AppHdr.BizMsgIdr
	creDt := b05.AppHdr.CreDt
	creDtTm := time.Now().Format("2006-01-02T15:04:05:123Z")
	uuidStr := strings.Replace(strings.ToUpper(uuid.New().String()), "-", "", -1)
	operationControlNumber := "b05res" + strconv.Itoa(rand.Intn(9999-1000)+1000) + "xq"

	log.Info().
		Str("transactionID", transactionID).
		Str("fromId", fromId).
		Str("toId", toId).
		Str("bizMsgIdr", bizMsgIdr).
		Str("creDt", creDt).
		Str("creDtTm", creDtTm).
		Str("uuid", uuidStr).
		Str("operationControlNumber", operationControlNumber).
		Msg("B06: формируется ответ")

	response := responseTemplateB06
	response = strings.Replace(response, "fromId", fromId, -1)
	response = strings.Replace(response, "toId", toId, -1)
	response = strings.Replace(response, "bizMsgIdr", bizMsgIdr, -1)
	response = strings.Replace(response, "time1", creDt, -1)
	response = strings.Replace(response, "time2", creDtTm, -1)
	//	response = strings.Replace(response, "ident", msgId, -1)
	response = strings.Replace(response, "transactionNumber", transactionID, -1)
	response = strings.Replace(response, "uuid", uuidStr, -1)
	response = strings.Replace(response, "operationControlNumber", operationControlNumber, -1)

	log.Info().
		Str("transactionID", transactionID).
		Msg("B06: ответ успешно сформирован")

	return response
}

type B06 struct {
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

// package entity

// import (
// 	"encoding/xml"
// 	"strings"
// 	"time"
// 	"github.com/google/uuid"
// 	"github.com/rs/zerolog/log"
// 	"strconv"
// 	"math/rand"
// )

// var (
// 	responseTemplateB06 string
// )

// func init() {
// 	b, err := embedXML.ReadFile("xml/B06.xml")
// 	if err != nil {
// 		log.Fatal().Err(err).Send()
// 	}

// 	responseTemplateB06 = string(b)
// }

// func NewB06(transactionID string, b05 *B05) string {
// 	fromId := b05.AppHdr.Fr.FIId.FinInstnId.Othr.ID
// 	toId := b05.AppHdr.To.FIId.FinInstnId.Othr.ID
// 	bizMsgIdr := b05.AppHdr.BizMsgIdr
// 	creDt := b05.AppHdr.CreDt
// 	creDtTm := time.Now().Format("2006-01-02T15:04:05:123Z")
// 	uuid := strings.Replace(strings.ToUpper(uuid.New().String()), "-", "", -1)
// 	operationControlNumber := "b05res" + strconv.Itoa(rand.Intn(9999-1000)+1000) + "xq"

// 	response := responseTemplateB06
// 	response = strings.Replace(response, "fromId", fromId, -1)
// 	response = strings.Replace(response, "toId", toId, -1)
// 	response = strings.Replace(response, "bizMsgIdr", bizMsgIdr, -1)
// 	response = strings.Replace(response, "time1", creDt, -1)
// 	response = strings.Replace(response, "time2", creDtTm, -1)
// //	response = strings.Replace(response, "ident", msgId, -1)
// 	response = strings.Replace(response, "transactionNumber", transactionID, -1)
// 	response = strings.Replace(response, "uuid", uuid, -1)
// 	response = strings.Replace(response, "operationControlNumber", operationControlNumber, -1)

// 	return response
// }

// // type B06 struct {
// // 	XMLName xml.Name `xml:"IPSEnvelope"`
// // 	Header  AppHdr   `xml:"AppHdr"`
// // 	Document Document `xml:"Document"`
// // }

// // type AppHdr struct {
// // 	Fr         FIId   `xml:"Fr"`
// // 	To         FIId   `xml:"To"`
// // 	BizMsgIdr  string `xml:"BizMsgIdr"`
// // 	MsgDefIdr  string `xml:"MsgDefIdr"`
// // 	BizSvc     string `xml:"BizSvc"`
// // 	CreDt      string `xml:"CreDt"`
// // 	Sgntr      string `xml:"Sgntr>Sign"`
// // 	Rltd       Rltd   `xml:"Rltd"`
// // }

// // type FIId struct {
// // 	FinInstnId FinInstnId `xml:"FIId>FinInstnId"`
// // }

// // type FinInstnId struct {
// // 	Othr Othr `xml:"Othr"`
// // }

// // type Othr struct {
// // 	Id string `xml:"Id"`
// // }

// // type Rltd struct {
// // 	Fr        FIId   `xml:"Fr"`
// // 	To        FIId   `xml:"To"`
// // 	BizMsgIdr string `xml:"BizMsgIdr"`
// // 	MsgDefIdr string `xml:"MsgDefIdr"`
// // 	BizSvc    string `xml:"BizSvc"`
// // 	CreDt     string `xml:"CreDt"`
// // }

// // type Document struct {
// // 	FIToFIPmtStsRpt FIToFIPmtStsRpt `xml:"FIToFIPmtStsRpt"`
// // }

// // type FIToFIPmtStsRpt struct {
// // 	GrpHdr              GrpHdr              `xml:"GrpHdr"`
// // 	OrgnlGrpInfAndSts   OrgnlGrpInfAndSts   `xml:"OrgnlGrpInfAndSts"`
// // 	TxInfAndSts         TxInfAndSts         `xml:"TxInfAndSts"`
// // }

// // type GrpHdr struct {
// // 	MsgId    string `xml:"MsgId"`
// // 	CreDtTm  string `xml:"CreDtTm"`
// // }

// // type OrgnlGrpInfAndSts struct {
// // 	OrgnlMsgId   string `xml:"OrgnlMsgId"`
// // 	OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
// // 	OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
// // }

// // type TxInfAndSts struct {
// // 	OrgnlEndToEndId string      `xml:"OrgnlEndToEndId"`
// // 	OrgnlTxId      string      `xml:"OrgnlTxId"`
// // 	TxSts          string      `xml:"TxSts"`
// // 	StsRsnInf      StsRsnInf   `xml:"StsRsnInf"`
// // 	AccptncDtTm    string      `xml:"AccptncDtTm"`
// // 	OrgnlTxRef     OrgnlTxRef  `xml:"OrgnlTxRef"`
// // 	SplmtryData    SplmtryData `xml:"SplmtryData"`
// // }

// // type StsRsnInf struct {
// // 	Orgtr Orgtr `xml:"Orgtr"`
// // 	Rsn   Rsn   `xml:"Rsn"`
// // 	AddtlInf string `xml:"AddtlInf"`
// // }

// // type Orgtr struct {
// // 	Id OrgId `xml:"Id"`
// // }

// // type OrgId struct {
// // 	Othr Othr `xml:"Othr"`
// // }

// // type Rsn struct {
// // 	Prtry string `xml:"Prtry"`
// // }

// // type OrgnlTxRef struct {
// // 	IntrBkSttlmAmt IntrBkSttlmAmt `xml:"IntrBkSttlmAmt"`
// // 	Dbtr           Dbtr           `xml:"Dbtr"`
// // 	DbtrAcct       DbtrAcct       `xml:"DbtrAcct"`
// // 	DbtrAgt        FinInstnId     `xml:"DbtrAgt"`
// // 	CdtrAgt        FinInstnId     `xml:"CdtrAgt"`
// // 	Cdtr           Cdtr           `xml:"Cdtr"`
// // 	CdtrAcct       CdtrAcct       `xml:"CdtrAcct"`
// // }

// // type IntrBkSttlmAmt struct {
// // 	Ccy string `xml:"Ccy,attr"`
// // 	Value string `xml:",chardata"`
// // }

// // type Dbtr struct {
// // 	Pty Pty `xml:"Pty"`
// // }

// // type Cdtr struct {
// // 	Pty Pty `xml:"Pty"`
// // }

// // type DbtrAcct struct {
// // 	Id Othr `xml:"Id>Othr"`
// // }

// // type CdtrAcct struct {
// // 	Id Othr `xml:"Id>Othr"`
// // }

// // type Pty struct {
// // 	Nm      string `xml:"Nm"`
// // 	Id      OrgId  `xml:"Id"`
// // }

// // type SplmtryData struct {
// // 	Envlp Envlp `xml:"Envlp"`
// // }

// // type Envlp struct {
// // 	IpsDt IpsDt `xml:"IpsDt"`
// // }

// // type IpsDt struct {
// // 	FrScr    string `xml:"FrScr"`
// // 	OperDate string `xml:"OperDate"`
// // }
// type B06 struct {
// 	XMLName xml.Name `xml:"IPSEnvelope"`
// 	Xsi     string   `xml:"xsi,attr"`
// 	Xmlns   string   `xml:"xmlns,attr"`
// 	AppHdr  struct {
// 		Xmlns     string `xml:"xmlns,attr"`
// 		Fr        struct {
// 			FIId struct {
// 				FinInstnId struct {
// 					Othr struct {
// 						ID string `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"Fr"`
// 		To struct {
// 			FIId struct {
// 				FinInstnId struct {
// 					Othr struct {
// 						ID string `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"To"`
// 		BizMsgIdr string `xml:"BizMsgIdr"`
// 		CreDt     string `xml:"CreDt"`
// 	} `xml:"AppHdr"`
// 	Document struct {
// 		Xmlns  string `xml:"xmlns,attr"`
// 		RspnData struct {
// 			MsgId   string `xml:"MsgId"`
// 			CreDtTm string `xml:"CreDtTm"`
// 			Status  string `xml:"Status"`
// 		} `xml:"RspnData"`
// 	} `xml:"Document"`
// }
