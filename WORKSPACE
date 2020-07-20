load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name="io_bazel_rules_go",
    sha256="2d536797707dd1697441876b2e862c58839f975c8fc2f0f96636cbd428f45866",
    urls=[
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.5/rules_go-v0.23.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.5/rules_go-v0.23.5.tar.gz",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
)

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name="bazel_gazelle",
    sha256="cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls=[
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# For claat.
# ggazelle:exclude WORKSPACE
# go_repository(
#    name="com_github_googlecodelabs_tools",
#    commit="286d7d8fa1f527fea803c819ca6e49c77a6b32db",
#    importpath="github.com/googlecodelabs/tools",
#    build_file_generation="on",
# )

git_repository(
    name="com_github_googlecodelabs_tools",
    branch="fixing-bazelbuild",
    remote="https://github.com/fabstu/tools",
)

# local_repository(
#    name="com_github_googlecodelabs_tools",
#    path="/Users/fabiansturm/Documents/projects/intprojects/claat",
# )

go_rules_dependencies()

# Auto-managed repos.

go_repository(
    name="co_honnef_go_tools",
    importpath="honnef.co/go/tools",
    sum="h1:XJP7lxbSxWLOMNdBE4B/STaqVy6L73o0knwj2vIlxnw=",
    version="v0.0.0-20190102054323-c2f93a96b099",
)

go_repository(
    name="com_github_alecthomas_template",
    importpath="github.com/alecthomas/template",
    sum="h1:cAKDfWh5VpdgMhJosfJnn5/FoN2SRZ4p7fJNX58YPaU=",
    version="v0.0.0-20160405071501-a0175ee3bccc",
)

go_repository(
    name="com_github_alecthomas_units",
    importpath="github.com/alecthomas/units",
    sum="h1:qet1QNfXsQxTZqLG4oE62mJzwPIB8+Tee4RNCL9ulrY=",
    version="v0.0.0-20151022065526-2efee857e7cf",
)

go_repository(
    name="com_github_armon_consul_api",
    importpath="github.com/armon/consul-api",
    sum="h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version="v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name="com_github_beorn7_perks",
    importpath="github.com/beorn7/perks",
    sum="h1:HWo1m869IqiPhD389kmkxeTalrjNbbJTC8LXupb+sl0=",
    version="v1.0.0",
)

go_repository(
    name="com_github_burntsushi_toml",
    importpath="github.com/BurntSushi/toml",
    sum="h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version="v0.3.1",
)

go_repository(
    name="com_github_cespare_xxhash",
    importpath="github.com/cespare/xxhash",
    sum="h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version="v1.1.0",
)

go_repository(
    name="com_github_client9_misspell",
    importpath="github.com/client9/misspell",
    sum="h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version="v0.3.4",
)

go_repository(
    name="com_github_coreos_bbolt",
    importpath="github.com/coreos/bbolt",
    sum="h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
    version="v1.3.2",
)

go_repository(
    name="com_github_coreos_etcd",
    importpath="github.com/coreos/etcd",
    sum="h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
    version="v3.3.10+incompatible",
)

go_repository(
    name="com_github_coreos_go_semver",
    importpath="github.com/coreos/go-semver",
    sum="h1:3Jm3tLmsgAYcjC+4Up7hJrFBPr+n7rAqYeSw/SZazuY=",
    version="v0.2.0",
)

go_repository(
    name="com_github_coreos_go_systemd",
    importpath="github.com/coreos/go-systemd",
    sum="h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
    version="v0.0.0-20190321100706-95778dfbb74e",
)

go_repository(
    name="com_github_coreos_pkg",
    importpath="github.com/coreos/pkg",
    sum="h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version="v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name="com_github_cpuguy83_go_md2man_v2",
    importpath="github.com/cpuguy83/go-md2man/v2",
    sum="h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
    version="v2.0.0",
)

