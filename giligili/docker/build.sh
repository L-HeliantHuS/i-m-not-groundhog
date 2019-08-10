npm run build
cp -r ../dist ./
docker build -t registry.cn-hangzhou.aliyuncs.com/helianthusihm/giligili:v0.0.1 ./
docker push registry.cn-hangzhou.aliyuncs.com/helianthusihm/giligili:v0.0.1
