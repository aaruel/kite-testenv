args="${@}"

KONTROL_PORT=6789 \
KONTROL_USERNAME="test" \
KONTROL_STORAGE="etcd" \
KONTROL_KONTROLURL="http://127.0.0.1:6789/kite" \
KONTROL_PUBLICKEYFILE=$(echo ~)/.kite/certs/key_pub.pem \
KONTROL_PRIVATEKEYFILE=$(echo ~)/.kite/certs/key.pem \
kontrol $args

