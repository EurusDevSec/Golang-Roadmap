# Bài tập 06 — File và config

Tạo code tại `Go/practice/06-files-config/`.

## Bài 1 — Parse endpoint JSON

Input `endpoints.json`:

```json
{"endpoints":[{"name":"api","url":"https://api.example.com"},{"name":"bad","url":":://"}]}
```

Output:

```text
error: endpoints[1].url: invalid URL
```

JSON có field lạ phải bị từ chối.

## Bài 2 — Thống kê log

Input:

```text
INFO server started
WARN disk usage 80%
ERROR database unavailable
INFO retrying
ERROR timeout
```

Output JSON chính xác:

```json
{"ERROR":2,"INFO":2,"WARN":1}
```

Đọc theo dòng, không dùng `os.ReadFile`.

## Bài 3 — Thứ tự ưu tiên config

Default port `8080`, file đặt `9000`, env `APP_PORT=7000`, flag `-port=6000`.

```text
output: port=6000
```

Bỏ flag thì output `7000`; bỏ env thì `9000`; bỏ file thì `8080`.

## Bài 4 — JSON sang CSV

Input:

```json
[{"name":"api","cpu":2},{"name":"db","cpu":4}]
```

Output:

```csv
name,cpu
api,2
db,4
```

Parser nhận `io.Reader`, writer nhận `io.Writer`.

