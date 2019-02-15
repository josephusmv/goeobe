package httpsample1

import "eobe/eobehttp"

//RespFetcherImpl ...
type RespFetcherImpl struct {
}

//FetchResponse As an sample, we only return the index.html
func (rf *RespFetcherImpl) FetchResponse(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	var tp eobehttp.TemplateData
	tp.Rows = [][]string{
		[]string{"Value1111", "Value1112", "Vao & bar斯卡拉lue1113"},
		[]string{"V卡alue1222", "Value2222", "Value22卡拉23"},
		[]string{"Value3卡拉331", "Value3卡332", "Value卡拉3333"},
	}

	rsp.ContentType = "text/html"
	rsp.HTMLTmpltName = "index.tmplt"
	rsp.HTMLTmpltData = tp
	rsp.Body = nil
	rsp.CookieList = req.CookieList //just copy paste for index.
	return
}
