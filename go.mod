module github.com/sapcc/kubernikus

require (
	github.com/BurntSushi/toml v0.0.0-20170626110600-a368813c5e64
	github.com/Masterminds/goutils v1.0.1
	github.com/Masterminds/semver v1.3.1 // indirect
	github.com/Masterminds/sprig v2.14.1+incompatible
	github.com/PuerkitoBio/purell v1.1.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/ajeddeloh/yaml v0.0.0-20160722214022-1072abfea311 // indirect
	github.com/alecthomas/units v0.0.0-20150109002421-6b4e7dc5e314 // indirect
	github.com/aokoli/goutils v1.0.1
	github.com/asaskevich/govalidator v0.0.0-20170903095215-73945b6115bf
	github.com/beorn7/perks v0.0.0-20160804104726-4c0e84591b9a // indirect
	github.com/cenkalti/backoff v0.0.0-20170918082838-80e08cb804a3 // indirect
	github.com/coreos/container-linux-config-transpiler v0.4.2
	github.com/coreos/go-semver v0.0.0-20170209201757-5e3acbb5668c // indirect
	github.com/coreos/go-systemd v0.0.0-20161114122254-48702e0da86b // indirect
	github.com/coreos/ignition v0.17.2
	github.com/danieljoos/wincred v1.0.1 // indirect
	github.com/databus23/goslo.policy v0.0.0-20170317131957-3ae74dd07ebf
	github.com/databus23/guttle v0.0.0-20180831113238-290306d52eb2
	github.com/databus23/keystone v0.0.0-20180111110916-350fd0e663cd
	github.com/databus23/requestutil v0.0.0-20160108082528-5ff8e981f38f
	github.com/davecgh/go-spew v0.0.0-20171005155431-ecdeabc65495 // indirect
	github.com/docker/go-units v0.0.0-20171221200356-d59758554a3d // indirect
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/emicklei/go-restful v0.0.0-20170410110728-ff4f55a20633 // indirect
	github.com/ghodss/yaml v0.0.0-20150909031657-73d445a93680
	github.com/go-kit/kit v0.6.0
	github.com/go-logfmt/logfmt v0.3.0 // indirect
	github.com/go-openapi/analysis v0.0.0-20170813233457-8ed83f2ea9f0 // indirect
	github.com/go-openapi/errors v0.0.0-20170426151106-03cfca65330d
	github.com/go-openapi/jsonpointer v0.0.0-20170102174223-779f45308c19 // indirect
	github.com/go-openapi/jsonreference v0.0.0-20161105162150-36d33bfe519e // indirect
	github.com/go-openapi/loads v0.0.0-20171020172901-c3e1ca4c0b61
	github.com/go-openapi/runtime v0.0.0-20171109160715-2da70401d8fa
	github.com/go-openapi/spec v0.0.0-20171105074921-a4fa9574c7aa
	github.com/go-openapi/strfmt v0.0.0-20170822153411-610b6cacdcde
	github.com/go-openapi/swag v0.0.0-20171111214437-cf0bdb963811
	github.com/go-openapi/validate v0.0.0-20171109223028-a52193aca9d5
	github.com/go-stack/stack v1.7.0
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20151105175453-c7fdd8b5cd55 // indirect
	github.com/gogo/protobuf v0.0.0-20170330071051-c0656edd0d9e // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20160516000752-02826c3e7903 // indirect
	github.com/golang/protobuf v0.0.0-20171021043952-1643683e1b54 // indirect
	github.com/google/btree v0.0.0-20160524151835-7d79101e329e // indirect
	github.com/google/gofuzz v0.0.0-20161122191042-44d81051d367 // indirect
	github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d // indirect
	github.com/gophercloud/gophercloud v0.0.0-20180623033157-1ad7a96d69a2
	github.com/gregjones/httpcache v0.0.0-20170728041850-787624de3eb7 // indirect
	github.com/hashicorp/golang-lru v0.0.0-20160207214719-a0d98a5f2880 // indirect
	github.com/hashicorp/yamux v0.0.0-20180826203732-cc6d2ea263b2 // indirect
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/huandu/xstrings v0.0.0-20171208101919-37469d0c81a7 // indirect
	github.com/imdario/mergo v0.0.0-20160216103600-3e95a51e0639
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/joeshaw/envdecode v0.0.0-20170502020559-6326cbed175e
	github.com/json-iterator/go v0.0.0-20170829155851-36b14963da70 // indirect
	github.com/juju/ratelimit v0.0.0-20170523012141-5b9ff8664717 // indirect
	github.com/justinas/alice v0.0.0-20171023064455-03f45bd4b7da
	github.com/kennygrant/sanitize v1.2.3
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515 // indirect
	github.com/mailru/easyjson v0.0.0-20171106100207-5f62e4f3afa2 // indirect
	github.com/matttproud/golang_protobuf_extensions v0.0.0-20150406173934-fc2b8d3a73c4 // indirect
	github.com/mitchellh/mapstructure v0.0.0-20171017171808-06020f85339e // indirect
	github.com/oklog/run v0.0.0-20180308005104-6934b124db28
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pmylund/go-cache v2.1.0+incompatible // indirect
	github.com/prometheus/client_golang v0.0.0-20180215121106-fcc130e101e7
	github.com/prometheus/client_model v0.0.0-20171117100541-99fa1f4be8e5
	github.com/prometheus/common v0.0.0-20170427095455-13ba4ddd0caa
	github.com/prometheus/procfs v0.0.0-20170519190837-65c1f6f8f0fc // indirect
	github.com/rs/cors v0.0.0-20170801073201-eabcc6af4bbe
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.0-20160722081547-f62e98d28ab7
	github.com/spf13/pflag v0.0.0-20171106142849-4c012f6dcd95
	github.com/stretchr/objx v0.0.0-20180129172003-8a3f7159479f // indirect
	github.com/stretchr/testify v1.2.2
	github.com/tredoe/osutil v0.0.0-20161130133508-7d3ee1afa71c
	github.com/tylerb/graceful v0.0.0-20170221171003-d72b0151351a
	github.com/vincent-petithory/dataurl v0.0.0-20160330182126-9a301d65acbb // indirect
	github.com/zalando/go-keyring v0.0.0-20180221093347-6d81c293b3fb
	golang.org/x/crypto v0.0.0-20170825220121-81e90905daef // indirect
	golang.org/x/net v0.0.0-20170809000501-1c05540f6879
	golang.org/x/sys v0.0.0-20171031081856-95c657629925
	golang.org/x/text v0.0.0-20170810154203-b19bf474d317 // indirect
	google.golang.org/genproto v0.0.0-20170731182057-09f6ed296fc6 // indirect
	google.golang.org/grpc v1.3.0
	gopkg.in/inf.v0 v0.9.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20160818020120-3f83fa500528 // indirect
	gopkg.in/yaml.v2 v2.0.0-20170812160011-eb3733d160e7
	k8s.io/api v0.0.0-20171207041203-11147472b7c9
	k8s.io/apiextensions-apiserver v0.0.0-20180103022501-097c49c1906b
	k8s.io/apimachinery v0.0.0-20180103014849-68f9c3a1feb3
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/code-generator v0.0.0-20180103022856-91d3f6a57905
	k8s.io/gengo v0.0.0-20171109050133-b58fc7edb82e // indirect
	k8s.io/helm v2.7.2+incompatible
	k8s.io/kube-openapi v0.0.0-20171101183504-39a7bf85c140 // indirect
	k8s.io/utils v0.0.0-20171122000934-aedf551cdb8b
)

replace github.com/golang/glog => github.com/bugroger/glognomore v0.0.0-20180702093556-abbb4165c641
