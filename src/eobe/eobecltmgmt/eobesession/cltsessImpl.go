package eobesession

import (
	"fmt"
	"net/http"
)

type ClientSessionImpl struct {
	clientData
}

//Interface implemetations for
func (csi *ClientSessionImpl) GetClientInfo() (sid string, user string) {
	return csi.sidStr(), csi.userName()
}

func (csi *ClientSessionImpl) GetAllCookies() map[string]*http.Cookie {
	return csi.copyAllCookies()
}

func (csi *ClientSessionImpl) GetBindData() BindDataInf {
	return csi.bd
}

func (csi *ClientSessionImpl) SetBindData(bd BindDataInf) error {
	if bd == nil {
		return fmt.Errorf(cStrErrorNULLBINDDATA)
	}

	csi.bd = bd
	return nil
}
