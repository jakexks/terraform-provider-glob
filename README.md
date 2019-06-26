# Glob (Directory listing) provider for terraform

## Why?

Because terraform cannot get a list of files matching a pattern
(https://github.com/hashicorp/terraform/issues/16697), despite hyper-specific
functions like template_dir existing, so I'm just going to write it.

This has happened before (https://github.com/jakexks/terraform-provider-gzip),
which thankfully eventually made it in to the main terraform code
(https://www.terraform.io/docs/configuration/functions/base64gzip.html)

## How?

Install the provider:

```
go get -u github.com/jakexks/terraform-provider-glob
```

Copy the provider to the user plugins directory:

```
mv ~/go/bin/terraform-provider-glob ~/.terraform.d/plugins # on Windows systems, move to %APPDATA%\terraform.d\plugins instead
```

Use the provider:

```
data "glob_filename_list" "some_directory" {
  pattern = "./*"
}

output "all_files_in_some_directory" {
  value = "${data.glob_filename_list.some_directory.matches}"
}

data "glob_filename_list" "go_files" {
  pattern = "./*.go"
}

data "glob_contents_list" "go_files" {
  pattern = "./*.go"
}

data "glob_map" "go_files" {
  pattern = "./*.go"
}

output "go_filenames" {
  value = "${data.glob_filename_list.go_files.matches}"
}

output "go_filecontents" {
  value = "${data.glob_contents_list.go_files.matches}"
}

output "map" {
  value = "${data.glob_map.go_files.matches}"
}
```