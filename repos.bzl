load("@bazel_gazelle//:deps.bzl", "go_repository")


def go_repositories():
    go_repository(
        name="com_github_googlecodelabs_tools",
        importpath="github.com/googlecodelabs/tools",
        sum="h1:pW6O2xWfr0k/kcaJlDY93jltyQnhDAagu9aX/pIoOKM=",
        version="v2.2.0+incompatible",
    )

    go_repository(
        name="cc_mvdan_gofumpt",
        importpath="mvdan.cc/gofumpt",
        sum="h1:gi7cb8HTDZ6q8VqsUpkdoFi3vxwHMneQ6+Q5Ap5hjPE=",
        version="v0.0.0-20200709182408-4fd085cb6d5f",
    )

    go_repository(
        name="cc_mvdan_interfacer",
        importpath="mvdan.cc/interfacer",
        sum="h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
        version="v0.0.0-20180901003855-c20040233aed",
    )

    go_repository(
        name="cc_mvdan_lint",
        importpath="mvdan.cc/lint",
        sum="h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
        version="v0.0.0-20170908181259-adc824a0674b",
    )

    go_repository(
        name="cc_mvdan_unparam",
        importpath="mvdan.cc/unparam",
        sum="h1:kAREL6MPwpsk1/PQPFD3Eg7WAQR5mPTWZJaBiG5LDbY=",
        version="v0.0.0-20200501210554-b37ab49443f7",
    )

    go_repository(
        name="cc_mvdan_xurls_v2",
        importpath="mvdan.cc/xurls/v2",
        sum="h1:NSZPykBXJFCetGZykLAxaL6SIpvbVy/UFEniIfHAa8A=",
        version="v2.2.0",
    )

    go_repository(
        name="co_honnef_go_tools",
        importpath="honnef.co/go/tools",
        sum="h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
        version="v0.0.1-2020.1.4",
    )

    go_repository(
        name="com_github_akavel_rsrc",
        importpath="github.com/akavel/rsrc",
        sum="h1:zjWn7ukO9Kc5Q62DOJCcxGpXC18RawVtYAGdz2aLlfw=",
        version="v0.8.0",
    )

    go_repository(
        name="com_github_alcortesm_tgz",
        importpath="github.com/alcortesm/tgz",
        sum="h1:uSoVVbwJiQipAclBbw+8quDsfcvFjOpI5iCf4p/cqCs=",
        version="v0.0.0-20161220082320-9c5fe88206d7",
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
        name="com_github_andybalholm_cascadia",
        importpath="github.com/andybalholm/cascadia",
        sum="h1:vuRCkM5Ozh/BfmsaTm26kbjm0mIOM3yS5Ek/F5h18aE=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_anmitsu_go_shlex",
        importpath="github.com/anmitsu/go-shlex",
        sum="h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
        version="v0.0.0-20161002113705-648efa622239",
    )

    go_repository(
        name="com_github_armon_circbuf",
        importpath="github.com/armon/circbuf",
        sum="h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version="v0.0.0-20150827004946-bbbad097214e",
    )

    go_repository(
        name="com_github_armon_consul_api",
        importpath="github.com/armon/consul-api",
        sum="h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version="v0.0.0-20180202201655-eb2c6b5be1b6",
    )

    go_repository(
        name="com_github_armon_go_metrics",
        importpath="github.com/armon/go-metrics",
        sum="h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
        version="v0.0.0-20180917152333-f0300d1749da",
    )

    go_repository(
        name="com_github_armon_go_radix",
        importpath="github.com/armon/go-radix",
        sum="h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
        version="v0.0.0-20180808171621-7fddfc383310",
    )

    go_repository(
        name="com_github_armon_go_socks5",
        importpath="github.com/armon/go-socks5",
        sum="h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
        version="v0.0.0-20160902184237-e75332964ef5",
    )

    go_repository(
        name="com_github_atotto_clipboard",
        importpath="github.com/atotto/clipboard",
        sum="h1:YZCtFu5Ie8qX2VmVTBnrqLSiU9XOWwqNRmdT3gIQzbY=",
        version="v0.1.2",
    )

    go_repository(
        name="com_github_bazelbuild_rules_webtesting",
        importpath="github.com/bazelbuild/rules_webtesting",
        sum="h1:nhqjA2IslEOLViRBF5djQCiOD//7VyyHNKrqAZ1AuYA=",
        version="v0.2.0",
    )

    go_repository(
        name="com_github_beorn7_perks",
        importpath="github.com/beorn7/perks",
        sum="h1:HWo1m869IqiPhD389kmkxeTalrjNbbJTC8LXupb+sl0=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_bgentry_speakeasy",
        importpath="github.com/bgentry/speakeasy",
        sum="h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version="v0.1.0",
    )

    go_repository(
        name="com_github_bketelsen_crypt",
        importpath="github.com/bketelsen/crypt",
        sum="h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
        version="v0.0.3-0.20200106085610-5cbc8cc4026c",
    )

    go_repository(
        name="com_github_blang_semver",
        importpath="github.com/blang/semver",
        sum="h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
        version="v3.5.1+incompatible",
    )

    go_repository(
        name="com_github_bombsimon_wsl",
        importpath="github.com/bombsimon/wsl",
        sum="h1:b+E/W7koicKBZDU+vEsw/hnQTN8026Gv1eMZDLUU/Wc=",
        version="v1.2.8",
    )

    go_repository(
        name="com_github_bombsimon_wsl_v3",
        importpath="github.com/bombsimon/wsl/v3",
        sum="h1:E5SRssoBgtVFPcYWUOFJEcgaySgdtTNYzsSKDOY7ss8=",
        version="v3.1.0",
    )

    go_repository(
        name="com_github_burntsushi_toml",
        importpath="github.com/BurntSushi/toml",
        sum="h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version="v0.3.1",
    )

    go_repository(
        name="com_github_burntsushi_xgb",
        importpath="github.com/BurntSushi/xgb",
        sum="h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version="v0.0.0-20160522181843-27f122750802",
    )

    go_repository(
        name="com_github_burntsushi_xgbutil",
        importpath="github.com/BurntSushi/xgbutil",
        sum="h1:4ZrkT/RzpnROylmoQL57iVUL57wGKTR5O6KpVnbm2tA=",
        version="v0.0.0-20160919175755-f7c97cef3b4e",
    )

    go_repository(
        name="com_github_census_instrumentation_opencensus_proto",
        importpath="github.com/census-instrumentation/opencensus-proto",
        sum="h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
        version="v0.2.1",
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
        sum="h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
        version="v3.3.13+incompatible",
    )

    go_repository(
        name="com_github_coreos_go_semver",
        importpath="github.com/coreos/go-semver",
        sum="h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version="v0.3.0",
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
        name="com_github_cosiner_argv",
        importpath="github.com/cosiner/argv",
        sum="h1:rIXlvz2IWiupMFlC45cZCXZFvKX/ExBcSLrDy2G0Lp8=",
        version="v0.0.0-20170225145430-13bacc38a0a5",
    )

    go_repository(
        name="com_github_cpuguy83_go_md2man",
        importpath="github.com/cpuguy83/go-md2man",
        sum="h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version="v1.0.10",
    )

    go_repository(
        name="com_github_cpuguy83_go_md2man_v2",
        importpath="github.com/cpuguy83/go-md2man/v2",
        sum="h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
        version="v2.0.0",
    )

    go_repository(
        name="com_github_creack_pty",
        importpath="github.com/creack/pty",
        sum="h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version="v1.1.9",
    )

    go_repository(
        name="com_github_davecgh_go_spew",
        importpath="github.com/davecgh/go-spew",
        sum="h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version="v1.1.1",
    )

    go_repository(
        name="com_github_denis_tingajkin_go_header",
        importpath="github.com/denis-tingajkin/go-header",
        sum="h1:ymEpSiFjeItCy1FOP+x0M2KdCELdEAHUsNa8F+hHc6w=",
        version="v0.3.1",
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
        name="com_github_djarvur_go_err113",
        importpath="github.com/Djarvur/go-err113",
        sum="h1:XTrzB+F8+SpRmbhAH8HLxhiiG6nYNwaBZjrFps1oWEk=",
        version="v0.0.0-20200511133814-5174e21577d5",
    )

    go_repository(
        name="com_github_emirpasic_gods",
        importpath="github.com/emirpasic/gods",
        sum="h1:QAUIPSaCu4G+POclxeqb3F+WPpdKqFGlw36+yOzGlrg=",
        version="v1.12.0",
    )

    go_repository(
        name="com_github_envoyproxy_go_control_plane",
        importpath="github.com/envoyproxy/go-control-plane",
        sum="h1:4cmBvAEBNJaGARUEs3/suWRyfyBfhf7I60WBZq+bv2w=",
        version="v0.9.1-0.20191026205805-5f8ba28d4473",
    )

    go_repository(
        name="com_github_envoyproxy_protoc_gen_validate",
        importpath="github.com/envoyproxy/protoc-gen-validate",
        sum="h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
        version="v0.1.0",
    )

    go_repository(
        name="com_github_fatih_color",
        importpath="github.com/fatih/color",
        sum="h1:8xPHl4/q1VyqGIPif1F+1V3Y3lSmrq01EabUW3CoW5s=",
        version="v1.9.0",
    )

    go_repository(
        name="com_github_flynn_go_shlex",
        importpath="github.com/flynn/go-shlex",
        sum="h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
        version="v0.0.0-20150515145356-3f9db97f8568",
    )

    go_repository(
        name="com_github_fsnotify_fsnotify",
        importpath="github.com/fsnotify/fsnotify",
        sum="h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
        version="v1.4.9",
    )

    go_repository(
        name="com_github_fyne_io_mobile",
        importpath="github.com/fyne-io/mobile",
        sum="h1:eGmCR5lkFxk0PnPafGppLFRD5QODJfSVdrjhLjanOVg=",
        version="v0.0.2",
    )

    go_repository(
        name="com_github_ghodss_yaml",
        importpath="github.com/ghodss/yaml",
        sum="h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_gliderlabs_ssh",
        importpath="github.com/gliderlabs/ssh",
        sum="h1:6zsha5zo/TWhRhwqCD3+EarCAgZ2yN28ipRnGPnwkI0=",
        version="v0.2.2",
    )

    go_repository(
        name="com_github_go_critic_go_critic",
        importpath="github.com/go-critic/go-critic",
        sum="h1:Ic2p5UCl5fX/2WX2w8nroPpPhxRNsNTMlJzsu/uqwnM=",
        version="v0.5.0",
    )

    go_repository(
        name="com_github_go_delve_delve",
        importpath="github.com/go-delve/delve",
        sum="h1:kZs0umEv+VKnK84kY9/ZXWrakdLTeRTyYjFdgLelZCQ=",
        version="v1.4.1",
    )

    go_repository(
        name="com_github_go_git_gcfg",
        importpath="github.com/go-git/gcfg",
        sum="h1:Q5ViNfGF8zFgyJWPqYwA7qGFoMTEiBmdlkcfRmpIMa4=",
        version="v1.5.0",
    )

    go_repository(
        name="com_github_go_git_go_billy",
        importpath="github.com/go-git/go-billy",
        sum="h1:Z6QtVXd5tjxUtcODLugkJg4WaZnGg13CD8qB9pr+7q0=",
        version="v4.2.0+incompatible",
    )

    go_repository(
        name="com_github_go_git_go_git",
        importpath="github.com/go-git/go-git",
        sum="h1:+W9rgGY4DOKKdX2x6HxSR7HNeTxqiKrOvKnuittYVdA=",
        version="v4.7.0+incompatible",
    )

    go_repository(
        name="com_github_go_gl_gl",
        importpath="github.com/go-gl/gl",
        sum="h1:SCYMcCJ89LjRGwEa0tRluNRiMjZHalQZrVrvTbPh+qw=",
        version="v0.0.0-20190320180904-bf2b1f2f34d7",
    )

    go_repository(
        name="com_github_go_gl_glfw",
        importpath="github.com/go-gl/glfw",
        sum="h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version="v0.0.0-20190409004039-e6da0acd62b1",
    )

    go_repository(
        name="com_github_go_gl_glfw_v3_3_glfw",
        importpath="github.com/go-gl/glfw/v3.3/glfw",
        sum="h1:q521PfSp5/z6/sD9FZZOWj4d1MLmfQW8PkRnI9M6PCE=",
        version="v0.0.0-20200625191551-73d3c3675aa3",
    )

    go_repository(
        name="com_github_go_kit_kit",
        importpath="github.com/go-kit/kit",
        sum="h1:Wz+5lgoB0kkuqLEc6NVmwRknTKP6dTGbSqvhZtBI/j0=",
        version="v0.8.0",
    )

    go_repository(
        name="com_github_go_lintpack_lintpack",
        importpath="github.com/go-lintpack/lintpack",
        sum="h1:DI5mA3+eKdWeJ40nU4d6Wc26qmdG8RCi/btYq0TuRN0=",
        version="v0.5.2",
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
        name="com_github_go_ole_go_ole",
        importpath="github.com/go-ole/go-ole",
        sum="h1:2lOsA72HgjxAuMlKpFiCbHTvu44PIVkZ5hqm3RSdI/E=",
        version="v1.2.1",
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
        name="com_github_go_toolsmith_astcast",
        importpath="github.com/go-toolsmith/astcast",
        sum="h1:JojxlmI6STnFVG9yOImLeGREv8W2ocNUM+iOhR6jE7g=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_astcopy",
        importpath="github.com/go-toolsmith/astcopy",
        sum="h1:OMgl1b1MEpjFQ1m5ztEO06rz5CUd3oBv9RF7+DyvdG8=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_astequal",
        importpath="github.com/go-toolsmith/astequal",
        sum="h1:4zxD8j3JRFNyLN46lodQuqz3xdKSrur7U/sr0SDS/gQ=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_astfmt",
        importpath="github.com/go-toolsmith/astfmt",
        sum="h1:A0vDDXt+vsvLEdbMFJAUBI/uTbRw1ffOPnxsILnFL6k=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_astinfo",
        importpath="github.com/go-toolsmith/astinfo",
        sum="h1:wP6mXeB2V/d1P1K7bZ5vDUO3YqEzcvOREOxZPEu3gVI=",
        version="v0.0.0-20180906194353-9809ff7efb21",
    )

    go_repository(
        name="com_github_go_toolsmith_astp",
        importpath="github.com/go-toolsmith/astp",
        sum="h1:alXE75TXgcmupDsMK1fRAy0YUzLzqPVvBKoyWV+KPXg=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_pkgload",
        importpath="github.com/go-toolsmith/pkgload",
        sum="h1:4DFWWMXVfbcN5So1sBNW9+yeiMqLFGl1wFLTL5R0Tgg=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_strparse",
        importpath="github.com/go-toolsmith/strparse",
        sum="h1:Vcw78DnpCAKlM20kSbAyO4mPfJn/lyYA4BJUDxe2Jb4=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_go_toolsmith_typep",
        importpath="github.com/go-toolsmith/typep",
        sum="h1:8xdsa1+FSIH/RhEkgnD1j2CJOy5mNllW1Q9tRiYwvlk=",
        version="v1.0.2",
    )

    go_repository(
        name="com_github_go_xmlfmt_xmlfmt",
        importpath="github.com/go-xmlfmt/xmlfmt",
        sum="h1:khEcpUM4yFcxg4/FHQWkvVRmgijNXRfzkIDHh23ggEo=",
        version="v0.0.0-20191208150333-d5b6f63a941b",
    )

    go_repository(
        name="com_github_gobwas_glob",
        importpath="github.com/gobwas/glob",
        sum="h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
        version="v0.2.3",
    )

    go_repository(
        name="com_github_gocolly_colly",
        importpath="github.com/gocolly/colly",
        sum="h1:qRz9YAn8FIH0qzgNUw+HT9UN7wm1oF9OBAilwEWpyrI=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_godbus_dbus_v5",
        importpath="github.com/godbus/dbus/v5",
        sum="h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
        version="v5.0.3",
    )

    go_repository(
        name="com_github_gofrs_flock",
        importpath="github.com/gofrs/flock",
        sum="h1:DP+LD/t0njgoPBvT5MJLeliUIVQR03hiKR6vezdwHlc=",
        version="v0.7.1",
    )

    go_repository(
        name="com_github_gogo_protobuf",
        importpath="github.com/gogo/protobuf",
        sum="h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
        version="v1.3.1",
    )

    go_repository(
        name="com_github_goki_freetype",
        importpath="github.com/goki/freetype",
        sum="h1:W71vTCKoxtdXgnm1ECDFkfQnpdqAO00zzGXLA5yaEX8=",
        version="v0.0.0-20181231101311-fa8a33aabaff",
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
        sum="h1:qGJ6qTW+x6xX/my+8YUVl4WNpX9B7+/l2tRsHGZ7f2s=",
        version="v1.3.1",
    )

    go_repository(
        name="com_github_golang_protobuf",
        importpath="github.com/golang/protobuf",
        sum="h1:ZFgWrT+bLgsYPirOnRfKLYJLvssAegOj/hgyMFdJZe0=",
        version="v1.4.1",
    )

    go_repository(
        name="com_github_golangci_check",
        importpath="github.com/golangci/check",
        sum="h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
        version="v0.0.0-20180506172741-cfe4005ccda2",
    )

    go_repository(
        name="com_github_golangci_dupl",
        importpath="github.com/golangci/dupl",
        sum="h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
        version="v0.0.0-20180902072040-3e9179ac440a",
    )

    go_repository(
        name="com_github_golangci_errcheck",
        importpath="github.com/golangci/errcheck",
        sum="h1:YYWNAGTKWhKpcLLt7aSj/odlKrSrelQwlovBpDuf19w=",
        version="v0.0.0-20181223084120-ef45e06d44b6",
    )

    go_repository(
        name="com_github_golangci_go_misc",
        importpath="github.com/golangci/go-misc",
        sum="h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
        version="v0.0.0-20180628070357-927a3d87b613",
    )

    go_repository(
        name="com_github_golangci_goconst",
        importpath="github.com/golangci/goconst",
        sum="h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
        version="v0.0.0-20180610141641-041c5f2b40f3",
    )

    go_repository(
        name="com_github_golangci_gocyclo",
        importpath="github.com/golangci/gocyclo",
        sum="h1:pXTK/gkVNs7Zyy7WKgLXmpQ5bHTrq5GDsp8R9Qs67g0=",
        version="v0.0.0-20180528144436-0a533e8fa43d",
    )

    go_repository(
        name="com_github_golangci_gofmt",
        importpath="github.com/golangci/gofmt",
        sum="h1:iR3fYXUjHCR97qWS8ch1y9zPNsgXThGwjKPrYfqMPks=",
        version="v0.0.0-20190930125516-244bba706f1a",
    )

    go_repository(
        name="com_github_golangci_golangci_lint",
        importpath="github.com/golangci/golangci-lint",
        sum="h1:0ufaO3l2R1R712cFC+KT3TtwO/IOcsloKZBavRtzrBk=",
        version="v1.29.0",
    )

    go_repository(
        name="com_github_golangci_ineffassign",
        importpath="github.com/golangci/ineffassign",
        sum="h1:gLLhTLMk2/SutryVJ6D4VZCU3CUqr8YloG7FPIBWFpI=",
        version="v0.0.0-20190609212857-42439a7714cc",
    )

    go_repository(
        name="com_github_golangci_lint_1",
        importpath="github.com/golangci/lint-1",
        sum="h1:MfyDlzVjl1hoaPzPD4Gpb/QgoRfSBR0jdhwGyAWwMSA=",
        version="v0.0.0-20191013205115-297bf364a8e0",
    )

    go_repository(
        name="com_github_golangci_maligned",
        importpath="github.com/golangci/maligned",
        sum="h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
        version="v0.0.0-20180506175553-b1d89398deca",
    )

    go_repository(
        name="com_github_golangci_misspell",
        importpath="github.com/golangci/misspell",
        sum="h1:pLzmVdl3VxTOncgzHcvLOKirdvcx/TydsClUQXTehjo=",
        version="v0.3.5",
    )

    go_repository(
        name="com_github_golangci_prealloc",
        importpath="github.com/golangci/prealloc",
        sum="h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
        version="v0.0.0-20180630174525-215b22d4de21",
    )

    go_repository(
        name="com_github_golangci_revgrep",
        importpath="github.com/golangci/revgrep",
        sum="h1:HVfrLniijszjS1aiNg8JbBMO2+E1WIQ+j/gL4SQqGPg=",
        version="v0.0.0-20180526074752-d9c87f5ffaf0",
    )

    go_repository(
        name="com_github_golangci_unconvert",
        importpath="github.com/golangci/unconvert",
        sum="h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
        version="v0.0.0-20180507085042-28b1c447d1f4",
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
        name="com_github_google_go_dap",
        importpath="github.com/google/go-dap",
        sum="h1:whjIGQRumwbR40qRU7CEKuFLmePUUc2s4Nt9DoXXxWk=",
        version="v0.2.0",
    )

    go_repository(
        name="com_github_google_go_github_v27",
        importpath="github.com/google/go-github/v27",
        sum="h1:N/EEqsvJLgqTbepTiMBz+12KhwLovv6YvwpRezd+4Fg=",
        version="v27.0.4",
    )

    go_repository(
        name="com_github_google_go_querystring",
        importpath="github.com/google/go-querystring",
        sum="h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_google_martian",
        importpath="github.com/google/martian",
        sum="h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version="v2.1.0+incompatible",
    )

    go_repository(
        name="com_github_google_pprof",
        importpath="github.com/google/pprof",
        sum="h1:Jnx61latede7zDD3DiiP4gmNz33uK0U5HDUaF0a/HVQ=",
        version="v0.0.0-20190515194954-54271f7e092f",
    )

    go_repository(
        name="com_github_google_renameio",
        importpath="github.com/google/renameio",
        sum="h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version="v0.1.0",
    )

    go_repository(
        name="com_github_google_uuid",
        importpath="github.com/google/uuid",
        sum="h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
        version="v1.1.1",
    )

    go_repository(
        name="com_github_googleapis_gax_go_v2",
        importpath="github.com/googleapis/gax-go/v2",
        sum="h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
        version="v2.0.5",
    )

    go_repository(
        name="com_github_googlecloudplatform_govanityurls",
        importpath="github.com/GoogleCloudPlatform/govanityurls",
        sum="h1:rx+UxBa3+GOKqHor2H8AKiovqV62FZWu2fO1qHAAhMk=",
        version="v0.1.0",
    )

    go_repository(
        name="com_github_gookit_color",
        importpath="github.com/gookit/color",
        sum="h1:xOYBan3Fwlrqj1M1UN2TlHOCRiek3bGzWf/vPnJ1roE=",
        version="v1.2.4",
    )

    go_repository(
        name="com_github_gopherjs_gopherjs",
        importpath="github.com/gopherjs/gopherjs",
        sum="h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version="v0.0.0-20181017120253-0766667cb4d1",
    )

    go_repository(
        name="com_github_gorilla_websocket",
        importpath="github.com/gorilla/websocket",
        sum="h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
        version="v1.4.2",
    )

    go_repository(
        name="com_github_gostaticanalysis_analysisutil",
        importpath="github.com/gostaticanalysis/analysisutil",
        sum="h1:iwp+5/UAyzQSFgQ4uR2sni99sJ8Eo9DEacKWM5pekIg=",
        version="v0.0.3",
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
        name="com_github_hashicorp_consul_api",
        importpath="github.com/hashicorp/consul/api",
        sum="h1:BNQPM9ytxj6jbjjdRPioQ94T6YXriSopn0i8COv6SRA=",
        version="v1.1.0",
    )

    go_repository(
        name="com_github_hashicorp_consul_sdk",
        importpath="github.com/hashicorp/consul/sdk",
        sum="h1:LnuDWGNsoajlhGyHJvuWW6FVqRl8JOTPqS6CPTsYjhY=",
        version="v0.1.1",
    )

    go_repository(
        name="com_github_hashicorp_errwrap",
        importpath="github.com/hashicorp/errwrap",
        sum="h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_cleanhttp",
        importpath="github.com/hashicorp/go-cleanhttp",
        sum="h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
        version="v0.5.1",
    )

    go_repository(
        name="com_github_hashicorp_go_immutable_radix",
        importpath="github.com/hashicorp/go-immutable-radix",
        sum="h1:AKDB1HM5PWEA7i4nhcpwOrO2byshxBjXVn/J/3+z5/0=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_msgpack",
        importpath="github.com/hashicorp/go-msgpack",
        sum="h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
        version="v0.5.3",
    )

    go_repository(
        name="com_github_hashicorp_go_multierror",
        importpath="github.com/hashicorp/go-multierror",
        sum="h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_net",
        importpath="github.com/hashicorp/go.net",
        sum="h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
        version="v0.0.1",
    )

    go_repository(
        name="com_github_hashicorp_go_rootcerts",
        importpath="github.com/hashicorp/go-rootcerts",
        sum="h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_sockaddr",
        importpath="github.com/hashicorp/go-sockaddr",
        sum="h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_syslog",
        importpath="github.com/hashicorp/go-syslog",
        sum="h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_go_uuid",
        importpath="github.com/hashicorp/go-uuid",
        sum="h1:fv1ep09latC32wFoVwnqcnKJGnMSdBanPczbHAYm1BE=",
        version="v1.0.1",
    )

    go_repository(
        name="com_github_hashicorp_golang_lru",
        importpath="github.com/hashicorp/golang-lru",
        sum="h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version="v0.5.4",
    )

    go_repository(
        name="com_github_hashicorp_hcl",
        importpath="github.com/hashicorp/hcl",
        sum="h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_logutils",
        importpath="github.com/hashicorp/logutils",
        sum="h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_mdns",
        importpath="github.com/hashicorp/mdns",
        sum="h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_hashicorp_memberlist",
        importpath="github.com/hashicorp/memberlist",
        sum="h1:EmmoJme1matNzb+hMpDuR/0sbJSUisxyqBGG676r31M=",
        version="v0.1.3",
    )

    go_repository(
        name="com_github_hashicorp_serf",
        importpath="github.com/hashicorp/serf",
        sum="h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
        version="v0.8.2",
    )

    go_repository(
        name="com_github_hpcloud_tail",
        importpath="github.com/hpcloud/tail",
        sum="h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
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
        name="com_github_jackmordaunt_icns",
        importpath="github.com/jackmordaunt/icns",
        sum="h1:NfuKjkj/Xc2z1xZIj+EmNCm5p1nKJPyw3F4E20usXvg=",
        version="v0.0.0-20181231085925-4f16af745526",
    )

    go_repository(
        name="com_github_jawher_mow_cli",
        importpath="github.com/jawher/mow.cli",
        sum="h1:NdtHXRc0CwZQ507wMvQ/IS+Q3W3x2fycn973/b8Zuk8=",
        version="v1.1.0",
    )

    go_repository(
        name="com_github_jbenet_go_context",
        importpath="github.com/jbenet/go-context",
        sum="h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=",
        version="v0.0.0-20150711004518-d14ea06fba99",
    )

    go_repository(
        name="com_github_jessevdk_go_flags",
        importpath="github.com/jessevdk/go-flags",
        sum="h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
        version="v1.4.0",
    )

    go_repository(
        name="com_github_jingyugao_rowserrcheck",
        importpath="github.com/jingyugao/rowserrcheck",
        sum="h1:GmsqmapfzSJkm28dhRoHz2tLRbJmqhU86IPgBtN3mmk=",
        version="v0.0.0-20191204022205-72ab7603b68a",
    )

    go_repository(
        name="com_github_jirfag_go_printf_func_name",
        importpath="github.com/jirfag/go-printf-func-name",
        sum="h1:jNYPNLe3d8smommaoQlK7LOA5ESyUJJ+Wf79ZtA7Vp4=",
        version="v0.0.0-20191110105641-45db9963cdd3",
    )

    go_repository(
        name="com_github_jmoiron_sqlx",
        importpath="github.com/jmoiron/sqlx",
        sum="h1:lrdPtrORjGv1HbbEvKWDUAy97mPpFm4B8hp77tcCUJY=",
        version="v1.2.1-0.20190826204134-d7d95172beb5",
    )

    go_repository(
        name="com_github_jonboulle_clockwork",
        importpath="github.com/jonboulle/clockwork",
        sum="h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
        version="v0.1.0",
    )

    go_repository(
        name="com_github_josephspurrier_goversioninfo",
        importpath="github.com/josephspurrier/goversioninfo",
        sum="h1:ozPUX9TKQZVek4lZWYRsQo7uS8vJ+q4OOHvRhHiCLfU=",
        version="v0.0.0-20200309025242-14b0ab84c6ca",
    )

    go_repository(
        name="com_github_json_iterator_go",
        importpath="github.com/json-iterator/go",
        sum="h1:MrUvLMLTMxbqFJ9kzlvat/rYZqZnW3u4wkLzWTaFwKs=",
        version="v1.1.6",
    )

    go_repository(
        name="com_github_jstemmer_go_junit_report",
        importpath="github.com/jstemmer/go-junit-report",
        sum="h1:rBMNdlhTLzJjJSDIjNEXX1Pz3Hmwmz91v+zycvx9PJc=",
        version="v0.0.0-20190106144839-af01ea7f8024",
    )

    go_repository(
        name="com_github_jtolds_gls",
        importpath="github.com/jtolds/gls",
        sum="h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version="v4.20.0+incompatible",
    )

    go_repository(
        name="com_github_julienschmidt_httprouter",
        importpath="github.com/julienschmidt/httprouter",
        sum="h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_kennygrant_sanitize",
        importpath="github.com/kennygrant/sanitize",
        sum="h1:gN25/otpP5vAsO2djbMhF/LQX6R7+O1TB4yv8NzpJ3o=",
        version="v1.2.4",
    )

    go_repository(
        name="com_github_kevinburke_ssh_config",
        importpath="github.com/kevinburke/ssh_config",
        sum="h1:Coekwdh0v2wtGp9Gmz1Ze3eVRAWJMLokvN3QjdzCHLY=",
        version="v0.0.0-20190725054713-01f96b0aa0cd",
    )

    go_repository(
        name="com_github_kisielk_errcheck",
        importpath="github.com/kisielk/errcheck",
        sum="h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_kisielk_gotool",
        importpath="github.com/kisielk/gotool",
        sum="h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_klauspost_compress",
        importpath="github.com/klauspost/compress",
        sum="h1:7q6vHIqubShURwQz8cQK6yIe/xC3IF0Vm7TGfqjewrc=",
        version="v1.10.5",
    )

    go_repository(
        name="com_github_kodeworks_golang_image_ico",
        importpath="github.com/Kodeworks/golang-image-ico",
        sum="h1:1ltqoej5GtaWF8jaiA49HwsZD459jqm9YFz9ZtMFpQA=",
        version="v0.0.0-20141118225523-73f0f4cfade9",
    )

    go_repository(
        name="com_github_konsorten_go_windows_terminal_sequences",
        importpath="github.com/konsorten/go-windows-terminal-sequences",
        sum="h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        version="v1.0.3",
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
        sum="h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version="v1.1.8",
    )

    go_repository(
        name="com_github_kr_text",
        importpath="github.com/kr/text",
        sum="h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version="v0.2.0",
    )

    go_repository(
        name="com_github_kyoh86_exportloopref",
        importpath="github.com/kyoh86/exportloopref",
        sum="h1:u+iHuTbkbTS2D/JP7fCuZDo/t3rBVGo3Hf58Rc+lQVY=",
        version="v0.1.7",
    )

    go_repository(
        name="com_github_lib_pq",
        importpath="github.com/lib/pq",
        sum="h1:X5PMW56eZitiTeO7tKzZxFCSpbFZJtkMMooicw2us9A=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_logrusorgru_aurora",
        importpath="github.com/logrusorgru/aurora",
        sum="h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
        version="v0.0.0-20181002194514-a7b3b318ed4e",
    )

    go_repository(
        name="com_github_lucor_goinfo",
        importpath="github.com/lucor/goinfo",
        sum="h1:4djPngMU3ttoFCf6DOgPNQYmxyNmRRmpLg4/uz2TTEg=",
        version="v0.0.0-20200401173949-526b5363a13a",
    )

    go_repository(
        name="com_github_magiconair_properties",
        importpath="github.com/magiconair/properties",
        sum="h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
        version="v1.8.1",
    )

    go_repository(
        name="com_github_maratori_testpackage",
        importpath="github.com/maratori/testpackage",
        sum="h1:QtJ5ZjqapShm0w5DosRjg0PRlSdAdlx+W6cCKoALdbQ=",
        version="v1.0.1",
    )

    go_repository(
        name="com_github_masterminds_semver",
        importpath="github.com/Masterminds/semver",
        sum="h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=",
        version="v1.5.0",
    )

    go_repository(
        name="com_github_matoous_godox",
        importpath="github.com/matoous/godox",
        sum="h1:RHba4YImhrUVQDHUCe2BNSOz4tVy2yGyXhvYDvxGgeE=",
        version="v0.0.0-20190911065817-5d6d842e92eb",
    )

    go_repository(
        name="com_github_mattn_go_colorable",
        importpath="github.com/mattn/go-colorable",
        sum="h1:bQGKb3vps/j0E9GfJQ03JyhRuxsvdAanXlT9BTw3mdw=",
        version="v0.1.7",
    )

    go_repository(
        name="com_github_mattn_go_isatty",
        importpath="github.com/mattn/go-isatty",
        sum="h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=",
        version="v0.0.12",
    )

    go_repository(
        name="com_github_mattn_go_sqlite3",
        importpath="github.com/mattn/go-sqlite3",
        sum="h1:pDRiWfl+++eC2FEFRy6jXmQlvp4Yh3z1MJKg4UeYM/4=",
        version="v1.9.0",
    )

    go_repository(
        name="com_github_mattn_goveralls",
        importpath="github.com/mattn/goveralls",
        sum="h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
        version="v0.0.2",
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
        name="com_github_mitchellh_cli",
        importpath="github.com/mitchellh/cli",
        sum="h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_mitchellh_go_homedir",
        importpath="github.com/mitchellh/go-homedir",
        sum="h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version="v1.1.0",
    )

    go_repository(
        name="com_github_mitchellh_go_ps",
        importpath="github.com/mitchellh/go-ps",
        sum="h1:i6ampVEEF4wQFF+bkYfwYgY+F/uYJDktmvLPf7qIgjc=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_mitchellh_go_testing_interface",
        importpath="github.com/mitchellh/go-testing-interface",
        sum="h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_mitchellh_gox",
        importpath="github.com/mitchellh/gox",
        sum="h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
        version="v0.4.0",
    )

    go_repository(
        name="com_github_mitchellh_iochan",
        importpath="github.com/mitchellh/iochan",
        sum="h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_mitchellh_mapstructure",
        importpath="github.com/mitchellh/mapstructure",
        sum="h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version="v1.1.2",
    )

    go_repository(
        name="com_github_modern_go_concurrent",
        importpath="github.com/modern-go/concurrent",
        sum="h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version="v0.0.0-20180306012644-bacd9c7ef1dd",
    )

    go_repository(
        name="com_github_modern_go_reflect2",
        importpath="github.com/modern-go/reflect2",
        sum="h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
        version="v1.0.1",
    )

    go_repository(
        name="com_github_mozilla_tls_observatory",
        importpath="github.com/mozilla/tls-observatory",
        sum="h1:1xJ+Xi9lYWLaaP4yB67ah0+548CD3110mCPWhVVjFkI=",
        version="v0.0.0-20200317151703-4fa42e1c2dee",
    )

    go_repository(
        name="com_github_mwitkow_go_conntrack",
        importpath="github.com/mwitkow/go-conntrack",
        sum="h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
        version="v0.0.0-20161129095857-cc309e4a2223",
    )

    go_repository(
        name="com_github_nakabonne_nestif",
        importpath="github.com/nakabonne/nestif",
        sum="h1:+yOViDGhg8ygGrmII72nV9B/zGxY188TYpfolntsaPw=",
        version="v0.3.0",
    )

    go_repository(
        name="com_github_nbutton23_zxcvbn_go",
        importpath="github.com/nbutton23/zxcvbn-go",
        sum="h1:AREM5mwr4u1ORQBMvzfzBgpsctsbQikCVpvC+tX285E=",
        version="v0.0.0-20180912185939-ae427f1e4c1d",
    )

    go_repository(
        name="com_github_nfnt_resize",
        importpath="github.com/nfnt/resize",
        sum="h1:zYyBkD/k9seD2A7fsi6Oo2LfFZAehjjQMERAvZLEDnQ=",
        version="v0.0.0-20180221191011-83c6a9932646",
    )

    go_repository(
        name="com_github_niemeyer_pretty",
        importpath="github.com/niemeyer/pretty",
        sum="h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version="v0.0.0-20200227124842-a10e7caefd8e",
    )

    go_repository(
        name="com_github_nishanths_exhaustive",
        importpath="github.com/nishanths/exhaustive",
        sum="h1:W3KBC2LFyfgd+wNudlfgCCsTo4q97MeNWrfz8/wSdSc=",
        version="v0.0.0-20200708172631-8866003e3856",
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
        name="com_github_onsi_ginkgo",
        importpath="github.com/onsi/ginkgo",
        sum="h1:Iw5WCbBcaAAd0fpRb1c9r5YCylv4XDoCSigm1zLevwU=",
        version="v1.12.0",
    )

    go_repository(
        name="com_github_onsi_gomega",
        importpath="github.com/onsi/gomega",
        sum="h1:R1uwffexN6Pr340GtYRIdZmAiN4J+iw6WG4wog1DUXg=",
        version="v1.9.0",
    )

    go_repository(
        name="com_github_openpeedeep_depguard",
        importpath="github.com/OpenPeeDeeP/depguard",
        sum="h1:VlW4R6jmBIv3/u1JNlawEvJMM4J+dPORPaZasQee8Us=",
        version="v1.0.1",
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
        name="com_github_pascaldekloe_goe",
        importpath="github.com/pascaldekloe/goe",
        sum="h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
        version="v0.0.0-20180627143212-57f6aae5913c",
    )

    go_repository(
        name="com_github_pborman_uuid",
        importpath="github.com/pborman/uuid",
        sum="h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_pelletier_go_buffruneio",
        importpath="github.com/pelletier/go-buffruneio",
        sum="h1:U4t4R6YkofJ5xHm3dJzuRpPZ0mr5MMCoAWooScCR7aA=",
        version="v0.2.0",
    )

    go_repository(
        name="com_github_pelletier_go_toml",
        importpath="github.com/pelletier/go-toml",
        sum="h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_peterh_liner",
        importpath="github.com/peterh/liner",
        sum="h1:8uaXtUkxiy+T/zdLWuxa/PG4so0TPZDZfafFNNSaptE=",
        version="v0.0.0-20170317030525-88609521dc4b",
    )

    go_repository(
        name="com_github_phayes_checkstyle",
        importpath="github.com/phayes/checkstyle",
        sum="h1:CdDQnGF8Nq9ocOS/xlSptM1N3BbrA6/kmaep5ggwaIA=",
        version="v0.0.0-20170904204023-bfd46e6a821d",
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
        name="com_github_posener_complete",
        importpath="github.com/posener/complete",
        sum="h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
        version="v1.1.1",
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
        sum="h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version="v0.0.0-20190812154241-14fe0d1b01d4",
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
        name="com_github_quasilyte_go_consistent",
        importpath="github.com/quasilyte/go-consistent",
        sum="h1:JoUA0uz9U0FVFq5p4LjEq4C0VgQ0El320s3Ms0V4eww=",
        version="v0.0.0-20190521200055-c6f3937de18c",
    )

    go_repository(
        name="com_github_quasilyte_go_ruleguard",
        importpath="github.com/quasilyte/go-ruleguard",
        sum="h1:8HyzMwa8iQE0hXsW50XzbG2JoUaRZWx412lbCP4ex2E=",
        version="v0.1.2",
    )

    go_repository(
        name="com_github_quasilyte_regex_syntax",
        importpath="github.com/quasilyte/regex/syntax",
        sum="h1:rjBjlam2Bbr6Dwp0T8HY2paibXTjMsNQU7vUH8hB+C4=",
        version="v0.0.0-20200419152657-af9db7f4a3ab",
    )

    go_repository(
        name="com_github_remyoudompheng_go_misc",
        importpath="github.com/remyoudompheng/go-misc",
        sum="h1:eTWZyPUnHcuGRDiryS/l2I7FfKjbU3IBx3IjqHPxuKU=",
        version="v0.0.0-20190427085024-2d6ac652a50e",
    )

    go_repository(
        name="com_github_rogpeppe_fastuuid",
        importpath="github.com/rogpeppe/fastuuid",
        sum="h1:gu+uRPtBe88sKxUCEXRoeCvVG90TJmwhiqRpvdhQFng=",
        version="v0.0.0-20150106093220-6724a57986af",
    )

    go_repository(
        name="com_github_rogpeppe_go_internal",
        importpath="github.com/rogpeppe/go-internal",
        sum="h1:IZRgg4sfrDH7nsAD1Y/Nwj+GzIfEwpJSLjCaNC3SbsI=",
        version="v1.6.0",
    )

    go_repository(
        name="com_github_rs_cors",
        importpath="github.com/rs/cors",
        sum="h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
        version="v1.7.0",
    )

    go_repository(
        name="com_github_russross_blackfriday",
        importpath="github.com/russross/blackfriday",
        sum="h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        version="v1.5.2",
    )

    go_repository(
        name="com_github_russross_blackfriday_v2",
        importpath="github.com/russross/blackfriday/v2",
        sum="h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
        version="v2.0.1",
    )

    go_repository(
        name="com_github_ryancurrah_gomodguard",
        importpath="github.com/ryancurrah/gomodguard",
        sum="h1:DWbye9KyMgytn8uYpuHkwf0RHqAYO6Ay/D0TbCpPtVU=",
        version="v1.1.0",
    )

    go_repository(
        name="com_github_ryanrolds_sqlclosecheck",
        importpath="github.com/ryanrolds/sqlclosecheck",
        sum="h1:AZx+Bixh8zdUBxUA1NxbxVAS78vTPq4rCb8OUZI9xFw=",
        version="v0.3.0",
    )

    go_repository(
        name="com_github_ryanuber_columnize",
        importpath="github.com/ryanuber/columnize",
        sum="h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
        version="v0.0.0-20160712163229-9b3edd62028f",
    )

    go_repository(
        name="com_github_saintfish_chardet",
        importpath="github.com/saintfish/chardet",
        sum="h1:NugYot0LIVPxTvN8n+Kvkn6TrbMyxQiuvKdEwFdR9vI=",
        version="v0.0.0-20120816061221-3af4cd4741ca",
    )

    go_repository(
        name="com_github_sean_seed",
        importpath="github.com/sean-/seed",
        sum="h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version="v0.0.0-20170313163322-e2103e2c3529",
    )

    go_repository(
        name="com_github_securego_gosec_v2",
        importpath="github.com/securego/gosec/v2",
        sum="h1:y/9mCF2WPDbSDpL3QDWZD3HHGrSYw0QSHnCqTfs4JPE=",
        version="v2.3.0",
    )

    go_repository(
        name="com_github_sergi_go_diff",
        importpath="github.com/sergi/go-diff",
        sum="h1:we8PVUC3FE2uYfodKH/nBHMSetSfHDR6scGdBi+erh0=",
        version="v1.1.0",
    )

    go_repository(
        name="com_github_shazow_go_diff",
        importpath="github.com/shazow/go-diff",
        sum="h1:W65qqJCIOVP4jpqPQ0YvHYKwcMEMVWIzWC5iNQQfBTU=",
        version="v0.0.0-20160112020656-b6b7b6733b8c",
    )

    go_repository(
        name="com_github_shirou_gopsutil",
        importpath="github.com/shirou/gopsutil",
        sum="h1:WokF3GuxBeL+n4Lk4Fa8v9mbdjlrl7bHuneF4N1bk2I=",
        version="v0.0.0-20190901111213-e4ec7b275ada",
    )

    go_repository(
        name="com_github_shirou_w32",
        importpath="github.com/shirou/w32",
        sum="h1:udFKJ0aHUL60LboW/A+DfgoHVedieIzIXE8uylPue0U=",
        version="v0.0.0-20160930032740-bb4de0191aa4",
    )

    go_repository(
        name="com_github_shurcool_go",
        importpath="github.com/shurcooL/go",
        sum="h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
        version="v0.0.0-20180423040247-9e1955d9fb6e",
    )

    go_repository(
        name="com_github_shurcool_go_goon",
        importpath="github.com/shurcooL/go-goon",
        sum="h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
        version="v0.0.0-20170922171312-37c2f522c041",
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
        sum="h1:UBcNElsrwanuuMsnGSlYmtmgbb23qDR5dG+6X6Oo89I=",
        version="v1.6.0",
    )

    go_repository(
        name="com_github_skratchdot_open_golang",
        importpath="github.com/skratchdot/open-golang",
        sum="h1:JIAuq3EEf9cgbU6AtGPK4CTG3Zf6CKMNqf0MHTggAUA=",
        version="v0.0.0-20200116055534-eef842397966",
    )

    go_repository(
        name="com_github_smartystreets_assertions",
        importpath="github.com/smartystreets/assertions",
        sum="h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
        version="v0.0.0-20180927180507-b2de0cb4f26d",
    )

    go_repository(
        name="com_github_smartystreets_goconvey",
        importpath="github.com/smartystreets/goconvey",
        sum="h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version="v1.6.4",
    )

    go_repository(
        name="com_github_soheilhy_cmux",
        importpath="github.com/soheilhy/cmux",
        sum="h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
        version="v0.1.4",
    )

    go_repository(
        name="com_github_sonatard_noctx",
        importpath="github.com/sonatard/noctx",
        sum="h1:VC1Qhl6Oxx9vvWo3UDgrGXYCeKCe3Wbw7qAWL6FrmTY=",
        version="v0.0.1",
    )

    go_repository(
        name="com_github_sourcegraph_go_diff",
        importpath="github.com/sourcegraph/go-diff",
        sum="h1:lhIKJ2nXLZZ+AfbHpYxTn0pXpNTTui0DX7DO3xeb1Zs=",
        version="v0.5.3",
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
        sum="h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version="v1.0.5",
    )

    go_repository(
        name="com_github_spf13_viper",
        importpath="github.com/spf13/viper",
        sum="h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
        version="v1.7.0",
    )

    go_repository(
        name="com_github_src_d_gcfg",
        importpath="github.com/src-d/gcfg",
        sum="h1:xXbNR5AlLSA315x2UO+fTSSAXCDf+Ar38/6oyGbDKQ4=",
        version="v1.4.0",
    )

    go_repository(
        name="com_github_srwiley_oksvg",
        importpath="github.com/srwiley/oksvg",
        sum="h1:HunZiaEKNGVdhTRQOVpMmj5MQnGnv+e8uZNu3xFLgyM=",
        version="v0.0.0-20200311192757-870daf9aa564",
    )

    go_repository(
        name="com_github_srwiley_rasterx",
        importpath="github.com/srwiley/rasterx",
        sum="h1:m59mIOBO4kfcNCEzJNy71UkeF4XIx2EVmL9KLwDQdmM=",
        version="v0.0.0-20200120212402-85cb7272f5e9",
    )

    go_repository(
        name="com_github_stackexchange_wmi",
        importpath="github.com/StackExchange/wmi",
        sum="h1:fLjPD/aNc3UIOA6tDi6QXUemppXK3P9BI7mr2hd6gx8=",
        version="v0.0.0-20180116203802-5d049714c4a6",
    )

    go_repository(
        name="com_github_stretchr_objx",
        importpath="github.com/stretchr/objx",
        sum="h1:Hbg2NidpLE8veEBkEZTL3CvlkUIVzuU9jDplZO54c48=",
        version="v0.2.0",
    )

    go_repository(
        name="com_github_stretchr_testify",
        importpath="github.com/stretchr/testify",
        sum="h1:hDPOHmpOpP40lSULcqw7IrRb/u7w6RpDC9399XyoNd0=",
        version="v1.6.1",
    )

    go_repository(
        name="com_github_subosito_gotenv",
        importpath="github.com/subosito/gotenv",
        sum="h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
        version="v1.2.0",
    )

    go_repository(
        name="com_github_tdakkota_asciicheck",
        importpath="github.com/tdakkota/asciicheck",
        sum="h1:Xr9gkxfOP0KQWXKNqmwe8vEeSUiUj4Rlee9CMVX2ZUQ=",
        version="v0.0.0-20200416190851-d7f85be797a2",
    )

    go_repository(
        name="com_github_tebeka_selenium",
        importpath="github.com/tebeka/selenium",
        sum="h1:cNziB+etNgyH/7KlNI7RMC1ua5aH1+5wUlFQyzeMh+w=",
        version="v0.9.9",
    )

    go_repository(
        name="com_github_temoto_robotstxt",
        importpath="github.com/temoto/robotstxt",
        sum="h1:Gh8RCs8ouX3hRSxxK7B1mO5RFByQ4CmJZDwgom++JaA=",
        version="v1.1.1",
    )

    go_repository(
        name="com_github_tetafro_godot",
        importpath="github.com/tetafro/godot",
        sum="h1:Dib7un+rYJFUi8vN0Bk6EHheKy6fv6ZzFURHw75g6m8=",
        version="v0.4.2",
    )

    go_repository(
        name="com_github_timakin_bodyclose",
        importpath="github.com/timakin/bodyclose",
        sum="h1:RumXZ56IrCj4CL+g1b9OL/oH0QnsF976bC8xQFYUD5Q=",
        version="v0.0.0-20190930140734-f7f2e9bca95e",
    )

    go_repository(
        name="com_github_tmc_grpc_websocket_proxy",
        importpath="github.com/tmc/grpc-websocket-proxy",
        sum="h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
        version="v0.0.0-20190109142713-0ad062ec5ee5",
    )

    go_repository(
        name="com_github_tommy_muehle_go_mnd",
        importpath="github.com/tommy-muehle/go-mnd",
        sum="h1:RC4maTWLKKwb7p1cnoygsbKIgNlJqSYBeAFON3Ar8As=",
        version="v1.3.1-0.20200224220436-e6f9a994e8fa",
    )

    go_repository(
        name="com_github_ugorji_go",
        importpath="github.com/ugorji/go",
        sum="h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
        version="v1.1.4",
    )

    go_repository(
        name="com_github_ultraware_funlen",
        importpath="github.com/ultraware/funlen",
        sum="h1:Av96YVBwwNSe4MLR7iI/BIa3VyI7/djnto/pK3Uxbdo=",
        version="v0.0.2",
    )

    go_repository(
        name="com_github_ultraware_whitespace",
        importpath="github.com/ultraware/whitespace",
        sum="h1:If7Va4cM03mpgrNH9k49/VOicWpGoG70XPBFFODYDsg=",
        version="v0.0.4",
    )

    go_repository(
        name="com_github_uudashr_gocognit",
        importpath="github.com/uudashr/gocognit",
        sum="h1:MoG2fZ0b/Eo7NXoIwCVFLG5JED3qgQz5/NEE+rOsjPs=",
        version="v1.0.1",
    )

    go_repository(
        name="com_github_valyala_bytebufferpool",
        importpath="github.com/valyala/bytebufferpool",
        sum="h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version="v1.0.0",
    )

    go_repository(
        name="com_github_valyala_fasthttp",
        importpath="github.com/valyala/fasthttp",
        sum="h1:TsB9qkSeiMXB40ELWWSRMjlsE+8IkqXHcs01y2d9aw0=",
        version="v1.12.0",
    )

    go_repository(
        name="com_github_valyala_quicktemplate",
        importpath="github.com/valyala/quicktemplate",
        sum="h1:JtrOtlRMP+pySF9dkPBfVfhZ3YHTitAJpMLNqT9ZjFk=",
        version="v1.5.1",
    )

    go_repository(
        name="com_github_valyala_tcplisten",
        importpath="github.com/valyala/tcplisten",
        sum="h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
        version="v0.0.0-20161114210144-ceec8f93295a",
    )

    go_repository(
        name="com_github_vektra_mockery",
        importpath="github.com/vektra/mockery",
        sum="h1:uc0Yn67rJpjt8U/mAZimdCKn9AeA97BOkjpmtBSlfP4=",
        version="v1.1.2",
    )

    go_repository(
        name="com_github_x1ddos_csslex",
        importpath="github.com/x1ddos/csslex",
        sum="h1:SX7lFdwn40ahL78CxofAh548P+dcWjdRNpirU7+sKiE=",
        version="v0.0.0-20160125172232-7894d8ab8bfe",
    )

    go_repository(
        name="com_github_xanzy_ssh_agent",
        importpath="github.com/xanzy/ssh-agent",
        sum="h1:TCbipTQL2JiiCprBWx9frJ2eJlCYT00NmctrHxVAr70=",
        version="v0.2.1",
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
        sum="h1:5tjfNdR2ki3yYQ842+eX2sQHeiwpKJ0RnHO4IYOc4V8=",
        version="v1.1.32",
    )

    go_repository(
        name="com_google_cloud_go",
        importpath="cloud.google.com/go",
        sum="h1:AVXDdKsrtX33oR9fbCMu/+c1o8Ofjq6Ku/MInaLVg5Y=",
        version="v0.46.3",
    )

    go_repository(
        name="com_google_cloud_go_bigquery",
        importpath="cloud.google.com/go/bigquery",
        sum="h1:hL+ycaJpVE9M7nLoiXb/Pn10ENE2u+oddxbD8uu0ZVU=",
        version="v1.0.1",
    )

    go_repository(
        name="com_google_cloud_go_datastore",
        importpath="cloud.google.com/go/datastore",
        sum="h1:Kt+gOPPp2LEPWp8CSfxhsM8ik9CcyE/gYu+0r+RnZvM=",
        version="v1.0.0",
    )

    go_repository(
        name="com_google_cloud_go_firestore",
        importpath="cloud.google.com/go/firestore",
        sum="h1:9x7Bx0A9R5/M9jibeJeZWqjeVEIxYW9fZYqB9a70/bY=",
        version="v1.1.0",
    )

    go_repository(
        name="com_google_cloud_go_pubsub",
        importpath="cloud.google.com/go/pubsub",
        sum="h1:W9tAK3E57P75u0XLLR82LZyw8VpAnhmyTOxW9qzmyj8=",
        version="v1.0.1",
    )

    go_repository(
        name="com_google_cloud_go_storage",
        importpath="cloud.google.com/go/storage",
        sum="h1:VV2nUM3wwLLGh9lSABFgZMjInyUbJeaRSE64WuAIQ+4=",
        version="v1.0.0",
    )

    go_repository(
        name="com_shuralyov_dmitri_gpu_mtl",
        importpath="dmitri.shuralyov.com/gpu/mtl",
        sum="h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version="v0.0.0-20190408044501-666a987793e9",
    )

    go_repository(
        name="com_sourcegraph_sqs_pbtypes",
        importpath="sourcegraph.com/sqs/pbtypes",
        sum="h1:JPJh2pk3+X4lXAkZIk2RuE/7/FoK9maXw+TNPJhVS/c=",
        version="v0.0.0-20180604144634-d3ebe8f20ae4",
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
        sum="h1:BLraFXnmrev5lT+xlilqcH8XK9/i0At2xKjWk4p6zsU=",
        version="v1.0.0-20200227125254-8fa46927fb4f",
    )

    go_repository(
        name="in_gopkg_errgo_v2",
        importpath="gopkg.in/errgo.v2",
        sum="h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version="v2.1.0",
    )

    go_repository(
        name="in_gopkg_fsnotify_v1",
        importpath="gopkg.in/fsnotify.v1",
        sum="h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version="v1.4.7",
    )

    go_repository(
        name="in_gopkg_ini_v1",
        importpath="gopkg.in/ini.v1",
        sum="h1:AQvPpx3LzTDM0AjnIRlVFwFFGC+npRopjZxLJj6gdno=",
        version="v1.51.0",
    )

    go_repository(
        name="in_gopkg_resty_v1",
        importpath="gopkg.in/resty.v1",
        sum="h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version="v1.12.0",
    )

    go_repository(
        name="in_gopkg_src_d_go_billy_v4",
        importpath="gopkg.in/src-d/go-billy.v4",
        sum="h1:0SQA1pRztfTFx2miS8sA97XvooFeNOmvUenF4o0EcVg=",
        version="v4.3.2",
    )

    go_repository(
        name="in_gopkg_src_d_go_git_fixtures_v3",
        importpath="gopkg.in/src-d/go-git-fixtures.v3",
        sum="h1:ivZFOIltbce2Mo8IjzUHAFoq/IylO9WHhNOAJK+LsJg=",
        version="v3.5.0",
    )

    go_repository(
        name="in_gopkg_src_d_go_git_v4",
        importpath="gopkg.in/src-d/go-git.v4",
        sum="h1:SRtFyV8Kxc0UP7aCHcijOMQGPxHSmMOPrzulQWolkYE=",
        version="v4.13.1",
    )

    go_repository(
        name="in_gopkg_tomb_v1",
        importpath="gopkg.in/tomb.v1",
        sum="h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version="v1.0.0-20141024135613-dd632973f1e7",
    )

    go_repository(
        name="in_gopkg_warnings_v0",
        importpath="gopkg.in/warnings.v0",
        sum="h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
        version="v0.1.2",
    )

    go_repository(
        name="in_gopkg_yaml_v2",
        importpath="gopkg.in/yaml.v2",
        sum="h1:clyUAQHOM3G0M3f5vQj7LuJrETvjVot3Z5el9nffUtU=",
        version="v2.3.0",
    )

    go_repository(
        name="in_gopkg_yaml_v3",
        importpath="gopkg.in/yaml.v3",
        sum="h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
        version="v3.0.0-20200313102051-9f266ea9e77c",
    )

    go_repository(
        name="io_etcd_go_bbolt",
        importpath="go.etcd.io/bbolt",
        sum="h1:Z/90sZLPOeCy2PwprqkFa25PdkusRzaj9P8zm/KNyvk=",
        version="v1.3.2",
    )

    go_repository(
        name="io_fyne_fyne",
        importpath="fyne.io/fyne",
        sum="h1:9MVvZeYeGKJa0I6PKsf0LlZfeurYpsQMEpTaM2O6UpA=",
        version="v1.3.2",
    )

    go_repository(
        name="io_k8s_klog",
        importpath="k8s.io/klog",
        sum="h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version="v1.0.0",
    )

    go_repository(
        name="io_k8s_klog_v2",
        importpath="k8s.io/klog/v2",
        sum="h1:XRvcwJozkgZ1UQJmfMGpvRthQHOvihEhYtDfAaxMz/A=",
        version="v2.2.0",
    )

    go_repository(
        name="io_opencensus_go",
        importpath="go.opencensus.io",
        sum="h1:C9hSCOW830chIVkdja34wa6Ky+IzWllkUinR+BtRZd4=",
        version="v0.22.0",
    )

    go_repository(
        name="io_rsc_binaryregexp",
        importpath="rsc.io/binaryregexp",
        sum="h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version="v0.2.0",
    )

    go_repository(
        name="io_rsc_pdf",
        importpath="rsc.io/pdf",
        sum="h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version="v0.1.1",
    )

    go_repository(
        name="net_starlark_go",
        importpath="go.starlark.net",
        sum="h1:lkYv5AKwvvduv5XWP6szk/bvvgO6aDeUujhZQXIFTes=",
        version="v0.0.0-20190702223751-32f345186213",
    )

    go_repository(
        name="org_golang_dl",
        importpath="golang.org/dl",
        sum="h1:qQYtBoXnj57L4UinmXmIC0LiTfUayhXi0ZIUSUtp//E=",
        version="v0.0.0-20200717001634-fec1f33ee347",
    )

    go_repository(
        name="org_golang_google_api",
        importpath="google.golang.org/api",
        sum="h1:Q3Ui3V3/CVinFWFiW39Iw0kMuVrRzYX0wN6OPFp0lTA=",
        version="v0.13.0",
    )

    go_repository(
        name="org_golang_google_appengine",
        importpath="google.golang.org/appengine",
        sum="h1:tycE03LOZYQNhDpS27tcQdAzLCVMaj7QT2SXxebnpCM=",
        version="v1.6.5",
    )

    go_repository(
        name="org_golang_google_genproto",
        importpath="google.golang.org/genproto",
        sum="h1:KoSCpgvphmtnpVycBCtfV9hNdHbInsqdx4gqnDPQCkg=",
        version="v0.0.0-20200720141249-1244ee217b7e",
    )

    go_repository(
        name="org_golang_google_grpc",
        importpath="google.golang.org/grpc",
        sum="h1:rRYRFMVgRv6E0D70Skyfsr28tDXIuuPZyWGMPdMcnXg=",
        version="v1.27.0",
    )
    go_repository(
        name="org_golang_google_protobuf",
        build_file_generation="on",
        build_file_proto_mode="disable",
        importpath="google.golang.org/protobuf",
        sum="h1:UhZDfRO8JRQru4/+LlLE0BRKGF8L+PICnvYZmx/fEGA=",
        version="v1.24.0",
    )

    go_repository(
        name="org_golang_x_arch",
        importpath="golang.org/x/arch",
        sum="h1:QlVATYS7JBoZMVaf+cNjb90WD/beKVHnIxFKT4QaHVI=",
        version="v0.0.0-20190927153633-4e8777c89be4",
    )

    go_repository(
        name="org_golang_x_crypto",
        importpath="golang.org/x/crypto",
        sum="h1:DZhuSZLsGlFL4CmhA8BcRA0mnthyA/nZ00AqCUo7vHg=",
        version="v0.0.0-20200709230013-948cd5f35899",
    )

    go_repository(
        name="org_golang_x_exp",
        importpath="golang.org/x/exp",
        sum="h1:A1gGSx58LAGVHUUsOf7IiR0u8Xb6W51gRwfDBhkdcaw=",
        version="v0.0.0-20191030013958-a1ab85dbe136",
    )

    go_repository(
        name="org_golang_x_image",
        importpath="golang.org/x/image",
        sum="h1:6WW6V3x1P/jokJBpRQYUJnMHRP6isStQwCozxnU7XQw=",
        version="v0.0.0-20200430140353-33d19683fad8",
    )

    go_repository(
        name="org_golang_x_lint",
        importpath="golang.org/x/lint",
        sum="h1:5hukYrvBGR8/eNkX5mdUezrA6JiaEZDtJb9Ei+1LlBs=",
        version="v0.0.0-20190930215403-16217165b5de",
    )

    go_repository(
        name="org_golang_x_mobile",
        importpath="golang.org/x/mobile",
        sum="h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version="v0.0.0-20190719004257-d2bd2a29d028",
    )

    go_repository(
        name="org_golang_x_mod",
        importpath="golang.org/x/mod",
        sum="h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
        version="v0.3.0",
    )

    go_repository(
        name="org_golang_x_net",
        importpath="golang.org/x/net",
        sum="h1:vGXIOMxbNfDTk/aXCmfdLgkrSV+Z2tcbze+pEc3v5W4=",
        version="v0.0.0-20200625001655-4c5254603344",
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
        sum="h1:qwRHBd0NqMbJxfbotnDhm2ByMI1Shq4Y6oRJo21SGJA=",
        version="v0.0.0-20200625203802-6e8e738ad208",
    )

    go_repository(
        name="org_golang_x_sys",
        importpath="golang.org/x/sys",
        sum="h1:TC0v2RSO1u2kn1ZugjrFXkRZAEaqMN/RW+OTZkBzmLE=",
        version="v0.0.0-20200327173247-9dae0f8f5775",
    )

    go_repository(
        name="org_golang_x_text",
        importpath="golang.org/x/text",
        sum="h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
        version="v0.3.2",
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
        sum="h1:rD1FcWVsRaMY+l8biE9jbWP5MS/CJJ/90a9TMkMgNrM=",
        version="v0.0.0-20200710042808-f1c4188a97a1",
    )

    go_repository(
        name="org_golang_x_tools_gopls",
        importpath="golang.org/x/tools/gopls",
        sum="h1:irz7Q+XdHNECamFKbNWKvMV2Ak6zBbwdwbZndG4545I=",
        version="v0.4.3",
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
