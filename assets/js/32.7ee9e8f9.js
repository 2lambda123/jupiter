(window.webpackJsonp=window.webpackJsonp||[]).push([[32],{354:function(s,a,t){"use strict";t.r(a);var e=t(10),n=Object(e.a)({},(function(){var s=this,a=s._self._c;return a("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[a("h1",{attrs:{id:"docker"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#docker"}},[s._v("#")]),s._v(" Docker")]),s._v(" "),a("h2",{attrs:{id:"_1-2-docker安装"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_1-2-docker安装"}},[s._v("#")]),s._v(" 1.2 docker安装")]),s._v(" "),a("h3",{attrs:{id:"_1-2-1-安装"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_1-2-1-安装"}},[s._v("#")]),s._v(" 1.2.1 安装")]),s._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" clone https://github.com/douyu/juno-install.git\n\n"),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" juno-install\n\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("docker")]),s._v(" build "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-t")]),s._v(" juno-install:v1  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-f")]),s._v(" ./docker/all-in-one/Dockerfile ./\n\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("docker")]),s._v(" run "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-itd")]),s._v("  "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--name")]),s._v(" juno-demo "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("50000")]),s._v(":50000 "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[s._v("50004")]),s._v(":50004 "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--privileged")]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("true juno-install:v1 /usr/sbin/init\n\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("docker")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("exec")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-it")]),s._v(" juno-demo /bin/bash\n\n./install.sh\n\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br"),a("span",{staticClass:"line-number"},[s._v("2")]),a("br"),a("span",{staticClass:"line-number"},[s._v("3")]),a("br"),a("span",{staticClass:"line-number"},[s._v("4")]),a("br"),a("span",{staticClass:"line-number"},[s._v("5")]),a("br"),a("span",{staticClass:"line-number"},[s._v("6")]),a("br"),a("span",{staticClass:"line-number"},[s._v("7")]),a("br"),a("span",{staticClass:"line-number"},[s._v("8")]),a("br"),a("span",{staticClass:"line-number"},[s._v("9")]),a("br"),a("span",{staticClass:"line-number"},[s._v("10")]),a("br"),a("span",{staticClass:"line-number"},[s._v("11")]),a("br"),a("span",{staticClass:"line-number"},[s._v("12")]),a("br")])]),a("h2",{attrs:{id:"_1-2-2-调试"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_1-2-2-调试"}},[s._v("#")]),s._v(" 1.2.2 调试")]),s._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("#查看mysql")]),s._v("\nmysql "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-uroot")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("#查看etcd里内容")]),s._v("\netcdctl get "),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('""')]),s._v(" "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--prefix")]),s._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("#查看go")]),s._v("\ngo version\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("#查看pprof")]),s._v("\ngo tool pprof "),a("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-http")]),a("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),a("span",{pre:!0,attrs:{class:"token string"}},[s._v('":8081"')]),s._v(" http://127.0.0.1:50004/debug/pprof/profile\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br"),a("span",{staticClass:"line-number"},[s._v("2")]),a("br"),a("span",{staticClass:"line-number"},[s._v("3")]),a("br"),a("span",{staticClass:"line-number"},[s._v("4")]),a("br"),a("span",{staticClass:"line-number"},[s._v("5")]),a("br"),a("span",{staticClass:"line-number"},[s._v("6")]),a("br"),a("span",{staticClass:"line-number"},[s._v("7")]),a("br"),a("span",{staticClass:"line-number"},[s._v("8")]),a("br")])])])}),[],!1,null,null,null);a.default=n.exports}}]);