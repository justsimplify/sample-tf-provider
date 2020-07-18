NOTES:
---
## Assumption:

- There is a REST API running to perform tasks on redis such as addition, deletion etc.
  - API is of the following format (All are `GET` method)
    - Required Headers:
      - `redis_host`- Redis Instance address to connect to
      - `redis_port`- Redis Instance port to connect to
    - `/add/:key/:value` - Adds key to the redis
    - `/delete/:key` - Deletes key from redis
    - `/get/:key` - Gets key from redis


## How to run:
- Build the provider in root project as 

  - ```bash
    # Format is terraform-<TYPE>-<NAME>
    # here type is `provider` and name is `redis-object` which we reference later
    go build -o terraform-provider-redis-object
    ```
    
- Create `variables.tfvars` as follows (example):

  - ```hcl
    redis_host = "http://0.0.0.0"
    redis_port = "6379"
    ```
    
- Create `variables.tf` as follows (example):

  - ```hcl
    variable "redis_host" {
      description = "redis host"
    }
    
    variable "redis_port" {
      description = "redis port"
    }
    ```
    
- Create `main.tf` as follows (example):
  - ```hcl
    resource "redis-object" "my-redis-object-1" {
      key   = "k1"
      value = "v1"
    }
    ```
    
- Run the following
  - ```bash
    # Initialize the terraform
    terraform init
    
    # Plan
    terraform plan -var-file=variables.tfvars
    
    # Apply
    terraform apply -var-file=variables.tfvars -auto-approve 
    
    # For cleanup
    terraform destroy -var-file=variables.tfvars -auto-approve
    ```
    
## Debug:
- Set `set TF_LOG=DEBUG`. As per the warning, `TRACE` will be replaced permanently and other log types will be removed. 

## References:
- https://github.com/justsimplify/sample-redis-api - (For sample Redis API)