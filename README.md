# Powerhusky

Toolset to manage servers

## Why this is created?

I personally use this toolset to enable/disable VPS-es on demand to save some money on running servers :)

For example, I have a powerful server to run Android app build jobs (on GitLab), which should be turned off as soon as build is created and uploaded. So, I set up server with Gitlab Runner and turn it on as soon as webhook with relevant job is invoked. As soon as job finishes corresponding webhook is invoked and VM will be stopped.

## Architecture

There are two key parts of this toolchain: daemon and webhook. **Note:** currently, only GCE and Gitlab are supported.

### Daemon

This is service running on server and handling all incoming requests from webhook. It is responsible for shutdown logic.

### Webhook

This service is being invoked by supported integrations and handles VPS power, cloud integration and authorization.

## TODO

1. Webhook
   - [x] GCE setup
   - [ ] Server setup
   - [x] VPS enabling
2. Daemon
   - [ ] Webhook-daemon communication
   - [ ] Power-off logic
