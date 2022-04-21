module kube_web

go 1.16

require github.com/beego/beego/v2 v2.0.2

require (
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/getsentry/raven-go v0.2.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/prometheus/common v0.28.0
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/spf13/cobra v1.2.1
	github.com/streadway/amqp v1.0.0
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/igm/sockjs-go.v2 v2.1.0
	gopkg.in/ldap.v2 v2.5.1
	k8s.io/api v0.23.5
	k8s.io/apiextensions-apiserver v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	k8s.io/kubectl v0.22.3
)
