# setup

1. Create an Google Cloud project on Cloud Console
2. Create an Cloud Storage bucket on Cloud Shell
    ```
    PROJECT_ID=$(gcloud config get-value project)
    gsutil mb gs://${PROJECT_ID}-tfstate
    gsutil versioning set on gs://${PROJECT_ID}-tfstate
    ```
3. Initialize the gcloud CLI
    ```
    docker run --rm --platform linux/x86_64 -v $HOME/.config/gcloud:/root/.config/gcloud -ti gcr.io/google.com/cloudsdktool/cloud-sdk gcloud init --console-only
    ```
4. Open the `gcp` directory in a container. Please see https://code.visualstudio.com/docs/remote/containers in details.

# Services not managed by terraform

- [Identity Platform](https://console.cloud.google.com/customer-identity/)
    - Email/Password