go_repository(
    name="com_github_davecgh_go_spew",
    importpath="github.com/davecgh/go-spew",
    sum="h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version="v1.1.1",
)

go_repository(
    name="com_github_desertbit_timer",
    importpath="github.com/desertbit/timer",
    sum="h1:U5y3Y5UE0w7amNe7Z5G/twsBW0KEalRQXZzf8ufSh9I=",
    version="v0.0.0-20180107155436-c41aec40b27f",
)

go_repository(
    name="com_github_dgrijalva_jwt_go",
    importpath="github.com/dgrijalva/jwt-go",
    sum="h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version="v3.2.0+incompatible",
)

go_repository(
    name="com_github_dgryski_go_sip13",
    importpath="github.com/dgryski/go-sip13",
    sum="h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version="v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name="com_github_fsnotify_fsnotify",
    importpath="github.com/fsnotify/fsnotify",
    sum="h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    version="v1.4.7",
)

go_repository(
    name="com_github_ghodss_yaml",
    importpath="github.com/ghodss/yaml",
    sum="h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version="v1.0.0",
)

go_repository(
    name="com_github_go_kit_kit",
    importpath="github.com/go-kit/kit",
    sum="h1:Wz+5lgoB0kkuqLEc6NVmwRknTKP6dTGbSqvhZtBI/j0=",
    version="v0.8.0",
)

go_repository(
    name="com_github_go_logfmt_logfmt",
    importpath="github.com/go-logfmt/logfmt",
    sum="h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
    version="v0.4.0",
)

go_repository(
    name="com_github_go_logr_logr",
    importpath="github.com/go-logr/logr",
    sum="h1:QvGt2nLcHH0WK9orKa+ppBPAxREcH364nPUedEpK0TY=",
    version="v0.2.0",
)

go_repository(
    name="com_github_go_sql_driver_mysql",
    importpath="github.com/go-sql-driver/mysql",
    sum="h1:7LxgVwFb2hIQtMm87NdgAVfXjnt4OePseqT1tKx+opk=",
    version="v1.4.0",
)

go_repository(
    name="com_github_go_stack_stack",
    importpath="github.com/go-stack/stack",
    sum="h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version="v1.8.0",
)

go_repository(
    name="com_github_gogo_protobuf",
    importpath="github.com/gogo/protobuf",
    sum="h1:/s5zKNz0uPFCZ5hddgPdo2TK2TVrUNMn0OOX8/aZMTE=",
    version="v1.2.1",
)

go_repository(
    name="com_github_golang_glog",
    importpath="github.com/golang/glog",
    sum="h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version="v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name="com_github_golang_groupcache",
    importpath="github.com/golang/groupcache",
    sum="h1:veQD95Isof8w9/WXiA+pa3tz3fJXkt5B7QaRBrM62gk=",
    version="v0.0.0-20190129154638-5b532d6fd5ef",
)

go_repository(
    name="com_github_golang_mock",
    importpath="github.com/golang/mock",
    sum="h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version="v1.1.1",
)

go_repository(
    name="com_github_golang_protobuf",
    importpath="github.com/golang/protobuf",
    sum="h1:P3YflyNX/ehuJFLhxviNdFxQPkGK5cDcApsge1SqnvM=",
    version="v1.2.0",
)

go_repository(
    name="com_github_google_btree",
    importpath="github.com/google/btree",
    sum="h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version="v1.0.0",
)

go_repository(
    name="com_github_google_go_cmp",
    importpath="github.com/google/go-cmp",
    sum="h1:/QaMHBdZ26BB3SSst0Iwl10Epc+xhTquomWX0oZEB6w=",
    version="v0.5.0",
)

go_repository(
    name="com_github_google_uuid",
    importpath="github.com/google/uuid",
    sum="h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
    version="v1.1.1",
)

go_repository(
    name="com_github_gorilla_websocket",
    importpath="github.com/gorilla/websocket",
    sum="h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
    version="v1.4.0",
)

