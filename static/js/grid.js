function mioGrid(containerId, url, pageVar, sortVar, dirVar, filterVar){
	this.initialize(containerId, url, pageVar, sortVar, dirVar, filterVar)
};
mioGrid.prototype = {
    initialize : function(containerId, url, pageVar, sortVar, dirVar, filterVar){
        this.containerId = containerId;
        this.url = url;
        this.pageVar = pageVar || false;
        this.sortVar = sortVar || false;
        this.dirVar  = dirVar || false;
        this.filterVar  = filterVar || false;
        this.tableSufix = '_table';
        this.useAjax = false;
        this.rowClickCallback = false;
        this.checkboxCheckCallback = false;
        this.preInitCallback = false;
        this.initCallback = false;
        this.initRowCallback = false;
        this.doFilterCallback = false;

        this.reloadParams = false;
		
        this.initGrid();
    },
    initGrid : function(){
        if(this.preInitCallback){
            this.preInitCallback(this);
        }
        if($(this.containerId+this.tableSufix)){
            this.rows = $('#'+this.containerId+this.tableSufix+' tbody tr');
            var row=0;
            this.rows.each(function(row,e){
            	if(row%2==0){
                    $(e).addClass('even');
                }
            });
            
        }
        if(this.initCallback){
            try {
                this.initCallback(this);
            }
            catch (e) {
                if(console) {
                    console.log(e);
                }
            }
        }
    },
    initGridAjax: function () {
        this.initGrid();
        this.initGridRows();
    },
    initGridRows: function() {
        if(this.initRowCallback){
            for (var row=0; row<this.rows.length; row++) {
                try {
                    this.initRowCallback(this, this.rows[row]);
                } catch (e) {
                    if(console) {
                        console.log(e);
                    }
                }
            }
        }
    },
    getContainerId : function(){
        return this.containerId;
    },
    rowMouseOver : function(event){
        var element = Event.findElement(event, 'tr');

        if (!element.title) return;

        Element.addClassName(element, 'on-mouse');

        if (!Element.hasClassName('pointer')
            && (this.rowClickCallback !== openGridRow || element.title)) {
            if (element.title) {
                Element.addClassName(element, 'pointer');
            }
        }
    },
    rowMouseOut : function(event){
        var element = Event.findElement(event, 'tr');
        Element.removeClassName(element, 'on-mouse');
    },
    rowMouseClick : function(event){
        if(this.rowClickCallback){
            try{
                this.rowClickCallback(this, event);
            }
            catch(e){}
        }
        varienGlobalEvents.fireEvent('gridRowClick', event);
    },
    rowMouseDblClick : function(event){
        varienGlobalEvents.fireEvent('gridRowDblClick', event);
    },
    keyPress : function(event){

    },
    doSort : function(event){
        var element = Event.findElement(event, 'a');

        if(element.name && element.title){
            this.addVarToUrl(this.sortVar, element.name);
            this.addVarToUrl(this.dirVar, element.title);
            this.reload(this.url);
        }
        Event.stop(event);
        return false;
    },
    loadByElement : function(element){
        if(element && element.name){
            this.reload(this.addVarToUrl(element.name, element.value));
        }
    },
    reload : function(url){
        if (!this.reloadParams) {
            this.reloadParams = {form_key: FORM_KEY};
        }
        else {
            this.reloadParams.form_key = FORM_KEY;
        }
        var $this = this;
        url = url || this.url;
		//alert(this.reloadParams);
        if (this.useAjax) {
        	//如果用AJAX
        	var url = url + (url.match(new RegExp('\\?')) ? '&ajax=true' : '?ajax=true' );
        	var params = this.reloadParams || {};
        	$.get(url,params,function(res){
        		if(res.match("^\{(.+:.+,*){1,}\}$")) {
        			res = $.parseJSON(res);
        			if(res.ajaxExpired && res.ajaxRedirect) {
                        setLocation(response.ajaxRedirect);
                    }
        		} else {
        			
        			var strObj = $(res);
        			$('#'+$this.containerId).html(strObj.find('.grid').parent());
        			
        			if(typeof(myGridSelected)=="function") {
        				myGridSelected();
        			}
        			
        		}
        	})
        } else {
        	for(var p in this.reloadParams ){
        		if(typeof(this.reloadParams[p])!="function") {
        			url = this.addVarToUrl(p, this.reloadParams[p]);
        		}
        	}
            location.href = url;
        }
    },
    _processFailure : function(transport){
        location.href = BASE_URL;
    },
    _addVarToUrl : function(url, varName, varValue){
    	
        var re = new RegExp('\/('+varName+'\/.*?\/)');
        var parts = url.split(new RegExp('\\?'));
        url = parts[0].replace(re, '/');
        url+= varName+'/'+varValue+'/';
        if(parts.length>1) {
            url+= '?' + parts[1];
        }
        return url;
    },
    addVarToUrl : function(varName, varValue){
        this.url = this._addVarToUrl(this.url, varName, varValue);
        return this.url;
    },
    doExport : function(){
        if($(this.containerId+'_export')){
            var exportUrl = $('#'+this.containerId+'_export').val();
            if(this.massaction && this.massaction.checkedString) {
                exportUrl = this._addVarToUrl(exportUrl, this.massaction.formFieldNameInternal, this.massaction.checkedString);
            }
            location.href = exportUrl;
        }
    },
    bindFilterFields : function(){
    	
    	var $this = this;
        $('#'+this.containerId+' .filter input, #'+this.containerId+' .filter select').each(function(){
        	//console.log('bindFilterFields');
        	$(this).keydown(function(event){
	        	
	        	if (event.which==13) {
			    	$this.doFilter();
				}
	        })
        
        })
	        
    },
    bindFieldsChange : function(){
        if (!$(this.containerId)) {
            return;
        }
        var dataElements = $(this.containerId+this.tableSufix).find('tbody input,tbody select');
    },
    filterKeyPress : function(event){
    	//console.log(event.keyCode);
        if(event.keyCode==13){
            this.doFilter();
        }
    },
    doFilter : function(){
        var filters = $('.grid-filter-element');
       
        var elements = [];
        filters.each(function(i,e){
            if($(e).val()) {
            	elements.push(encodeURIComponent($(e).attr('name'))+'='+encodeURIComponent($(e).val()));
            }
        })
        
        //console.log(elements);return;
        elements = elements.join('&');
        if (!this.doFilterCallback || (this.doFilterCallback && this.doFilterCallback())) {
            this.reload(this.addVarToUrl(this.filterVar, encode_base64(elements)));
        }
    },
    resetFilter : function(){
        this.reload(this.addVarToUrl(this.filterVar, ''));
    },
    
    //直接输入分页
    inputPage : function(event, maxNum){
        var element = event.target;
        var keyCode = event.keyCode || event.which;
        if(keyCode==13){
            this.setPage(element.value);
        }
    },
    setPage : function(pageNumber){
    	$('.page[name=page]').val(pageNumber);
        this.reload(this.addVarToUrl(this.pageVar, pageNumber));
    }
};

