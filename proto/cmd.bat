for %%f in ("D:\goproject\super-send-tool\proto\*.proto") do (
    protoc --go_opt=paths=source_relative --proto_path=D:\goproject\super-send-tool\proto --go_out=D:\goproject\super-send-tool\proto --go-grpc_out=D:\goproject\super-send-tool\proto "%%f"
)
