// Vue 更换变量分隔符
Vue.options.delimiters = ['${', '}'];
var mage = mage || {};
mage.admin = (function() {
    function goPage(url) {
        history.replaceState(null,null, url);
        $.showLoading();
        $.ajax({
            type:'GET',
            url:url,
            beforeSend:function(req) {
                req.setRequestHeader("ajax-menu",url);
            },
            success:function(res){
                $('#menu').html($(res).find('#menu').val());
                $('#menucur').html($(res).find('#menucur').val());
                $('#main').html($(res).find('#main').html());
                getHistoryMenu(url);
                resizeBoard();
            },
            complete:function(res) {
                $.removeLoading();
            }
        })
    }
    function delPage(url,urlReplace) {
        $.post('/admin/page/menuDel',{url:url},function(res){
            if (res.code==1) {
                vueHeader.delHistory(url);
                if (urlReplace) {
                    goPage(urlReplace)
                }
            }
        },'json')
    }
    function getHistoryMenu(url) {
        $.get('/admin/page/menuHistory',{},function(res){
            if (res.code==1) {
                vueHeader.renderHistory(res.data,url);
            }
        },'json')
    }
    function resizeBoard() {
        let h = document.documentElement.clientHeight-$('.app-wrap > .header').height();
        if (h<300) {
            h = 300;
        }
        $('.inner-wrap').css('height',h+'px').css('overflow','hidden');
        if ($('.hor-scroll').length) {
            $('.hor-scroll').css('height',(h-$('.hor-scroll').offset().top-75)+'px').css('overflow-y','scroll');
        }
    }
    function init() {
        resizeBoard();
        $(window).resize(function () {
            resizeBoard();
        })
    }
    return {
        init:init,
        goPage: goPage,
        delPage: delPage,
        resizeBoard:resizeBoard,
    }
})();
$(function(){
    mage.admin.init();
})

String.prototype.replaceAll  = function(s1,s2){
    return this.replace(new RegExp(s1,"gm"),s2);
}
if (!Array.isArray) {
    Array.isArray = function(arg) {
        return Object.prototype.toString.call(arg) === '[object Array]';
    };
}
function timestampToTime(timestamp,onlyDay) {
    var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    var D = date.getDate() + ' ';
    var h =(date.getHours() < 10 ? '0':'') +date.getHours() + ':';
    var m = date.getMinutes() + ':';
    var s = date.getSeconds();
    return Y+M+D+(onlyDay ? "" : h+m+s );
}
