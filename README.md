# Rate-Limiter


## Introduction
Welcome to our rate limiter service! This powerful solution is built using the sliding window counter algorithm, offering efficient rate limiting capabilities. Our service is designed to integrate seamlessly with gRPC, providing you with a flexible and customizable rate limiting middleware.

## Key Features
- Efficient rate limiting using the sliding window counter algorithm
- Seamless integration with gRPC middleware
- Customizable rate limiting parameters for fine-tuning behavior
- Easy-to-use interface for configuring and managing rate limiting settings

## Getting Started
To start using our rate limiter service, follow these steps:
1. Install the necessary dependencies.
2. Setup the deployments.
3. Customize the rate limiting parameters to match your requirements.
4. Run your gRPC server and witness the power of rate limiting in action.


## Deployments
You can deploy the needed services(like postgres and redis) using following command with docker-compose:
```bash
make deployment
```





## Example Usage

### Configuration

By editting configs/service-configs.yaml you can add rate limiter to your service:

```yaml
ratelimit:
  - rpcName: /service.server.Server/GetUser
    requestsPerMinute: 20

  - rpcName: /service.server.Server/UpdateProfile
    requestsPerMinute: 10

```

### Database Migration
```bash
./rate_limiter migrate
```

### Create sample user
This command create user object with username=`ali`, password=`12345`, email=`ali.sadafi.dev@gmail.com`, and first_name=`Ali`:
```bash
./rate_limiter create user ali 12345 ali.sadafi.dev@gmail.com Ali
```

### Sample requests for test

#### Login
Use POST `/login` request with `username` and `password` parameters.

#### Logout
Use POST `/logout` request.

#### GetUser
Use GET `/user` request. (rate limited request)



