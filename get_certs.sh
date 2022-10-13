#!/bin/bash
TOKEN=`curl -s -X POST http://10.128.228.8/v1/auth -d '{"username":"doomsday","password":"password"}' | jq -r '.token'`
TIME_NOW=`date +%s`
TIME_30DAYS_FROM_NOW=`date -v+30d +%s`

CERTS=`curl -s -X GET --cookie "doomsday-token=$TOKEN" http://10.128.228.8/v1/cache`

CERTS_EXPIRED_NOW=`echo $CERTS | jq '.content[] | select(.not_after < '"$TIME_NOW"')'`
CERTS_EXPIRE_IN_NEXT_30DAYS=`echo $CERTS | jq '.content[] | select((.not_after > '"$TIME_NOW"') and (.not_after < '"$TIME_30DAYS_FROM_NOW"'))'`

echo $CERTS_EXPIRED_NOW | jq -r '.common_name'
echo
echo "Number of certs that are expired"
echo "###############"
echo $CERTS_EXPIRED_NOW | jq -r '.common_name' | wc -l
echo
echo $CERTS_EXPIRE_IN_NEXT_30DAYS | jq -r '.common_name'
echo
echo "Number of certs that will expire in next 30 days"
echo "###############"
echo $CERTS_EXPIRE_IN_NEXT_30DAYS | jq -r '.common_name' | wc -l
