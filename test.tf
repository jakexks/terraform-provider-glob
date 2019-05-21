data "glob_list" "some-directory" {
  pattern = "./*"
}

output "filelist" {
  value = "${data.glob_list.some-directory.matches}"
}

data "glob_map" "some-other-directory" {
  pattern = "./*.go"
}

output "filemap" {
  value = "${data.glob_map.some-other-directory.matches}"
}
