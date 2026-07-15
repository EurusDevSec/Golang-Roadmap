# Bài tập 15 — Kubernetes

Tạo code tại `Go/practice/15-kubernetes/`. Bắt đầu bằng fake client; chỉ dùng cluster local/sandbox cho integration test.

## Bài 1 — Liệt kê Pod

CLI:

```text
kubetool pods --namespace=dev --selector=app=api --output=json
```

Output sắp theo name:

```json
[{"name":"api-1","phase":"Running"},{"name":"api-2","phase":"Pending"}]
```

Namespace bắt buộc; context timeout 5 giây.

## Bài 2 — Deployment drift

Input trạng thái:

| Name | Desired | Available |
|---|---:|---:|
| api | 3 | 2 |
| worker | 2 | 2 |

Output:

```text
api desired=3 available=2 status=degraded
```

Exit `2` nếu có drift, `0` nếu tất cả đạt desired.

## Bài 3 — Reconcile ConfigMap

Desired data `{"MODE":"prod"}`, actual `{"MODE":"dev"}`.

```text
first reconcile:  action=updated requeue=false
second reconcile: action=none requeue=false
```

Chứng minh idempotent bằng test gọi reconcile hai lần.

## Bài 4 — RBAC tối thiểu

Viết Role cho controller chỉ được `get,list,watch,create,update,patch` ConfigMap trong một namespace. Không dùng wildcard, không có quyền Secret/Delete.

