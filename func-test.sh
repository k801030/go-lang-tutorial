HTTP_CODE=$(curl -s -o /dev/null -w %{http_code} "http://localhost:80/ping?name=vison")
if [ $HTTP_CODE -eq 200 ]; then
  echo "[Functional test] pass (1/1)"
  exit 0
else 
  echo "[Functional test] pass (0/1)"
  echo "http code error: expected 200, but get $HTTP_CODE"
  exit 1
fi