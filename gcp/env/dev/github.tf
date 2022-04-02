provider "github" {
  owner = var.github["username"]
  token = var.github["token"]
}

data "github_user" "landin" {
  username = var.github["username"]
}

data "github_repository" "landin" {
  full_name = join("/", [data.github_user.landin.username, var.github["repository"]])
}
