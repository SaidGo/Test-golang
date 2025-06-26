#!/bin/bash

echo "[*] Генерация OpenAPI через oapi-codegen..."

oapi-codegen -config openapi/.openapi openapi/openapi.yaml > internal/web/tasks/api.gen.go

if [ $? -ne 0 ]; then
    echo "[!] Ошибка генерации!"
    exit 1
else
    echo "[+] Генерация прошла успешно: internal/web/tasks/api.gen.go"
fi
