for %%f in ("D:\goproject\super-send-tool\proto\*.proto") do (
    protoc --proto_path=D:\goproject\super-send-tool\proto --go_out=D:\goproject\super-send-tool\proto --go-grpc_out=D:\goproject\super-send-tool\proto "%%f"
)
