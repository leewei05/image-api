# Image-api(Work in progress)

## Local Testing & Build

```
// Run PostgreSQL
docker run -it --rm -p 5432:5432 -e POSTGRES_PASSWORD=123456 postgres
// Create database
psql -h localhost -U postgres -f ./sql/create_table.sql
// Connect to PostgreSQL
psql -h localhost -U postgres

// Run Redis
docker run --rm -p 6379:6379 redis
// Connect to Redis
redis-cli -h localhost -p 6379

// Create test GCS bucket
gsutil mb gs://[BUCKETNAME]
// Copy image 'cat.jpg' to the GCS bucket
gsutil cp ./assets/cat.jpg gs://lee-image-test

// Run go test in root directory
go test -v ./...

// Build Docker image
docker build -t image-api .
```

## Infrastructure setup

### Create GCP Cloud SQL instace

```
gcloud sql instances create pg-sql \
  --zone us-west1-a \
  --database-version POSTGRES_10 \
  --root-password 123456 \
  --cpu=2 --memory=8GiB
```

### Create Google Kubernetes Engine

```
gcloud container clusters create image-k8s \
  --zone us-west1-a \
  --machine-type n1-standard-1 \
  --num-nodes 2

// Connect to GKE cluster, in order to use kubectl
gcloud container clusters get-credentials image-k8s
```
### Create GKE Ingress(HTTP(S) Load Balancer)

### Create GCP Cloud Storage Bucket

Create GCS bucket and make it public for serving
```
gsutil mb gs://image-api-v1
gsutil defacl set public-read gs://image-api-v1
```

## Deploy to GKE