go_repository(
    name="com_github_grpc_ecosystem_go_grpc_middleware",
    importpath="github.com/grpc-ecosystem/go-grpc-middleware",
    sum="h1:Iju5GlWwrvL6UBg4zJJt3btmonfrMlCDdsejg4CZE7c=",
    version="v1.0.0",
)

go_repository(
    name="com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath="github.com/grpc-ecosystem/go-grpc-prometheus",
    sum="h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version="v1.2.0",
)

go_repository(
    name="com_github_grpc_ecosystem_grpc_gateway",
    importpath="github.com/grpc-ecosystem/grpc-gateway",
    sum="h1:bM6ZAFZmc/wPFaRDi0d5L7hGEZEx/2u+Tmr2evNHDiI=",
    version="v1.9.0",
)

go_repository(
    name="com_github_hashicorp_hcl",
    importpath="github.com/hashicorp/hcl",
    sum="h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version="v1.0.0",
)

go_repository(
    name="com_github_improbable_eng_grpc_web",
    importpath="github.com/improbable-eng/grpc-web",
    sum="h1:GlCS+lMZzIkfouf7CNqY+qqpowdKuJLSLLcKVfM1oLc=",
    version="v0.12.0",
)

go_repository(
    name="com_github_inconshreveable_mousetrap",
    importpath="github.com/inconshreveable/mousetrap",
    sum="h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version="v1.0.0",
)

go_repository(
    name="com_github_jmoiron_sqlx",
    importpath="github.com/jmoiron/sqlx",
    sum="h1:41Ip0zITnmWNR/vHV+S4m+VoUivnWY5E4OJfLZjCJMA=",
    version="v1.2.0",
)

go_repository(
    name="com_github_jonboulle_clockwork",
    importpath="github.com/jonboulle/clockwork",
    sum="h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    version="v0.1.0",
)

go_repository(
    name="com_github_julienschmidt_httprouter",
    importpath="github.com/julienschmidt/httprouter",
    sum="h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version="v1.2.0",
)

go_repository(
    name="com_github_kisielk_errcheck",
    importpath="github.com/kisielk/errcheck",
    sum="h1:ZqfnKyx9KGpRcW04j5nnPDgRgoXUeLh2YFBeFzphcA0=",
    version="v1.1.0",
)

go_repository(
    name="com_github_kisielk_gotool",
    importpath="github.com/kisielk/gotool",
    sum="h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version="v1.0.0",
)

go_repository(
    name="com_github_konsorten_go_windows_terminal_sequences",
    importpath="github.com/konsorten/go-windows-terminal-sequences",
    sum="h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
    version="v1.0.1",
)

go_repository(
    name="com_github_kr_logfmt",
    importpath="github.com/kr/logfmt",
    sum="h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version="v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name="com_github_kr_pretty",
    importpath="github.com/kr/pretty",
    sum="h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
    version="v0.1.0",
)

go_repository(
    name="com_github_kr_pty",
    importpath="github.com/kr/pty",
    sum="h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version="v1.1.1",
)

go_repository(
    name="com_github_kr_text",
    importpath="github.com/kr/text",
    sum="h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version="v0.1.0",
)

go_repository(
    name="com_github_lib_pq",
    importpath="github.com/lib/pq",
    sum="h1:X5PMW56eZitiTeO7tKzZxFCSpbFZJtkMMooicw2us9A=",
    version="v1.0.0",
)

go_repository(
    name="com_github_magiconair_properties",
    importpath="github.com/magiconair/properties",
    sum="h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
    version="v1.8.0",
)

go_repository(
    name="com_github_mattn_go_sqlite3",
    importpath="github.com/mattn/go-sqlite3",
    sum="h1:pDRiWfl+++eC2FEFRy6jXmQlvp4Yh3z1MJKg4UeYM/4=",
    version="v1.9.0",
)

