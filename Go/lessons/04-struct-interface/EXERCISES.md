# Bài tập 04 — Struct và interface

Tạo code tại `Go/practice/04-struct-interface/`.

## Bài 1 — Server address

Tạo `Server{Name string, Host string, Port int}` và constructor validate input. Method `Address()` trả `host:port`.

```text
input:  name=api host=10.0.0.8 port=8080
output: api -> 10.0.0.8:8080
```

Host rỗng hoặc port ngoài 1–65535 trả lỗi.

## Bài 2 — Resource summary

Tạo interface `Resource` có `ID() string`, `Kind() string`. Implement `VM` và `Bucket`.

```text
input:  VM(vm-01), Bucket(logs), VM(vm-02)
output: bucket=1 vm=2
```

## Bài 3 — Checker thay thế được

Tạo interface `Checker { Check(context.Context, string) error }` và `FakeChecker` nhận sẵn map kết quả.

```text
api.local   -> healthy
db.local    -> error: connection refused
missing     -> error: target not configured
```

Viết test không gọi network thật.

