data "glob_list" "some-directory" {
  pattern = "./*"
}

output "filelist" {
  value = "${data.glob_list.some-directory.matches}"
}
