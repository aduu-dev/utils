# gazelle:repository_macro repos.bzl%go_repositories
# gazelle:repo bazel_gazelle
workspace(name = "dev_aduu_utils")

local_repository(
    name = "dev_aduu_repo_infra",
    path = "../aduu_dev_repo_infra",
)

load("@dev_aduu_repo_infra//:load.bzl", "repositories")

repositories()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("//:repos.bzl", "go_repositories")

go_repositories()
