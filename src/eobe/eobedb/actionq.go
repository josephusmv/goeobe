package eobedb

type queItem struct {
	qryDefn QueryDefn
	exData  exchgnData
}

type dbActionQueue struct {
	itemQueue        []queItem
	chanDBActionDone chan bool
	dbWroker         databaseWorker
}

func (q *dbActionQueue) addToQueue(qryDefn QueryDefn, exData exchgnData) {
	q.itemQueue = append(q.itemQueue, queItem{qryDefn, exData})
	//fmt.Printf("addToQueue: Force to print for DEV: Queue current length = %d\n", len(q.itemQueue))
}

func (q *dbActionQueue) processNextOne() (hasNewRoutine bool) {
	//fmt.Printf("processNextOne: Force to print for DEV: Queue current length = %d\n", len(q.itemQueue))
	if len(q.itemQueue) <= 0 {
		return false
		//This is a test for version sync on MAC!!
	}

	item := queItem{qryDefn: q.itemQueue[0].qryDefn, exData: q.itemQueue[0].exData}
	q.itemQueue = append(q.itemQueue[:0], q.itemQueue[1:]...)

	//fmt.Printf("processNextOne: Force to print for DEV: Queue current length = %d\n", len(q.itemQueue))
	go q.dbWroker.doDBAction(item.qryDefn, item.exData, q.chanDBActionDone)
	return true
}