go_repository(
    name="com_github_matttproud_golang_protobuf_extensions",
    importpath="github.com/matttproud/golang_protobuf_extensions",
    sum="h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version="v1.0.1",
)

go_repository(
    name="com_github_miekg_dns",
    importpath="github.com/miekg/dns",
    sum="h1:xHBEhR+t5RzcFJjBLJlax2daXOrTYtr9z4WdKEfWFzg=",
    version="v1.1.29",
)

go_repository(
    name="com_github_mitchellh_go_homedir",
    importpath="github.com/mitchellh/go-homedir",
    sum="h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version="v1.1.0",
)

go_repository(
    name="com_github_mitchellh_mapstructure",
    importpath="github.com/mitchellh/mapstructure",
    sum="h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version="v1.1.2",
)

go_repository(
    name="com_github_mwitkow_go_conntrack",
    importpath="github.com/mwitkow/go-conntrack",
    sum="h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version="v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name="com_github_oklog_ulid",
    importpath="github.com/oklog/ulid",
    sum="h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version="v1.3.1",
)

go_repository(
    name="com_github_oneofone_xxhash",
    importpath="github.com/OneOfOne/xxhash",
    sum="h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version="v1.2.2",
)

go_repository(
    name="com_github_otiai10_copy",
    importpath="github.com/otiai10/copy",
    sum="h1:PH7IFlRQ6Fv9vYmuXbDRLdgTHoP1w483kPNUP2bskpo=",
    version="v1.1.1",
)

go_repository(
    name="com_github_otiai10_curr",
    importpath="github.com/otiai10/curr",
    sum="h1:TJIWdbX0B+kpNagQrjgq8bCMrbhiuX73M2XwgtDMoOI=",
    version="v1.0.0",
)

go_repository(
    name="com_github_otiai10_mint",
    importpath="github.com/otiai10/mint",
    sum="h1:BCmzIS3n71sGfHB5NMNDB3lHYPz8fWSkCAErHed//qc=",
    version="v1.3.1",
)

go_repository(
    name="com_github_pelletier_go_toml",
    importpath="github.com/pelletier/go-toml",
    sum="h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version="v1.2.0",
)

go_repository(
    name="com_github_pkg_errors",
    importpath="github.com/pkg/errors",
    sum="h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version="v0.9.1",
)

go_repository(
    name="com_github_pmezard_go_difflib",
    importpath="github.com/pmezard/go-difflib",
    sum="h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version="v1.0.0",
)

go_repository(
    name="com_github_prometheus_client_golang",
    importpath="github.com/prometheus/client_golang",
    sum="h1:9iH4JKXLzFbOAdtqv/a+j8aewx2Y8lAjAydhbaScPF8=",
    version="v0.9.3",
)

go_repository(
    name="com_github_prometheus_client_model",
    importpath="github.com/prometheus/client_model",
    sum="h1:S/YWwWx/RA8rT8tKFRuGUZhuA90OyIBpPCXkcbwU8DE=",
    version="v0.0.0-20190129233127-fd36f4220a90",
)

go_repository(
    name="com_github_prometheus_common",
    importpath="github.com/prometheus/common",
    sum="h1:7etb9YClo3a6HjLzfl6rIQaU+FDfi0VSX39io3aQ+DM=",
    version="v0.4.0",
)

go_repository(
    name="com_github_prometheus_procfs",
    importpath="github.com/prometheus/procfs",
    sum="h1:sofwID9zm4tzrgykg80hfFph1mryUeLRsUfoocVVmRY=",
    version="v0.0.0-20190507164030-5867b95ac084",
)

go_repository(
    name="com_github_prometheus_tsdb",
    importpath="github.com/prometheus/tsdb",
    sum="h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
    version="v0.7.1",
)

