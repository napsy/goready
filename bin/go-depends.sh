FILE=requirements.txt
echo "Downloading and installing Go dependencies ..."
while read CMD; do
    echo "'$CMD' ..."
    go get $CMD
done < "$FILE"
