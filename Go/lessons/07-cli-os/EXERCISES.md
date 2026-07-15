# Bài tập 07 — CLI và OS

Tạo code tại `Go/practice/07-cli-os/`.

## Bài 1 — envcheck

CLI: `envcheck -required DB_HOST,DB_PORT,API_TOKEN`.

```text
environment: DB_HOST=db.local, DB_PORT=5432
stdout:      (rỗng)
stderr:      missing environment variables: API_TOKEN
exit code:   2
```

Không in giá trị biến môi trường.

## Bài 2 — Command timeout

CLI nhận `-timeout` rồi chạy command sau `--`.

```text
> runner -timeout=1s -- long-command
error: command timed out after 1s
```

Dùng `exec.CommandContext`, không dùng `sh -c`. Thành công phải giữ nguyên stdout của process con.

## Bài 3 — Output stream

Với config hợp lệ:

```text
stdout: config valid
stderr: (rỗng)
exit:   0
```

Với config sai:

```text
stdout: (rỗng)
stderr: error: port must be between 1 and 65535
exit:   2
```

