package topsql

import (
	"net/http"

	"github.com/pingcap/ng-monitoring/component/subscriber"
	"github.com/pingcap/ng-monitoring/component/topology"
	"github.com/pingcap/ng-monitoring/component/topsql/query"
	"github.com/pingcap/ng-monitoring/component/topsql/service"
	"github.com/pingcap/ng-monitoring/component/topsql/store"
	sub "github.com/pingcap/ng-monitoring/component/topsql/subscriber"
	"github.com/pingcap/ng-monitoring/config"
	"github.com/pingcap/ng-monitoring/config/pdvariable"
	"github.com/pingcap/ng-monitoring/database/docdb"

	"github.com/gin-gonic/gin"
)

var (
	defStore      store.Store
	defQuery      query.Query
	defSubscriber *subscriber.Subscriber
	defService    *service.Service
)

func Init(
	cfgSub config.Subscriber,
	docDB docdb.DocDB,
	insertHdr, selectHdr http.HandlerFunc,
	topSub topology.Subscriber,
	varSub pdvariable.Subscriber,
	metaRetentionSecs int64,
) (err error) {
	defStore = store.NewDefaultStore(insertHdr, docDB, metaRetentionSecs)
	defQuery = query.NewDefaultQuery(selectHdr, docDB)
	defSubscriber = sub.NewSubscriber(topSub, varSub, cfgSub, defStore)
	defService = service.NewService(defQuery)
	return nil
}

func HTTPService(g *gin.RouterGroup) {
	defService.HTTPService(g)
}

func Stop() {
	defSubscriber.Close()
	defQuery.Close()
	defStore.Close()
}
