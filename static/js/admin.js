// Vue 更换变量分隔符
Vue.options.delimiters = ['${', '}'];
var mage = mage || {};
mage.admin = (function() {
    function goPage(url,showLast) {
        history.replaceState(null,null, url);
        var url = mage.admin.formatUrl(url);
        $.showLoading();
        $.ajax({
            type:'GET',
            url:url,
            beforeSend:function(req) {
                req.setRequestHeader("ajax-menu",url);
            },
            success:function(res){
                $('#main').html($(res).find('#main').html());
                $('#cmenu').html($(res).find('#cmenu').html());
                _renderMenu();
                mage.admin.reset();
                if (showLast) {
                    $('.cmenu-ctl-r').click();
                }
                _resizeBoard();
            },
            complete:function(res) {
                $.removeLoading();
            }
        })
    }
    function formatUrl(url) {
        url = url.replace(URL,'');
        if (url.indexOf('/')!=0) {
            url = '/'+url;
        }
        return url;
    }
    function _renderMenu() {
        $('a.v-ajax').off('click').click(function(){
            var url = $(this).attr('href') ? $(this).attr('href') : $(this).data('url');
            $('li.level0').removeClass("active");
            $(this).parents('i.level0').addClass('active');
            mage.admin.goPage(url);
            return false;
        })
        $('.m-ajax').off('click').click(function(){
            mage.admin.goPage(URL+$(this).data('url'));
        })
        //menus resize
        var _cmenuLen = $('#cmenu').width();
        var _cmenuBoxLen = $('#cmenu-box').width();
        if (_cmenuBoxLen>_cmenuLen) {
            $('#cmenu').addClass('scroll');
        } else {
            $('#cmenu').removeClass('scroll');
        }
        $('.cmenu-ctl-l').off('click').click(function(){
            $('#cmenu-box').css('transform','translate3d(0, 0px, 0px)');
        })
        $('.cmenu-ctl-r').off('click').click(function(){
            if (_cmenuBoxLen>_cmenuLen) {
                $('#cmenu-box').css('transform','translate3d(-'+(_cmenuBoxLen-_cmenuLen+15)+'px, 0px, 0px)');
            }
        })
        $('.m-ajax i').off('click').click(function(event){
            var _this = $(this);
            event.stopPropagation();
            //删除当前&激活最后一个
            $.post(URL+'b/ajax/delMenu',{form_key:FORM_KEY,'ajax-menu':$(this).parent().data('url')},function(res){
                if (res.code==1) {
                    $.showLoading();

                    _this.parent().remove();
                    if ($('.m-ajax').length>1) {
                        var _url = $('.m-ajax:last-child').data('url');
                        mage.admin.goPage(URL+_url,true);
                    } else {
                        window.location.href = URL+'b/dashboard/';
                    }
                }
            },'json');

        })
    }
    function _resizeBoard() {
        var h = document.documentElement.clientHeight-$('.app-wrap > .header').height();
        if (h<300) {
            h = 300;
        }
        $('.inner-wrap').css('height',h+'px').css('overflow','hidden');
        if ($('.hor-scroll').length) {
            $('.hor-scroll').css('height',(h-$('.hor-scroll').offset().top-75)+'px').css('overflow-y','scroll');
        }

    }
    function init() {
        _renderMenu();
        _resizeBoard();
        $(window).resize(function () {
            _resizeBoard();
        })
    }
    return {
        goPage: goPage,
        formatUrl:formatUrl,
        init:init,
        resizeBoard:_resizeBoard,
    }
})();
$(function(){
    mage.admin.init();
})
String.prototype.padStart = function padStart(targetLength,padString) {
    targetLength = targetLength>>0;
    padString = String((typeof padString !== 'undefined' ? padString : ' '));
    if (this.length > targetLength) {
        return String(this);
    }
    else {
        targetLength = targetLength-this.length;
        if (targetLength > padString.length) {
            padString += padString.repeat(targetLength/padString.length);
        }
        return padString.slice(0,targetLength) + String(this);
    }
};
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
//后台表单图片上传
function formRenderBase64(htmlId) {
    var img = document.getElementById(htmlId)
    var imgFile = new FileReader();
    imgFile.readAsDataURL(img.files[0]);
    imgFile.onload = function () {
        if (this.result) {
            if ($('#'+htmlId+'_image').length>0) {
                $('#'+htmlId+'_image').attr('src',this.result);
            } else {
                if ($('#'+htmlId).parent().siblings('#'+htmlId+'_image').length>0) {
                    $('#'+htmlId+'_image').attr('src',this.result);
                } else {
                    $('#'+htmlId).parent().before('<img id="'+htmlId+'_image" src="'+this.result+'" class="small-image-preview v-middle" height="40">');
                }
            }
        }
    }
}