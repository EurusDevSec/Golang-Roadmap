# Roadmap Go cho DevOps & Cloud

## Đích đến

Bạn được xem là thành thạo theo hướng này khi tự xây dựng được một công cụ có:

- CLI và cấu hình rõ ràng.
- Gọi API với timeout, retry có giới hạn và cancellation.
- Chạy nhiều tác vụ đồng thời nhưng kiểm soát tài nguyên.
- Log có cấu trúc, metrics và health check.
- Unit test, integration test và graceful shutdown.
- Docker image nhỏ, chạy non-root.
- Tương tác an toàn với Kubernetes hoặc một cloud SDK.

## Cách học

- 70% viết code, 20% đọc và sửa code, 10% ghi chú.
- Mỗi buổi học một khái niệm rồi áp dụng ngay vào bài tập.
- Không chép lời giải. Tự viết test cho các trường hợp lỗi.
- Luôn chạy `go fmt ./...`, `go vet ./...`, `go test ./...`.
- Không tối ưu sớm; ưu tiên code đúng, rõ và đo được.

## Lộ trình đề xuất

### Chặng 1 — Nền tảng Go

1. [Cú pháp, function, pointer](Go/lessons/00-go-basics/README.md)
2. [Array](Go/lessons/01-array/README.md)
3. [Slice](Go/lessons/02-slice/README.md)
4. [Map](Go/lessons/03-map/README.md)
5. [Struct, method, interface](Go/lessons/04-struct-interface/README.md)
6. [Error, package, module](Go/lessons/05-errors-packages/README.md)

Checkpoint: viết chương trình inventory tài nguyên server bằng struct, slice, map; tách package và xử lý lỗi.

### Chặng 2 — Automation

7. [File, JSON, YAML, config](Go/lessons/06-files-config/README.md)
8. [CLI, environment, process, signal](Go/lessons/07-cli-os/README.md)
9. [HTTP client và API](Go/lessons/08-http-api/README.md)
10. [Context, timeout, retry](Go/lessons/09-context-retry/README.md)

Checkpoint: CLI đọc config, kiểm tra nhiều endpoint, xuất JSON và trả exit code đúng.

### Chặng 3 — Production quality

11. [Goroutine, channel, synchronization](Go/lessons/10-concurrency/README.md)
12. [Testing](Go/lessons/11-testing/README.md)
13. [Logging, metrics, health](Go/lessons/12-observability/README.md)
14. [Service và graceful shutdown](Go/lessons/13-service-runtime/README.md)

Checkpoint: health checker đồng thời có worker limit, test, structured log, metrics và shutdown sạch.

### Chặng 4 — Cloud native

15. [Container best practices](Go/lessons/14-container/README.md)
16. [Kubernetes client và controller](Go/lessons/15-kubernetes/README.md)
17. [Cloud SDK và automation an toàn](Go/lessons/16-cloud-automation/README.md)
18. [Capstone projects](Go/lessons/17-capstone/README.md)

Checkpoint cuối: hoàn thành một capstone có README, tests, Dockerfile, CI và tài liệu vận hành.

## Không cần ưu tiên lúc đầu

- Thuật toán competitive programming.
- Reflection và unsafe nâng cao.
- Tối ưu assembly hoặc memory vi mô.
- Framework web lớn khi `net/http` đã đủ.
- Microservices phức tạp trước khi làm tốt một service đơn.

## Nhịp học tham khảo

- 8–10 giờ/tuần: khoảng 5–7 tháng.
- 15 giờ/tuần: khoảng 3–4 tháng.

Không chạy theo số tuần. Chỉ chuyển chặng khi hoàn thành checkpoint.

