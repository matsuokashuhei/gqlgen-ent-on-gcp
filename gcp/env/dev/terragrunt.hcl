remote_state {
  backend = "gcs"
  config = {
    bucket = "studio-landin-dev-tfstate"
    prefix = "${path_relative_to_include()}"
  }
}

terraform {
  extra_arguments "common_vars" {
    commands = get_terraform_commands_that_need_vars()

    arguments = [
      "-var-file=../shared/variables.auto.tfvars",
    ]
  }
}

# inputs = {
#   project = {
#     id     = "studio-landin-dev"
#     region = "asia-northeast1"
#     zone   = "asia-northeast1-a"
#   }
# }
