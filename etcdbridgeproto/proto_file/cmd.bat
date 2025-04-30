for %%f in ("D:\goproject\super-send-tool\etcdbridgeproto\proto_file\*.proto") do (
    protoc --go_opt=paths=source_relative --proto_path=D:\goproject\super-send-tool\etcdbridgeproto\proto_file --go_out=D:\goproject\super-send-tool\etcdbridgeproto --go-grpc_out=D:\goproject\super-send-tool\etcdbridgeproto "%%f"
)