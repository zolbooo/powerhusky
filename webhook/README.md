# Webhook

This service is deployed to Google Cloud Functions.

## Deployment

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
   - GCP_PROJECT
   - GCE_INSTANCE_ID
   - GCE_INSTANCE_REGION
