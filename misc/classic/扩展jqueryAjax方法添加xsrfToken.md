###   来自beego 群记录：

~~~
// [青岛]slene(309394958)  23:47:20
var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});
// [青岛]slene(309394958)  23:48:15
// 我在提供个每个站内 ajax 请求都增加 xsrf token 的方法。
~~~

### 重要：
以上思路可以用在任何语言中 不限于go web编程！！！
上面的xsrf 是go beego框架的  其他语言的实现可能不一样 比如php 的YII框架