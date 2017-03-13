// Automatically generated by MockGen. DO NOT EDIT!
// Source: pkg/kubelego_const/interfaces.go

package mocks

import (
	logrus "github.com/Sirupsen/logrus"
	gomock "github.com/golang/mock/gomock"
	. "github.com/jetstack/kube-lego/pkg/kubelego_const"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/pkg/api/v1"
	v1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	net "net"
	time "time"
)

// Mock of KubeLego interface
type MockKubeLego struct {
	ctrl     *gomock.Controller
	recorder *_MockKubeLegoRecorder
}

// Recorder for MockKubeLego (not exported)
type _MockKubeLegoRecorder struct {
	mock *MockKubeLego
}

func NewMockKubeLego(ctrl *gomock.Controller) *MockKubeLego {
	mock := &MockKubeLego{ctrl: ctrl}
	mock.recorder = &_MockKubeLegoRecorder{mock}
	return mock
}

func (_m *MockKubeLego) EXPECT() *_MockKubeLegoRecorder {
	return _m.recorder
}

func (_m *MockKubeLego) KubeClient() *kubernetes.Clientset {
	ret := _m.ctrl.Call(_m, "KubeClient")
	ret0, _ := ret[0].(*kubernetes.Clientset)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) KubeClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KubeClient")
}

func (_m *MockKubeLego) Log() *logrus.Entry {
	ret := _m.ctrl.Call(_m, "Log")
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) Log() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log")
}

func (_m *MockKubeLego) AcmeClient() Acme {
	ret := _m.ctrl.Call(_m, "AcmeClient")
	ret0, _ := ret[0].(Acme)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) AcmeClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcmeClient")
}

func (_m *MockKubeLego) LegoHTTPPort() intstr.IntOrString {
	ret := _m.ctrl.Call(_m, "LegoHTTPPort")
	ret0, _ := ret[0].(intstr.IntOrString)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoHTTPPort() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoHTTPPort")
}

func (_m *MockKubeLego) LegoEmail() string {
	ret := _m.ctrl.Call(_m, "LegoEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoEmail() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoEmail")
}

func (_m *MockKubeLego) LegoURL() string {
	ret := _m.ctrl.Call(_m, "LegoURL")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoURL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoURL")
}

func (_m *MockKubeLego) LegoNamespace() string {
	ret := _m.ctrl.Call(_m, "LegoNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoNamespace")
}

func (_m *MockKubeLego) LegoIngressNameNginx() string {
	ret := _m.ctrl.Call(_m, "LegoIngressNameNginx")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoIngressNameNginx() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoIngressNameNginx")
}

func (_m *MockKubeLego) LegoServiceNameNginx() string {
	ret := _m.ctrl.Call(_m, "LegoServiceNameNginx")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoServiceNameNginx() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoServiceNameNginx")
}

func (_m *MockKubeLego) LegoServiceNameGce() string {
	ret := _m.ctrl.Call(_m, "LegoServiceNameGce")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoServiceNameGce() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoServiceNameGce")
}

func (_m *MockKubeLego) LegoDefaultIngressClass() string {
	ret := _m.ctrl.Call(_m, "LegoDefaultIngressClass")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoDefaultIngressClass() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoDefaultIngressClass")
}

func (_m *MockKubeLego) LegoSupportedIngressClass() []string {
	ret := _m.ctrl.Call(_m, "LegoSupportedIngressClass")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoSupportedIngressClass() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoSupportedIngressClass")
}

func (_m *MockKubeLego) LegoCheckInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "LegoCheckInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoCheckInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoCheckInterval")
}

func (_m *MockKubeLego) LegoMinimumValidity() time.Duration {
	ret := _m.ctrl.Call(_m, "LegoMinimumValidity")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoMinimumValidity() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoMinimumValidity")
}

func (_m *MockKubeLego) LegoPodIP() net.IP {
	ret := _m.ctrl.Call(_m, "LegoPodIP")
	ret0, _ := ret[0].(net.IP)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoPodIP() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoPodIP")
}

func (_m *MockKubeLego) IngressProvider(_param0 string) (IngressProvider, error) {
	ret := _m.ctrl.Call(_m, "IngressProvider", _param0)
	ret0, _ := ret[0].(IngressProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKubeLegoRecorder) IngressProvider(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IngressProvider", arg0)
}

func (_m *MockKubeLego) Version() string {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockKubeLego) AcmeUser() (map[string][]byte, error) {
	ret := _m.ctrl.Call(_m, "AcmeUser")
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKubeLegoRecorder) AcmeUser() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcmeUser")
}

