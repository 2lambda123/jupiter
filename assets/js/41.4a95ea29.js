(window.webpackJsonp=window.webpackJsonp||[]).push([[41],{367:function(t,s,a){"use strict";a.r(s);var e=a(10),n=Object(e.a)({},(function(){var t=this,s=t._self._c;return s("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[s("h1",{attrs:{id:"_13-2-系统消息事件配置"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#_13-2-系统消息事件配置"}},[t._v("#")]),t._v(" 13.2 系统消息事件配置")]),t._v(" "),s("p",[t._v("默认情况下，"),s("code",[t._v("Juno")]),t._v(" 系统会将事件写入名为 "),s("code",[t._v("appevent")]),t._v(" 的数据表中。如果你希望通过 "),s("code",[t._v("MQ")]),t._v(" 获取 "),s("code",[t._v("Juno")]),t._v(" 系统的事件，那么可以通过下述配置来实现。目前系统事件消息只支持了 "),s("code",[t._v("RocketMQ")]),t._v("。")]),t._v(" "),s("h2",{attrs:{id:"配置示例"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#配置示例"}},[t._v("#")]),t._v(" 配置示例:")]),t._v(" "),s("div",{staticClass:"language-toml line-numbers-mode"},[s("pre",{pre:!0,attrs:{class:"language-toml"}},[s("code",[s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),s("span",{pre:!0,attrs:{class:"token table class-name"}},[t._v("junoevent.rocketmq")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("enable")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("false")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# 开关.如果为false，则系统事件不写MQ.")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("addr")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"127.0.0.1:9876"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# mq地址")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("topic")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"juno_test_job"')]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# mq topic")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("group")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"juno_devops_go"')]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# mq producer group")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("retry")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("3")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# MQ写重试次数")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token key property"}},[t._v("dialTimeout")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"3s"')]),t._v(" "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# MQ连接超时时间")]),t._v("\n")])]),t._v(" "),s("div",{staticClass:"line-numbers-wrapper"},[s("span",{staticClass:"line-number"},[t._v("1")]),s("br"),s("span",{staticClass:"line-number"},[t._v("2")]),s("br"),s("span",{staticClass:"line-number"},[t._v("3")]),s("br"),s("span",{staticClass:"line-number"},[t._v("4")]),s("br"),s("span",{staticClass:"line-number"},[t._v("5")]),s("br"),s("span",{staticClass:"line-number"},[t._v("6")]),s("br"),s("span",{staticClass:"line-number"},[t._v("7")]),s("br")])]),s("h3",{attrs:{id:"配置项说明"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#配置项说明"}},[t._v("#")]),t._v(" 配置项说明:")]),t._v(" "),s("table",[s("thead",[s("tr",[s("th",{staticStyle:{"text-align":"left"}},[t._v("配置项")]),t._v(" "),s("th",{staticStyle:{"text-align":"left"}},[t._v("字段说明")]),t._v(" "),s("th",{staticStyle:{"text-align":"right"}},[t._v("示例值")])])]),t._v(" "),s("tbody",[s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("enable")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("开关。是否打开本功能，如果为 true ， JUNO 系统事件会发送消息到配置的 MQ")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[s("code",[t._v("true")])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("addr")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("MQ 地址")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[s("code",[t._v('["127.0.0.1:9876"]')])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("topic")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("RocketMQ 的Topic名称")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[s("code",[t._v('"xxxx"')])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("group")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("Producer Group 名称")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[s("code",[t._v("xxx")])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("retry")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("重试次数")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[t._v("3")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("dialTimeout")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("连接超时时间")]),t._v(" "),s("td",{staticStyle:{"text-align":"right"}},[s("code",[t._v('"3s"')])])])])]),t._v(" "),s("h2",{attrs:{id:"消息格式"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#消息格式"}},[t._v("#")]),t._v(" 消息格式:")]),t._v(" "),s("div",{staticClass:"language-javascript line-numbers-mode"},[s("pre",{pre:!0,attrs:{class:"language-javascript"}},[s("code",[s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"id"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("                "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 事件ID")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"app_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("         "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 应用名称")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"aid"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("              "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 应用ID")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"zone_code"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("        "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// Zone Code")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"env"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("              "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 环境")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"host_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("        "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 主机名称")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"operator_type"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("    "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 用户类型 (user | openapi)")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"user_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("        "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 用户名")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"uid"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("              "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 用户ID")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"operation"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("        "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 事件类型")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"create_time"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("      "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 事件创建时间")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"source"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("           "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 来源")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"metadata"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("         "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 事件详情 (与事件类型有关)")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"operation_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("   "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 事件操作名称")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"source_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),t._v("       "),s("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 来源名称")]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])]),t._v(" "),s("div",{staticClass:"line-numbers-wrapper"},[s("span",{staticClass:"line-number"},[t._v("1")]),s("br"),s("span",{staticClass:"line-number"},[t._v("2")]),s("br"),s("span",{staticClass:"line-number"},[t._v("3")]),s("br"),s("span",{staticClass:"line-number"},[t._v("4")]),s("br"),s("span",{staticClass:"line-number"},[t._v("5")]),s("br"),s("span",{staticClass:"line-number"},[t._v("6")]),s("br"),s("span",{staticClass:"line-number"},[t._v("7")]),s("br"),s("span",{staticClass:"line-number"},[t._v("8")]),s("br"),s("span",{staticClass:"line-number"},[t._v("9")]),s("br"),s("span",{staticClass:"line-number"},[t._v("10")]),s("br"),s("span",{staticClass:"line-number"},[t._v("11")]),s("br"),s("span",{staticClass:"line-number"},[t._v("12")]),s("br"),s("span",{staticClass:"line-number"},[t._v("13")]),s("br"),s("span",{staticClass:"line-number"},[t._v("14")]),s("br"),s("span",{staticClass:"line-number"},[t._v("15")]),s("br"),s("span",{staticClass:"line-number"},[t._v("16")]),s("br"),s("span",{staticClass:"line-number"},[t._v("17")]),s("br")])]),s("h3",{attrs:{id:"消息字段说明"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#消息字段说明"}},[t._v("#")]),t._v(" 消息字段说明:")]),t._v(" "),s("table",[s("thead",[s("tr",[s("th",{staticStyle:{"text-align":"left"}},[t._v("字段名")]),t._v(" "),s("th",{staticStyle:{"text-align":"left"}},[t._v("说明")])])]),t._v(" "),s("tbody",[s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("id")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("事件ID")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("app_name")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("应用名称")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("aid")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("应用ID")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("zone_code")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("Zone Code")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("env")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("环境")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("host_name")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("主机名称")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("operator_type")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("用户类型 (user: 用户操作; openapi: 通过OpenAPI操作)")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("user_name")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("用户名")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("uid")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("用户ID")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("operation")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("事件类型，见 "),s("a",{attrs:{href:"#%E4%BA%8B%E4%BB%B6%E7%B1%BB%E5%9E%8B"}},[t._v("事件类型")])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("create_time")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("事件创建时间")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("source")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("来源，见 "),s("a",{attrs:{href:"#%E6%9D%A5%E6%BA%90%E7%B1%BB%E5%9E%8B"}},[t._v("来源")])])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("metadata")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("事件详情 (与事件类型有关，不同的事件类型数据格式不同)")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("operation_name")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("事件操作名称")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("source_name")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("来源名称")])])])]),t._v(" "),s("p",[t._v("事件消息示例:")]),t._v(" "),s("div",{staticClass:"language-javascript line-numbers-mode"},[s("pre",{pre:!0,attrs:{class:"language-javascript"}},[s("code",[s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"id"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("11979")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"app_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"bj-im-srv-tencent-callback-go"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"aid"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("14185")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"zone_code"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"ALIYUN-HB2-G"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"env"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"prod"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"host_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('""')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"user_name"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"杜旻翔_gitlab"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"uid"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"operation"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"confgo_file_update"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"create_time"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token number"}},[t._v("1603942730")]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"source"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"confgo"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"metadata"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"{\\"access_token_id\\":0,\\"change_log\\":\\"1\\",\\"configuration_id\\":988,\\"format\\":\\"toml\\",\\"id\\":0,\\"name\\":\\"test\\",\\"uid\\":1,\\"version\\":\\"50347a3f14aea923e9f8eac867fd3bb1\\"}"')]),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n    "),s("span",{pre:!0,attrs:{class:"token string-property property"}},[t._v('"operator_type"')]),s("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),s("span",{pre:!0,attrs:{class:"token string"}},[t._v('"user"')]),t._v("\n"),s("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])]),t._v(" "),s("div",{staticClass:"line-numbers-wrapper"},[s("span",{staticClass:"line-number"},[t._v("1")]),s("br"),s("span",{staticClass:"line-number"},[t._v("2")]),s("br"),s("span",{staticClass:"line-number"},[t._v("3")]),s("br"),s("span",{staticClass:"line-number"},[t._v("4")]),s("br"),s("span",{staticClass:"line-number"},[t._v("5")]),s("br"),s("span",{staticClass:"line-number"},[t._v("6")]),s("br"),s("span",{staticClass:"line-number"},[t._v("7")]),s("br"),s("span",{staticClass:"line-number"},[t._v("8")]),s("br"),s("span",{staticClass:"line-number"},[t._v("9")]),s("br"),s("span",{staticClass:"line-number"},[t._v("10")]),s("br"),s("span",{staticClass:"line-number"},[t._v("11")]),s("br"),s("span",{staticClass:"line-number"},[t._v("12")]),s("br"),s("span",{staticClass:"line-number"},[t._v("13")]),s("br"),s("span",{staticClass:"line-number"},[t._v("14")]),s("br"),s("span",{staticClass:"line-number"},[t._v("15")]),s("br")])]),s("h3",{attrs:{id:"事件类型"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#事件类型"}},[t._v("#")]),t._v(" 事件类型")]),t._v(" "),s("table",[s("thead",[s("tr",[s("th",{staticStyle:{"text-align":"left"}},[t._v("事件类型")]),t._v(" "),s("th",{staticStyle:{"text-align":"left"}},[t._v("说明")])])]),t._v(" "),s("tbody",[s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("confgo_file_create")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置创建")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("confgo_file_update")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置更新")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("confgo_file_delete")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置删除")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("confgo_file_publish")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置发布")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("confgo_file_rollback")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置回滚")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("app_node_restart")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("应用重启")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("grafana_alert_notification")])]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[s("code",[t._v("Grafana")]),t._v(" 监控告警")])])])]),t._v(" "),s("h3",{attrs:{id:"来源类型"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#来源类型"}},[t._v("#")]),t._v(" 来源类型")]),t._v(" "),s("table",[s("thead",[s("tr",[s("th",{staticStyle:{"text-align":"left"}},[t._v("Source")]),t._v(" "),s("th",{staticStyle:{"text-align":"left"}},[t._v("说明")])])]),t._v(" "),s("tbody",[s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("confgo")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("配置中心")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("git")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("git 事件")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("devops")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("发布事件")])]),t._v(" "),s("tr",[s("td",{staticStyle:{"text-align":"left"}},[t._v("grafana")]),t._v(" "),s("td",{staticStyle:{"text-align":"left"}},[t._v("监控告警")])])])])])}),[],!1,null,null,null);s.default=n.exports}}]);