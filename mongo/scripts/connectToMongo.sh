function connectToMongoDBandRunJS(){
    script=$1
    source ../../.env
    mongosh -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD < $script
}

read -p "Connecting to mongoDB and running script: $1. Continue(Y/n): " userConfirm
if [ "$userConfirm" != "Y"  ]; then
    exit 0
fi

echo ""
echo "Running script..."

connectToMongoDBandRunJS $1
