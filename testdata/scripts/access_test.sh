echo Testing health:
curl http://localhost:8080/api/health
CERT_NAME="prod-content-web.sdccd.edu"
echo Testing anonymous access:
curl http://localhost:8080/api/certificates/${CERT_NAME}/metadata
echo Testing bad token access:
curl -H 'Authorization: Bearer wrong-token' http://localhost:8080/api/certificates/${CERT_NAME}/metadata
echo Testing good token access:
curl -H 'Authorization: Bearer ff49ca3ea21edf591ff796f60dccba666a50992a0479cee78adc528beb75a228' http://localhost:8080/api/certificates/${CERT_NAME}/metadata

