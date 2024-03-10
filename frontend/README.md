# Frontend

## Lang

- vue
- vuetify
- vite

## Structure

- dist
- node_modules
- public
- src
  - asserts
    _Used to put the static files, like css, img_
  - components
    _Used to put the .vue components, which can be used repeatedly._
  - hooks
    _Used to put the public functions, which can be used repeatedly._
    - checkValid.ts
      _The functions used to check the inputs in form is valid_
  - pages
    _Used to struct the pages. There are two files in each dir, App.vue and main.js._
    - login
    - member
    - signup
  - plugins

## Tips

When you git push your code to github, it is not be recommended to push the node_modules dir (abspath: frontend/vue/node_modules). The reason is on [https://segmentfault.com/q/1010000018517286](https://segmentfault.com/q/1010000018517286 "https://segmentfault.com/q/1010000018517286"). Just following the .gitignore will be fine.

> node_modules 不是你自己的源代码，而是存放你在 package.json 中指定的依赖的外部库和框架的文件，这些依赖库少的有几百 M 多的好几 G，而且又不你的源代码，其内容已经存储在其他服务器上，没必要放到 github，你只要把 package.json 放上去，别人下载你的 package.json，运行个 npm install 就会自动把 node_modules 文件夹建立起来，和你自己的 node_modules 文件夹里面的内容一模一样。
