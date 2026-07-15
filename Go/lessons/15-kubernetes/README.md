# 15 — Kubernetes client và controller

## Học trọng tâm

- REST config: in-cluster và kubeconfig cho local development.
- Typed client, dynamic client, informer/cache.
- Reconciliation: desired state → observe → act → requeue.
- Idempotency, owner reference, finalizer, status và event.
- RBAC tối thiểu.

## Best practices

- Không poll API liên tục; dùng watch/informer/cache.
- Reconcile phải idempotent và chịu được chạy lặp.
- Không giả định event chỉ đến một lần hoặc đúng thứ tự.
- Cập nhật status riêng, ghi condition rõ reason/message.
- Chỉ xin RBAC verb/resource cần thiết.
- Tôn trọng context, rate limit và API conflict.
- Dùng envtest/fake client đúng phạm vi; test integration cho hành vi API thật.

## Bài tập

1. CLI liệt kê Pod theo namespace và label selector.
2. Tìm Deployment chưa đạt desired replicas.
3. Watch ConfigMap và hủy bằng context.
4. Viết reconciler đồng bộ một ConfigMap từ custom resource giả lập.

Checkpoint: controller nhỏ idempotent, có status, RBAC và test reconciliation.

