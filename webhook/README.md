# Webhook

## Deployment

Run following command:

```bash
gcloud functions deploy powerhusky \
 --project=<your project>
 --gen2 \
 --runtime=go116 \
 --region=asia-east2 \
 --source=. \
 --entry-point=powerhusky \
 --trigger-http \
 --allow-unauthenticated
```