go_repository(
    name="com_github_rogpeppe_fastuuid",
    importpath="github.com/rogpeppe/fastuuid",
    sum="h1:gu+uRPtBe88sKxUCEXRoeCvVG90TJmwhiqRpvdhQFng=",
    version="v0.0.0-20150106093220-6724a57986af",
)

go_repository(
    name="com_github_rs_cors",
    importpath="github.com/rs/cors",
    sum="h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
    version="v1.7.0",
)

go_repository(
    name="com_github_russross_blackfriday_v2",
    importpath="github.com/russross/blackfriday/v2",
    sum="h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version="v2.0.1",
)

go_repository(
    name="com_github_shurcool_sanitized_anchor_name",
    importpath="github.com/shurcooL/sanitized_anchor_name",
    sum="h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version="v1.0.0",
)

go_repository(
    name="com_github_sirupsen_logrus",
    importpath="github.com/sirupsen/logrus",
    sum="h1:juTguoYk5qI21pwyTXY3B3Y5cOTH3ZUyZCg1v/mihuo=",
    version="v1.2.0",
)

go_repository(
    name="com_github_skratchdot_open_golang",
    importpath="github.com/skratchdot/open-golang",
    sum="h1:JIAuq3EEf9cgbU6AtGPK4CTG3Zf6CKMNqf0MHTggAUA=",
    version="v0.0.0-20200116055534-eef842397966",
)

go_repository(
    name="com_github_soheilhy_cmux",
    importpath="github.com/soheilhy/cmux",
    sum="h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version="v0.1.4",
)

go_repository(
    name="com_github_spaolacci_murmur3",
    importpath="github.com/spaolacci/murmur3",
    sum="h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version="v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name="com_github_spf13_afero",
    importpath="github.com/spf13/afero",
    sum="h1:m8/z1t7/fwjysjQRYbP0RD+bUIF/8tJwPdEZsI83ACI=",
    version="v1.1.2",
)

go_repository(
    name="com_github_spf13_cast",
    importpath="github.com/spf13/cast",
    sum="h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version="v1.3.0",
)

go_repository(
    name="com_github_spf13_cobra",
    importpath="github.com/spf13/cobra",
    sum="h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
    version="v1.0.0",
)

go_repository(
    name="com_github_spf13_jwalterweatherman",
    importpath="github.com/spf13/jwalterweatherman",
    sum="h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version="v1.0.0",
)

go_repository(
    name="com_github_spf13_pflag",
    importpath="github.com/spf13/pflag",
    sum="h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg=",
    version="v1.0.3",
)

go_repository(
    name="com_github_spf13_viper",
    importpath="github.com/spf13/viper",
    sum="h1:yXHLWeravcrgGyFSyCgdYpXQ9dR9c/WED3pg1RhxqEU=",
    version="v1.4.0",
)

go_repository(
    name="com_github_stretchr_objx",
    importpath="github.com/stretchr/objx",
    sum="h1:2vfRuCMp5sSVIDSqO8oNnWJq7mPa6KVP3iPIwFBuy8A=",
    version="v0.1.1",
)

go_repository(
    name="com_github_stretchr_testify",
    importpath="github.com/stretchr/testify",
    sum="h1:nOGnQDM7FYENwehXlg/kFVnos3rEvtKTjRvOWSzb6H4=",
    version="v1.5.1",
)

go_repository(
    name="com_github_tmc_grpc_websocket_proxy",
    importpath="github.com/tmc/grpc-websocket-proxy",
    sum="h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
    version="v0.0.0-20190109142713-0ad062ec5ee5",
)

go_repository(
    name="com_github_ugorji_go",
    importpath="github.com/ugorji/go",
    sum="h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
    version="v1.1.4",
)

go_repository(
    name="com_github_xiang90_probing",
    importpath="github.com/xiang90/probing",
    sum="h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version="v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name="com_github_xordataexchange_crypt",
    importpath="github.com/xordataexchange/crypt",
    sum="h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version="v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name="com_github_yuin_goldmark",
    importpath="github.com/yuin/goldmark",
    sum="h1:nqDD4MMMQA0lmWq03Z2/myGPYLQoXtmi0rGVs95ntbo=",
    version="v1.1.27",
)

