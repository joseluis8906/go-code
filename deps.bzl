load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:ytYb4rOqyp1TSa2EPvNVwtPQJctSELKaMyLfqNP4+34=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_template",
        importpath = "github.com/alecthomas/template",
        sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
        version = "v0.0.0-20190718012654-fb15b899a751",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_anmitsu_go_shlex",
        importpath = "github.com/anmitsu/go-shlex",
        sum = "h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
        version = "v0.0.0-20161002113705-648efa622239",
    )
    go_repository(
        name = "com_github_antihax_optional",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:hR91U9KYmb6bLBYLQjyM+3j+rcd/UhE+G78SFnF8gJA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        importpath = "github.com/armon/go-radix",
        sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:ip6w0uFQkncKQ979AypyG0ER7mqUSBdKLOgAle/AT8A=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_bmizerany_perks",
        importpath = "github.com/bmizerany/perks",
        sum = "h1:AP/Y7sqYicnjGDfD5VcY4CIfh1hRXBUavxrvELjTiOE=",
        version = "v0.0.0-20141205001514-d9a9656a3a4b",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:GaI7EiDXDRfa8VshkTj7Fym7ha+y8/XxIgD2okUIjLw=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cactus_go_statsd_client_statsd",
        importpath = "github.com/cactus/go-statsd-client/statsd",
        sum = "h1:bXxS5/Z3/dfc8iFniQfgogNBomo0u+1//9eP+jl8GVo=",
        version = "v0.0.0-20191106001114-12b4e2b38748",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
        version = "v4.2.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_circonus_labs_circonus_gometrics",
        importpath = "github.com/circonus-labs/circonus-gometrics",
        sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
        version = "v2.3.1+incompatible",
    )
    go_repository(
        name = "com_github_circonus_labs_circonusllhist",
        importpath = "github.com/circonus-labs/circonusllhist",
        sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:7To3pQ+pZo0i3dsWEbinPNFs5gPSBOsJtx3wTT94VBY=",
        version = "v0.0.0-20231109132714-523115ebc101",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )
    go_repository(
        name = "com_github_crossdock_crossdock_go",
        importpath = "github.com/crossdock/crossdock-go",
        sum = "h1:WR1qVJzbvrVywhAk4kMQKRPx09AZVI0NdEdYs59iHcA=",
        version = "v0.0.0-20160816171116-049aabb0122b",
    )
    go_repository(
        name = "com_github_datadog_datadog_go",
        importpath = "github.com/DataDog/datadog-go",
        sum = "h1:qSG2N4FghB1He/r2mFrWKCaL7dXCilEuNEeAn20fdD4=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_farm",
        importpath = "github.com/dgryski/go-farm",
        sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
        version = "v0.0.0-20200201041132-a6ae2369ad13",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
        version = "v0.11.1",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:QkIBuU5k+x7/QXPvPPnWXWlCdaBFApVqftFV6k087DA=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_evalphobia_logrus_fluent",
        importpath = "github.com/evalphobia/logrus_fluent",
        sum = "h1:G4BSBTm7+L+oanWfFtA/A5Y3pvL2OMxviczyZPYO5xc=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:qfhVLaG5s+nCROl1zJsZRxFeYrHLqWroPOQ8BWiNb4w=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_fatih_structtag",
        importpath = "github.com/fatih/structtag",
        sum = "h1:/OdNE99OxoI/PqaW/SuSK9uxxT3f/tcSZgon/ssNSx4=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_fluent_fluent_logger_golang",
        importpath = "github.com/fluent/fluent-logger-golang",
        sum = "h1:zUdY44CHX2oIUc7VTNZc+4m+ORuO/mldQDA7czhWXEg=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_flynn_go_shlex",
        importpath = "github.com/flynn/go-shlex",
        sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
        version = "v0.0.0-20150515145356-3f9db97f8568",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:7Xjx+VpznH+oBnejlPUj8oUpdxnVs4f8XU8WnHkI4W8=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_gdamore_encoding",
        importpath = "github.com/gdamore/encoding",
        sum = "h1:+7OoQ1Bc6eTm5niUzBa0Ctsh6JbMW6Ra+YNuAtDBdko=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gdamore_tcell_v2",
        importpath = "github.com/gdamore/tcell/v2",
        sum = "h1:I5LiGTQuwrysAt1KS9wg1yFfOI3arI3ucFrxtd/xqaA=",
        version = "v2.7.0",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
        version = "v5.0.4",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:kX1es4djPJrsDhY7aZKJy7aZasdcB5oSOEphMjSB53c=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_gogo_status",
        importpath = "github.com/gogo/status",
        sum = "h1:+eIkrewn5q6b30y+g/BJINVVdi2xH7je5MPJ3ZPK3JA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:DVjP2PbBOzHyzA+dn3WhHIq4NdVu3Q+pvivFICf/7fo=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:Rd1kQnQu0Hq3qvJppYSG0HtP+f5LPPUiDswTLiEegLg=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:OF1IPgv+F4NmqmJ98KTjdN97Vs1JxDPB3vbmYzV2dpk=",
        version = "v0.2.1-0.20230907215043-c6f79328ddf9",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:IqNFLAmvJOgVlpdEBiQbDc2EwKW77amAycfTuWKdfvw=",
        version = "v3.3.2",
    )
    go_repository(
        name = "com_github_google_renameio",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:1p67kYwdtXjb0gL0BPiP1Av9wiZPo5A8z2cWkTZ+eyU=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:A+gCJKdRfqXkr+BIRGtZLibNXf0m1f9E4HG56etFpas=",
        version = "v2.12.0",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:zC34cGQu69FG7qzJ3WiKW244WfhDC3xxYMeNOX2gtUQ=",
        version = "v0.0.0-20210719221736-1c9a4c676720",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:Wqo399gCIufwto+VfwCSvsnfGpF/w5E9CNxSwbpD6No=",
        version = "v2.19.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:CqrdhYzc8XZuPnhIYZWH45toM0LB9ZeYr/gvpLVI3PE=",
        version = "v1.25.1",
    )
    go_repository(
        name = "com_github_hashicorp_consul_sdk",
        importpath = "github.com/hashicorp/consul/sdk",
        sum = "h1:ZiwE2bKb+zro68sWzZ1SgHF3kRMBZ94TwOCFRF4ylPs=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:bI2ocEMgcVlz55Oj1xZNBsVi900c7II+fWDyV9o+13c=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_retryablehttp",
        importpath = "github.com/hashicorp/go-retryablehttp",
        sum = "h1:QlWt0KvWT0lq8MFppF9tsJGF+ynG7ztc2KIPhzRGk7s=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_sockaddr",
        importpath = "github.com/hashicorp/go-sockaddr",
        sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_syslog",
        importpath = "github.com/hashicorp/go-syslog",
        sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_version",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:zEfKbn2+PDgroKdiOzqiE8rsmLqU2uwi5PB5pBJ3TkI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_logutils",
        importpath = "github.com/hashicorp/logutils",
        sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_mdns",
        importpath = "github.com/hashicorp/mdns",
        sum = "h1:sY0CMhFmjIPDMlTB+HfymFHCaYLhgifZ0QhjaYKD/UQ=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_hashicorp_memberlist",
        importpath = "github.com/hashicorp/memberlist",
        sum = "h1:EtYPN8DpAURiapus508I4n9CzHs2W+8NZGbmmR/prTM=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:Z1H2J60yRKvfDYAOZLd2MU0ND4AH/WDz7xYHDWQsIPY=",
        version = "v0.10.1",
    )
    go_repository(
        name = "com_github_hdrhistogram_hdrhistogram_go",
        importpath = "github.com/HdrHistogram/hdrhistogram-go",
        sum = "h1:dpujRju0R4M/QZzcnR1LH1qm+TVG3UzkWdp5tH1WMcg=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_jessevdk_go_flags",
        importpath = "github.com/jessevdk/go-flags",
        sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:RlWWUY/Dr4fL8qk9YG7DTZ7PDgME2V4csBXA8L/ixi4=",
        version = "v1.17.2",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:DB17ag19krx9CFsz4o3enTrPXyIXCl+2iCXH/aMAp9s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_lucasb_eyer_go_colorful",
        importpath = "github.com/lucasb-eyer/go-colorful",
        sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
        version = "v0.0.19",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=",
        version = "v0.0.15",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        importpath = "github.com/mattn/go-shellwords",
        sum = "h1:Y7Xqm8piKOO3v10Thp7Z36h4FYFjt5xB//6XvOrs2Gw=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:I0XW9+e1XWDxdcEniV4rQAIOPUGDq67JSCiRCgGCZLI=",
        version = "v1.0.2-0.20181231171920-c182affec369",
    )
    go_repository(
        name = "com_github_miekg_dns",
        importpath = "github.com/miekg/dns",
        sum = "h1:WMszZWJG0XmzbK9FEmzH2TVcqYzFesusSIB41b8KHxY=",
        version = "v1.1.41",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:tEElEatulEHDeedTxwckzyYMA5c86fbmNIUL1hBIiTg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:iruDEfMl2E6fbMZ9s0scYfZQ84/6SPL6zC8ACM2oIL0=",
        version = "v0.0.0-20171201202039-1bf9dbcd8cbe",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:Bx9BZS+aXYlxW08k8Gd3yR2s73pV5XSoAQUyp1Kwvp0=",
        version = "v1.32.0",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:RwNJbbIdYCoClSDNY7QVKZlyb/wfT6ugvFCiKy6vDvI=",
        version = "v0.4.7",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_pascaldekloe_goe",
        importpath = "github.com/pascaldekloe/goe",
        sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:FnwAJ4oYMvbT/34k9zzHuZNrhlz48GB3/s6at6/MHO4=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:bnDivRJ1EWPjUIRXV5KfORO897HTbpFAQddBdE8t7Gw=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_pkg_errors",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:JFZT4XbOU7l77xGSpOdW+pwIMqP044IyjXX6FGyEKFo=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_github_posener_complete",
        importpath = "github.com/posener/complete",
        sum = "h1:NP0eAhjcjImqslEwo/1hq7gpajME0fTLTezBKDqfXqo=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_prashantv_protectmem",
        importpath = "github.com/prashantv/protectmem",
        sum = "h1:AA9vgIBDjMHPC2McaGPojgV2dcI78ZC0TLNhYCXEKH8=",
        version = "v0.0.0-20171002184600-e20412882b3a",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:ygXvpU1AoN1MhdzckN+PyD9QJOSD4x7kmXYlnfbA6JU=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:VQw1hfvPvk3Uv6Qf29VrPF32JB6rtbgI6cYPYQjL0Qw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:QO8U2CdOzSn1BBsmXJXduaaW+dY/5QLjfB8svtSzKKE=",
        version = "v0.48.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:jluTpSng7V9hY0O2R9DzzJHYb2xULk9VTR1V1R/k6Bo=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:utMvzDsuh3suAEnhH0RdHmoPbU648o6CvXxTx4SBMOw=",
        version = "v0.4.3",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:mKX4bl4iPYJtEIxp6CYiUuLQ/8DYMoz0PUdtGgMFRVc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        importpath = "github.com/rs/zerolog",
        sum = "h1:FcTR3NnLWW+NnTwwhFWiJSZr4ECLpqCm6QsEnyvbV4A=",
        version = "v1.31.0",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
        version = "v0.0.0-20160712163229-9b3edd62028f",
    )
    go_repository(
        name = "com_github_sagikazarmark_crypt",
        importpath = "github.com/sagikazarmark/crypt",
        sum = "h1:ZA/7pXyjkHoK4bW4mIdnCLvL8hd+Nrbiw7Dqk7D4qUk=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_sagikazarmark_locafero",
        importpath = "github.com/sagikazarmark/locafero",
        sum = "h1:HApY1R9zGo4DBgr7dqsTH/JJxLTTsOt7u6keLGt6kNQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_sagikazarmark_slog_shim",
        importpath = "github.com/sagikazarmark/slog-shim",
        sum = "h1:diDBnUNK9N/354PgrxMywXnAwEr1QZcOr6gto+ugjYE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_samber_lo",
        importpath = "github.com/samber/lo",
        sum = "h1:j2XEAqXKb09Am4ebOg31SpvzUTTs6EN3VfgeLUhPdXM=",
        version = "v1.38.1",
    )
    go_repository(
        name = "com_github_samber_slog_fluentd",
        importpath = "github.com/samber/slog-fluentd",
        sum = "h1:mQYVCwiI5kzlRDyJJVwPAdNmw2FooTRYAQ7R8OBQQao=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_samuel_go_thrift",
        importpath = "github.com/samuel/go-thrift",
        sum = "h1:EiWVfh8mr40yFZEui2oF0d45KgH48PkB2H0Z0GANvSI=",
        version = "v0.0.0-20191111193933-5165175b40af",
    )
    go_repository(
        name = "com_github_sean_seed",
        importpath = "github.com/sean-/seed",
        sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version = "v0.0.0-20170313163322-e2103e2c3529",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_sourcegraph_conc",
        importpath = "github.com/sourcegraph/conc",
        sum = "h1:OQTbbt6P72L20UqAkXXuLOj79LfEanQ+YQFNpLA9ySo=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:WJQKhtpdm3v2IzqG8VMqrr6Rf3UYpEF239Jy9wNepM8=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:GEiTHELF+vaR5dhz3VqZfFSzZjYbgeKDpBxQVS4GYJ0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:LUXCnvUvSM6FXAsj6nnfc8Q2tp1dIgUfY9Kc8GsSOiQ=",
        version = "v1.18.2",
    )
    go_repository(
        name = "com_github_streadway_quantile",
        importpath = "github.com/streadway/quantile",
        sum = "h1:X4+kt6zM/OVO6gbJdAfJR60MGPsqCzbtXNnjoGqdfAs=",
        version = "v0.0.0-20220407130108-4246515d968d",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:9NlTDc1FTs4qu0DDq7AEtTPNw6SVm7uBMsUCUjABIf8=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:SHf3yoO2sGA0veCJeCBYLHuttAVFHGm2RHgNodW7wQU=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_tv42_httpunix",
        importpath = "github.com/tv42/httpunix",
        sum = "h1:G3dpKMzFDjgEh2q1Z7zUUtKa8ViPtH+ocF0bE0g00O8=",
        version = "v0.0.0-20150427012821-b75d8614f926",
    )
    go_repository(
        name = "com_github_uber_common_bark",
        importpath = "github.com/uber-common/bark",
        sum = "h1:cREJ9b7CpTjwZr0/5wV82fXlitoCIEHHnt9WkQ4lIk0=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_uber_go_mapdecode",
        importpath = "github.com/uber-go/mapdecode",
        sum = "h1:euUEFM9KnuCa1OBixz1xM+FIXmpixyay5DLymceOVrU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_uber_go_tally",
        importpath = "github.com/uber-go/tally",
        sum = "h1:9hLSgNBP28CjIaDmAuRTq9qV+UZY+9PcvAkXO4nNMwg=",
        version = "v3.3.15+incompatible",
    )
    go_repository(
        name = "com_github_uber_jaeger_client_go",
        importpath = "github.com/uber/jaeger-client-go",
        sum = "h1:NHcubEkVbahf9t3p75TOCR83gdUHXjRJvjoBh1yACsM=",
        version = "v2.22.1+incompatible",
    )
    go_repository(
        name = "com_github_uber_jaeger_lib",
        importpath = "github.com/uber/jaeger-lib",
        sum = "h1:td4jdvLcExb4cBISKIpHuGoVXh+dVKhn2Um6rjCsSsg=",
        version = "v2.4.1+incompatible",
    )
    go_repository(
        name = "com_github_uber_ringpop_go",
        importpath = "github.com/uber/ringpop-go",
        sum = "h1:aBa/SHmmFRcAXA63k7uBheoTL8tCmH7L+OgktB1AF/o=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_github_uber_tchannel_go",
        importpath = "github.com/uber/tchannel-go",
        sum = "h1:jq5HdA35SqXeRpSFmfLFARanNLvKeku0om4g2LZDbm0=",
        version = "v1.33.0",
    )
    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:FHX5I5B4i4hKRVRBCFRxq1iQRej7WO3hhBuJf+UUySY=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:XLI/Ng3O1Atzq0oBs3TWm+5ZVgkq2aqdlvP9JtoZ6c8=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
        version = "v0.0.0-20181117223130-1be2e3e5546d",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:YHLKNupSD1KqjDbQ3+LVdQ81h/UJbJyZG203cEfnQgM=",
        version = "v0.111.0",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:ZvLvJ952zK8pFHINjpMBY5k7LTAp/6pBf50RDMRgBUI=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:Yo4g2XrBETBCqyWIibN3NHNPQKUfQqti0lI+70rubeE=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:xyCAfpI4yUMOQ4VtHN/bdmxPQ8xoEkTwFM1nbVmuQhs=",
        version = "v1.58.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:fnV7B8lqyEYxCU0LKk+vUL7mTlqRAq4uFlIthIdr/iA=",
        version = "v0.21.6",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:VVIxCtVerchHienSlaGzV6XJGtEM9828Erzyr3miUGs=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:jSoGITWKgAj/ssVogNE9SdsTqcXnryPzsulENSRlusI=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:DSaD1iiqvELag+lV4VnnqUUFd8GXELu01tKVdWZrviE=",
        version = "v0.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:Qub3fqR7iA1daJWdzjp/Q0Jz0fUG0JbMc7Ui4E9IX/E=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:YnSO8m02pOIo6AEOgiOoUDVbw4pf+bg2KLHi4rky320=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:/hQaadYytMdA5zBh+RciIrXZQBWK4vN7EUsrQHG+/t8=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:VjwWNtEVsbpXfJqZbb2RLOBzSgAjN69vf2UJADHnkxk=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:FsLSkmYYeNuzDm8L4YPfLWV+lQaUrJmH5OuD37t1k20=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:i9tOKXX+1gE7+rHpWKjiuPfGBVIYoWvLNIGpWgPtF58=",
        version = "v1.13.4",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:oQiFYYCe0vwp7J8ZmF6siVKEumWtiPFJMJcGuyDVRUk=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:AxuSPoL2fWn/rUyvWeNCNd0V2WCr+iHRCU9QO1PUmpY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:VXf9SnrnSmj2BF2cHkoTHvOUp8gjsz1KJFOMW7czdsY=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:FiULdbbzUxWD0Y4ZGPSVCDLvqRSyCIO6zKV7E2nf5uA=",
        version = "v1.57.1",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:GvKy4xLy1zF1XPbwP5NJb2HjRxhnhxjjXxvyZ1S/IAo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:PHS89lcFayWIEe0/s2jTBiEOtqghCxzc7y7bRNlifBs=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:5YMQ3Q+dqGpwUZ9X5sipsOQ1fLPsxod9HNq0+nrqc6I=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:Rd4+fBrjiN6tZ4TR8R/38elkyEkz6oogGDr7jDyjmMY=",
        version = "v1.17.3",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:9IHfEMWdCklJ1cwouoiQrnxmP0q3pH7JUt8Hqx4Qbck=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:xe/wJKz55VO1+L891a1EG9lVUgfHr9Ju/I3xh1nwF84=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:5xXuFfAjg0Z5Wb81j2GAbB3e0bwroCeSF+5jBn/L650=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:6sVlXXBmbd7jNX0Ipq0trII3e4n1/MsADLK6a+aiVlk=",
        version = "v1.23.3",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:EiGBeejtDDtr3JXt9W7xlhXyZ+REB5k2tBgVPVtmNb0=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:jIltU529R2zBFvP8rhiG1mgeTcnT27KhU0H/1d6SQRg=",
        version = "v1.29.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:5rhYLX+3a01drpREqBZVXR9YmWH45RnML++8NsCtuD8=",
        version = "v0.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:rbYNmHwvAOOwnW2FPXYkaK3Mf1MmGqRzK0mMiIEyLdo=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:7VmCNWcPJBS/srN2QnStTB6nu4Eb5TMcpkmtaPVhRt4=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:jV+EsDamGX6cE127+QAcCR/lergVeeZdEQ6DdrxW3sQ=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:Q90alBEYlMi66zL5gMSGQHfbZLB55mOAg03DhwTTfsk=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:zrq4uMmunf2KFDl/7dS6iCDBBAxBnKVDyw6+ajz3yu0=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:ACVOuxwe7gP0SqEso9SLyXbcZNk5l8hjcTX+XLntI5s=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:tTVP9tTxmc8fixxOd/8s6Q6Pz/+yzn7r7XdZHretQH0=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:NJnu1kAPamZDs/if3bJ3+Wb6tjADHKL83NUWsaIp2zg=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:0P9WcsQeTWjuD1H14JIY7XQscIPQ4Laje8ti96IC5vg=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:Z2sKPIB7bT2kMW5Uhxy44ZgdJzxzE5uKjavoW+EuHEE=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:5OVjzm8MPC5kP+Ywbs0mdE0O7AXvAUXksSyHAyMFyMg=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:tLCWad8HZhlyUNfDzDP5m+oH6h/1Uvw/ei7B9AnsWMk=",
        version = "v1.47.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:OFlXedmPP/5//X1hBEeq3D9kUVm9fb6ywYANlpv/EsQ=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:hlYieOXUwiJ7HpBR/vEPfr8nfSxveLVzbqbUkSK0c/4=",
        version = "v1.23.7",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:ua4GvsDztZ5F3xqjeLKVRDeOvJshf5QFgWGg1CKti3A=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:Szy3Q/N6bqgQGyxqjI+6xJZbmvPvnFHp3UZr95DKcQ0=",
        version = "v1.1.4",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:S2if6wkjR4JCEAfDtIiYtD+sTz/oXjh2NUG4cgT1y/Q=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:+pFmO4eu4dOVipSaFBLkmqrRYG94Xl/TQZFOeohkuqU=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:/+wUEGwk3x3Kxomi2cP5dsR8+SIXxo7M0THDjreFSYo=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:8aLcKnMPoldYU3YHgu4t2exrKhLQkqaXAGqT0ljrFVw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:ZjdiV3MyumRM6++1Ixu6N0VV9LAGlCX4AhW6Yjr1t+U=",
        version = "v1.15.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:KhnOrr9A1tXYIYeXKqCKbCI8TL2ZNGiD3dm+d7BDUBg=",
        version = "v1.3.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:1JLpZl31YhQDQeJ98tK6QiwTpgHFYRJwpntggpQQWis=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:J5tYUtb3r0cl2mM7+YHvV32eL+uZQ7lONyUZnPikCEo=",
        version = "v0.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:NmJsNX9uQ2CT78957xnjXZb26TDIMvv+d5W2vVUt0Pg=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:uuw2Xd37yHftViSI8J2hUcCS8S7SH3ZWH09sUDLW30Q=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:1jTsCu4bcsNsE4iiqNT5SHwrDRCfRmIaaaVFhRveTJI=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:M4vDbQ4TLXdaljXVZSwW7XtxpwXUUarY2lIs66m0aCM=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:VuFqv2ctf/A7AyKlNxVvlHTzjrEvumWaZflUzBPz/M4=",
        version = "v1.4.4",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:m1WljtkZnvLTIRYW1YTOv5A6H1yKgLHR6nU7O8yf27w=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:pj1sRfut2eRbD9pFRjNnPNg/CzJPuQAzUujMIM1vVeM=",
        version = "v1.15.5",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:zg9uq2yS9PGIOdc0Kz/l+zMtOlxKWonZjjo5w5YPG2A=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:rZEI/UxcxVKEzyoRS/kdJ1VoolNItRWjNN0Uk9tfexg=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:iEIOXFO9EmSiTjDmfpbRjOxECO7R8C7b8IXUGOj7xZw=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:w8xEcbZodnA2BbW6sVirkkoC+1gP8wS57EUUgGS0GVg=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:SF/u1IJduMqQQdJA4MDyivlIQ4SrV5qAawkr/ZEREkY=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:WxxLo//b60nNFESefLgaBQevu8QGUmRV3+noOjCfIHs=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:VRCQfZB4s6jN0CSy7+cO3m4ewNwgVnaePanVCQh/9Z4=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:cdex/ayDd294XBj2cGeMe6Y+H1JvhN8y78B9UW7pxuQ=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:94l/Yxg9oBZjin2bzI79oK05feYefieDq0o5fjLSkC8=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:blrdvF0MkPPivSO041ihul7rFMhXdVp8Uq7F59DKXTU=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:e9lUkCe2BexsqsUc2bjV8+gFBpQa54J+/F3qKVtW+wA=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:HsQk4FNKJUX04k3OI6gUsoveiHMGvDRqlaFM2xGyvqU=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:947tNIPnj1bMGTIEBo3fc4QrrFKS5hh0bFVsHmFm4Vo=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:eTOTfNL1yM6L/PCtquJwjWg7ZZGR0URFaFgbs8kllbM=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:iFsoexcp13cGT3k/Hv8PA5aK+FP7FnbhwDO9llnruas=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:kgwZ2f6qMMYIVBtUGGoU8yjYWwMTHDanLwM/CQCFaoQ=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:RWuXQDr9GDYhjmrredQJC7aY7cbyqP9ZuLbq5GJGves=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:OrRCIYEAbrbXdhm13/JINn9pQchvTTIzgmOCA7uJw8I=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:NP/KgsD9+0r9hmHC5wKye0vJXVwdciv219DtYKYjgqE=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:sPLUQkHq6b4AL0czSJZ0jd6vL55GSTHz2B3Md+TCZI0=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:sq+ScLP83d7GJy9+wpwYJVnY+q6xNTXwOdRIuYjvHT4=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:Vo10IpWKbNvc/z/QZPVXgCiwfjpWoZ/wbgful4Uh/4E=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:6SPCPvWav64tj0sVX/+npCBKhUi/UjJehy9op/V3p2g=",
        version = "v1.33.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:Zrd4LvT9PaW91X/Z13H0i5RKEv9suCLuk8zp+bfOpN4=",
        version = "v2.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:JRiwe4hvu3auuh2hujiTc2qNgPPfVp+Q8KOpsXlEzKQ=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:tC+ljmCCbuZ/ybt43odTFlay91n/HLIhflvaOeb0Dh4=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:J9cEHxG9YLmA9o4jTSvWt/RuVEn6MTrPlYSCRHujxDQ=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:JwZ7Ggle54XQ/FVYSBrMLOQIKoIT/uer8mmNvNLK51k=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:yTIL2CsZswmMfFyx2Ic77oLVzfBFoWBYgpkgiSPnC4Y=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:geqdX1FNqqL2p0ADXjPpw8lq986iv5GrVcieTYafuJQ=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:qdfZteAm+vgzN1iXzILo3nJFQbzziudkJrvd9wCf3FQ=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:eMEettHlFhG5pXsoHouIM5nRT+k+zU4+GUvRtnxhuVI=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:krnX9qpG2kR2fJ+u+uNyNo+ACVhplIAS4Pu7u+4gd+k=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:sdnh4Islb1ljaNhpIXlIPgb3eYj70QWgPVDKOUYvzJc=",
        version = "v1.15.4",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:crdn2Z2rFIy8WffmmhdlX3CwZJusqCiShtnrGFRwpeE=",
        version = "v1.24.3",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:5niCMfkw+jifmFtbBrtRedbXkJm3fubSR/KHbxSJZVM=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:nurhlJcSVFZneoRZgkBEHumTYf/kFJptCK2eBUq/88M=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:ttU+lhARPF/iZE3OkCpmfsemCz9mLaqBhGPd3Qub2sQ=",
        version = "v1.54.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:qkxNao58oF8ghAHE1Eghen7XepawYEN5zuZXYWaUTA4=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:B59ahL//eDfx2IIKFBeT5Atm9wnNmj3+8xG/W4WB//w=",
        version = "v1.35.1",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:YM1dnj5gLjfL6aDldO2s4GeU8JoAvH1xyIwXre63KmI=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:LnRJhhYkODDBoTwf6BeYkiJHFw9k+1mAFNyArwZUZAs=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:ahrzTgr7uAbvebuhkBAAVU6kRwVD0HWsmDsvMhtad5Q=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:XIEH5c0WeYGaVy9H+UueiTaf3NI6XNdB4/v6TFQJxtE=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:2qOAuAzNezwW3QN+t41BtkDJOG42HywL73q8x/f6fnM=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:t5WXTqlrk8VVJu/i3WrYQACjzYJiff5szARHiyqqPzI=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:Xrpbm2S9UFQ1pZEeJt9Vqm5t2T/z9y/M3rNXhFoo8Is=",
        version = "v1.20.3",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:YS4j7lY0zxYyneTFXjBJUj2r4CFe/UoIi/PJG0Zt/Rg=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:T/ujUghvEaTb+YnFY/jiYwVAkMbIC8EieK0CJo6B4vg=",
        version = "v2.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:qPNdab4aGgtaRX+51jCOtJxlJp6P26qua4o1xxUDjpc=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:WY526PqM6QNmFHSqe2sRfK6gRpzWjmL98UFkql2+JDM=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:zbs3V+9ux45KYq8lxxn/wgXole6SlBHHKKyZhNJoS+8=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:iceR3k0BCRZgf2D/NiKviVMFfuNC9LmeNLtxUFRB/wI=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:5Gp7h5j7jywxLUp6NTpjNPkgZb3ngl0tUSw6ICWvtJQ=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:qocsqETmLAl34mSa01hKZjcqAvt699gaoFbooGGMvaM=",
        version = "v1.12.3",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:szRajuUUbLyppkhs9K6BRtjY37l66XQQmw7oZRANE4k=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:kfYIdQftBnbAq8pUWFXfpuuxFSKzlmM5cSn76JByiT0=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:MrmRktzv/XF8CvtQt+P6wLUlURaNpSDJHFZhe//2QE4=",
        version = "v2.305.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:W9TXNZ+oB3MCd/8UjxHTWK5J9Nquw9fQBLJd5ne5/Ao=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
        version = "v1.2.0",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:P+/g8GpuJGYbOp2tAdKrIPUX9JO02q8Q0YNlHolpibA=",
        version = "v0.48.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:sv9kVfal0MK0wBMCOGr+HeJm9v803BkJxGrk2au7j08=",
        version = "v0.47.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:Za4UzOqJYS+MUczKI320AtqZHZb7EqxO00jAHE0jmQY=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:o8iWeVFa1BcLtVEV0LzrCxV2/55tB3xLxADr6Kyoey4=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
        sum = "h1:p3A5+f5l9e/kuEBwLOrnpkIDHQFlHmbiVxMURWRK6gQ=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:PQJmqJ9u2QaJLBOELl1cxIdPcpbwzbkjfEyelTl2rlo=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:O7JmZw0h76if63LQdsBMKQDWNb5oEcOThG9IrxscV+E=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:4LrmmEd8AU2rFvU1zegmvqW7+kWarxtNOPyeL6HmYY8=",
        version = "v1.23.1",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:2Di21piLrCqJ3U3eXGCTPHE9R8Nh+0uglSnOyxikMeI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:SEspjXHVqE1m5a1fRy8JFB+5jSu+V0GEDKDghF3ttO4=",
        version = "v0.160.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:nz5NESFLZbJGPFxDT/HCn+V1mZ8JGNoY4nUpmW/Y2eg=",
        version = "v0.0.0-20240102182953-50ed04b92917",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:rcS6EyEaoCO52hQDupoSfrxI3R6C2Tq741is7X8OvnM=",
        version = "v0.0.0-20240102182953-50ed04b92917",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:QXtV4qU5zS94SeHJhPqxJQF0XyxssnVrEZOUgp1+NuY=",
        version = "v0.0.0-20240116215550-a9fa1716bcac",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:nUQEQmH/csSvFECKYRv6HWEyypysidKl2I6Qpsglq/0=",
        version = "v0.0.0-20240116215550-a9fa1716bcac",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:TOvOcuXn30kRao+gfcvsebNEa5iZIiLkisYEkf7R7o0=",
        version = "v1.61.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:pPC6BG5ex8PDFnkbrGU3EixyhKcQ2aDuBS36lqK/C7I=",
        version = "v1.32.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:PGVlW0xEltQnzFZ55hkuX5+KLyrMYhHld1YHO4AKcdc=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:GoHiUyI/Tp2nVkLI2mCxVkOjsbSXD66ic0XW0js0R9g=",
        version = "v0.0.0-20230905200255-921286631fa9",
    )
    go_repository(
        name = "org_golang_x_exp_typeparams",
        importpath = "golang.org/x/exp/typeparams",
        sum = "h1:qyrTQ++p1afMkO4DPEeLGq/3oTsdlvdH4vqZUBWzUKM=",
        version = "v0.0.0-20220218215828-6cf2b201936e",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:0IiAsCRByjO2QjX7ZPkw5oU9x+n1YqRL802rjC0c3Aw=",
        version = "v0.0.0-20200130185559-910be7a94367",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:I/DsJXRlw/8l/0c24sM9yb0T4z9liZTduXvdAWYiysY=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:aCL9BSgETF1k+blQaYUBx9hJ9LOGP3gAVemcZlf1Kpo=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:aDkGMBSYxElaoP81NpoUoz2oo2R2wHdZpGToUxfyQrQ=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/sys",
        sum = "h1:xWw16ngr6ZMtmxDyKyIgsE93KNKz5HKmMa3b8ALHidU=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:m+B6fahuftsE9qjo0VWp2FW0mB3MTJvR0BaMQrq0pmE=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:jvNa2pY0M4r62jkRQ6RwEZZyPcymeL9XZMLBbV7U2nc=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
        version = "v0.0.0-20220907171357-04be3eba64a2",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:YIc7HTYsKndGK4RFzJ3covLz1byri52x0IoMB0Pt/vk=",
        version = "v1.13.1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/atomic",
        sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_dig",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/dig",
        sum = "h1:5Chju+tUvcC+N7N6EV08BJz41UZuO3BmHcN4A287ZLI=",
        version = "v1.17.0",
    )
    go_repository(
        name = "org_uber_go_fx",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/fx",
        sum = "h1:zVwVQGS8zYvhh9Xxcu4w1M6ESyeMzebzj2NbSayZ4Mk=",
        version = "v1.20.1",
    )
    go_repository(
        name = "org_uber_go_goleak",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/goleak",
        sum = "h1:2K3zAYmnTNqV73imy9J1T3WC+gmCePx2hEGkimedGto=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_net_metrics",
        importpath = "go.uber.org/net/metrics",
        sum = "h1:iRLPuVecNYf/wIV+mQaA4IgN8ghifu3q1B4IT6HfwyY=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_thriftrw",
        importpath = "go.uber.org/thriftrw",
        sum = "h1:pRuFLzbGvTcnYwGSjizWRHlbJUzGhu84sRiL1h1kUd8=",
        version = "v1.29.2",
    )
    go_repository(
        name = "org_uber_go_tools",
        importpath = "go.uber.org/tools",
        sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
        version = "v0.0.0-20190618225709-2cfd321de3ee",
    )
    go_repository(
        name = "org_uber_go_yarpc",
        importpath = "go.uber.org/yarpc",
        sum = "h1:smtf+0Q9xZorST5+SKYhOaBsjMO4V2heqf14BK6yIdQ=",
        version = "v1.70.4",
    )
    go_repository(
        name = "org_uber_go_zap",
        build_file_proto_mode = "disable_global",
        importpath = "go.uber.org/zap",
        sum = "h1:OjGQ5KQDEUawVHxNwQgPpiypGHOxo2mNZsOqTak4fFY=",
        version = "v1.23.0",
    )
