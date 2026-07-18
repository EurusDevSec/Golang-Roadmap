# 18 — Gin thực chiến tối giản

## Mục tiêu

- Làm được API backend nhỏ bằng Gin mà không cần kiến trúc phức tạp.
- Tập trung vào thứ dùng thật trong công việc: route, JSON, middleware, config, health check, logging và Docker.
- Giữ code rõ ràng để sau này gắn vào DevOps, container và deployment dễ hơn.

## Học trọng tâm

- Router, group route, path param, query param.
- Bind JSON vào struct và trả JSON có format ổn định.
- Middleware cho log, recover và auth đơn giản.
- Đọc config từ biến môi trường.
- Health check và readiness check.
- CRUD nhỏ, typed request/response, status code đúng.
- Chuẩn bị cho Docker và deploy.

## Best practices vừa đủ dùng

1. Handler chỉ nhận request, gọi service, trả response.
2. Dùng struct cho request/response, không dùng map linh tinh khi schema đã rõ.
3. Trả status code đúng nghĩa: `200`, `201`, `400`, `401`, `404`, `500`.
4. Không log secret như token, password, cookie.
5. Đọc port, mode, secret từ env thay vì hardcode.
6. Có `/healthz` và `/readyz` cho môi trường deploy.
7. Middleware càng ít càng tốt nhưng phải rõ mục đích.
8. Chưa cần microservice, message queue hay clean architecture quá nặng.

## Lộ trình học

1. Tạo router và route đầu tiên.
2. Nhận JSON và trả JSON bằng struct.
3. Validate input và chuẩn hóa lỗi.
4. Viết middleware log và recover.
5. Đọc config từ env, thêm health check.
6. Làm một CRUD nhỏ in-memory hoặc với DB đơn giản.
7. Chạy thử với Docker sau khi API đã ổn.

## Mốc hoàn thành

- Bạn tạo được API nhỏ chạy được bằng Gin.
- Bạn biết vì sao cần middleware, JSON typed struct và health check.
- Bạn có thể đọc log để debug khi chạy local hoặc trong container.
- Bạn không bị cuốn vào kiến trúc phức tạp trước khi cần.
