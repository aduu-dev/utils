# gazelle:repository_macro repos.bzl%go_repositories
workspace(name = "dev_aduu_utils")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "com_google_protobuf",
    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
    strip_prefix = "protobuf-3.11.4",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# rules_go

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2d536797707dd1697441876b2e862c58839f975c8fc2f0f96636cbd428f45866",
    urls = [
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
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

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
    name = "com_github_googlecodelabs_tools",
    commit = "9f77c56996154e04ecad00606e8699be0b4fb76c",
    remote = "https://github.com/fabstu/tools",
)

# local_repository(
#    name="com_github_googlecodelabs_tools",
#    path="/Users/fabiansturm/Documents/projects/intprojects/claat",
# )

go_rules_dependencies()

load("//:repos.bzl", "go_repositories")

go_repositories()
