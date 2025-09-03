package entity

import (
	"embed"
	"encoding/xml"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	responseTemplateB21 string
)

//go:embed xml/B21.xml
var embedXMLB21 embed.FS

func init() {
	b, err := embedXMLB21.ReadFile("xml/B21.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	responseTemplateB21 = string(b)
}

func NewB21(transactionNumber string, b05 *B05Request) string {
	template := string(responseTemplateB21)

	// Очищаем значения перед подстановкой
	fromId := strings.TrimSpace(b05.AppHdr.Fr.FIId.FinInstnId.Othr.ID)
	toId := strings.TrimSpace(b05.AppHdr.To.FIId.FinInstnId.Othr.ID)
	bizMsgIdr := strings.TrimSpace(b05.AppHdr.BizMsgIdr)
	creDt := strings.TrimSpace(b05.AppHdr.CreDt)
	creDtTm := time.Now().Format("2006-01-02T15:04:05:123Z")
	//ident13 := b05.AppHdr.CreDt
	//ident79 := b05.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId
	amount := strings.TrimSpace(b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Text)
	//dbtrNm := b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.Nm
	//dbtrID := b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.ID.PrvtId.Othr.ID
	//dbtrAcct := b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.ID.Othr.ID
	//crdtID := b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.ID.PrvtId.Othr.ID

	// // Идентификаторы участников
	// var dbtrID string
	// if len(b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.ID.PrvtId.Othr) > 0 {
	//     dbtrID = b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.ID.PrvtId.Othr[0].ID
	// }

	// var cdtrID string
	// if len(b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.ID.OrgId.Othr) > 0 {
	//     cdtrID = b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.ID.OrgId.Othr[0].ID
	// }

	// Финансовые данные
	currency := strings.TrimSpace(b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy)

	// Генерация уникальных идентификаторов

	uuid := strings.Replace(strings.ToUpper((uuid.New()).String()), "-", "", -1)

	// Формируем ответ
	template = strings.Replace(template, "fromId", fromId, -1)
	template = strings.Replace(template, "toId", toId, -1)
	template = strings.Replace(template, "bizMsgIdr", bizMsgIdr, -1)
	template = strings.Replace(template, "creDt", creDt, -1)
	template = strings.Replace(template, "creDtTm", creDtTm, -1)
	// response = strings.Replace(response, "dbtrID", dbtrID, -1)
	// response = strings.Replace(response, "cdtrID", cdtrID, -1)
	template = strings.Replace(template, "amount", amount, -1)
	template = strings.Replace(template, "currency", currency, -1)
	template = strings.Replace(template, "transactionNumber", transactionNumber, -1)
	template = strings.Replace(template, "uuid", uuid, -1)

	return template
}

// 	// Извлечение информации из структуры B21
// 	amount := b05.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Text
// 	var dbtrID string
// 	// Поиск ID должника
// 	for _, id := range b21.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.Dbtr.ID.PrvtId.Othr {
// 		switch id.SchmeNm.Prtry {
// 		case "MTEL":
// 			dbtrID = id.ID
// 		}
// 	}

// 	var crdtID string
// 	// Поиск ID кредитора
// 	for _, id := range b21.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.Cdtr.ID.PrvtId.Othr {
// 		switch id.SchmeNm.Prtry {
// 		case "MTEL":
// 			crdtID = id.ID
// 		}
// 	}

// 	// Генерация уникального идентификатора
// 	uuid := strings.Replace(strings.ToUpper((uuid.New()).String()), "-", "", -1)

// 	// Заполнение шаблона данными
// 	response := responseTemplateB21
// 	response = strings.Replace(response, "transactionNumber", transactionNumber, -1)
// 	response = strings.Replace(response, "uuid", uuid, -1)
// 	response = strings.Replace(response, "amount", amount, -1)
// //	response = strings.Replace(response, "dbtrID", dbtrID, -1)
// 	response = strings.Replace(response, "crdtID", crdtID, -1)

// 	return response
// }

type B21 struct {
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
		BizMsgIdr string    `xml:"BizMsgIdr"`
		MsgDefIdr string    `xml:"MsgDefIdr"`
		BizSvc    string    `xml:"BizSvc"`
		CreDt     time.Time `xml:"CreDt"`
		Sgntr     struct {
			Text string `xml:",chardata"`
			Sign struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
			} `xml:"Sign"`
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
							Text string `xml:",chardata"`
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
							OperProcDate string `xml:"OperProcDate"`
							OperDate     string `xml:"OperDate"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}

// type B21 struct {
// 	XMLName        xml.Name `xml:"IPSEnvelope"`
// 	Text           string   `xml:",chardata"`
// 	Sign           string   `xml:"sign,attr"`
// 	Xsi            string   `xml:"xsi,attr"`
// 	SchemaLocation string   `xml:"schemaLocation,attr"`
// 	Xmlns          string   `xml:"xmlns,attr"`
// 	AppHdr         struct {
// 		Text  string `xml:",chardata"`
// 		Xmlns string `xml:"xmlns,attr"`
// 		Fr    struct {
// 			Text string `xml:",chardata"`
// 			FIId struct {
// 				Text       string `xml:",chardata"`
// 				FinInstnId struct {
// 					Text string `xml:",chardata"`
// 					Othr struct {
// 						Text string `xml:",chardata"`
// 						ID   string `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"Fr"`
// 		To struct {
// 			Text string `xml:",chardata"`
// 			FIId struct {
// 				Text       string `xml:",chardata"`
// 				FinInstnId struct {
// 					Text string `xml:",chardata"`
// 					Othr struct {
// 						Text string `xml:",chardata"`
// 						ID   string `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"To"`
// 		BizMsgIdr string `xml:"BizMsgIdr"`
// 		MsgDefIdr string `xml:"MsgDefIdr"`
// 		BizSvc    string `xml:"BizSvc"`
// 		CreDt     string `xml:"CreDt"`
// 	} `xml:"AppHdr"`
// 	Document struct {
// 		Text        string `xml:",chardata"`
// 		Xmlns       string `xml:"xmlns,attr"`
// 		IdVrfctnReq struct {
// 			Text    string `xml:",chardata"`
// 			Assgnmt struct {
// 				Text    string `xml:",chardata"`
// 				MsgId   string `xml:"MsgId"`
// 				CreDtTm string `xml:"CreDtTm"`
// 				Cretr   struct {
// 					Text string `xml:",chardata"`
// 					Pty  struct {
// 						Text string `xml:",chardata"`
// 						ID   struct {
// 							Text   string `xml:",chardata"`
// 							PrvtId struct {
// 								Text string `xml:",chardata"`
// 								Othr struct {
// 									Text    string `xml:",chardata"`
// 									ID      string `xml:"Id"`
// 									SchmeNm struct {
// 										Text  string `xml:",chardata"`
// 										Prtry string `xml:"Prtry"`
// 									} `xml:"SchmeNm"`
// 								} `xml:"Othr"`
// 							} `xml:"PrvtId"`
// 						} `xml:"Id"`
// 					} `xml:"Pty"`
// 				} `xml:"Cretr"`
// 				Assgnr struct {
// 					Text string `xml:",chardata"`
// 					Agt  struct {
// 						Text       string `xml:",chardata"`
// 						FinInstnId struct {
// 							Text string `xml:",chardata"`
// 							Othr struct {
// 								Text string `xml:",chardata"`
// 								ID   string `xml:"Id"`
// 							} `xml:"Othr"`
// 						} `xml:"FinInstnId"`
// 					} `xml:"Agt"`
// 				} `xml:"Assgnr"`
// 				Assgne struct {
// 					Text string `xml:",chardata"`
// 					Agt  struct {
// 						Text       string `xml:",chardata"`
// 						FinInstnId struct {
// 							Text string `xml:",chardata"`
// 							Othr struct {
// 								Text string `xml:",chardata"`
// 								ID   string `xml:"Id"`
// 							} `xml:"Othr"`
// 						} `xml:"FinInstnId"`
// 					} `xml:"Agt"`
// 				} `xml:"Assgne"`
// 			} `xml:"Assgnmt"`
// 			Vrfctn struct {
// 				Text         string `xml:",chardata"`
// 				ID           string `xml:"Id"`
// 				PtyAndAcctId struct {
// 					Text string `xml:",chardata"`
// 					Pty  struct {
// 						Text string `xml:",chardata"`
// 						ID   struct {
// 							Text   string `xml:",chardata"`
// 							PrvtId struct {
// 								Text string `xml:",chardata"`
// 								Othr struct {
// 									Text    string `xml:",chardata"`
// 									ID      string `xml:"Id"`
// 									SchmeNm struct {
// 										Text  string `xml:",chardata"`
// 										Prtry string `xml:"Prtry"`
// 									} `xml:"SchmeNm"`
// 								} `xml:"Othr"`
// 							} `xml:"PrvtId"`
// 						} `xml:"Id"`
// 					} `xml:"Pty"`
// 				} `xml:"PtyAndAcctId"`
// 			} `xml:"Vrfctn"`
// 			SplmtryData struct {
// 				Text  string `xml:",chardata"`
// 				Envlp struct {
// 					Text  string `xml:",chardata"`
// 					IpsDt struct {
// 						Text     string `xml:",chardata"`
// 						Document string `xml:"document,attr"`
// 						TrnTp    string `xml:"TrnTp"`
// 					} `xml:"IpsDt"`
// 				} `xml:"Envlp"`
// 			} `xml:"SplmtryData"`
// 		} `xml:"IdVrfctnReq"`
// 	} `xml:"Document"`
// }
