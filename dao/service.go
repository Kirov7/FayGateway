package dao

import (
	"github.com/e421083458/FayGateway/dto"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sync"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
}

var ServiceMangerHandler *ServiceManager

//加载dao的时候,直接执行init,使ServiceMangerHandler初始化
func init() {
	ServiceMangerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap   map[string]*ServiceDetail
	ServiceSlice []*ServiceDetail
	Locker       sync.RWMutex
	init         sync.Once
	err          error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:   map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker:       sync.RWMutex{},
		init:         sync.Once{},
	}
}

func (s *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail, error) {
	//1. 前缀匹配 /abc ==> serviceSlice.rule
	//2. 域名 www.abc.com ==> serviceSlice.rule

	// host c.Request.Host
	// path c.Request.URL.Path
	return nil, nil
}

func (s *ServiceManager) LoadOnce() error {
	s.init.Do(func() {
		serviceInfo := &ServiceInfo{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		tx, err := lib.GetGormPool("default")
		if err != nil {
			s.err = err
			return
		}
		params := &dto.ServiceListInput{PageNo: 1, PageSize: 99999}
		list, _, err := serviceInfo.PageList(c, tx, params)
		s.Locker.Lock()
		defer s.Locker.Unlock()
		for _, listItem := range list {
			serviceDetail, err := listItem.ServiceDetail(c, tx, &listItem)
			if err != nil {
				s.err = err
				return
			}
			s.ServiceMap[listItem.ServiceName] = serviceDetail
			s.ServiceSlice = append(s.ServiceSlice, serviceDetail)
		}
	})
	return s.err
}
