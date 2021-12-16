#!/bin/sh

sed -i "s/PROFILER_API_URL/${PROFILER_API_URL}/g" /etc/nginx/nginx.conf

nginx & pid1="$!"
echo "nginx started with pid $pid1"
./profiler  & pid2="$!"
echo "profiler started with pid $pid2"

handle_sigterm() {
  echo "[INFO] Received SIGTERM"
  kill -SIGTERM $pid1 $pid2
  wait $pid1 $pid2
}
trap handle_sigterm SIGTERM

wait