function openGridRow(grid, event){
    var element = Event.findElement(event, 'tr');
    if(['a', 'input', 'select', 'option'].indexOf(Event.element(event).tagName.toLowerCase())!=-1) {
        return;
    }

    if(element.title){
        setLocation(element.title);
    }
}
function mioGridMassaction(containerId, grid, checkedValues, formFieldNameInternal, formFieldName){
	this.initialize(containerId, grid, checkedValues, formFieldNameInternal, formFieldName)
};
mioGridMassaction.prototype = {
    /* Predefined vars */
    checkedValues: {},
    checkedString: '',
    oldCallbacks: {},
    errorText:'',
    items: {},
    gridIds: [],
    useSelectAll: false,
    currentItem: false,
    lastChecked: { left: false, top: false, checkbox: false },
    fieldTemplate: '<input type="hidden" name="#{name}" value="#{value}" />',
    initialize: function (containerId, grid, checkedValues, formFieldNameInternal, formFieldName) {
        this.useAjax        = false;
        this.grid           = grid;
        this.grid.massaction = this;
        this.containerId    = containerId;
        this.initMassactionElements();

        this.checkedString          = checkedValues;
        this.formFieldName          = formFieldName;
        this.initCheckboxes();
        this.checkCheckboxes();
        
    },
    setUseAjax: function(flag) {
        this.useAjax = flag;
    },
    setUseSelectAll: function(flag) {
        this.useSelectAll = flag;
    },
    initMassactionElements: function() {
    	
        this.container      = $('#'+this.containerId);
        this.count          = $('#'+this.containerId + '-count');
        this.formHiddens    = $('#'+this.containerId + '-form-hiddens');
        this.formAdditional = $('#'+this.containerId + '-form-additional');
        this.select         = $('#'+this.containerId + '-select');
        this.form           = this.prepareForm();
        this.validator      = new Validation(this.form.id);
        this.lastChecked    = { left: false, top: false, checkbox: false };
        this.initMassSelect();
        var $this = this;
        $("#"+this.containerId+' select').bind('change',function(){
        	//console.log($this.formAdditional);
        	var item = $(this);
        	//console.log($this.containerId + '-item-' + item.val() + '-block');
        	if(item.val()) {
	            $this.formAdditional.html($('#'+$this.containerId + '-item-' + item.val() + '-block').html());
	        } else {
	            $this.formAdditional.html('');
	        }
        })
    },
    prepareForm: function() {
        var form = $('#'+this.containerId + '-form'), formPlace = null,
            formElement = this.formHiddens || this.formAdditional;

        if (!formElement) {
            formElement = this.container.getElementsByTagName('button')[0];
            formElement && formElement.parentNode;
        }
        if (!form && formElement) {
            /* fix problem with rendering form in FF through innerHTML property */
            form = document.createElement('form');
            form.setAttribute('method', 'post');
            form.setAttribute('action', '');
            form.id = this.containerId + '-form';
            formPlace = formElement.parentNode.parentNode;
            formPlace.parentNode.appendChild(form);
            form.appendChild(formPlace);
        }

        return form;
    },
    setGridIds: function(gridIds) {
        this.gridIds = gridIds;
        this.updateCount();
    },
    getGridIds: function() {
        return this.gridIds;
    },
    setItems: function(items) {
        this.items = items;
        this.updateCount();
    },
    getItems: function() {
        return this.items;
    },
    getItem: function(itemId) {
    	
        if(this.items[itemId]) {
            return this.items[itemId];
        }
        return false;
    },

    initCheckboxes: function() {
       
    },
    checkCheckboxes: function() {
    	var curArr = this.getCheckedValues().split(',');
        $('.massaction-checkbox').each(function(){
        	if ($.inArray($(this).val(),curArr)>=0) {
        		$(this).prop("checked",true);
        		
        	} else {
        		$(this).removeAttr('checked');
        	}
        	
    	})
    },
    selectAll: function() {
    	var allArr = [];
    	$('.massaction-checkbox').each(function(){
    		allArr.push($(this).val());
    	})
    	
        this.setCheckedValues(allArr.join(','));
        this.checkCheckboxes();
        this.updateCount();
        return false;
    },
    unselectAll: function() {
        this.setCheckedValues('');
        this.checkCheckboxes();
        this.updateCount();
        
        return false;
    },
    
    setCheckedValues: function(values) {
        this.checkedString = values;
    },
    getCheckedValues: function() {
        return this.checkedString;
    },
    
    getCheckboxesValues: function() {
        var result = [];
      
        $('.massaction-checkbox').each(function() {
            result.push($(this).val());
        });
        return result;
    },
    getCheckboxesValuesAsString: function()
    {
        return this.getCheckboxesValues().join(',');
    },
    
    updateCount: function() {
    
        this.count.html(varienStringArray.count(this.checkedString));
        if(!this.grid.reloadParams) {
            this.grid.reloadParams = {};
        }
        
    },
    getSelectedItem: function() {
        if(this.getItem(this.select.val())) {
            return this.getItem(this.select.val());
        } else {
            return false;
        }
    },
    apply: function() {
        if(varienStringArray.count(this.checkedString) == 0) {
                alert(this.errorText);
                return;
            }

        var item = this.getSelectedItem();
        console.log(item);
        if(!item) {
            this.validator.validate();
            return;
        }
        
        this.currentItem = item;
        var fieldName = (item.field ? item.field : this.formFieldName);
        var fieldsHtml = '';

        if(this.currentItem.confirm && !window.confirm(this.currentItem.confirm)) {
            return;
        }

        this.formHiddens.html('');
        
        //设置隐藏表单域
        this.fieldTemplate = this.fieldTemplate.replace('#{name}',fieldName);
        this.fieldTemplate = this.fieldTemplate.replace('#{value}',this.checkedString);
		this.formHiddens.html(this.fieldTemplate);
        if(!this.validator.validate()) {
            return;
        }
		
        if(item.url) {
            this.form.attr('action',item.url);
            this.form.submit();
        }
    },
    collectCheck: function() {
    	var allArr = [];
    	$('.massaction-checkbox:checked').each(function(){
    		allArr.push($(this).val());
    	})
        this.setCheckedValues(allArr.join(','));
        this.updateCount();
    },
   
    initMassSelect: function() {
    	var _this = this
        $('.massaction-checkbox').click(function(){
        	_this.collectCheck();
        });
        
    },
};

