# Bài tập 16 — Cloud automation

Tạo code tại `Go/practice/16-cloud-automation/`. Dùng fake API trước, không cần tài khoản cloud để làm ba bài đầu.

## Bài 1 — Pagination

Fake API trả:

```text
page 1: [vm-1 vm-2], next=token-2
page 2: [vm-3], next=""
```

Output:

```text
[vm-1 vm-2 vm-3]
calls=2
```

Nếu page 2 lỗi, trả partial result hay không phải được quy định rõ và test.

## Bài 2 — Tag auditor

Tag bắt buộc: `owner`, `environment`.

Input:

```json
[{"id":"vm-1","tags":{"owner":"ops"}},{"id":"vm-2","tags":{"owner":"dev","environment":"staging"}}]
```

Output:

```json
[{"id":"vm-1","missing_tags":["environment"]}]
```

Exit `2` khi có vi phạm.

## Bài 3 — Cleanup dry-run

```text
> cleanup --environment=dev --older-than=30d
DRY-RUN delete vm-1
summary: matched=1 deleted=0
```

Chỉ khi thêm `--apply --confirm=dev` mới xóa. `--confirm` khác environment phải bị từ chối.

## Bài 4 — Bulk operation

Xử lý 50 resource với tối đa 5 request đồng thời. Retry `429` tối đa 3 attempts, không retry `403`. Output report phải có success/error riêng cho từng resource.

