HTTP_CODE=$(curl -s -o /tmp/server.log -w %{http_code} "http://localhost:80/health-check")

if [ $HTTP_CODE -eq 200 ]; then
  echo "ok"
  exit 0
else 
  echo "fail"
  exit 1
fi