/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
 function AddNewAPIEntry(){
    var apiName = $("#apiname").val()
    var apidesc = $("#apidesc").val()
    var formData = new FormData();
    formData.append("apiname", apiName);
    formData.append("apidesc", apidesc);
    
    var xhr = new XMLHttpRequest();
    xhr.open('POST', "AddNewAPIEntry", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
          alert(xhr.responseText);
        }
    };
    xhr.send(formData);
}
// allowed data source, 0xE(14) for variables only, 0x1 for literal only, 0xF for both
var allwsrclist = {
   '0xE': 'variables',
   '0x1':'literal',
   '0xF':'both'
}
/* 0x3(0011): string,  0xC(1100): int*/
var datatypelist = {
    '0x3':'string',
    '0xC':'int'
}

var ParamIndex = 0
//AddNewAPIParam
//apiname, paramname, allwsrc, datatype, posnum, paramdesc
function AddParam() {
    if (ParamIndex > 6){
        return;
    }
    var divName = "newParam"+ParamIndex;
    var $div = $("<div>", {id: divName, class:"parambox"});
    ParamIndex+=1;
    
    var $paramName = $("<input>", {id: "paramname", "type":"text", class:"inputsitems"});
   $div.append($paramName);
  $div.append("<br/>");
    
    var $allwsrc = $('<select />', {id:"allwsrc", class:"inputsitems"});    
    for(var val in allwsrclist) {
        $('<option />', {value: val, text: allwsrclist[val]}).appendTo($allwsrc);
    }
   $div.append($allwsrc);
    
    var $datatype = $('<select />', {id:"datatype", class:"inputsitems"});    
    for(var val in datatypelist) {
        $('<option />', {value: val, text: datatypelist[val]}).appendTo($datatype);
    }
   $div.append($datatype);
    $div.append("<br/>");
    
    var $posnum = $("<input>", {"id": "posnum", "type":"number", "min":"0", "max":"5", class:"inputsitems"});
    $posnum.val(ParamIndex);
   $div.append($posnum);
   $div.append("<br/>");
    
    var $paramdesc = $("<textarea>", {id: "paramdesc", class:"inputsitems"});
    $div.append($paramdesc);  
   $div.append("<br/>");
    var $addbtn = $("<button>", { "type":"button", "onclick":"AddNewAPIParam(\""+divName+"\")", class:"inputsitems"});
    $addbtn.text("Add");
    $div.append($addbtn);  
    $("#paramlst").append($div);
    $("#paramlst").append("<br/>");
}

function AddNewAPIParam(divName){
    var apiname = $("#apiname").val();
    var paramname = $("#"+divName).children("#paramname").val();
    var allwsrc = $("#"+divName).children("#allwsrc").val();
    var datatype = $("#"+divName).children("#datatype").val();
    var posnum = $("#"+divName).children("#posnum").val();
    var paramdesc = $("#"+divName).children("#paramdesc").val();
    alert(apiname+", " + paramname + ", " + allwsrc+ ", " + datatype+ ", " + posnum + ", " + paramdesc );
}


function AddRslt() {
    var fileInput = document.getElementById('myFile');
    var file = fileInput.files[0];
    var formData = new FormData();
    formData.append('HasUploadFileInKey', fileInput.files[0].name)
    formData.append(fileInput.files[0].name, file);

    var xhr = new XMLHttpRequest();
    xhr.open('POST', "UploadFile", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
          alert(xhr.responseText);
        }
    };
    xhr.send(formData);
}