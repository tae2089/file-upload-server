openssl genrsa -out app.rsa 4096
openssl rsa -in app.rsa -pubout > app.rsa.pub