var mioGridAction = {
    execute: function(select) {
        if(!select.value) {
            return;
        }

        var config = $.parseJSON(select.value);
        if(config.confirm && !window.confirm(config.confirm)) {
            select.options[0].selected = true;
            return;
        }

        if(config.popup) {
            var win = window.open(config.href, 'action_window', 'width=500,height=600,resizable=1,scrollbars=1');
            win.focus();
            select.options[0].selected = true;
        } else {
            setLocation(config.href);
        }
    }
};

var varienStringArray = {
    remove: function(str, haystack)
    {
        haystack = ',' + haystack + ',';
        haystack = haystack.replace(new RegExp(',' + str + ',', 'g'), ',');
        return this.trimComma(haystack);
    },
    add: function(str, haystack)
    {
        haystack = ',' + haystack + ',';
        if (haystack.search(new RegExp(',' + str + ',', 'g'), haystack) === -1) {
            haystack += str + ',';
        }
        return this.trimComma(haystack);
    },
    has: function(str, haystack)
    {
        haystack = ',' + haystack + ',';
        if (haystack.search(new RegExp(',' + str + ',', 'g'), haystack) === -1) {
            return false;
        }
        return true;
    },
    count: function(haystack)
    {
        if (typeof haystack != 'string') {
            return 0;
        }
        if (match = haystack.match(new RegExp(',', 'g'))) {
            return match.length + 1;
        } else if (haystack.length != 0) {
            return 1;
        }
        return 0;
    },
    each: function(haystack, fnc)
    {
        var haystack = haystack.split(',');
        for (var i=0; i<haystack.length; i++) {
            fnc(haystack[i]);
        }
    },
    trimComma: function(string)
    {
        string = string.replace(new RegExp('^(,+)','i'), '');
        string = string.replace(new RegExp('(,+)$','i'), '');
        return string;
    }
};


