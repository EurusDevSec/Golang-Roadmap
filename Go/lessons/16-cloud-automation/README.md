# 16 — Cloud SDK và automation an toàn

## Học trọng tâm

- Credential chain mặc định, region/project/subscription config.
- Pagination, rate limit, retry và eventual consistency.
- Tag/label, inventory và dry-run.
- Idempotent create/update/delete và audit trail.

## Best practices

- Không hard-code access key; ưu tiên workload identity/managed identity/role.
- Quyền tối thiểu và credential ngắn hạn.
- Mọi thao tác phá hủy cần `--dry-run`, xác nhận và scope rõ.
- Liệt kê qua pagination đầy đủ; không giả định một page chứa tất cả.
- Gắn tag ownership, environment và cost center.
- Ghi log resource ID và operation, không ghi secret.
- SDK client được inject để dễ fake/test.
- Với bulk operation: concurrency limit, retry có backoff và báo cáo từng item.

## Bài tập

1. Inventory resource từ file giả lập theo format của cloud API.
2. Viết paginator tổng quát cho danh sách nhiều page.
3. Công cụ tìm resource thiếu tag bắt buộc.
4. Cleanup tool có dry-run, allowlist environment và báo cáo JSON.

Checkpoint: chọn AWS/GCP/Azure, thay fake bằng SDK thật trong sandbox account với quyền read-only trước.

