```bash
cd ../
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/tiktok" -table="video" -dir="./videoModel" -c
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/tiktok" -table="follow" -dir="./followModel" -c
```