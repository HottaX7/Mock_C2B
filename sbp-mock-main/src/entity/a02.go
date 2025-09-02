package entity

import (
	"embed"
	"encoding/xml"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	responseTemplateA02 string
)

//go:embed xml/*.xml
var embedXML embed.FS

func init() {
	b, err := embedXML.ReadFile("xml/A02.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	responseTemplateA02 = string(b)
}

func NewA02(transactionID string, a01 *A01) string {
	fromId := a01.AppHdr.Fr.FIId.FinInstnId.Othr.ID
	toId := a01.AppHdr.To.FIId.FinInstnId.Othr.ID
	bizMsgIdr := a01.AppHdr.BizMsgIdr
	creDt := a01.AppHdr.CreDt
	creDtTm := time.Now().Format("2006-01-02T15:04:05:123Z")
	msgId := a01.Document.IdVrfctnReq.Assgnmt.MsgId
	cretrID := a01.Document.IdVrfctnReq.Assgnmt.Cretr.Pty.ID.PrvtId.Othr.ID
	prvtID := a01.Document.IdVrfctnReq.Vrfctn.PtyAndAcctId.Pty.ID.PrvtId.Othr.ID
	uuid := strings.Replace(strings.ToUpper(uuid.New().String()), "-", "", -1)
	operationControlNumber := "q3rew" + strconv.Itoa(rand.Intn(9999-1000)+1000) + "eu"

	response := responseTemplateA02
	response = strings.Replace(response, "fromId", fromId, -1)
	response = strings.Replace(response, "toId", toId, -1)
	response = strings.Replace(response, "bizMsgIdr", bizMsgIdr, -1)
	response = strings.Replace(response, "time1", creDt, -1)
	response = strings.Replace(response, "time2", creDtTm, -1)
	response = strings.Replace(response, "ident", msgId, -1)
	response = strings.Replace(response, "cretrID", cretrID, -1)
	response = strings.Replace(response, "prvtID", prvtID, -1)
	response = strings.Replace(response, "transactionNumber", transactionID, -1)
	response = strings.Replace(response, "uuid", uuid, -1)
	response = strings.Replace(response, "operationControlNumber", operationControlNumber, -1)

	return response
}

type A02 struct {
	XMLName xml.Name `xml:"IPSEnvelope"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	AppHdr  struct {
		Xmlns string `xml:"xmlns,attr"`
		Fr    struct {
			FIId struct {
				FinInstnId struct {
					Othr struct {
						ID string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To struct {
			FIId struct {
				FinInstnId struct {
					Othr struct {
						ID string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMsgIdr string `xml:"BizMsgIdr"`
		MsgDefIdr string `xml:"MsgDefIdr"`
		BizSvc    string `xml:"BizSvc"`
		CreDt     string `xml:"CreDt"`
		Rltd      struct {
			Fr struct {
				FIId struct {
					FinInstnId struct {
						Othr struct {
							ID string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"FIId"`
			} `xml:"Fr"`
			To struct {
				FIId struct {
					FinInstnId struct {
						Othr struct {
							ID string `xml:"Id"`
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
		Xmlns       string `xml:"xmlns,attr"`
		IdVrfctnRpt struct {
			Assgnmt struct {
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
				Cretr   struct {
					Pty struct {
						ID struct {
							PrvtId struct {
								Othr struct {
									ID      string `xml:"Id"`
									SchmeNm struct {
										Prtry string `xml:"Prtry"`
									} `xml:"SchmeNm"`
								} `xml:"Othr"`
							} `xml:"PrvtId"`
						} `xml:"Id"`
					} `xml:"Pty"`
				} `xml:"Cretr"`
				Assgnr struct {
					Agt struct {
						FinInstnId struct {
							Othr struct {
								ID string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"Agt"`
				} `xml:"Assgnr"`
				Assgne struct {
					Agt struct {
						FinInstnId struct {
							Othr struct {
								ID string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"Agt"`
				} `xml:"Assgne"`
			} `xml:"Assgnmt"`
			OrgnlAssgnmt struct {
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
			} `xml:"OrgnlAssgnmt"`
			Rpt struct {
				OrgnlId string `xml:"OrgnlId"`
				Vrfctn  string `xml:"Vrfctn"`
				RSN     struct {
					Prtry string `xml:"Prtry"`
				} `xml:"Rsn"`
				OrgnlPtyAndAcctId struct {
					Pty struct {
						ID struct {
							PrvtId struct {
								Othr struct {
									ID      string `xml:"Id"`
									SchmeNm struct {
										Prtry string `xml:"Prtry"`
									} `xml:"SchmeNm"`
								} `xml:"Othr"`
							} `xml:"PrvtId"`
						} `xml:"Id"`
					} `xml:"Pty"`
				} `xml:"OrgnlPtyAndAcctId"`
			} `xml:"Rpt"`
			SplmtryData struct {
				Envlp struct {
					IpsDt struct {
						Document string `xml:"document,attr"`
						TrnTp    string `xml:"TrnTp"`
						TrnCv    string `xml:"TrnCv"`
						OperDate string `xml:"OperDate"`
					} `xml:"IpsDt"`
				} `xml:"Envlp"`
			} `xml:"SplmtryData"`
		} `xml:"IdVrfctnRpt"`
	} `xml:"Document"`
}
