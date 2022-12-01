(window.webpackJsonp=window.webpackJsonp||[]).push([[72],{417:function(t,e,a){"use strict";a.r(e);var l=a(10),s=Object(l.a)({},(function(){var t=this,e=t._self._c;return e("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[e("h1",{attrs:{id:"_6-5-etcd"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#_6-5-etcd"}},[t._v("#")]),t._v(" 6.5 ETCD")]),t._v(" "),e("h2",{attrs:{id:"范式"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#范式"}},[t._v("#")]),t._v(" 范式")]),t._v(" "),e("p",[e("a",{attrs:{href:"https://github.com/douyu/jupiter/tree/master/client/etcd/config.go",target:"_blank",rel:"noopener noreferrer"}},[t._v("参考地址"),e("OutboundLink")],1)]),t._v(" "),e("table",[e("thead",[e("tr",[e("th",{staticStyle:{"text-align":"left"}},[t._v("名称")]),t._v(" "),e("th",{staticStyle:{"text-align":"left"}},[t._v("类型")]),t._v(" "),e("th",{staticStyle:{"text-align":"left"}},[t._v("描述")])])]),t._v(" "),e("tbody",[e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("endpoints")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("[]string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("etcd连接地址")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("certFile")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("证书")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("keyFile")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("私钥")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("caCert")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("认证授权文件")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("basicAuth")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("bool")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("是否开启简单用户名/密码验证")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("userName")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("用户名")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("password")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("string")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("密码")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("connectTimeout")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("time")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("连接超时时间")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("secure")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("bool")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("是否开启与ETCD服务器证书验证，默认false")])]),t._v(" "),e("tr",[e("td",{staticStyle:{"text-align":"left"}},[e("code",[t._v("autoAsyncInterval")])]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("time")]),t._v(" "),e("td",{staticStyle:{"text-align":"left"}},[t._v("自动同步etcd member list的间隔")])])])]),t._v(" "),e("h2",{attrs:{id:"示例"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#示例"}},[t._v("#")]),t._v(" 示例")]),t._v(" "),e("div",{staticClass:"language-toml line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-toml"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),e("span",{pre:!0,attrs:{class:"token table class-name"}},[t._v("jupiter.etcdv3.myetcd")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n    "),e("span",{pre:!0,attrs:{class:"token key property"}},[t._v("endpoints")]),t._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),e("span",{pre:!0,attrs:{class:"token string"}},[t._v('"127.0.0.1:2379"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n    "),e("span",{pre:!0,attrs:{class:"token key property"}},[t._v("connectTimeout")]),t._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[t._v('"10s"')]),t._v("\n")])]),t._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[t._v("1")]),e("br"),e("span",{staticClass:"line-number"},[t._v("2")]),e("br"),e("span",{staticClass:"line-number"},[t._v("3")]),e("br")])])])}),[],!1,null,null,null);e.default=s.exports}}]);