go_repository(
    name="com_google_cloud_go",
    importpath="cloud.google.com/go",
    sum="h1:eOI3/cP2VTU6uZLDYAoic+eyzzB9YyGmJ7eIjl8rOPg=",
    version="v0.34.0",
)

go_repository(
    name="in_gopkg_alecthomas_kingpin_v2",
    importpath="gopkg.in/alecthomas/kingpin.v2",
    sum="h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version="v2.2.6",
)

go_repository(
    name="in_gopkg_check_v1",
    importpath="gopkg.in/check.v1",
    sum="h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=",
    version="v1.0.0-20180628173108-788fd7840127",
)

go_repository(
    name="in_gopkg_resty_v1",
    importpath="gopkg.in/resty.v1",
    sum="h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version="v1.12.0",
)

go_repository(
    name="in_gopkg_yaml_v2",
    importpath="gopkg.in/yaml.v2",
    sum="h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version="v2.2.2",
)

go_repository(
    name="io_etcd_go_bbolt",
    importpath="go.etcd.io/bbolt",
    sum="h1:Z/90sZLPOeCy2PwprqkFa25PdkusRzaj9P8zm/KNyvk=",
    version="v1.3.2",
)

go_repository(
    name="io_k8s_klog_v2",
    importpath="k8s.io/klog/v2",
    sum="h1:XRvcwJozkgZ1UQJmfMGpvRthQHOvihEhYtDfAaxMz/A=",
    version="v2.2.0",
)

go_repository(
    name="org_golang_google_appengine",
    importpath="google.golang.org/appengine",
    sum="h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
    version="v1.4.0",
)

go_repository(
    name="org_golang_google_genproto",
    importpath="google.golang.org/genproto",
    sum="h1:Nw54tB0rB7hY/N0NQvRW8DG4Yk3Q6T9cu9RcFQDu1tc=",
    version="v0.0.0-20180817151627-c66870c02cf8",
)

go_repository(
    name="org_golang_google_grpc",
    importpath="google.golang.org/grpc",
    sum="h1:G+97AoqBnmZIT91cLG/EkCoK9NSelj64P8bOHHNmGn0=",
    version="v1.21.0",
)

go_repository(
    name="org_golang_x_crypto",
    importpath="golang.org/x/crypto",
    sum="h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=",
    version="v0.0.0-20200622213623-75b288015ac9",
)

go_repository(
    name="org_golang_x_lint",
    importpath="golang.org/x/lint",
    sum="h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
    version="v0.0.0-20190313153728-d0100b6bd8b3",
)

go_repository(
    name="org_golang_x_mod",
    importpath="golang.org/x/mod",
    sum="h1:KU7oHjnv3XNWfa5COkzUifxZmxp1TyI7ImMXqFxLwvQ=",
    version="v0.2.0",
)

go_repository(
    name="org_golang_x_net",
    importpath="golang.org/x/net",
    sum="h1:VXak5I6aEWmAXeQjA+QSZzlgNrpq9mjcfDemuexIKsU=",
    version="v0.0.0-20200707034311-ab3426394381",
)

go_repository(
    name="org_golang_x_oauth2",
    importpath="golang.org/x/oauth2",
    sum="h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=",
    version="v0.0.0-20200107190931-bf48bf16ab8d",
)

go_repository(
    name="org_golang_x_sync",
    importpath="golang.org/x/sync",
    sum="h1:vcxGaoTs7kV8m5Np9uUNQin4BrLOthgV7252N8V+FwY=",
    version="v0.0.0-20190911185100-cd5d95a43a6e",
)

