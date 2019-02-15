function initModuleNaviLinks(){
    //send request to get module info
    gModules = new CModuleInfo();    //should be global
    console.log(gModules.moduleNames);
    //store module to sessionStorage
    updateModuleSideNavBar();
    $(".sideNaveCollapsible").on('show.bs.collapse', function () {
        $(".sideNaviBar").find(".sideNaveCollapsible").collapse('hide');
    });
}

function updateModuleSideNavBar(){
    var moduleName = readModulesArrayFromSessionStorage();
    for (x in moduleName) { 
        console.log(moduleName[x]);
        var liEle = createModuleEntryInNaveBar(moduleName[x]);
        $(liEle).insertAfter("#sideNavConfigItem");
    }
}

const moduleHTMLStr = "<a class=\"nav-link nav-link-collapse collapsed collapseBtnItem\" data-toggle=\"collapse\" href=\"#collapseComponents\" data-parent=\"#exampleAccordion\"><i class=\"fas fa-caret-down\"></i><span class=\"nav-link-text moduleTitleItem\">*ModuleNames*</span> </a>   <ul class=\"sidenav-second-level collapse collapseTargetItem sideNaveCollapsible\" id=\"collapseComponents\">  <li>  <a href=\"resdefine\">HTTPActions</a> </li>   <li><a href=\"dbactions\">DBActions</a> </li> <li> <a href=\"userconst\">User Constants</a></li> </ul>";

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