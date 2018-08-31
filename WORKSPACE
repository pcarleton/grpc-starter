load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.1/rules_go-0.15.1.tar.gz"],
    sha256 = "5f3b0304cdf0c505ec9e5b3c4fc4a87b5ca21b13d8ecc780c97df3d1809b9ce6",
)
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

go_repository(
    name = "org_golang_x_oauth2",
    remote = "git@github.com:golang/oauth2.git",
    vcs = "git",
    commit = "fdc9e635145ae97e6c2cb777c48305600cf515cb",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "com_github_nmrshll_oauth2_noserver",
    remote = "git@github.com:nmrshll/oauth2-noserver.git",
    vcs = "git",
    commit = "16b622b98a45a4bb67e24e61e76892580026c9c9",
    importpath = "github.com:nmrshll/oauth2-noserver",
)
