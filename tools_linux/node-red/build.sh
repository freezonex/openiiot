docker build -t eco-registry.supos.com/supvxcen/node-red:1.0.0 --platform linux/amd64 . 
docker login eco-registry.supos.com --username=zhanghongzhi@freezonex.io
docker push eco-registry.supos.com/supvxcen/node-red:1.0.0