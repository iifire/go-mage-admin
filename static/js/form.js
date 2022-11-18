//列表页直接创建ajax表单


function mioForm(formId, validationUrl, options){
	this.initialize(formId, validationUrl, options)
};
mioForm.prototype = {
    initialize : function(formId, validationUrl, options){
        this.formId = formId;
        this.validationUrl = validationUrl;
        this.submitUrl = $('#'+this.formId).attr('action');

		this.submitCallback = (options && options.submitCallback) ? options.submitCallback : false;
        if($('#'+this.formId)){
            this.validator  = new Validation(this.formId);
        }
        this.errorSections = {};
    },
    serialize : function() {
    	
    	//for(instance in CKEDITOR.instances){
			//CKEDITOR.instances[instance].updateElement();
		//}
    	return $('#'+this.formId).serialize();
    },
    validate : function(){
        if(this.validator && this.validator.validate()){
            if(this.validationUrl){
                this._validate();
            }
            return true;
        }
        return false;
    },
    submit : function(url){
        this.errorSections = {};
        this.canShowError = true;
        this.submitUrl = url;
        CKupdate();
        if(this.validator && this.validator.validate()){
            if(this.validationUrl){
                this._validate();
            }
            else{
            	
                this._submit();
            }
            return true;
        }
        return false;
    },
	//异步验证
    _validate : function(){
    	$.ajax({
    		type: "POST",
    		url: this.validationUrl,
    		data: $('#'+this.formId).serialize(),
    		dataType: "JSON",
    		success: function(response) {
    			this._processValidationResult(response);
    		},
    		error: function(response) {
    			this._processFailure(response);
    		}
    	})
    },

	//验证返回成功
    _processValidationResult : function(response){
        if(response.error){
            if($('#messages')){
                $('#messages').innerHTML = response.message;
            }
        }
        else{
            this._submit();
        }
    },
	
	//验证请求失败
    _processFailure : function(response){
    	console.log(response);
        //location.href = BASE_URL;
    },
    _submit : function(){
    	$this = this;
        var $form = $('#'+this.formId);
        //判断是否Ajax提交
        if ($form.hasClass('ajax-submit')) {
        	$.ajax({
	    		type: "POST",
	    		url: $('#'+$this.formId).attr('action'),
	    		data: $('#'+$this.formId).serialize(),
	    		dataType: "JSON",
	    		success: function(response) {
	    			if($this.submitCallback){
	    				$this.submitCallback(response);
    				}
	    		},
	    		error: function(response) {
	    			$this._processFailure(response);
	    		}
	    	})
        } else {
        	if(this.submitUrl){
	            $form.action = $this.submitUrl;
	        }
	        $form.submit();
        }
    }
}

//表单元素依赖关系
function FormElementDependenceController(elementsMap){
	this.initialize(elementsMap)
};

FormElementDependenceController.prototype = {
    initialize : function(elementsMap){
    	
        for (var idTo in elementsMap) {
            for (var idFrom in elementsMap[idTo]) {
                if ($('input[name='+idFrom+']').length) {
                   var v = $('input[name='+idFrom+']').val();
                   //alert(v);
                   //alert(elementsMap[idTo][idFrom]);
                   if (v!=elementsMap[idTo][idFrom]) {
                   		$('#'+idTo).parent().parent().hide();
                   }
                   $('input[name='+idFrom+']').change(function(){
                   	   var v = $(this).val();
	                   if (v!=elementsMap[idTo][idFrom]) {
	                   		$('#'+idTo).parent().parent().hide();
	                   } else {
	                   		$('#'+idTo).parent().parent().show();
	                   }
                   })
                }
            }
        }
    }
}

function Validator(){
	this.initialize();
}

Validator.prototype = {
    initialize : function() {
        this.error = '验证失败。';
    },
    //测试
    test : function(elm) {
    	var obj = Validator.methods;
    	elm.siblings('.validation-advice').remove();
    	for(var m in obj){
    		if (elm.hasClass(obj[m][0])) {
    			if (obj[m][2](elm.val())) {
    				//验证通过
    			} else {
    				//验证不通过
    				var adviceHtml = '<div class="validation-advice" id="advice-' + obj[m][0] + '-' + elm.id +'" style="display:block">' + obj[m][1] + '</div>'
    				elm.after(adviceHtml);
    				//tips(obj[m][1],0);
    				break;
    			}
    		}  
      	}
    }
}
function $F(elm) {
	return $('#'+elm).val();
}
function $IsEmpty(v) {
 	return  (v == '' || (v == null) || (v.length == 0) || /^\s+$/.test(v));
}
Validator.methods = [
	['required-entry', '必填项', function(v) {
            return !$IsEmpty(v);
        }
    ]
];
function Validation(form, options) {
	this.initialize(form, options);
}
Validation.defaultOptions = {
    containerClassName: '.input-box',
};

Validation.prototype = {
    initialize : function(form, options){
        this.form = $('#'+form);
        if (!this.form) {
            return;
        }
    },
    onChange : function (ev) {
        Validation.isOnChange = false;
    },
    onSubmit :  function(ev){
        if(!this.validate()) {
        	return false;
        }
    },
    validate : function() {
        var result = false;
        try {
        	//遍历元素
        	this.form.find('input,select,textarea,file').each(function(index,element){
        		var $this = $(this);
        		//开始验证
        		var v = new Validator();
        		v.test($this);
        	})
        	if (this.form.find('.validation-advice').length==0) {
        		result = true;
        	}
        	
        } catch (e) {
        	console.log(e);
        }
        return result;
    },
    reset : function() {
        //重置
    },
    
    isVisible : function(elm) {
    	//元素是否可见
        return true;
    },
    createAdvice : function(name, elm, useTitle, customError) {
       //显示错误提示
    },
    get : function(name) {
        return  Validation.methods[name] ? Validation.methods[name] : null;
    },
    addAllThese : function(validators) {
    	//加入验证器
        var nv = {};
        return;
    }
}

function removeDelimiters (v) {
    v = v.replace(/\s/g, '');
    v = v.replace(/\-/g, '');
    return v;
}
function parseNumber(v)
{
    if (typeof v != 'string') {
        return parseFloat(v);
    }

    var isDot  = v.indexOf('.');
    var isComa = v.indexOf(',');

    if (isDot != -1 && isComa != -1) {
        if (isComa > isDot) {
            v = v.replace('.', '').replace(',', '.');
        }
        else {
            v = v.replace(',', '');
        }
    }
    else if (isComa != -1) {
        v = v.replace(',', '.');
    }

    return parseFloat(v);
}


