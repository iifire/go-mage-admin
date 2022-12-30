// Vue 更换变量分隔符
Vue.options.delimiters = ['${', '}'];
var mage = mage || {};
mage.admin = (function() {
    function goPage(url,showLast) {
        history.replaceState(null,null, url);
        $.showLoading();
        $.ajax({
            type:'GET',
            url:url,
            beforeSend:function(req) {
                req.setRequestHeader("ajax-menu",url);
            },
            success:function(res){
                $('#menu').val($(res).find('#menu').val());
                $('#menucur').val($(res).find('#menucur').val());

                $('#main').html($(res).find('#main').html());
                //处理历史菜单
                getHistoryMenu(url,showLast);
                resizeBoard();
            },
            complete:function(res) {
                $.removeLoading();
            }
        })
    }
    function getHistoryMenu(url,showLast) {
        $.get('/admin/page/menuHistory',{},function(res){
            if (res.code==1) {
                vueHeader.renderHistory(res.data,url);
                vueHeader.switchMenuCopy($('#menu').val(),url);
            }
        },'json')
        if (showLast) {
            $('.cmenu-ctl-r').click();
        }
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
