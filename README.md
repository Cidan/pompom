Pom Pom
=======
Pom Pom is a highly configurable and modular bridge for Google Pub/Sub and Google Cloud Storage that accepts inputs via REST, gRPC, and flat files on disk and submits that data to Google Pub/Sub or Google Cloud Storage

## Getting Started
tbd

## Local Cache
Pom Pom can optionally be configured to hold a local cache of data that is accepted and submit data to Pub/Sub or Cloud Storage asynchronously. This is a useful feature for remote deployments where connections to Google's infrastructure might be spotty or otherwise runs the risk of disconnection, but you don't want to apply back pressure to the submitting interface.

### Configuring

## Tests
Run `make test` for unit and integration tests. Tests require the Google Cloud Pub/Sub emulator running via `gcloud beta emulators pubsub start` using the GCP Cloud SDK

## License
MIT