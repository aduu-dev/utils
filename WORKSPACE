# gazelle:repo bazel_gazelle
workspace(name = "dev_aduu_utils")

# load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# git_repository(
#     name = "dev_aduu_repo_infra",
#     commit = "ac2a5de3760a9cae7ded0fd6c30d3a286acf32d0",
#     remote = "https://github.com/aduu-dev/repo-infra",
#     shallow_since = "1595747467 +0200",
# )

local_repository(
    name = "dev_aduu_repo_infra",
    path = "../aduu_dev_repo_infra",
)

load("@dev_aduu_repo_infra//:load.bzl", "shared_repositories")

shared_repositories()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//:repos.bzl", "repositories")

# gazelle:repository_macro repos.bzl%repositories
repositories()

# load("@bazel_gazelle//:deps.bzl", "go_repository")

# go_repository(
#     name = "com_github_googlecodelabs_tools",
#     build_file_generation = "on",
#     importpath = "github.com/googlecodelabs/tools",
#     sum = "h1:pW6O2xWfr0k/kcaJlDY93jltyQnhDAagu9aX/pIoOKM=",
#     version = "v2.2.0+incompatible",
# )

# local_repository(
#     name = "com_github_googlecodelabs_tools",
#     path = "../claat",
# )

# load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
# git_repository(
#     name = "com_github_googlecodelabs_tools",
#     commit = "91ad58b3adcebfefaf26617d486ede833828e214",
#     remote = "http://github.com/fabstu/claat",
#     shallow_since = "1595759420 +0200",
# )

# load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# http_archive(
#     name = "com_github_googlecodelabs_tools",
#     sha256 = "8992b65ec5d98725352601b6fbbbf9585fa784a03c02e8dd927b23da3a84d39f",
#     strip_prefix = "claat-2.2.1",
#     urls = ["https://github.com/fabstu/claat/archive/v2.2.1.zip"],
# )

# load("@com_github_googlecodelabs_tools//claat:claat.bzl", "claat_dependencies")

# claat_dependencies()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "dev_aduu_rules_claat",
    sha256 = "6b72bebeff6ff2280ff08512fa92fb1b95120c2e9d18dceee8288ecd19123dcf",
    strip_prefix = "rules_claat-0.1.1",
    urls = ["https://github.com/aduu-dev/rules_claat/archive/v0.1.1.zip"],
)

load("@dev_aduu_rules_claat//:claat.bzl", "claat_dependencies")

claat_dependencies()
