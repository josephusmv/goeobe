

function initControls(){
    initRadiobtns();  
    bindCollapsDetection();
    bindSideBarRelatedBtns();
}

function initRadiobtns(){
    createDateTimeRadioButtons();
    createPriorityRadioButtons();
    createUrgencyRadioButtons();
    bindDateTimeRadioBtn();
    bindPriorityRadioBtn();
    bindurgencyRadioBtn();
}

var crtdttmRBTNGrp = "crtdttmRBTN"
var crtdttmValueArry = ["OneDay", "OneWeek", "OneMonth", "OneYear", "All"];
var crtdttmTextArry = ["Within 1 Day", "Within 1 Week", "Within 1 Month", "Within 1 Year", "Show all"];


function createDateTimeRadioButtons(){
    createRadioButtons( $("#filterByDTTM"), crtdttmRBTNGrp, crtdttmValueArry, crtdttmTextArry);
}

function bindDateTimeRadioBtn(){
    var holderName = "input:radio[name=" + crtdttmRBTNGrp + "]";
    $(holderName).change(function() {
        radioBtnOnChange(this, crtdttmValueArry, crtdttmTextArry);   
    });
}

function createRadioButtons(holder, name, valueArray, textArray){
    var formEle = $("<form><form>");
    
    for(i=0; i < valueArray.length; i++) {
        var radioOption = $("<input></input>");
        radioOption.attr('id', valueArray[i]);
        radioOption.attr('type','radio');
        radioOption.attr('name',name);
        radioOption.attr('value',valueArray[i]);
        radioOption.addClass("filter-radio-btn");
        formEle.append(radioOption);
        
        var labeltext = $("<label></label>");
        labeltext.attr('for', valueArray[i]);
        labeltext.addClass("filter-radio-btn");
        labeltext.text(" " + textArray[i]);
        radioOption.after("<br>");
        radioOption.after(labeltext);
        //radioOption.after("&nbsp;" + textArray[i] + "<br>");
    }
    
    holder.append(formEle);
}

function radioBtnOnChange(rbtngrp, valueArray, textArray){
    for(i=0; i < valueArray.length; i++) {
        if (rbtngrp.value == valueArray[i]){
            alert(valueArray[i] + " : " + textArray[i]);
        }
    }
}


var priorityBtnGrp = "priorityBtnGrp"
var priorityIDArry = ["TopPriority", "HighPriority", "MediumPriority", "LowPriority", "AllPriority"];
var btncolorArray = ["btn-danger", "btn-warning", "btn-primary", "btn-secondary", "btn-success"];
var priorityTextArry = ["Top", "H", "M", "L", "All"];


function createPriorityRadioButtons(){ 
    createButtonGroups( $("#filterByPriority"), priorityBtnGrp, priorityIDArry, btncolorArray, priorityTextArry, "prty-btn");  
}

function bindPriorityRadioBtn(){
    for(i=0; i < priorityIDArry.length; i++) {     
        var id=priorityIDArry[i]
        $("#"+ id).click(function(){            
            priorityBtnOnclick(this.id);
        });
    }    
}

function priorityBtnOnclick(id){
    alert(id);
}

var urgencyBtnGrp = "urgencyBtnGrp"
var urgencyIDArry = ["ImediateUrg", "WeeksUrg", "SeasonPlan", "LongtimerPlan", "AllUrg"];
var btnUrgColorArray = ["btn-danger", "btn-warning", "btn-primary", "btn-info", "btn-success"];
var urgencyTextArry = ["Urg", "H", "M", "L", "All"];


function createUrgencyRadioButtons(){ 
    createButtonGroups( $("#filterByUrgency"), urgencyBtnGrp, urgencyIDArry, btnUrgColorArray, urgencyTextArry, "prty-btn");  
}

function bindurgencyRadioBtn(){
    for(i=0; i < urgencyIDArry.length; i++) {     
        var id=urgencyIDArry[i]
        $("#"+ id).click(function(){            
            urgencyBtnOnclick(this.id);
        });
    }    
}

function urgencyBtnOnclick(id){
    alert(id);
}


function createButtonGroups(holder, name, idArray, btncolorArray, textArray, classS2){
    var btngrpEle = $("<div></div>");
    //btngrpEle.addClass("btn-group");
    btngrpEle.addClass("mr-2");
    btngrpEle.attr("role", "group");
    btngrpEle.attr("aria-label", name);
    
    for(i=0; i < idArray.length; i++) {
        var btn = $("<button></button>");
        btn.attr("id", idArray[i]);
        btn.attr("type", "button");
        btn.addClass("btn");
        btn.addClass(btncolorArray[i]);
        btn.addClass(classS2);
        btn.text(textArray[i]);
        btngrpEle.append(btn);
    }
    holder.append(btngrpEle);
}


//Side bar statues and change class...
var gIsSideBarOpen = true;

function bindSideBarRelatedBtns(){
    $("#closeSideBar").click(function(){
        if(gIsSideBarOpen){
            changeOpenFilterImageAfterCollapsed(false);
            changeMainDivAfterCllapsed(false);
        } else {
            changeOpenFilterImageAfterCollapsed(true);
            changeMainDivAfterCllapsed(true);
        }
    });
    
    $("#clpsTrigger").click(function(){
        if(gIsSideBarOpen){
            changeOpenFilterImageAfterCollapsed(false);
            changeMainDivAfterCllapsed(false);
        } else {
            changeOpenFilterImageAfterCollapsed(true);
            changeMainDivAfterCllapsed(true);
        }
    });
}

function changeOpenFilterImageAfterCollapsed(collapsed){
    var closeImg = $("#closeImg");
    if (collapsed) {
        closeImg.removeClass("openimage");
        closeImg.addClass("closeimage");
    } else {
        closeImg.removeClass("closeimage");
        closeImg.addClass("openimage");
    }
}

function changeMainDivAfterCllapsed(collapsed){
    var mccnter = $("#MainContentContainer");
    if (collapsed) {
        mccnter.removeClass("col-md-12");
        mccnter.addClass("col-md-9");
    } else {
        mccnter.removeClass("col-md-9");
        mccnter.addClass("col-md-12");
    }
}

function sideBarCollapseEventtriggered(collapsed){
    gIsSideBarOpen = collapsed;
    calculateFrameDivSize(collapsed);
}

function bindCollapsDetection(){
    $('#sidebar').on('shown.bs.collapse', function () {
        sideBarCollapseEventtriggered(true);
    });

    $('#sidebar').on('hidden.bs.collapse', function () {
        sideBarCollapseEventtriggered(false);
    });
}