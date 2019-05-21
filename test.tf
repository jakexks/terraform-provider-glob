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
