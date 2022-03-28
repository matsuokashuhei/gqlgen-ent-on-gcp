# gcloud init

```
docker run --rm --platform linux/x86_64 -v $HOME/.config/gcloud:/root/.config/gcloud -ti gcr.io/google.com/cloudsdktool/cloud-sdk gcloud init --console-only
```

# auth login

```
docker run --rm --platform linux/x86_64 -v $HOME/.config/gcloud:/root/.config/gcloud -ti gcr.io/google.com/cloudsdktool/cloud-sdk gcloud auth application-default login --no-launch-browser
```

[Cloud Storage バケットに状態を保存するように Terraform を構成する](https://cloud.google.com/architecture/managing-infrastructure-as-code?utm_source=youtube&utm_medium=unpaidsoc&utm_campaign=CDR_mao_gcp_ce93fpqrkck_ServerlessExpeditions_040821&utm_content=description#configuring_terraform_to_store_state_in_a_cloud_storage_bucket)
