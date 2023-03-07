DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# install go grpc generation plugin
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get github.com/golang/mock/mockgen@v1.6.0

# set go path so protoc can use plugin binaries
export PATH="$PATH:$(go env GOPATH)/bin"

cd ${DIR}
 TEMP_DIR=${DIR}/temp
 PROTOC_ROOT=${TEMP_DIR}/protoc
 PROTOC_BIN=${PROTOC_ROOT}/bin
 PROTOC_INCLUDE=${PROTOC_ROOT}/include
 PROTO_VERSION=$1
 if [[ -z "${PROTO_VERSION}" ]]; then
     PROTO_VERSION=3.6.1
 fi;

 # download and extract protoc
 if [[ ! -e ${PROTOC_ROOT} ]]; then
     PLATFORM="linux-x86_64"
     OS=$(uname -a)
     OS=${OS:0:5}
     echo "OS: $OS"
     if [[ "$OS" == "MINGW" ]]; then
         # windows
         PLATFORM="win32"
         echo "Windows detected"
     fi;
     rm -rf ${TEMP_DIR}
     mkdir ${TEMP_DIR}
     curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/protoc-${PROTO_VERSION}-${PLATFORM}.zip \
         -o ${TEMP_DIR}/protoc.zip

     unzip ${TEMP_DIR}/protoc.zip -d ${PROTOC_ROOT}
 fi;


# create output directory if it doesnt exist
PROTO_SRC_DIR=${DIR}/proto
OUT_DIR=${DIR}/src/generated
if [[ ! -e ${OUT_DIR} ]]; then
    mkdir ${OUT_DIR} -p
fi


# compile protobuf code (rule_api)
${PROTOC_BIN}/protoc --go_out=${OUT_DIR} --proto_path="${PROTO_SRC_DIR}"  \
    rule_api.proto

# compile protobuf code (rule_data)
${PROTOC_BIN}/protoc --go_out=${OUT_DIR} --proto_path="${PROTO_SRC_DIR}"  \
    rule_data.proto


# compile protobuf code (core)
${PROTOC_BIN}/protoc --go_out=${OUT_DIR} --proto_path="${PROTO_SRC_DIR}"  \
    core.proto

# compile protobuf code (middleware)
${PROTOC_BIN}/protoc --go_out=${OUT_DIR} --proto_path="${PROTO_SRC_DIR}"  \
    middleware.proto

# compile grpc code (middleware)
${PROTOC_BIN}/protoc --go-grpc_out=${OUT_DIR}  --proto_path="${PROTO_SRC_DIR}"  \
    middleware.proto


# compile grpc code (emotional state)
${PROTOC_BIN}/protoc --go-grpc_out=${OUT_DIR}  --proto_path="${PROTO_SRC_DIR}"  \
    rule_api.proto