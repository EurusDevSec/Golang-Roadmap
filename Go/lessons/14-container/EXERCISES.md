# Bài tập 14 — Container

Tạo project tại `Go/practice/14-container/` gồm `main.go`, `Dockerfile`, `.dockerignore` và README.

## Bài 1 — Multi-stage image

Build binary ở stage Go, runtime chỉ chứa binary cần chạy.

```text
docker build -t health-service:test .
docker run --rm health-service:test --version
version=dev
```

Runtime image không chứa Go compiler và source `.go`.

## Bài 2 — Non-root và read-only

Container phải chạy được với:

```text
docker run --rm --read-only --cap-drop=ALL health-service:test
```

Process UID không được là `0`.

## Bài 3 — Config runtime

Mount `/config/app.json` read-only. Thiếu file:

```text
stderr: error: read config "/config/app.json": ...
exit: 1
```

Không `COPY` config môi trường hoặc secret vào image.

## Bài 4 — SIGTERM

Chạy container, gửi `docker stop --time=5`. Log cuối phải có `shutdown complete` và container thoát trước 5 giây.

