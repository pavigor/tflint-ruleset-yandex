# TFLint Ruleset for Yandex Cloud (terraform-provider-yandex)
[![Build Status](https://github.com/pavigor/tflint-ruleset-yandex-cloud/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

TFLint ruleset plugin for [Terraform Yandex Provider](https://github.com/yandex-cloud/terraform-provider-yandex)

## Requirements

- TFLint v0.40+
- Go v1.19

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "yandex" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/pavigor/tflint-ruleset-yandex"
}
```

## Rules

Check current list of [rules](https://github.com/pavigor/tflint-ruleset-yandex/tree/main/rules) 

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "yandex" {
  enabled = true
}
EOS
$ tflint
```
