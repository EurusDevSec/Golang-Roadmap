# 14 — Container hóa chương trình Go

## Học trọng tâm

- Multi-stage build.
- Static binary khi phù hợp, `CGO_ENABLED=0`.
- PID 1, signal và graceful shutdown.
- Read-only filesystem, non-root user và resource limit.

## Best practices

- Pin phiên bản base image; cập nhật có quy trình.
- Build reproducible và chỉ copy binary/config cần thiết.
- Không đưa source, credential hoặc toolchain vào runtime image.
- Chạy non-root, drop capability và dùng filesystem read-only nếu có thể.
- Không nhúng config theo môi trường vào image.
- Có health endpoint nhưng cân nhắc probe từ orchestrator thay vì HEALTHCHECK phức tạp.
- Quét vulnerability và tạo SBOM trong CI khi project đi production.

## Bài tập

1. Viết multi-stage Dockerfile cho health checker.
2. Chạy container bằng non-root user.
3. Gửi SIGTERM và xác nhận service shutdown sạch.
4. Mount config read-only và truyền secret qua cơ chế runtime.

