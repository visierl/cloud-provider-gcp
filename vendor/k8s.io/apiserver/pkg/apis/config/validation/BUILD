load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["validation.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiserver/pkg/apis/config/validation",
    importpath = "k8s.io/apiserver/pkg/apis/config/validation",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/util/validation/field:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/apis/config:go_default_library",
    ],
)
