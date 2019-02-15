const ModuleNamesTag="MODULE_NAMES";
class CModuleInfo {
    constructor(){
        this.moduleNameArray = [];
        this.loadModuleInfoAjax();
        this.storeToSessionStorage();
    }
    
    // Method : load from server with AJAX.
    loadModuleInfoAjax(){
        //Test codes
        this.moduleNameArray = ["Module#1", "Module#2"];
    }
    
    storeToSessionStorage(){    
        sessionStorage.setItem(ModuleNamesTag, JSON.stringify(this.moduleNameArray));        
    }
    
    // Getter
    get moduleNames() {
        return this.moduleNameArray;
    }
}

function readModulesArrayFromSessionStorage(){
    var mnamesStored = sessionStorage.getItem(ModuleNamesTag);
    if(typeof(mnamesStored) == "undefined" || mnamesStored == null ){
        return null;
    }
    return JSON.parse(mnamesStored);
}
