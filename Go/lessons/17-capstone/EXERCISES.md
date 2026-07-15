# Bài tập 17 — Capstone acceptance criteria

Chọn đúng một project. Đây là đặc tả đầu ra, không phải lời giải.

## A — Endpoint Guardian

Input `config.json`:

```json
{"concurrency":2,"timeout":"1s","endpoints":[{"name":"api","url":"http://api:8080/healthz"},{"name":"db-proxy","url":"http://db:9090/healthz"}]}
```

Lệnh và output:

```text
> guardian check --config=config.json --output=json
[{"name":"api","status":200,"healthy":true},{"name":"db-proxy","status":503,"healthy":false}]
exit code: 2
```

Exit `0` khi tất cả healthy, `2` khi có endpoint unhealthy, `1` khi lỗi chương trình/config. Test bắt buộc: 200, 503, timeout, cancellation, worker limit.

## B — Kubernetes Drift Reporter

```text
> drift-report --namespace=prod --selector=team=platform --output=json
[{"kind":"Deployment","name":"api","desired":3,"available":2}]
exit code: 2
```

Không có drift trả `[]` và exit `0`. Lỗi Kubernetes API trả exit `1`. Output phải ổn định theo kind/name. RBAC read-only.

## C — Cloud Resource Auditor

```text
> cloud-audit tags --required=owner,environment --output=json
[{"id":"vm-01","missing_tags":["owner"]}]
exit code: 2
```

Phải hỗ trợ pagination, timeout, concurrency limit và fake client test. Mọi remediation mặc định dry-run và cần `--apply` cùng xác nhận environment.

## Điều kiện nộp

```text
go fmt ./...       PASS
go vet ./...       PASS
go test ./...      PASS
go test -race ./... PASS
docker build       PASS
SIGTERM shutdown   PASS
```

README project phải có: kiến trúc, cách chạy, config reference, exit codes, metrics/logs, security notes và troubleshooting.