func (_m *MockKubeLego) SaveAcmeUser(_param0 map[string][]byte) error {
	ret := _m.ctrl.Call(_m, "SaveAcmeUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) SaveAcmeUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAcmeUser", arg0)
}

// Mock of Acme interface
type MockAcme struct {
	ctrl     *gomock.Controller
	recorder *_MockAcmeRecorder
}

// Recorder for MockAcme (not exported)
type _MockAcmeRecorder struct {
	mock *MockAcme
}

func NewMockAcme(ctrl *gomock.Controller) *MockAcme {
	mock := &MockAcme{ctrl: ctrl}
	mock.recorder = &_MockAcmeRecorder{mock}
	return mock
}

func (_m *MockAcme) EXPECT() *_MockAcmeRecorder {
	return _m.recorder
}

func (_m *MockAcme) ObtainCertificate(domains []string) (map[string][]byte, error) {
	ret := _m.ctrl.Call(_m, "ObtainCertificate", domains)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAcmeRecorder) ObtainCertificate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ObtainCertificate", arg0)
}

// Mock of Tls interface
type MockTls struct {
	ctrl     *gomock.Controller
	recorder *_MockTlsRecorder
}

// Recorder for MockTls (not exported)
type _MockTlsRecorder struct {
	mock *MockTls
}

func NewMockTls(ctrl *gomock.Controller) *MockTls {
	mock := &MockTls{ctrl: ctrl}
	mock.recorder = &_MockTlsRecorder{mock}
	return mock
}

func (_m *MockTls) EXPECT() *_MockTlsRecorder {
	return _m.recorder
}

func (_m *MockTls) Hosts() []string {
	ret := _m.ctrl.Call(_m, "Hosts")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockTlsRecorder) Hosts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hosts")
}

func (_m *MockTls) SecretMetadata() *v1.ObjectMeta {
	ret := _m.ctrl.Call(_m, "SecretMetadata")
	ret0, _ := ret[0].(*v1.ObjectMeta)
	return ret0
}

func (_mr *_MockTlsRecorder) SecretMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecretMetadata")
}

func (_m *MockTls) IngressMetadata() *v1.ObjectMeta {
	ret := _m.ctrl.Call(_m, "IngressMetadata")
	ret0, _ := ret[0].(*v1.ObjectMeta)
	return ret0
}

func (_mr *_MockTlsRecorder) IngressMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IngressMetadata")
}

func (_m *MockTls) Process() error {
	ret := _m.ctrl.Call(_m, "Process")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTlsRecorder) Process() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Process")
}

// Mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceRecorder
}

// Recorder for MockService (not exported)
type _MockServiceRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &_MockServiceRecorder{mock}
	return mock
}

func (_m *MockService) EXPECT() *_MockServiceRecorder {
	return _m.recorder
}

func (_m *MockService) Object() *v1.Service {
	ret := _m.ctrl.Call(_m, "Object")
	ret0, _ := ret[0].(*v1.Service)
	return ret0
}

func (_mr *_MockServiceRecorder) Object() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Object")
}

func (_m *MockService) SetKubeLegoSpec() {
	_m.ctrl.Call(_m, "SetKubeLegoSpec")
}

func (_mr *_MockServiceRecorder) SetKubeLegoSpec() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetKubeLegoSpec")
}

func (_m *MockService) SetEndpoints(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetEndpoints", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) SetEndpoints(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEndpoints", arg0)
}

func (_m *MockService) Save() error {
	ret := _m.ctrl.Call(_m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) Save() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save")
}

func (_m *MockService) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

// Mock of Secret interface
type MockSecret struct {
	ctrl     *gomock.Controller
	recorder *_MockSecretRecorder
}

// Recorder for MockSecret (not exported)
type _MockSecretRecorder struct {
	mock *MockSecret
}

func NewMockSecret(ctrl *gomock.Controller) *MockSecret {
	mock := &MockSecret{ctrl: ctrl}
	mock.recorder = &_MockSecretRecorder{mock}
	return mock
}

func (_m *MockSecret) EXPECT() *_MockSecretRecorder {
	return _m.recorder
}

func (_m *MockSecret) Object() *v1.Secret {
	ret := _m.ctrl.Call(_m, "Object")
	ret0, _ := ret[0].(*v1.Secret)
	return ret0
}

func (_mr *_MockSecretRecorder) Object() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Object")
}

