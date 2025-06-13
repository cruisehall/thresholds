# Signals
## Aws Cloudwatch Metrics

### `lambda_errors`
Counts the number of Lambda invocation errors.

#### Scope
- `function_name`: The name of the Lambda function.

### `lambda_latency`
Measures the latency of Lambda function invocations.

#### Scope
- `function_name`: The name of the Lambda function.

### `lambda_invocations`
Counts the number of times a Lambda function is invoked.

#### Scope
- `function_name`: The name of the Lambda function.

### `lambda_throttles`
Counts the number of throttled Lambda invocations.

#### Scope
- `function_name`: The name of the Lambda function.

### `lambda_concurrent_executions`
Tracks the number of concurrent Lambda executions.

#### Scope
- `function_name`: The name of the Lambda function.

### `api_gateway_5xx_errors`
Counts the number of 5xx errors returned by API Gateway.

#### Scope
- `api_id`: The ID of the API Gateway.

### `api_gateway_4xx_errors`
Counts the number of 4xx errors returned by API Gateway.

#### Scope
- `api_id`: The ID of the API Gateway.

### `api_gateway_latency`
Measures the latency of API Gateway requests.

#### Scope
- `api_id`: The ID of the API Gateway.

### `sqs_approximate_number_of_messages_visible`
Tracks the approximate number of visible messages in an SQS queue.

#### Scope
- `queue_name`: The name of the SQS queue.

### `sqs_approximate_age_of_oldest_message`
Measures the approximate age of the oldest message in an SQS queue.

#### Scope
- `queue_name`: The name of the SQS queue.

### `dynamodb_read_throttle_events`
Counts the number of read throttle events in DynamoDB.

#### Scope
- `table_name`: The name of the DynamoDB table.

### `dynamodb_write_throttle_events`
Counts the number of write throttle events in DynamoDB.

#### Scope
- `table_name`: The name of the DynamoDB table.
## Kubernetes

### `k8s_pod_cpu_usage`
Measures the CPU usage of a Kubernetes pod.

#### Scope
- `namespace`: The namespace of the pod.
- `pod_name`: The name of the pod.

### `k8s_pod_memory_usage`
Measures the memory usage of a Kubernetes pod.

#### Scope
- `namespace`: The namespace of the pod.
- `pod_name`: The name of the pod.

### `k8s_node_cpu_usage`
Measures the CPU usage of a Kubernetes node.

#### Scope
- `node_name`: The name of the node.

### `k8s_node_memory_usage`
Measures the memory usage of a Kubernetes node.

#### Scope
- `node_name`: The name of the node.
