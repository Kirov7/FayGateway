package dao

import (
	"fmt"
	"github.com/e421083458/FayGateway/public"
	"github.com/e421083458/FayGateway/reverse_proxy/load_balance"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type LoadBalance struct {
	ID            int64  `json:"id" gorm:"primary_key"`
	ServiceID     int64  `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	CheckMethod   int    `json:"check_method" gorm:"column:check_method" description:"检查方法 tcpchk=检测端口是否握手成功	"`
	CheckTimeout  int    `json:"check_timeout" gorm:"column:check_timeout" description:"check超时时间	"`
	CheckInterval int    `json:"check_interval" gorm:"column:check_interval" description:"检查间隔, 单位s		"`
	RoundType     int    `json:"round_type" gorm:"column:round_type" description:"轮询方式 round/weight_round/random/ip_hash"`
	IpList        string `json:"ip_list" gorm:"column:ip_list" description:"ip列表"`
	WeightList    string `json:"weight_list" gorm:"column:weight_list" description:"权重列表"`
	ForbidList    string `json:"forbid_list" gorm:"column:forbid_list" description:"禁用ip列表"`

	UpstreamConnectTimeout int `json:"upstream_connect_timeout" gorm:"column:upstream_connect_timeout" description:"下游建立连接超时, 单位s"`
	UpstreamHeaderTimeout  int `json:"upstream_header_timeout" gorm:"column:upstream_header_timeout" description:"下游获取header超时, 单位s	"`
	UpstreamIdleTimeout    int `json:"upstream_idle_timeout" gorm:"column:upstream_idle_timeout" description:"下游链接最大空闲时间, 单位s	"`
	UpstreamMaxIdle        int `json:"upstream_max_idle" gorm:"column:upstream_max_idle" description:"下游最大空闲链接数"`
}

func (t *LoadBalance) TableName() string {
	return "gateway_service_load_balance"
}

func (t *LoadBalance) Find(c *gin.Context, tx *gorm.DB, search *LoadBalance) (*LoadBalance, error) {
	model := &LoadBalance{}
	err := tx.WithContext(c).Where(search).Find(model).Error
	return model, err
}

func (t *LoadBalance) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.WithContext(c).Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (t LoadBalance) GetIPListByModel() []string {
	return strings.Split(t.IpList, ",")
}

func (t LoadBalance) GetWeightListByModel() []string {
	return strings.Split(t.WeightList, ",")
}

var LoadBalancerHandler *LoadBalancer

type LoadBalancer struct {
	LoadBalanceMap   map[string]*LoadBalancerItem
	LoadBalanceSlice []*LoadBalancerItem
	Locker           sync.RWMutex
}

type LoadBalancerItem struct {
	LoadBalance load_balance.LoadBalance
	ServiceName string
	Locker      sync.RWMutex
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		LoadBalanceMap:   map[string]*LoadBalancerItem{},
		LoadBalanceSlice: []*LoadBalancerItem{},
		Locker:           sync.RWMutex{},
	}
}

func init() {
	LoadBalancerHandler = NewLoadBalancer()
}

func (lbr *LoadBalancer) GetLoadBalancer(service *ServiceDetail) (load_balance.LoadBalance, error) {
	for _, lbrItem := range lbr.LoadBalanceSlice {
		if lbrItem.ServiceName == service.Info.ServiceName {
			return lbrItem.LoadBalance, nil
		}
	}
	schema := "http"
	if service.HTTPRule.NeedHttps == 1 {
		schema = "https"
	}
	prefix := ""
	if service.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL {
		prefix = service.HTTPRule.Rule
	}
	ipList := service.LoadBalance.GetIPListByModel()
	weightList := service.LoadBalance.GetWeightListByModel()
	ipConf := map[string]string{}
	for ipIndex, ipItem := range ipList {
		ipConf[ipItem] = weightList[ipIndex]
	}
	mConf, err := load_balance.NewLoadBalanceCheckConf(
		fmt.Sprintf("%s://%s%s", schema, "%s", prefix), ipConf)
	if err != nil {
		return nil, err
	}
	lb := load_balance.LoadBanlanceFactorWithConf(load_balance.LbType(service.LoadBalance.RoundType), mConf)

	//save to map and slice
	lbItem := &LoadBalancerItem{
		LoadBalance: lb,
		ServiceName: service.Info.ServiceName,
	}
	lbr.LoadBalanceSlice = append(lbr.LoadBalanceSlice, lbItem)

	//加入服务多的时候使用map可以优化加速
	lbr.Locker.Lock()
	defer lbr.Locker.Unlock()
	lbr.LoadBalanceMap[service.Info.ServiceName] = lbItem
	return lb, nil
}

var TransporterHandler *Transporter

type Transporter struct {
	TransportMap   map[string]*TransportItem
	TransportSlice []*TransportItem
	Locker         sync.RWMutex
}

type TransportItem struct {
	Trans       *http.Transport
	ServiceName string
}

func NewTransporter() *Transporter {
	return &Transporter{
		TransportMap:   map[string]*TransportItem{},
		TransportSlice: []*TransportItem{},
		Locker:         sync.RWMutex{},
	}
}

func init() {
	TransporterHandler = NewTransporter()
}

func (t *Transporter) GetTrans(service *ServiceDetail) (*http.Transport, error) {
	for _, transItem := range t.TransportSlice {
		if transItem.ServiceName == service.Info.ServiceName {
			return transItem.Trans, nil
		}
	}
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: time.Duration(service.LoadBalance.UpstreamConnectTimeout) * time.Second, //连接超时
		}).DialContext,
		MaxIdleConns:          service.LoadBalance.UpstreamMaxIdle,                                    //最大空闲连接
		IdleConnTimeout:       time.Duration(service.LoadBalance.UpstreamIdleTimeout) * time.Second,   //空闲超时时间
		ResponseHeaderTimeout: time.Duration(service.LoadBalance.UpstreamHeaderTimeout) * time.Second, //100-continue超时时间
	}
	transItem := &TransportItem{
		Trans:       trans,
		ServiceName: service.Info.ServiceName,
	}
	t.TransportSlice = append(t.TransportSlice, transItem)

	//加入服务多的时候使用map可以优化加速
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.TransportMap[service.Info.ServiceName] = transItem
	return trans, nil
}