function encode_base64( what )
{
    var base64_encodetable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    var result = "";
    var len = what.length;
    var x, y;
    var ptr = 0;
   

    while( len-- > 0 )
    {
        x = what.charCodeAt( ptr++ );
        result += base64_encodetable.charAt( ( x >> 2 ) & 63 );

        if( len-- <= 0 )
        {
            result += base64_encodetable.charAt( ( x << 4 ) & 63 );
            result += "==";
            break;
        }

        y = what.charCodeAt( ptr++ );
        result += base64_encodetable.charAt( ( ( x << 4 ) | ( ( y >> 4 ) & 15 ) ) & 63 );

        if ( len-- <= 0 )
        {
            result += base64_encodetable.charAt( ( y << 2 ) & 63 );
            result += "=";
            break;
        }

        x = what.charCodeAt( ptr++ );
        result += base64_encodetable.charAt( ( ( y << 2 ) | ( ( x >> 6 ) & 3 ) ) & 63 );
        result += base64_encodetable.charAt( x & 63 );

    }

    return result;
}


function decode_base64( what )
{
    var base64_decodetable = new Array (
        255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
        255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
        255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,  62, 255, 255, 255,  63,
         52,  53,  54,  55,  56,  57,  58,  59,  60,  61, 255, 255, 255, 255, 255, 255,
        255,   0,   1,   2,   3,   4,   5,   6,   7,   8,   9,  10,  11,  12,  13,  14,
         15,  16,  17,  18,  19,  20,  21,  22,  23,  24,  25, 255, 255, 255, 255, 255,
        255,  26,  27,  28,  29,  30,  31,  32,  33,  34,  35,  36,  37,  38,  39,  40,
         41,  42,  43,  44,  45,  46,  47,  48,  49,  50,  51, 255, 255, 255, 255, 255
    );
    var result = "";
    var len = what.length;
    var x, y;
    var ptr = 0;

    while( !isNaN( x = what.charCodeAt( ptr++ ) ) )
    {
        if( x == 13 || x == 10 )
            continue;

        if( ( x > 127 ) || (( x = base64_decodetable[x] ) == 255) )
            return false;
        if( ( isNaN( y = what.charCodeAt( ptr++ ) ) ) || (( y = base64_decodetable[y] ) == 255) )
            return false;

        result += String.fromCharCode( (x << 2) | (y >> 4) );

        if( (x = what.charCodeAt( ptr++ )) == 61 )
        {
            if( (what.charCodeAt( ptr++ ) != 61) || (!isNaN(what.charCodeAt( ptr ) ) ) )
                return false;
        }
        else
        {
            if( ( x > 127 ) || (( x = base64_decodetable[x] ) == 255) )
                return false;
            result += String.fromCharCode( (y << 4) | (x >> 2) );
            if( (y = what.charCodeAt( ptr++ )) == 61 )
            {
                if( !isNaN(what.charCodeAt( ptr ) ) )
                    return false;
            }
            else
            {
                if( (y > 127) || ((y = base64_decodetable[y]) == 255) )
                    return false;
                result += String.fromCharCode( (x << 6) | y );
            }
        }
    }
    return result;
}

function wrap76( what )
{
    var result = "";
    var i;

    for(i=0; i < what.length; i+=76)
    {
        result += what.substring(i, i+76) + String.fromCharCode(13) + String.fromCharCode(10);
    }
    return result;
}