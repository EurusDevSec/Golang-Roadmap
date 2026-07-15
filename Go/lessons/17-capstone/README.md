# 17 — Capstone projects

Chọn một project. Không cần làm cả ba.

## Project A — Endpoint Guardian

CLI/service kiểm tra endpoint theo lịch:

- Config JSON/YAML và environment override.
- Worker pool, timeout, retry có giới hạn.
- Structured logs, metrics, health/readiness.
- JSON report, graceful shutdown, Docker image non-root.
- Unit/integration tests bằng `httptest`.

## Project B — Kubernetes Drift Reporter

Công cụ phát hiện workload không đạt desired state:

- Namespace/label filters.
- Kubernetes client với context và pagination.
- Output text/JSON, exit code dùng trong CI.
- RBAC read-only, container image và manifests.
- Tests cho logic đánh giá trạng thái.

## Project C — Cloud Resource Auditor

Công cụ kiểm tra tag, public exposure hoặc resource cũ:

- SDK client được inject.
- Pagination, concurrency limit và rate-limit handling.
- Read-only mặc định; remediation bắt buộc dry-run trước.
- Báo cáo JSON/CSV và exit code cho policy violation.
- Không lưu credential trong source/config.

## Definition of Done

- README có kiến trúc, cài đặt, config, ví dụ và troubleshooting.
- `go fmt`, `go vet`, `go test` và race detector thành công.
- Error path, timeout, cancellation đều có test.
- Không leak secret; dependency được khóa version.
- Docker chạy non-root và xử lý SIGTERM.
- CI build, test và scan cơ bản.
- Có runbook: cách quan sát, rollback và xử lý lỗi thường gặp.

## Tự review

1. Nếu dependency chậm hoặc chết, chương trình phản ứng thế nào?
2. Nếu nhận SIGTERM giữa công việc, dữ liệu có nhất quán không?
3. Nếu chạy lại cùng lệnh, kết quả có idempotent không?
4. Nếu input tăng 100 lần, concurrency có bị mất kiểm soát không?
5. Operator có thể chẩn đoán lỗi bằng log/metrics hiện có không?
