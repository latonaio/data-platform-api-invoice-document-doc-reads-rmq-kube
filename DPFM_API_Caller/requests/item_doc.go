package requests

type ItemDoc struct {
	InvoiceDocument		 	 int	`json:"InvoiceDocument"`
	InvoiceDocumentItem      int	`json:"InvoiceDocumentItem"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
