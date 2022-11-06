# Webhook

This service is deployed to Google Cloud Functions.

## Deployment

### Github Actions flow

1. [Create service account](https://cloud.google.com/functions/docs/securing/function-identity#individual) for the function.

   You should provide least privileges for function to start GCE instances, for example: `Compute Instance Admin (beta)`. (If you find some role with less permissions, which allows to start instances, feel free to raise an issue/PR!)

2. Configure [Workload Identity Federation](https://github.com/google-github-actions/auth#setting-up-workload-identity-federation)

3. Create `powerhusky-gitlab-token`, `powerhusky-daemon-secret` secrets in [Secret Manager](https://cloud.google.com/secret-manager/docs/create-secret)

4. Configure [Github environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment). You need following secrets:

   - SERVICE_ACCOUNT - service account email
   - APP_SERVICE_ACCOUNT - service account email for Cloud Function. You created it on the first step
   - APP_REGION
   - WORKLOAD_IDENTITY_POOL_ID
   - GCP_PROJECT - GCP project ID
   - APP_PROJECT_ID - project ID where GCE instance is located
   - GCE_INSTANCE_REGION
   - GCE_INSTANCE_ID

5. Run deploy-webhook workflow. It could raise timeout error, so that make sure to check Cloud Console/use CLI to check whether if function was deployed.

### Bare flow

1. [Create service account](https://cloud.google.com/functions/docs/securing/function-identity#individual) for the function.

   You should provide least privileges for function to start GCE instances, for example: `Compute Instance Admin (beta)`. (If you find some role with less permissions, which allows to start instances, feel free to raise an issue/PR!)

2. Run following command:

```bash
gcloud functions deploy powerhusky \
 --project=<your project>
 --gen2 \
 --runtime=go116 \
 --region=<your region> \
 --source=. \
 --entry-point=powerhusky \
 --trigger-http \
 --allow-unauthenticated
```

3. Configure environment variables:

   - GITLAB_TOKEN
   - DAEMON_SECRET
   - GCP_PROJECT
   - GCE_INSTANCE_ID
   - GCE_INSTANCE_REGION
