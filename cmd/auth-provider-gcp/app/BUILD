package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)
load("//defs:version.bzl", "version_x_defs")

go_library(
    name = "go_default_library",
    srcs = ["getcredentials.go"],
    importpath = "k8s.io/cloud-provider-gcp/cmd/auth-provider-gcp/app",
    deps = [
        "//cmd/auth-provider-gcp/provider:go_default_library",
        "//pkg/apis/credentialprovider:go_default_library",
        "//pkg/credentialconfig:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/k8s.io/klog/v2:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
    ],
)
