# 11 — Testing cho công cụ hạ tầng

## Học trọng tâm

- Unit test, table-driven test, subtest.
- Fake/stub nhỏ qua interface; `httptest` cho HTTP.
- Temp directory bằng `t.TempDir()`.
- Integration test tách bằng build tag hoặc environment rõ ràng.
- Fuzz test hữu ích cho parser/config input.

## Best practices

- Test hành vi công khai, không khóa chặt implementation detail.
- Mỗi test độc lập, không phụ thuộc thứ tự hoặc cloud account thật.
- Dùng `t.Cleanup` để hoàn nguyên tài nguyên.
- Không dùng `time.Sleep` dài để chờ concurrency; dùng signal/channel/event.
- Test error path, timeout, cancellation và input rỗng.
- Coverage là tín hiệu, không phải mục tiêu duy nhất.

## Bài tập

1. Viết table test cho `parsePort`.
2. Test config loader bằng `t.TempDir()`.
3. Test HTTP client bằng `httptest.Server`.
4. Test cancellation mà không chờ timeout dài.
5. Chạy `go test -race -cover ./...`.

Checkpoint: project có test cho happy path và ít nhất ba failure path.

