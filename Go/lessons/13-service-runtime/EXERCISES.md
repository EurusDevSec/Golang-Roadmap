# Bài tập 13 — Service runtime

Tạo code tại `Go/practice/13-service-runtime/`.

## Bài 1 — HTTP graceful shutdown

Server có request kéo dài 200ms. Gửi SIGTERM sau khi request bắt đầu.

```text
new requests: rejected
in-flight request: completes with 200
process exits: within 2s
```

## Bài 2 — Worker drain

Queue có 5 job, worker đang xử lý job 1 thì context bị hủy.

```text
job 1: completed
jobs 2-5: not started
worker: exited
```

## Bài 3 — Cleanup order

Dependencies khởi tạo theo thứ tự `database -> metrics -> server`. Khi shutdown, output phải là:

```text
close server
close metrics
close database
```

Ngay cả khi `close metrics` trả lỗi, database vẫn phải được đóng và lỗi được tổng hợp.

## Bài 4 — Startup failure

Port đã bị chiếm. Service phải trả lỗi trong dưới 1 giây, không để background worker tiếp tục chạy.

