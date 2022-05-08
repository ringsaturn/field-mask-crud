PROTO_FILES=$(shell find api -name *.proto)

.PHONY:pb
pb:
	protoc  --proto_path=. \
			--go_out=paths=source_relative:. \
			$(PROTO_FILES)