func (_m *MockSecret) KubeLego() KubeLego {
	ret := _m.ctrl.Call(_m, "KubeLego")
	ret0, _ := ret[0].(KubeLego)
	return ret0
}

func (_mr *_MockSecretRecorder) KubeLego() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KubeLego")
}

func (_m *MockSecret) Exists() bool {
	ret := _m.ctrl.Call(_m, "Exists")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockSecretRecorder) Exists() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exists")
}

func (_m *MockSecret) Save() error {
	ret := _m.ctrl.Call(_m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSecretRecorder) Save() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save")
}

func (_m *MockSecret) TlsDomains() ([]string, error) {
	ret := _m.ctrl.Call(_m, "TlsDomains")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSecretRecorder) TlsDomains() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TlsDomains")
}

func (_m *MockSecret) TlsDomainsInclude(domains []string) bool {
	ret := _m.ctrl.Call(_m, "TlsDomainsInclude", domains)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockSecretRecorder) TlsDomainsInclude(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TlsDomainsInclude", arg0)
}

func (_m *MockSecret) TlsExpireTime() (time.Time, error) {
	ret := _m.ctrl.Call(_m, "TlsExpireTime")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSecretRecorder) TlsExpireTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TlsExpireTime")
}

// Mock of Ingress interface
type MockIngress struct {
	ctrl     *gomock.Controller
	recorder *_MockIngressRecorder
}

// Recorder for MockIngress (not exported)
type _MockIngressRecorder struct {
	mock *MockIngress
}

func NewMockIngress(ctrl *gomock.Controller) *MockIngress {
	mock := &MockIngress{ctrl: ctrl}
	mock.recorder = &_MockIngressRecorder{mock}
	return mock
}

func (_m *MockIngress) EXPECT() *_MockIngressRecorder {
	return _m.recorder
}

func (_m *MockIngress) Object() *v1beta1.Ingress {
	ret := _m.ctrl.Call(_m, "Object")
	ret0, _ := ret[0].(*v1beta1.Ingress)
	return ret0
}

func (_mr *_MockIngressRecorder) Object() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Object")
}

func (_m *MockIngress) KubeLego() KubeLego {
	ret := _m.ctrl.Call(_m, "KubeLego")
	ret0, _ := ret[0].(KubeLego)
	return ret0
}

func (_mr *_MockIngressRecorder) KubeLego() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KubeLego")
}

func (_m *MockIngress) Log() *logrus.Entry {
	ret := _m.ctrl.Call(_m, "Log")
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockIngressRecorder) Log() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log")
}

func (_m *MockIngress) Save() error {
	ret := _m.ctrl.Call(_m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIngressRecorder) Save() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save")
}

func (_m *MockIngress) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIngressRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockIngress) IngressClass() string {
	ret := _m.ctrl.Call(_m, "IngressClass")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIngressRecorder) IngressClass() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IngressClass")
}

func (_m *MockIngress) Tls() []Tls {
	ret := _m.ctrl.Call(_m, "Tls")
	ret0, _ := ret[0].([]Tls)
	return ret0
}

func (_mr *_MockIngressRecorder) Tls() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tls")
}

func (_m *MockIngress) Ignore() bool {
	ret := _m.ctrl.Call(_m, "Ignore")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockIngressRecorder) Ignore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ignore")
}

// Mock of IngressProvider interface
type MockIngressProvider struct {
	ctrl     *gomock.Controller
	recorder *_MockIngressProviderRecorder
}

// Recorder for MockIngressProvider (not exported)
type _MockIngressProviderRecorder struct {
	mock *MockIngressProvider
}

func NewMockIngressProvider(ctrl *gomock.Controller) *MockIngressProvider {
	mock := &MockIngressProvider{ctrl: ctrl}
	mock.recorder = &_MockIngressProviderRecorder{mock}
	return mock
}

func (_m *MockIngressProvider) EXPECT() *_MockIngressProviderRecorder {
	return _m.recorder
}

func (_m *MockIngressProvider) Log() *logrus.Entry {
	ret := _m.ctrl.Call(_m, "Log")
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockIngressProviderRecorder) Log() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log")
}

func (_m *MockIngressProvider) Process(_param0 Ingress) error {
	ret := _m.ctrl.Call(_m, "Process", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIngressProviderRecorder) Process(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Process", arg0)
}

func (_m *MockIngressProvider) Reset() error {
	ret := _m.ctrl.Call(_m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIngressProviderRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockIngressProvider) Finalize() error {
	ret := _m.ctrl.Call(_m, "Finalize")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIngressProviderRecorder) Finalize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Finalize")
}
