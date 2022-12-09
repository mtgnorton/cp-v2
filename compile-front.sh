frontProjectPath="/Users/mtgnorton/Coding/front/cp-v2-admin-ui"
backendProjectPath="/Users/mtgnorton/Coding/go/src/github.com/mtgnorton/cp-v2"

#前端文件打包并且移动到后端目录中
cd $frontProjectPath || exit
npm run build:prod
rm -rf ${backendProjectPath}/public/resource-admin
rm -f ${backendProjectPath}/template/admin/index.html
#rm -f ${backendProjectPath}/public/favicon.ico

mv ${frontProjectPath}/dist/index.html ${backendProjectPath}/template/admin/admin.html
mv ${frontProjectPath}/dist/resource-admin ${backendProjectPath}/public/