go_repository(
    name="org_golang_x_sys",
    importpath="golang.org/x/sys",
    sum="h1:Ih9Yo4hSPImZOpfGuA4bR/ORKTAbhZo2AbWNRCnevdo=",
    version="v0.0.0-20200625212154-ddb9806d33ae",
)

go_repository(
    name="org_golang_x_text",
    importpath="golang.org/x/text",
    sum="h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version="v0.3.0",
)

go_repository(
    name="org_golang_x_time",
    importpath="golang.org/x/time",
    sum="h1:SvFZT6jyqRaOeXpc5h/JSfZenJ2O330aBsf7JfSUXmQ=",
    version="v0.0.0-20190308202827-9d24e82272b4",
)

go_repository(
    name="org_golang_x_tools",
    importpath="golang.org/x/tools",
    sum="h1:FkAkwuYWQw+IArrnmhGlisKHQF4MsZ2Nu/fX4ttW55o=",
    version="v0.0.0-20190122202912-9c309ee22fab",
)

go_repository(
    name="org_golang_x_xerrors",
    importpath="golang.org/x/xerrors",
    sum="h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
    version="v0.0.0-20191204190536-9bdfabe68543",
)

go_repository(
    name="org_uber_go_atomic",
    importpath="go.uber.org/atomic",
    sum="h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
    version="v1.4.0",
)

go_repository(
    name="org_uber_go_multierr",
    importpath="go.uber.org/multierr",
    sum="h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
    version="v1.1.0",
)

go_repository(
    name="org_uber_go_zap",
    importpath="go.uber.org/zap",
    sum="h1:ORx85nbTijNz8ljznvCMR1ZBIPKFn3jQrag10X2AsuM=",
    version="v1.10.0",
)

go_repository(
    name="tools_gotest",
    importpath="gotest.tools",
    sum="h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version="v2.2.0+incompatible",
)

go_repository(
    name="com_github_bazelbuild_bazel_gazelle",
    importpath="github.com/bazelbuild/bazel-gazelle",
    sum="h1:buszGdD9d/Z691sxFDgOdcEUWli0ZT2tBXUxfbLMrb4=",
    version="v0.21.1",
)

go_repository(
    name="com_github_bazelbuild_buildtools",
    importpath="github.com/bazelbuild/buildtools",
    sum="h1:OfyUN/Msd8yqJww6deQ9vayJWw+Jrbe6Qp9giv51QQI=",
    version="v0.0.0-20190731111112-f720930ceb60",
)

go_repository(
    name="com_github_bazelbuild_rules_go",
    importpath="github.com/bazelbuild/rules_go",
    sum="h1:wzbawlkLtl2ze9w/312NHZ84c7kpUCtlkD8HgFY27sw=",
    version="v0.0.0-20190719190356-6dae44dc5cab",
)

go_repository(
    name="com_github_bmatcuk_doublestar",
    importpath="github.com/bmatcuk/doublestar",
    sum="h1:oC24CykoSAB8zd7XgruHo33E0cHJf/WhQA/7BeXj+x0=",
    version="v1.2.2",
)

go_repository(
    name="com_github_chzyer_logex",
    importpath="github.com/chzyer/logex",
    sum="h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version="v1.1.10",
)

go_repository(
    name="com_github_chzyer_readline",
    importpath="github.com/chzyer/readline",
    sum="h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version="v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name="com_github_chzyer_test",
    importpath="github.com/chzyer/test",
    sum="h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version="v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name="com_github_x1ddos_csslex",
    importpath="github.com/x1ddos/csslex",
    sum="h1:SX7lFdwn40ahL78CxofAh548P+dcWjdRNpirU7+sKiE=",
    version="v0.0.0-20160125172232-7894d8ab8bfe",
)

go_repository(
    name="net_starlark_go",
    importpath="go.starlark.net",
    sum="h1:uFqwFYlX7d5ZSp+IqhXxct0SybXrTzEBDvb2CkEhPBs=",
    version="v0.0.0-20200707032745-474f21a9602d",
)
