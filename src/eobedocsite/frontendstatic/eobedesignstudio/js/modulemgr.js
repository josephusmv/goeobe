function loadModules(){
    //send request to get module info
    //store module to sessionStorage
}

const ModuleNamesTag="MODULE_NAMES";
function updateModules(moduleNames){
    if(typeof(moduleNames) == "undefined" || moduleNames == null || !Array.isArray(moduleNames)){
        alert("Input parameter error for moduleNames");
        return;
    }
    
    sessionStorage.setItem(ModuleNamesTag, JSON.stringify(moduleNames));
}

function readModulesArray(){
    var mnamesStored = sessionStorage.getItem(ModuleNamesTag);
    if(typeof(mnamesStored) == "undefined" || mnamesStored == null ){
        return null;
    }
    return JSON.parse(mnamesStored);
}

function updateModuleSideNavBar(){
    var moduleName = readModulesArray();
    for (x in moduleName) { 
        console.log(moduleName[x]);
        var liEle = createModuleEntryInNaveBar(moduleName[x]);
        $(liEle).insertAfter("#sideNavConfigItem");
    }
}

const moduleHTMLStr = "<a class=\"nav-link nav-link-collapse collapsed collapseBtnItem\" data-toggle=\"collapse\" href=\"#collapseComponents\" data-parent=\"#exampleAccordion\"><i class=\"fas fa-caret-down\"></i><span class=\"nav-link-text moduleTitleItem\">*ModuleNames*</span> </a>   <ul class=\"sidenav-second-level collapse collapseTargetItem sideNaveCollapsible\" id=\"collapseComponents\">  <li>  <a href=\"httpactions.html\">HTTPActions</a> </li>   <li><a href=\"dbactions.html\">DBActions</a> </li> <li> <a href=\"userconst.html\">User Constants</a></li> </ul>";

// <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Components">
function createModuleEntryInNaveBar(moduleName){    
  var liEle = $("<li></li>");
  liEle.addClass("nav-item");
  liEle.attr("id", "snav-"+moduleName);
  liEle.attr("data-toggle", "tooltip");
  liEle.attr("data-placement", "right");
  liEle.attr("title", "Components");
  liEle.html(moduleHTMLStr);
  
  liEle.find(".collapseBtnItem").attr("href","#collapseModule"+moduleName);
  liEle.find(".collapseTargetItem").attr("id","collapseModule"+moduleName);
  liEle.find(".moduleTitleItem").text(moduleName);
  liEle.on('show.bs.collapse', function () {
    $(".sideNaviBar").find(".sideNaveCollapsible").collapse('hide');
  });

  // Change this to div.childNodes to support multiple top-level nodes
  return liEle; 
}
