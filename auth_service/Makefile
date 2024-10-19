# Указываем путь к proto-файлам и целевые директории
PROTO_SRC := ./pb/auth.proto
PROTO_OUT := ./

# Пути к зависимостям Google APIs
PROTO_INCLUDE := -I . -I C:/protoc-28.2-win64/include/googleapis

# Команды для генерации Go и gRPC файлов
.PHONY: all clean proto

# Запускаем все команды генерации
all: proto

# Генерируем Go-код из proto-файлов
proto:
	protoc $(PROTO_INCLUDE) \
		--go_out=$(PROTO_OUT) \
		--go-grpc_out=$(PROTO_OUT) \
		--grpc-gateway_out=$(PROTO_OUT) \
		$(PROTO_SRC)

# Удаление сгенерированных файлов
clean:
	rm -f $(PROTO_OUT)/*.go
