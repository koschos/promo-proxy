```bash
aws ecr create-repository --repository-name promo-proxy --region eu-west-1
```

```bash
aws iam create-role --role-name ecsTaskExecutionRole --assume-role-policy-document file://$HOME/go/src/github.com/koschos/promo-proxy/aws/fargate/ecs-tasks-trust-policy.json
```

```bash
aws iam attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
```

```bash
aws ecs register-task-definition --region eu-west-1 --cli-input-json file://$HOME/go/src/github.com/koschos/promo-proxy/aws/fargate/task.json
```