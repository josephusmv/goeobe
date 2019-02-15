/*Package eobecltmgmt  Manage the user info. Including:
 *		1. Cookie managment
 *		2. client IP managment
 *		3. Login User managment
 *		4. User singin/signout/signup
 *		5. Directly access eobedb package for query/add/delete/update	?
 *		6. Directly access eobedb package for user permission query	?
 *		7. Directly access eobehttp	?
 */
package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
)

const CStrKeyWordSessionID = eobesession.CStrKeyWordSessionID
const CStrKeyWordCurrentUser = eobesession.CStrKeyWordCurrentUser
const CStrKeyWordClientID = eobesession.CStrKeyWordClientID

//ClientManager Open manager interface for Client Management
//	This manager will be called by different goroutines.
//	Commands wil be sent and handled serializly
type ClientManager struct {
	worker
}

func NewClientManager(logPath string) *ClientManager {
	var cltmgr ClientManager
	cltmgr.wlogger = pckgLogger{logPath: logPath}
	cltmgr.wlogger.initLogger()
	return &cltmgr
}

func (cm *ClientManager) StartClientManagerServer() {
	cm.init()
}

func (cm *ClientManager) StopClientManagerServer() {
	cm.sendQuit()
}

//NewHerald For each different request handler instance, they should share one ClientManager.
//		When a new request comes, the ServerHTTP handler should get their own Herald from shared ClientManager
//		And any command should be send through Herald to ClientManager, rather than throguh ClientManager directly.
func (cm *ClientManager) NewHerald() *Herald {
	return &Herald{cm: cm}
}
